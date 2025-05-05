package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	ValuesFilePath string
}

// PCRMeasurements represents the PCR values from a commit message
type PCRMeasurements struct {
	HashAlgorithm string `json:"HashAlgorithm"`
	PCR0          string `json:"PCR0"`
	PCR1          string `json:"PCR1"`
	PCR2          string `json:"PCR2"`
}

// Commit represents a GitHub commit
type Commit struct {
	Message string `json:"message"`
}

// GitHubCommit represents the GitHub API response for a commit
type GitHubCommit struct {
	Commit Commit `json:"commit"`
}

// UpdateInfo contains the information needed to update values.yaml
type UpdateInfo struct {
	ImageVersion string
	PCRValues    PCRMeasurements
	NewIndex     int
}

// ValuesYAML represents the structure of values.yaml
type ValuesYAML struct {
	Bridge struct {
		Env map[string]string `yaml:"env"`
	} `yaml:"bridge"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <commit_url>", os.Args[0])
	}

	config := Config{
		ValuesFilePath: "charts/enclave-signer-registry/values.yaml",
	}

	commitURL := os.Args[1]
	commitMessage, err := fetchCommitMessage(commitURL)
	if err != nil {
		return fmt.Errorf("failed to fetch commit message: %w", err)
	}

	updateInfo, err := parseCommitMessage(commitMessage)
	if err != nil {
		return fmt.Errorf("failed to parse commit message: %w", err)
	}

	if err := updateValuesYAML(config.ValuesFilePath, updateInfo); err != nil {
		return fmt.Errorf("failed to update values.yaml: %w", err)
	}

	fmt.Printf("Successfully updated values.yaml with:\n")
	fmt.Printf("New PCR values added with index: %d\n", updateInfo.NewIndex)
	return nil
}

// parseCommitMessage extracts the image version and PCR measurements from a commit message
func parseCommitMessage(message string) (*UpdateInfo, error) {
	imageVersion, err := extractImageVersion(message)
	if err != nil {
		return nil, err
	}

	pcrMeasurements, err := extractPCRMeasurements(message)
	if err != nil {
		return nil, err
	}

	return &UpdateInfo{
		ImageVersion: imageVersion,
		PCRValues:    *pcrMeasurements,
	}, nil
}

// extractImageVersion extracts the image version from a commit message
func extractImageVersion(message string) (string, error) {
	re := regexp.MustCompile(`Update Image Version to (\w+)`)
	matches := re.FindStringSubmatch(message)
	if len(matches) < 2 {
		return "", fmt.Errorf("could not find image version in commit message")
	}
	return matches[1], nil
}

// extractPCRMeasurements extracts PCR measurements from a commit message
func extractPCRMeasurements(message string) (*PCRMeasurements, error) {
	// Find the first { and last } to extract the JSON
	start := strings.Index(message, "{")
	end := strings.LastIndex(message, "}")
	if start == -1 || end == -1 || start >= end {
		return nil, fmt.Errorf("could not find valid JSON in commit message")
	}

	jsonStr := message[start : end+1]
	var measurements PCRMeasurements
	if err := json.Unmarshal([]byte(jsonStr), &measurements); err != nil {
		fmt.Println(string(message))
		return nil, fmt.Errorf("failed to parse PCR measurements: %w", err)
	}

	return &measurements, nil
}

// updateValuesYAML updates the values.yaml file with new PCR values
func updateValuesYAML(filePath string, info *UpdateInfo) error {
	// Read the YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read values.yaml: %w", err)
	}

	// Parse the YAML to find the highest index
	var values ValuesYAML
	if err := yaml.Unmarshal(data, &values); err != nil {
		return fmt.Errorf("failed to parse values.yaml: %w", err)
	}

	// Find the highest existing PCR index
	maxIndex := -1
	for key := range values.Bridge.Env {
		if strings.HasPrefix(key, "PCRS_") {
			var index int
			if _, err := fmt.Sscanf(key, "PCRS_%d_", &index); err == nil && index > maxIndex {
				maxIndex = index
			}
		}
	}

	info.NewIndex = maxIndex + 1

	// Print commit message with PCR values
	commitURL := os.Args[1]
	// Extract owner/repo@sha format from URL
	// Example: https://github.com/owner/repo/commit/abc123 -> owner/repo@abc123
	re := regexp.MustCompile(`github\.com/([^/]+)/([^/]+)/commit/([0-9a-f]+)`)
	matches := re.FindStringSubmatch(commitURL)
	if len(matches) == 4 {
		fmt.Printf("\nAdding PCRs for %s/%s@%s:\n", matches[1], matches[2], matches[3])
	} else {
		fmt.Printf("\nAdding PCRs for %s:\n", commitURL)
	}

	// Convert YAML content to string
	content := string(data)

	// Find the bridge.env section
	envSectionStart := strings.Index(content, "  env:")
	if envSectionStart == -1 {
		return fmt.Errorf("could not find bridge.env section in values.yaml")
	}

	// Find the end of the env section (either next section or end of file)
	nextSectionStart := strings.Index(content[envSectionStart:], "\n  ")
	if nextSectionStart == -1 {
		nextSectionStart = len(content) - envSectionStart
	}

	// Create the new PCR entries
	newPCRs := fmt.Sprintf("    PCRS_%d_PCR0: '0x%s'\n    PCRS_%d_PCR1: '0x%s'\n    PCRS_%d_PCR2: '0x%s'",
		info.NewIndex, info.PCRValues.PCR0,
		info.NewIndex, info.PCRValues.PCR1,
		info.NewIndex, info.PCRValues.PCR2)

	// Insert the new PCR entries at the end of the env section
	insertPos := envSectionStart + nextSectionStart
	updatedContent := content[:insertPos] + "\n" + newPCRs + content[insertPos:]

	// Write back to file
	if err := os.WriteFile(filePath, []byte(updatedContent), 0644); err != nil {
		return fmt.Errorf("failed to write values.yaml: %w", err)
	}

	return nil
}

// fetchCommitMessage fetches the commit message from a GitHub URL
func fetchCommitMessage(url string) (string, error) {
	apiURL := convertToAPIURL(url)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch commit: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch commit: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var commit GitHubCommit
	if err := json.Unmarshal(body, &commit); err != nil {
		return "", fmt.Errorf("failed to parse commit data: %w", err)
	}

	return commit.Commit.Message, nil
}

// convertToAPIURL converts a GitHub URL to its API URL
func convertToAPIURL(url string) string {
	url = strings.Replace(url, "github.com", "api.github.com/repos", 1)
	return strings.Replace(url, "/commit/", "/commits/", 1)
}
