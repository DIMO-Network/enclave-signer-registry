package config

import (
	"github.com/DIMO-Network/enclave-bridge/pkg/config"
	"github.com/ethereum/go-ethereum/common"
)

// Settings contains the application config.
type Settings struct {
	Environment string `env:"ENVIRONMENT" yaml:"environment"`
	LogLevel    string `env:"LOG_LEVEL"   yaml:"logLevel"`
	Port        int    `env:"PORT"        yaml:"port"`
	MonPort     int    `env:"MON_PORT"    yaml:"monPort"`

	DeveloperLicense common.Address `env:"DEVELOPER_LICENSE" yaml:"developerLicense"`
	PrivateKey       string         `env:"PRIVATE_KEY"       yaml:"privateKey"`
	PCRs             []PCRValues    `env:"PCRS"              yaml:"pcrs"`

	// Ethereum settings
	EthereumRPCURL     string         `env:"ETHEREUM_RPC_URL"     yaml:"ethereumRpcUrl"`
	DevLicenseContract common.Address `env:"DEV_LICENSE_CONTRACT" yaml:"devLicenseContract"`

	TLS config.TLSConfig `envPrefix:"TLS_"`
}

// PCRValues contains the values for the PCR0, PCR1, and PCR2.
type PCRValues struct {
	PCR0 string `env:"PCR0"  yaml:"pcr0"`
	PCR1 string `env:"PCR1"  yaml:"pcr1"`
	PCR2 string `env:"PCR2"  yaml:"pcr2"`
}
