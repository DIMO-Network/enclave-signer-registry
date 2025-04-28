package app

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/DIMO-Network/enclave-signer-registry/internal/client/devlicense"
	"github.com/DIMO-Network/enclave-signer-registry/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/hf/nitrite"
	"github.com/rs/zerolog"
)

type Controller struct {
	privateKey        *ecdsa.PrivateKey
	devLicense        common.Address
	devLicenseTokenID *big.Int
	validPCRs         []config.PCRValues
	devLicenseClient  *devlicense.Client
}

func NewController(
	settings *config.Settings,
	devLicenseClient *devlicense.Client,
) (*Controller, error) {
	if settings.PrivateKey == "" {
		return nil, errors.New("private key is not set")
	}
	privateKey, err := crypto.HexToECDSA(settings.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hex to ecdsa private key: %w", err)
	}
	tokenID, err := devLicenseClient.GetTokenID(settings.DeveloperLicense)
	if err != nil {
		return nil, fmt.Errorf("failed to get token ID From DevLicense '%s': %w", settings.DeveloperLicense.Hex(), err)
	}
	if tokenID == nil || tokenID.Sign() == 0 {
		return nil, fmt.Errorf("token ID is nil or zero for DevLicense '%s'", settings.DeveloperLicense.Hex())
	}
	owner, err := devLicenseClient.GetOwner(tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to get owner from DevLicense '%s': %w", settings.DeveloperLicense.Hex(), err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	if owner.Hex() != publicAddress.Hex() {
		return nil, fmt.Errorf("'%s' is not owner of dev license '%s'", publicAddress.Hex(), settings.DeveloperLicense.Hex())
	}
	return &Controller{
		privateKey:        privateKey,
		devLicense:        settings.DeveloperLicense,
		devLicenseTokenID: tokenID,
		validPCRs:         settings.PCRs,
		devLicenseClient:  devLicenseClient,
	}, nil
}

type DeveloperLicenseResponse struct {
	ClientId  string             `json:"clientId"`
	TokenId   string             `json:"tokenId"`
	ValidPCRs []config.PCRValues `json:"validPCRs"`
}

// GetDeveloperLicense godoc
// @Summary Get developer license
// @Description Get the developer license of the controller
// @Tags keys
// @Accept json
// @Produce json
// @Success 200 {object} DeveloperLicenseResponse
// @Router /Developer-license [get]
func (c *Controller) GetDeveloperLicense(ctx *fiber.Ctx) error {
	developerLicenseResponse := DeveloperLicenseResponse{
		ClientId:  c.devLicense.String(),
		TokenId:   c.devLicenseTokenID.String(),
		ValidPCRs: c.validPCRs,
	}
	return ctx.JSON(developerLicenseResponse)
}

type AddSignerRequest struct {
	SignerAddress       string `json:"signerAddress"`
	AttestationDocument []byte `json:"attestationDocument"`
}
type AddSignerResponse struct {
	TxHash       string `json:"txHash"`
	AlreadyAdded bool   `json:"alreadyAdded"`
}

// AddSigner godoc
// @Summary Add a new signer
// @Description Add a new signer by verifying their NSM attestation
// @Tags signer
// @Accept json
// @Produce json
// @Param attestation body AddSignerRequest true "NSM attestation document"
// @Success 200 {object} AddSignerResponse
// @Failure 400 {object} codeResp
// @Failure 500 {object} codeResp
// @Router /add-signer [post]
func (c *Controller) AddSigner(ctx *fiber.Ctx) error {
	logger := zerolog.Ctx(ctx.UserContext())
	var request AddSignerRequest
	if err := ctx.BodyParser(&request); err != nil {
		logger.Error().Err(err).Msg("Failed to parse attestation document")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse attestation document")
	}

	// Verify the attestation document
	opts := nitrite.VerifyOptions{
		CurrentTime: time.Now(),
	}
	result, err := nitrite.Verify(request.AttestationDocument, opts)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to verify attestation document")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to verify attestation document")
	}
	if !result.SignatureOK {
		logger.Error().Msg("Attestation document signature verification failed")
		return fiber.NewError(fiber.StatusBadRequest, "Attestation document signature verification failed")
	}

	var valid bool
	for i := range c.validPCRs {
		if bytes.Equal(c.validPCRs[i].PCR0, result.Document.PCRs[0]) &&
			bytes.Equal(c.validPCRs[i].PCR1, result.Document.PCRs[1]) &&
			bytes.Equal(c.validPCRs[i].PCR2, result.Document.PCRs[2]) {
			valid = true
			break
		}
	}
	if !valid {
		// make PCRs easier to read
		attestationPCRs := config.PCRValues{
			PCR0: hexutil.Bytes(result.Document.PCRs[0]),
			PCR1: hexutil.Bytes(result.Document.PCRs[1]),
			PCR2: hexutil.Bytes(result.Document.PCRs[2]),
		}
		logger.Error().Any("attestationPCRs", attestationPCRs).Any("validPCRs", c.validPCRs).Msg("Attestation document PCR values do not match")
		return fiber.NewError(fiber.StatusBadRequest, "Attestation document PCR values do not match")
	}

	pubKey, err := crypto.UnmarshalPubkey(result.Document.PublicKey)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal public key")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to unmarshal public key")
	}
	docAddress := crypto.PubkeyToAddress(*pubKey)
	if request.SignerAddress != docAddress.String() {
		logger.Error().Any("signerAddress", request.SignerAddress).Any("docAddress", docAddress).Msg("Signer address does not match document address")
		return fiber.NewError(fiber.StatusBadRequest, "Signer address does not match document address")
	}

	logger.Info().Str("signerAddress", request.SignerAddress).Msg("Registering new signer")
	txHash, err := c.registerSigner(docAddress)
	alreadyAdded := false
	if err != nil {
		if !errors.Is(err, devlicense.ErrAlreadyEnabled) {
			logger.Error().Err(err).Msg("Failed to register signer")
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to register signer")
		}
		alreadyAdded = true
	}
	logger.Debug().Str("signerAddress", request.SignerAddress).Msg("Signer registered")
	return ctx.JSON(AddSignerResponse{
		TxHash:       txHash,
		AlreadyAdded: alreadyAdded,
	})
}

// registerSigner enables a signer for the dev license token ID.
func (c *Controller) registerSigner(signerAddress common.Address) (string, error) {
	tx, err := c.devLicenseClient.EnableSigner(c.privateKey, c.devLicenseTokenID, signerAddress)
	if err != nil {
		return "", fmt.Errorf("failed to enable signer: %w", err)
	}
	return tx.Hash().Hex(), nil
}
