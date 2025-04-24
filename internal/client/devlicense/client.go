package devlicense

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/DIMO-Network/enclave-signer-registry/internal/config"
	"github.com/DIMO-Network/enclave-signer-registry/internal/contracts/devlicensedimo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var ErrAlreadyEnabled = errors.New("signer already enabled")

// Client represents a client for interacting with the DevLicense contract.
// It handles enabling signers and other contract operations.
type Client struct {
	client   *ethclient.Client
	contract *devlicensedimo.Devlicensedimo
}

// NewClient creates a new DevLicense client using the provided configuration.
// It establishes a connection to the Ethereum network and initializes the contract instance.
// Returns an error if any of the initialization steps fail.
func NewClient(ctx context.Context, cfg *config.Settings, httpClient *http.Client) (*Client, error) {
	// Connect to the Ethereum client
	rpcClient, err := rpc.DialOptions(ctx, cfg.EthereumRPCURL, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	client := ethclient.NewClient(rpcClient)

	// Create the contract instance
	contract, err := devlicensedimo.NewDevlicensedimo(cfg.DevLicenseContract, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	return &Client{
		client:   client,
		contract: contract,
	}, nil
}

// EnableSigner enables a signer address for the developer license.
// It creates and sends a transaction to the contract to enable the specified signer.
// Returns an error if the transaction fails or cannot be mined.
func (c *Client) EnableSigner(privateKey *ecdsa.PrivateKey, devLicenseTokenID *big.Int, signerAddress common.Address) (*types.Transaction, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key is nil")
	}
	if devLicenseTokenID == nil {
		return nil, fmt.Errorf("dev license token ID is nil")
	}

	// Get the chain ID
	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create the transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction signer: %w", err)
	}
	// skip if the signer is already enabled
	enabled, err := c.contract.IsSigner(nil, devLicenseTokenID, signerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check if signer is enabled: %w", err)
	}
	if enabled {
		return nil, ErrAlreadyEnabled
	}
	// Call enableSigner
	tx, err := c.contract.EnableSigner(auth, devLicenseTokenID, signerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to enable signer: %w", err)
	}

	return tx, nil
}

// WaitForTransaction waits for a transaction to be mined and returns the transaction receipt.
func (c *Client) WaitForTransaction(tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction: %w", err)
	}
	return receipt, nil
}

// GetTokenID returns the token ID for the developer license.
func (c *Client) GetTokenID(devLicenseClientID common.Address) (*big.Int, error) {
	return c.contract.ClientIdToTokenId(nil, devLicenseClientID)
}

// GetOwner returns the owner of the developer license.
func (c *Client) GetOwner(devLicenseTokenID *big.Int) (common.Address, error) {
	return c.contract.OwnerOf(nil, devLicenseTokenID)
}
