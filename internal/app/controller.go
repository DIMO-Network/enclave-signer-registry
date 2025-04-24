package app

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/DIMO-Network/enclave-bridge/pkg/attest"
	"github.com/DIMO-Network/enclave-signer-registry/internal/client/devlicense"
	"github.com/DIMO-Network/enclave-signer-registry/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/hf/nitrite"
	"github.com/hf/nsm/request"
	"github.com/rs/zerolog"
)

type Controller struct {
	logger            *zerolog.Logger
	privateKey        *ecdsa.PrivateKey
	nsmResult         *nitrite.Result
	devLicense        common.Address
	devLicenseTokenID *big.Int
	validPCRs         []config.PCRValues
	getCertFunc       func(*tls.ClientHelloInfo) (*tls.Certificate, error)
	devLicenseClient  *devlicense.Client
}

func NewController(
	settings *config.Settings,
	logger *zerolog.Logger,
	getCertFunc func(*tls.ClientHelloInfo) (*tls.Certificate, error),
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
		logger:            logger,
		privateKey:        privateKey,
		devLicense:        settings.DeveloperLicense,
		devLicenseTokenID: tokenID,
		validPCRs:         settings.PCRs,
		getCertFunc:       getCertFunc,
		devLicenseClient:  devLicenseClient,
	}, nil
}

type NSMAttestationResponse struct {
	Result    *nitrite.Result `json:"attestation"`
	RawResult []byte          `json:"document"`
}

// GetNSMAttestations godoc
// @Summary Get NSM attestation
// @Description Get the Nitro Security Module attestation
// @Tags attestation
// @Accept json
// @Produce json
// @Param nonce query string false "Optional nonce for attestation"
// @Success 200 {object} NSMAttestationResponse
// @Failure 500 {object} codeResp
// @Router /.well-known/nsm-attestation [get]
func (c *Controller) GetNSMAttestations(ctx *fiber.Ctx) error {
	certBytes, err := c.getCert()
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to marshal certificate")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to marshal certificate")
	}
	nonce := ctx.Query("nonce")

	// if nonce is empty, check if the cached result is valid
	if nonce == "" && c.nsmResult != nil && len(c.nsmResult.Certificates) > 0 &&
		c.nsmResult.Certificates[0] != nil &&
		c.nsmResult.Certificates[0].NotAfter.After(time.Now()) &&
		bytes.Equal(c.nsmResult.Document.UserData, certBytes) {
		return ctx.JSON(c.nsmResult)
	}

	// I want to pass in nil when nonce is empty not sure if it matters
	var nonceBytes []byte
	if nonce != "" {
		nonceBytes = []byte(nonce)
	}

	req := &request.Attestation{
		UserData: certBytes,
		Nonce:    nonceBytes,
	}
	docuemnt, nsmResult, err := attest.GetNSMAttestation(req)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to get NSM attestation")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get NSM attestation")
	}
	if nonce == "" {
		// only cache the result if nonce is empty
		c.nsmResult = nsmResult
	}
	return ctx.JSON(NSMAttestationResponse{
		Result:    nsmResult,
		RawResult: docuemnt,
	})
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
	var request AddSignerRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.logger.Error().Err(err).Msg("Failed to parse attestation document")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse attestation document")
	}

	// Verify the attestation document
	opts := nitrite.VerifyOptions{
		CurrentTime: time.Now(),
	}
	result, err := nitrite.Verify(request.AttestationDocument, opts)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to verify attestation document")
		return fiber.NewError(fiber.StatusBadRequest, "Failed to verify attestation document")
	}
	if !result.SignatureOK {
		c.logger.Error().Msg("Attestation document signature verification failed")
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
		c.logger.Error().Any("attestationPCRs", attestationPCRs).Any("validPCRs", c.validPCRs).Msg("Attestation document PCR values do not match")
		return fiber.NewError(fiber.StatusBadRequest, "Attestation document PCR values do not match")
	}

	pubKey, err := crypto.UnmarshalPubkey(result.Document.PublicKey)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to unmarshal public key")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to unmarshal public key")
	}
	docAddress := crypto.PubkeyToAddress(*pubKey)
	if request.SignerAddress != docAddress.String() {
		c.logger.Error().Any("signerAddress", request.SignerAddress).Any("docAddress", docAddress).Msg("Signer address does not match document address")
		return fiber.NewError(fiber.StatusBadRequest, "Signer address does not match document address")
	}

	c.logger.Info().Str("signerAddress", request.SignerAddress).Msg("Registering new signer")
	txHash, err := c.registerSigner(docAddress)
	alreadyAdded := false
	if err != nil {
		if !errors.Is(err, devlicense.ErrAlreadyEnabled) {
			c.logger.Error().Err(err).Msg("Failed to register signer")
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to register signer")
		}
		alreadyAdded = true
	}
	c.logger.Debug().Str("signerAddress", request.SignerAddress).Msg("Signer registered")
	return ctx.JSON(AddSignerResponse{
		TxHash:       txHash,
		AlreadyAdded: alreadyAdded,
	})
}

func (c *Controller) getCert() ([]byte, error) {
	if c.getCertFunc == nil {
		// No certificate function configured, return nil
		return nil, nil
	}
	cert, err := c.getCertFunc(nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to get certificate")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get certificate")
	}
	if cert == nil {
		return nil, nil
	}

	certBytes, err := x509.MarshalPKIXPublicKey(cert.Certificate[0])
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to marshal certificate")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to marshal certificate")
	}
	return certBytes, nil
}

// registerSigner enables a signer for the dev license token ID.
func (c *Controller) registerSigner(signerAddress common.Address) (string, error) {
	tx, err := c.devLicenseClient.EnableSigner(c.privateKey, c.devLicenseTokenID, signerAddress)
	if err != nil {
		return "", fmt.Errorf("failed to enable signer: %w", err)
	}
	return tx.Hash().Hex(), nil
}
