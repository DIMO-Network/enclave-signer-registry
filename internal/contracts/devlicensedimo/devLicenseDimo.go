// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package devlicensedimo

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DevlicensedimoMetaData contains all meta data concerning the Devlicensedimo contract.
var DevlicensedimoMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LICENSE_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REVOKER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"adminBurnLockedFunds\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminFreeze\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"frozen_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminReallocate\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminWithdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clientIdToTokenId\",\"inputs\":[{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractDescription\",\"inputs\":[],\"outputs\":[{\"name\":\"contractDescription_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractImage\",\"inputs\":[],\"outputs\":[{\"name\":\"contractImage_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractURI\",\"inputs\":[],\"outputs\":[{\"name\":\"contractURI_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dimoCredit\",\"inputs\":[],\"outputs\":[{\"name\":\"dimoCredit_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dimoToken\",\"inputs\":[],\"outputs\":[{\"name\":\"dimoToken_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableSigner\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableSigner\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"frozen\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"frozen_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseAliasByTokenId\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenIdByLicenseAlias\",\"inputs\":[{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenIdByLicenseAlias\",\"inputs\":[{\"name\":\"licenseAlias\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"receiver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"licenseAccountFactory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dimoTokenAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dimoCreditAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"licenseCostInUsd_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"image_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"description_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSigner\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSigner_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issueInDc\",\"inputs\":[{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"issueInDc\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"issueInDimo\",\"inputs\":[{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"issueInDimo\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"licenseAccountFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"licenseAccountFactory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"licenseCostInUsd1e18\",\"inputs\":[],\"outputs\":[{\"name\":\"licenseCostInUsd1e18_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"locked_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periodValidity\",\"inputs\":[],\"outputs\":[{\"name\":\"periodValidity_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provider\",\"inputs\":[],\"outputs\":[{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"receiver_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redirectUriStatus\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revoke\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractDescription\",\"inputs\":[{\"name\":\"description_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractImage\",\"inputs\":[{\"name\":\"image_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDimoCreditAddress\",\"inputs\":[{\"name\":\"dimoCreditAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDimoTokenAddress\",\"inputs\":[{\"name\":\"dimoTokenAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLicenseAlias\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLicenseCost\",\"inputs\":[{\"name\":\"licenseCostInUsd1e18_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLicenseFactoryAddress\",\"inputs\":[{\"name\":\"licenseAccountFactory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeriodValidity\",\"inputs\":[{\"name\":\"periodValidity_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceProviderAddress\",\"inputs\":[{\"name\":\"providerAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReceiverAddress\",\"inputs\":[{\"name\":\"receiver_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRedirectUri\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenDescription\",\"inputs\":[{\"name\":\"description_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenImage\",\"inputs\":[{\"name\":\"image_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signers\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeTotal\",\"inputs\":[],\"outputs\":[{\"name\":\"stakeTotal_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakedBalance\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stakedBalance_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDescription\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenDescription_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenIdToClientId\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"clientId\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenImage\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenImage_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenURI_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStaked\",\"inputs\":[],\"outputs\":[{\"name\":\"totalStaked_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetForfeit\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetFreezeUpdate\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"frozen\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Issued\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"clientId\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LicenseAliasSet\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"licenseAlias\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Locked\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedirectUriDisabled\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RedirectUriEnabled\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerDisabled\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerEnabled\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposit\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdraw\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateDimoCreditAddress\",\"inputs\":[{\"name\":\"dimoCredit_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateDimoTokenAddress\",\"inputs\":[{\"name\":\"dimoToken_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateLicenseAccountFactoryAddress\",\"inputs\":[{\"name\":\"licenseAccountFactory_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateLicenseCost\",\"inputs\":[{\"name\":\"licenseCost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePeriodValidity\",\"inputs\":[{\"name\":\"periodValidity\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePriceProviderAddress\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateReceiverAddress\",\"inputs\":[{\"name\":\"receiver_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AliasAlreadyInUse\",\"inputs\":[{\"name\":\"licenseAlias\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"AliasExceedsMaxLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FrozenToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientStakedFunds\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakedFunds\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// DevlicensedimoABI is the input ABI used to generate the binding from.
// Deprecated: Use DevlicensedimoMetaData.ABI instead.
var DevlicensedimoABI = DevlicensedimoMetaData.ABI

// Devlicensedimo is an auto generated Go binding around an Ethereum contract.
type Devlicensedimo struct {
	DevlicensedimoCaller     // Read-only binding to the contract
	DevlicensedimoTransactor // Write-only binding to the contract
	DevlicensedimoFilterer   // Log filterer for contract events
}

// DevlicensedimoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DevlicensedimoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevlicensedimoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DevlicensedimoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevlicensedimoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DevlicensedimoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DevlicensedimoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DevlicensedimoSession struct {
	Contract     *Devlicensedimo   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DevlicensedimoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DevlicensedimoCallerSession struct {
	Contract *DevlicensedimoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DevlicensedimoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DevlicensedimoTransactorSession struct {
	Contract     *DevlicensedimoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DevlicensedimoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DevlicensedimoRaw struct {
	Contract *Devlicensedimo // Generic contract binding to access the raw methods on
}

// DevlicensedimoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DevlicensedimoCallerRaw struct {
	Contract *DevlicensedimoCaller // Generic read-only contract binding to access the raw methods on
}

// DevlicensedimoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DevlicensedimoTransactorRaw struct {
	Contract *DevlicensedimoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDevlicensedimo creates a new instance of Devlicensedimo, bound to a specific deployed contract.
func NewDevlicensedimo(address common.Address, backend bind.ContractBackend) (*Devlicensedimo, error) {
	contract, err := bindDevlicensedimo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Devlicensedimo{DevlicensedimoCaller: DevlicensedimoCaller{contract: contract}, DevlicensedimoTransactor: DevlicensedimoTransactor{contract: contract}, DevlicensedimoFilterer: DevlicensedimoFilterer{contract: contract}}, nil
}

// NewDevlicensedimoCaller creates a new read-only instance of Devlicensedimo, bound to a specific deployed contract.
func NewDevlicensedimoCaller(address common.Address, caller bind.ContractCaller) (*DevlicensedimoCaller, error) {
	contract, err := bindDevlicensedimo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoCaller{contract: contract}, nil
}

// NewDevlicensedimoTransactor creates a new write-only instance of Devlicensedimo, bound to a specific deployed contract.
func NewDevlicensedimoTransactor(address common.Address, transactor bind.ContractTransactor) (*DevlicensedimoTransactor, error) {
	contract, err := bindDevlicensedimo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoTransactor{contract: contract}, nil
}

// NewDevlicensedimoFilterer creates a new log filterer instance of Devlicensedimo, bound to a specific deployed contract.
func NewDevlicensedimoFilterer(address common.Address, filterer bind.ContractFilterer) (*DevlicensedimoFilterer, error) {
	contract, err := bindDevlicensedimo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoFilterer{contract: contract}, nil
}

// bindDevlicensedimo binds a generic wrapper to an already deployed contract.
func bindDevlicensedimo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DevlicensedimoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Devlicensedimo *DevlicensedimoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Devlicensedimo.Contract.DevlicensedimoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Devlicensedimo *DevlicensedimoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.DevlicensedimoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Devlicensedimo *DevlicensedimoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.DevlicensedimoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Devlicensedimo *DevlicensedimoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Devlicensedimo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Devlicensedimo *DevlicensedimoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Devlicensedimo *DevlicensedimoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.DEFAULTADMINROLE(&_Devlicensedimo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.DEFAULTADMINROLE(&_Devlicensedimo.CallOpts)
}

// LICENSEADMINROLE is a free data retrieval call binding the contract method 0x6d1de63d.
//
// Solidity: function LICENSE_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) LICENSEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "LICENSE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LICENSEADMINROLE is a free data retrieval call binding the contract method 0x6d1de63d.
//
// Solidity: function LICENSE_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) LICENSEADMINROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.LICENSEADMINROLE(&_Devlicensedimo.CallOpts)
}

// LICENSEADMINROLE is a free data retrieval call binding the contract method 0x6d1de63d.
//
// Solidity: function LICENSE_ADMIN_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) LICENSEADMINROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.LICENSEADMINROLE(&_Devlicensedimo.CallOpts)
}

// REVOKERROLE is a free data retrieval call binding the contract method 0x7c4acabf.
//
// Solidity: function REVOKER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) REVOKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "REVOKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REVOKERROLE is a free data retrieval call binding the contract method 0x7c4acabf.
//
// Solidity: function REVOKER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) REVOKERROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.REVOKERROLE(&_Devlicensedimo.CallOpts)
}

// REVOKERROLE is a free data retrieval call binding the contract method 0x7c4acabf.
//
// Solidity: function REVOKER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) REVOKERROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.REVOKERROLE(&_Devlicensedimo.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) UPGRADERROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.UPGRADERROLE(&_Devlicensedimo.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Devlicensedimo.Contract.UPGRADERROLE(&_Devlicensedimo.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Devlicensedimo *DevlicensedimoCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Devlicensedimo *DevlicensedimoSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Devlicensedimo.Contract.UPGRADEINTERFACEVERSION(&_Devlicensedimo.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Devlicensedimo *DevlicensedimoCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Devlicensedimo.Contract.UPGRADEINTERFACEVERSION(&_Devlicensedimo.CallOpts)
}

// ClientIdToTokenId is a free data retrieval call binding the contract method 0xce052244.
//
// Solidity: function clientIdToTokenId(address clientId) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCaller) ClientIdToTokenId(opts *bind.CallOpts, clientId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "clientIdToTokenId", clientId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClientIdToTokenId is a free data retrieval call binding the contract method 0xce052244.
//
// Solidity: function clientIdToTokenId(address clientId) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoSession) ClientIdToTokenId(clientId common.Address) (*big.Int, error) {
	return _Devlicensedimo.Contract.ClientIdToTokenId(&_Devlicensedimo.CallOpts, clientId)
}

// ClientIdToTokenId is a free data retrieval call binding the contract method 0xce052244.
//
// Solidity: function clientIdToTokenId(address clientId) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCallerSession) ClientIdToTokenId(clientId common.Address) (*big.Int, error) {
	return _Devlicensedimo.Contract.ClientIdToTokenId(&_Devlicensedimo.CallOpts, clientId)
}

// ContractDescription is a free data retrieval call binding the contract method 0x872db889.
//
// Solidity: function contractDescription() view returns(string contractDescription_)
func (_Devlicensedimo *DevlicensedimoCaller) ContractDescription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "contractDescription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractDescription is a free data retrieval call binding the contract method 0x872db889.
//
// Solidity: function contractDescription() view returns(string contractDescription_)
func (_Devlicensedimo *DevlicensedimoSession) ContractDescription() (string, error) {
	return _Devlicensedimo.Contract.ContractDescription(&_Devlicensedimo.CallOpts)
}

// ContractDescription is a free data retrieval call binding the contract method 0x872db889.
//
// Solidity: function contractDescription() view returns(string contractDescription_)
func (_Devlicensedimo *DevlicensedimoCallerSession) ContractDescription() (string, error) {
	return _Devlicensedimo.Contract.ContractDescription(&_Devlicensedimo.CallOpts)
}

// ContractImage is a free data retrieval call binding the contract method 0x8eab84ee.
//
// Solidity: function contractImage() view returns(string contractImage_)
func (_Devlicensedimo *DevlicensedimoCaller) ContractImage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "contractImage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractImage is a free data retrieval call binding the contract method 0x8eab84ee.
//
// Solidity: function contractImage() view returns(string contractImage_)
func (_Devlicensedimo *DevlicensedimoSession) ContractImage() (string, error) {
	return _Devlicensedimo.Contract.ContractImage(&_Devlicensedimo.CallOpts)
}

// ContractImage is a free data retrieval call binding the contract method 0x8eab84ee.
//
// Solidity: function contractImage() view returns(string contractImage_)
func (_Devlicensedimo *DevlicensedimoCallerSession) ContractImage() (string, error) {
	return _Devlicensedimo.Contract.ContractImage(&_Devlicensedimo.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string contractURI_)
func (_Devlicensedimo *DevlicensedimoCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string contractURI_)
func (_Devlicensedimo *DevlicensedimoSession) ContractURI() (string, error) {
	return _Devlicensedimo.Contract.ContractURI(&_Devlicensedimo.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string contractURI_)
func (_Devlicensedimo *DevlicensedimoCallerSession) ContractURI() (string, error) {
	return _Devlicensedimo.Contract.ContractURI(&_Devlicensedimo.CallOpts)
}

// DimoCredit is a free data retrieval call binding the contract method 0xdb60ded9.
//
// Solidity: function dimoCredit() view returns(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoCaller) DimoCredit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "dimoCredit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DimoCredit is a free data retrieval call binding the contract method 0xdb60ded9.
//
// Solidity: function dimoCredit() view returns(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoSession) DimoCredit() (common.Address, error) {
	return _Devlicensedimo.Contract.DimoCredit(&_Devlicensedimo.CallOpts)
}

// DimoCredit is a free data retrieval call binding the contract method 0xdb60ded9.
//
// Solidity: function dimoCredit() view returns(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoCallerSession) DimoCredit() (common.Address, error) {
	return _Devlicensedimo.Contract.DimoCredit(&_Devlicensedimo.CallOpts)
}

// DimoToken is a free data retrieval call binding the contract method 0x0524f28c.
//
// Solidity: function dimoToken() view returns(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoCaller) DimoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "dimoToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DimoToken is a free data retrieval call binding the contract method 0x0524f28c.
//
// Solidity: function dimoToken() view returns(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoSession) DimoToken() (common.Address, error) {
	return _Devlicensedimo.Contract.DimoToken(&_Devlicensedimo.CallOpts)
}

// DimoToken is a free data retrieval call binding the contract method 0x0524f28c.
//
// Solidity: function dimoToken() view returns(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoCallerSession) DimoToken() (common.Address, error) {
	return _Devlicensedimo.Contract.DimoToken(&_Devlicensedimo.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0xe9ac0440.
//
// Solidity: function frozen(uint256 tokenId) view returns(bool frozen_)
func (_Devlicensedimo *DevlicensedimoCaller) Frozen(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "frozen", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Frozen is a free data retrieval call binding the contract method 0xe9ac0440.
//
// Solidity: function frozen(uint256 tokenId) view returns(bool frozen_)
func (_Devlicensedimo *DevlicensedimoSession) Frozen(tokenId *big.Int) (bool, error) {
	return _Devlicensedimo.Contract.Frozen(&_Devlicensedimo.CallOpts, tokenId)
}

// Frozen is a free data retrieval call binding the contract method 0xe9ac0440.
//
// Solidity: function frozen(uint256 tokenId) view returns(bool frozen_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Frozen(tokenId *big.Int) (bool, error) {
	return _Devlicensedimo.Contract.Frozen(&_Devlicensedimo.CallOpts, tokenId)
}

// GetLicenseAliasByTokenId is a free data retrieval call binding the contract method 0x3aea1a10.
//
// Solidity: function getLicenseAliasByTokenId(uint256 tokenId) view returns(string licenseAlias)
func (_Devlicensedimo *DevlicensedimoCaller) GetLicenseAliasByTokenId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "getLicenseAliasByTokenId", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLicenseAliasByTokenId is a free data retrieval call binding the contract method 0x3aea1a10.
//
// Solidity: function getLicenseAliasByTokenId(uint256 tokenId) view returns(string licenseAlias)
func (_Devlicensedimo *DevlicensedimoSession) GetLicenseAliasByTokenId(tokenId *big.Int) (string, error) {
	return _Devlicensedimo.Contract.GetLicenseAliasByTokenId(&_Devlicensedimo.CallOpts, tokenId)
}

// GetLicenseAliasByTokenId is a free data retrieval call binding the contract method 0x3aea1a10.
//
// Solidity: function getLicenseAliasByTokenId(uint256 tokenId) view returns(string licenseAlias)
func (_Devlicensedimo *DevlicensedimoCallerSession) GetLicenseAliasByTokenId(tokenId *big.Int) (string, error) {
	return _Devlicensedimo.Contract.GetLicenseAliasByTokenId(&_Devlicensedimo.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Devlicensedimo.Contract.GetRoleAdmin(&_Devlicensedimo.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Devlicensedimo.Contract.GetRoleAdmin(&_Devlicensedimo.CallOpts, role)
}

// GetTokenIdByLicenseAlias is a free data retrieval call binding the contract method 0x6eac1c40.
//
// Solidity: function getTokenIdByLicenseAlias(string licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCaller) GetTokenIdByLicenseAlias(opts *bind.CallOpts, licenseAlias string) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "getTokenIdByLicenseAlias", licenseAlias)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdByLicenseAlias is a free data retrieval call binding the contract method 0x6eac1c40.
//
// Solidity: function getTokenIdByLicenseAlias(string licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoSession) GetTokenIdByLicenseAlias(licenseAlias string) (*big.Int, error) {
	return _Devlicensedimo.Contract.GetTokenIdByLicenseAlias(&_Devlicensedimo.CallOpts, licenseAlias)
}

// GetTokenIdByLicenseAlias is a free data retrieval call binding the contract method 0x6eac1c40.
//
// Solidity: function getTokenIdByLicenseAlias(string licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCallerSession) GetTokenIdByLicenseAlias(licenseAlias string) (*big.Int, error) {
	return _Devlicensedimo.Contract.GetTokenIdByLicenseAlias(&_Devlicensedimo.CallOpts, licenseAlias)
}

// GetTokenIdByLicenseAlias0 is a free data retrieval call binding the contract method 0xf5c7223d.
//
// Solidity: function getTokenIdByLicenseAlias(bytes32 licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCaller) GetTokenIdByLicenseAlias0(opts *bind.CallOpts, licenseAlias [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "getTokenIdByLicenseAlias0", licenseAlias)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenIdByLicenseAlias0 is a free data retrieval call binding the contract method 0xf5c7223d.
//
// Solidity: function getTokenIdByLicenseAlias(bytes32 licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoSession) GetTokenIdByLicenseAlias0(licenseAlias [32]byte) (*big.Int, error) {
	return _Devlicensedimo.Contract.GetTokenIdByLicenseAlias0(&_Devlicensedimo.CallOpts, licenseAlias)
}

// GetTokenIdByLicenseAlias0 is a free data retrieval call binding the contract method 0xf5c7223d.
//
// Solidity: function getTokenIdByLicenseAlias(bytes32 licenseAlias) view returns(uint256 tokenId)
func (_Devlicensedimo *DevlicensedimoCallerSession) GetTokenIdByLicenseAlias0(licenseAlias [32]byte) (*big.Int, error) {
	return _Devlicensedimo.Contract.GetTokenIdByLicenseAlias0(&_Devlicensedimo.CallOpts, licenseAlias)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Devlicensedimo *DevlicensedimoCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Devlicensedimo *DevlicensedimoSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Devlicensedimo.Contract.HasRole(&_Devlicensedimo.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Devlicensedimo *DevlicensedimoCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Devlicensedimo.Contract.HasRole(&_Devlicensedimo.CallOpts, role, account)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 tokenId, address signer) view returns(bool isSigner_)
func (_Devlicensedimo *DevlicensedimoCaller) IsSigner(opts *bind.CallOpts, tokenId *big.Int, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "isSigner", tokenId, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 tokenId, address signer) view returns(bool isSigner_)
func (_Devlicensedimo *DevlicensedimoSession) IsSigner(tokenId *big.Int, signer common.Address) (bool, error) {
	return _Devlicensedimo.Contract.IsSigner(&_Devlicensedimo.CallOpts, tokenId, signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 tokenId, address signer) view returns(bool isSigner_)
func (_Devlicensedimo *DevlicensedimoCallerSession) IsSigner(tokenId *big.Int, signer common.Address) (bool, error) {
	return _Devlicensedimo.Contract.IsSigner(&_Devlicensedimo.CallOpts, tokenId, signer)
}

// LicenseAccountFactory is a free data retrieval call binding the contract method 0xddfaf5c2.
//
// Solidity: function licenseAccountFactory() view returns(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoCaller) LicenseAccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "licenseAccountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseAccountFactory is a free data retrieval call binding the contract method 0xddfaf5c2.
//
// Solidity: function licenseAccountFactory() view returns(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoSession) LicenseAccountFactory() (common.Address, error) {
	return _Devlicensedimo.Contract.LicenseAccountFactory(&_Devlicensedimo.CallOpts)
}

// LicenseAccountFactory is a free data retrieval call binding the contract method 0xddfaf5c2.
//
// Solidity: function licenseAccountFactory() view returns(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoCallerSession) LicenseAccountFactory() (common.Address, error) {
	return _Devlicensedimo.Contract.LicenseAccountFactory(&_Devlicensedimo.CallOpts)
}

// LicenseCostInUsd1e18 is a free data retrieval call binding the contract method 0x5d79373d.
//
// Solidity: function licenseCostInUsd1e18() view returns(uint256 licenseCostInUsd1e18_)
func (_Devlicensedimo *DevlicensedimoCaller) LicenseCostInUsd1e18(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "licenseCostInUsd1e18")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LicenseCostInUsd1e18 is a free data retrieval call binding the contract method 0x5d79373d.
//
// Solidity: function licenseCostInUsd1e18() view returns(uint256 licenseCostInUsd1e18_)
func (_Devlicensedimo *DevlicensedimoSession) LicenseCostInUsd1e18() (*big.Int, error) {
	return _Devlicensedimo.Contract.LicenseCostInUsd1e18(&_Devlicensedimo.CallOpts)
}

// LicenseCostInUsd1e18 is a free data retrieval call binding the contract method 0x5d79373d.
//
// Solidity: function licenseCostInUsd1e18() view returns(uint256 licenseCostInUsd1e18_)
func (_Devlicensedimo *DevlicensedimoCallerSession) LicenseCostInUsd1e18() (*big.Int, error) {
	return _Devlicensedimo.Contract.LicenseCostInUsd1e18(&_Devlicensedimo.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool locked_)
func (_Devlicensedimo *DevlicensedimoCaller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool locked_)
func (_Devlicensedimo *DevlicensedimoSession) Locked(tokenId *big.Int) (bool, error) {
	return _Devlicensedimo.Contract.Locked(&_Devlicensedimo.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool locked_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _Devlicensedimo.Contract.Locked(&_Devlicensedimo.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string name_)
func (_Devlicensedimo *DevlicensedimoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string name_)
func (_Devlicensedimo *DevlicensedimoSession) Name() (string, error) {
	return _Devlicensedimo.Contract.Name(&_Devlicensedimo.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string name_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Name() (string, error) {
	return _Devlicensedimo.Contract.Name(&_Devlicensedimo.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Devlicensedimo *DevlicensedimoCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Devlicensedimo *DevlicensedimoSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Devlicensedimo.Contract.OwnerOf(&_Devlicensedimo.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Devlicensedimo *DevlicensedimoCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Devlicensedimo.Contract.OwnerOf(&_Devlicensedimo.CallOpts, tokenId)
}

// PeriodValidity is a free data retrieval call binding the contract method 0xb30afb81.
//
// Solidity: function periodValidity() view returns(uint256 periodValidity_)
func (_Devlicensedimo *DevlicensedimoCaller) PeriodValidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "periodValidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodValidity is a free data retrieval call binding the contract method 0xb30afb81.
//
// Solidity: function periodValidity() view returns(uint256 periodValidity_)
func (_Devlicensedimo *DevlicensedimoSession) PeriodValidity() (*big.Int, error) {
	return _Devlicensedimo.Contract.PeriodValidity(&_Devlicensedimo.CallOpts)
}

// PeriodValidity is a free data retrieval call binding the contract method 0xb30afb81.
//
// Solidity: function periodValidity() view returns(uint256 periodValidity_)
func (_Devlicensedimo *DevlicensedimoCallerSession) PeriodValidity() (*big.Int, error) {
	return _Devlicensedimo.Contract.PeriodValidity(&_Devlicensedimo.CallOpts)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address provider_)
func (_Devlicensedimo *DevlicensedimoCaller) Provider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address provider_)
func (_Devlicensedimo *DevlicensedimoSession) Provider() (common.Address, error) {
	return _Devlicensedimo.Contract.Provider(&_Devlicensedimo.CallOpts)
}

// Provider is a free data retrieval call binding the contract method 0x085d4883.
//
// Solidity: function provider() view returns(address provider_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Provider() (common.Address, error) {
	return _Devlicensedimo.Contract.Provider(&_Devlicensedimo.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoSession) ProxiableUUID() ([32]byte, error) {
	return _Devlicensedimo.Contract.ProxiableUUID(&_Devlicensedimo.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Devlicensedimo *DevlicensedimoCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Devlicensedimo.Contract.ProxiableUUID(&_Devlicensedimo.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address receiver_)
func (_Devlicensedimo *DevlicensedimoCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address receiver_)
func (_Devlicensedimo *DevlicensedimoSession) Receiver() (common.Address, error) {
	return _Devlicensedimo.Contract.Receiver(&_Devlicensedimo.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address receiver_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Receiver() (common.Address, error) {
	return _Devlicensedimo.Contract.Receiver(&_Devlicensedimo.CallOpts)
}

// RedirectUriStatus is a free data retrieval call binding the contract method 0x2877fed1.
//
// Solidity: function redirectUriStatus(uint256 tokenId, string uri) view returns(bool enabled)
func (_Devlicensedimo *DevlicensedimoCaller) RedirectUriStatus(opts *bind.CallOpts, tokenId *big.Int, uri string) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "redirectUriStatus", tokenId, uri)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RedirectUriStatus is a free data retrieval call binding the contract method 0x2877fed1.
//
// Solidity: function redirectUriStatus(uint256 tokenId, string uri) view returns(bool enabled)
func (_Devlicensedimo *DevlicensedimoSession) RedirectUriStatus(tokenId *big.Int, uri string) (bool, error) {
	return _Devlicensedimo.Contract.RedirectUriStatus(&_Devlicensedimo.CallOpts, tokenId, uri)
}

// RedirectUriStatus is a free data retrieval call binding the contract method 0x2877fed1.
//
// Solidity: function redirectUriStatus(uint256 tokenId, string uri) view returns(bool enabled)
func (_Devlicensedimo *DevlicensedimoCallerSession) RedirectUriStatus(tokenId *big.Int, uri string) (bool, error) {
	return _Devlicensedimo.Contract.RedirectUriStatus(&_Devlicensedimo.CallOpts, tokenId, uri)
}

// Signers is a free data retrieval call binding the contract method 0xc920ba60.
//
// Solidity: function signers(uint256 tokenId, address signer) view returns(uint256 timestamp)
func (_Devlicensedimo *DevlicensedimoCaller) Signers(opts *bind.CallOpts, tokenId *big.Int, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "signers", tokenId, signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0xc920ba60.
//
// Solidity: function signers(uint256 tokenId, address signer) view returns(uint256 timestamp)
func (_Devlicensedimo *DevlicensedimoSession) Signers(tokenId *big.Int, signer common.Address) (*big.Int, error) {
	return _Devlicensedimo.Contract.Signers(&_Devlicensedimo.CallOpts, tokenId, signer)
}

// Signers is a free data retrieval call binding the contract method 0xc920ba60.
//
// Solidity: function signers(uint256 tokenId, address signer) view returns(uint256 timestamp)
func (_Devlicensedimo *DevlicensedimoCallerSession) Signers(tokenId *big.Int, signer common.Address) (*big.Int, error) {
	return _Devlicensedimo.Contract.Signers(&_Devlicensedimo.CallOpts, tokenId, signer)
}

// StakeTotal is a free data retrieval call binding the contract method 0x4b5c8bdf.
//
// Solidity: function stakeTotal() view returns(uint256 stakeTotal_)
func (_Devlicensedimo *DevlicensedimoCaller) StakeTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "stakeTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTotal is a free data retrieval call binding the contract method 0x4b5c8bdf.
//
// Solidity: function stakeTotal() view returns(uint256 stakeTotal_)
func (_Devlicensedimo *DevlicensedimoSession) StakeTotal() (*big.Int, error) {
	return _Devlicensedimo.Contract.StakeTotal(&_Devlicensedimo.CallOpts)
}

// StakeTotal is a free data retrieval call binding the contract method 0x4b5c8bdf.
//
// Solidity: function stakeTotal() view returns(uint256 stakeTotal_)
func (_Devlicensedimo *DevlicensedimoCallerSession) StakeTotal() (*big.Int, error) {
	return _Devlicensedimo.Contract.StakeTotal(&_Devlicensedimo.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0xf256cad5.
//
// Solidity: function stakedBalance(uint256 tokenId) view returns(uint256 stakedBalance_)
func (_Devlicensedimo *DevlicensedimoCaller) StakedBalance(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "stakedBalance", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0xf256cad5.
//
// Solidity: function stakedBalance(uint256 tokenId) view returns(uint256 stakedBalance_)
func (_Devlicensedimo *DevlicensedimoSession) StakedBalance(tokenId *big.Int) (*big.Int, error) {
	return _Devlicensedimo.Contract.StakedBalance(&_Devlicensedimo.CallOpts, tokenId)
}

// StakedBalance is a free data retrieval call binding the contract method 0xf256cad5.
//
// Solidity: function stakedBalance(uint256 tokenId) view returns(uint256 stakedBalance_)
func (_Devlicensedimo *DevlicensedimoCallerSession) StakedBalance(tokenId *big.Int) (*big.Int, error) {
	return _Devlicensedimo.Contract.StakedBalance(&_Devlicensedimo.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Devlicensedimo *DevlicensedimoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Devlicensedimo *DevlicensedimoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Devlicensedimo.Contract.SupportsInterface(&_Devlicensedimo.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Devlicensedimo *DevlicensedimoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Devlicensedimo.Contract.SupportsInterface(&_Devlicensedimo.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol_)
func (_Devlicensedimo *DevlicensedimoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol_)
func (_Devlicensedimo *DevlicensedimoSession) Symbol() (string, error) {
	return _Devlicensedimo.Contract.Symbol(&_Devlicensedimo.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string symbol_)
func (_Devlicensedimo *DevlicensedimoCallerSession) Symbol() (string, error) {
	return _Devlicensedimo.Contract.Symbol(&_Devlicensedimo.CallOpts)
}

// TokenDescription is a free data retrieval call binding the contract method 0x25e6f516.
//
// Solidity: function tokenDescription() view returns(string tokenDescription_)
func (_Devlicensedimo *DevlicensedimoCaller) TokenDescription(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "tokenDescription")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenDescription is a free data retrieval call binding the contract method 0x25e6f516.
//
// Solidity: function tokenDescription() view returns(string tokenDescription_)
func (_Devlicensedimo *DevlicensedimoSession) TokenDescription() (string, error) {
	return _Devlicensedimo.Contract.TokenDescription(&_Devlicensedimo.CallOpts)
}

// TokenDescription is a free data retrieval call binding the contract method 0x25e6f516.
//
// Solidity: function tokenDescription() view returns(string tokenDescription_)
func (_Devlicensedimo *DevlicensedimoCallerSession) TokenDescription() (string, error) {
	return _Devlicensedimo.Contract.TokenDescription(&_Devlicensedimo.CallOpts)
}

// TokenIdToClientId is a free data retrieval call binding the contract method 0x5df1453e.
//
// Solidity: function tokenIdToClientId(uint256 tokenId) view returns(address clientId)
func (_Devlicensedimo *DevlicensedimoCaller) TokenIdToClientId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "tokenIdToClientId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenIdToClientId is a free data retrieval call binding the contract method 0x5df1453e.
//
// Solidity: function tokenIdToClientId(uint256 tokenId) view returns(address clientId)
func (_Devlicensedimo *DevlicensedimoSession) TokenIdToClientId(tokenId *big.Int) (common.Address, error) {
	return _Devlicensedimo.Contract.TokenIdToClientId(&_Devlicensedimo.CallOpts, tokenId)
}

// TokenIdToClientId is a free data retrieval call binding the contract method 0x5df1453e.
//
// Solidity: function tokenIdToClientId(uint256 tokenId) view returns(address clientId)
func (_Devlicensedimo *DevlicensedimoCallerSession) TokenIdToClientId(tokenId *big.Int) (common.Address, error) {
	return _Devlicensedimo.Contract.TokenIdToClientId(&_Devlicensedimo.CallOpts, tokenId)
}

// TokenImage is a free data retrieval call binding the contract method 0xc1aef4f2.
//
// Solidity: function tokenImage() view returns(string tokenImage_)
func (_Devlicensedimo *DevlicensedimoCaller) TokenImage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "tokenImage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenImage is a free data retrieval call binding the contract method 0xc1aef4f2.
//
// Solidity: function tokenImage() view returns(string tokenImage_)
func (_Devlicensedimo *DevlicensedimoSession) TokenImage() (string, error) {
	return _Devlicensedimo.Contract.TokenImage(&_Devlicensedimo.CallOpts)
}

// TokenImage is a free data retrieval call binding the contract method 0xc1aef4f2.
//
// Solidity: function tokenImage() view returns(string tokenImage_)
func (_Devlicensedimo *DevlicensedimoCallerSession) TokenImage() (string, error) {
	return _Devlicensedimo.Contract.TokenImage(&_Devlicensedimo.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string tokenURI_)
func (_Devlicensedimo *DevlicensedimoCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string tokenURI_)
func (_Devlicensedimo *DevlicensedimoSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Devlicensedimo.Contract.TokenURI(&_Devlicensedimo.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string tokenURI_)
func (_Devlicensedimo *DevlicensedimoCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Devlicensedimo.Contract.TokenURI(&_Devlicensedimo.CallOpts, tokenId)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256 totalStaked_)
func (_Devlicensedimo *DevlicensedimoCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Devlicensedimo.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256 totalStaked_)
func (_Devlicensedimo *DevlicensedimoSession) TotalStaked() (*big.Int, error) {
	return _Devlicensedimo.Contract.TotalStaked(&_Devlicensedimo.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256 totalStaked_)
func (_Devlicensedimo *DevlicensedimoCallerSession) TotalStaked() (*big.Int, error) {
	return _Devlicensedimo.Contract.TotalStaked(&_Devlicensedimo.CallOpts)
}

// AdminBurnLockedFunds is a paid mutator transaction binding the contract method 0xde06ce84.
//
// Solidity: function adminBurnLockedFunds(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) AdminBurnLockedFunds(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "adminBurnLockedFunds", tokenId, amount)
}

// AdminBurnLockedFunds is a paid mutator transaction binding the contract method 0xde06ce84.
//
// Solidity: function adminBurnLockedFunds(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoSession) AdminBurnLockedFunds(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminBurnLockedFunds(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// AdminBurnLockedFunds is a paid mutator transaction binding the contract method 0xde06ce84.
//
// Solidity: function adminBurnLockedFunds(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) AdminBurnLockedFunds(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminBurnLockedFunds(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// AdminFreeze is a paid mutator transaction binding the contract method 0x0a98e154.
//
// Solidity: function adminFreeze(uint256 tokenId, bool frozen_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) AdminFreeze(opts *bind.TransactOpts, tokenId *big.Int, frozen_ bool) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "adminFreeze", tokenId, frozen_)
}

// AdminFreeze is a paid mutator transaction binding the contract method 0x0a98e154.
//
// Solidity: function adminFreeze(uint256 tokenId, bool frozen_) returns()
func (_Devlicensedimo *DevlicensedimoSession) AdminFreeze(tokenId *big.Int, frozen_ bool) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminFreeze(&_Devlicensedimo.TransactOpts, tokenId, frozen_)
}

// AdminFreeze is a paid mutator transaction binding the contract method 0x0a98e154.
//
// Solidity: function adminFreeze(uint256 tokenId, bool frozen_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) AdminFreeze(tokenId *big.Int, frozen_ bool) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminFreeze(&_Devlicensedimo.TransactOpts, tokenId, frozen_)
}

// AdminReallocate is a paid mutator transaction binding the contract method 0x92d937c7.
//
// Solidity: function adminReallocate(uint256 tokenId, uint256 amount, address to) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) AdminReallocate(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "adminReallocate", tokenId, amount, to)
}

// AdminReallocate is a paid mutator transaction binding the contract method 0x92d937c7.
//
// Solidity: function adminReallocate(uint256 tokenId, uint256 amount, address to) returns()
func (_Devlicensedimo *DevlicensedimoSession) AdminReallocate(tokenId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminReallocate(&_Devlicensedimo.TransactOpts, tokenId, amount, to)
}

// AdminReallocate is a paid mutator transaction binding the contract method 0x92d937c7.
//
// Solidity: function adminReallocate(uint256 tokenId, uint256 amount, address to) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) AdminReallocate(tokenId *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminReallocate(&_Devlicensedimo.TransactOpts, tokenId, amount, to)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa28835b6.
//
// Solidity: function adminWithdraw(address to) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) AdminWithdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "adminWithdraw", to)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa28835b6.
//
// Solidity: function adminWithdraw(address to) returns()
func (_Devlicensedimo *DevlicensedimoSession) AdminWithdraw(to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminWithdraw(&_Devlicensedimo.TransactOpts, to)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0xa28835b6.
//
// Solidity: function adminWithdraw(address to) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) AdminWithdraw(to common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.AdminWithdraw(&_Devlicensedimo.TransactOpts, to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Approve(&_Devlicensedimo.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Approve(&_Devlicensedimo.TransactOpts, arg0, arg1)
}

// DisableSigner is a paid mutator transaction binding the contract method 0xde9cc84d.
//
// Solidity: function disableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) DisableSigner(opts *bind.TransactOpts, tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "disableSigner", tokenId, signer)
}

// DisableSigner is a paid mutator transaction binding the contract method 0xde9cc84d.
//
// Solidity: function disableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoSession) DisableSigner(tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.DisableSigner(&_Devlicensedimo.TransactOpts, tokenId, signer)
}

// DisableSigner is a paid mutator transaction binding the contract method 0xde9cc84d.
//
// Solidity: function disableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) DisableSigner(tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.DisableSigner(&_Devlicensedimo.TransactOpts, tokenId, signer)
}

// EnableSigner is a paid mutator transaction binding the contract method 0x3b1c393b.
//
// Solidity: function enableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) EnableSigner(opts *bind.TransactOpts, tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "enableSigner", tokenId, signer)
}

// EnableSigner is a paid mutator transaction binding the contract method 0x3b1c393b.
//
// Solidity: function enableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoSession) EnableSigner(tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.EnableSigner(&_Devlicensedimo.TransactOpts, tokenId, signer)
}

// EnableSigner is a paid mutator transaction binding the contract method 0x3b1c393b.
//
// Solidity: function enableSigner(uint256 tokenId, address signer) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) EnableSigner(tokenId *big.Int, signer common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.EnableSigner(&_Devlicensedimo.TransactOpts, tokenId, signer)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.GrantRole(&_Devlicensedimo.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.GrantRole(&_Devlicensedimo.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1dccc525.
//
// Solidity: function initialize(address receiver_, address licenseAccountFactory_, address provider_, address dimoTokenAddress_, address dimoCreditAddress_, uint256 licenseCostInUsd_, string image_, string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) Initialize(opts *bind.TransactOpts, receiver_ common.Address, licenseAccountFactory_ common.Address, provider_ common.Address, dimoTokenAddress_ common.Address, dimoCreditAddress_ common.Address, licenseCostInUsd_ *big.Int, image_ string, description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "initialize", receiver_, licenseAccountFactory_, provider_, dimoTokenAddress_, dimoCreditAddress_, licenseCostInUsd_, image_, description_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1dccc525.
//
// Solidity: function initialize(address receiver_, address licenseAccountFactory_, address provider_, address dimoTokenAddress_, address dimoCreditAddress_, uint256 licenseCostInUsd_, string image_, string description_) returns()
func (_Devlicensedimo *DevlicensedimoSession) Initialize(receiver_ common.Address, licenseAccountFactory_ common.Address, provider_ common.Address, dimoTokenAddress_ common.Address, dimoCreditAddress_ common.Address, licenseCostInUsd_ *big.Int, image_ string, description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Initialize(&_Devlicensedimo.TransactOpts, receiver_, licenseAccountFactory_, provider_, dimoTokenAddress_, dimoCreditAddress_, licenseCostInUsd_, image_, description_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1dccc525.
//
// Solidity: function initialize(address receiver_, address licenseAccountFactory_, address provider_, address dimoTokenAddress_, address dimoCreditAddress_, uint256 licenseCostInUsd_, string image_, string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) Initialize(receiver_ common.Address, licenseAccountFactory_ common.Address, provider_ common.Address, dimoTokenAddress_ common.Address, dimoCreditAddress_ common.Address, licenseCostInUsd_ *big.Int, image_ string, description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Initialize(&_Devlicensedimo.TransactOpts, receiver_, licenseAccountFactory_, provider_, dimoTokenAddress_, dimoCreditAddress_, licenseCostInUsd_, image_, description_)
}

// IssueInDc is a paid mutator transaction binding the contract method 0x69054339.
//
// Solidity: function issueInDc(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactor) IssueInDc(opts *bind.TransactOpts, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "issueInDc", licenseAlias)
}

// IssueInDc is a paid mutator transaction binding the contract method 0x69054339.
//
// Solidity: function issueInDc(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoSession) IssueInDc(licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDc(&_Devlicensedimo.TransactOpts, licenseAlias)
}

// IssueInDc is a paid mutator transaction binding the contract method 0x69054339.
//
// Solidity: function issueInDc(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactorSession) IssueInDc(licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDc(&_Devlicensedimo.TransactOpts, licenseAlias)
}

// IssueInDc0 is a paid mutator transaction binding the contract method 0x6cca0c45.
//
// Solidity: function issueInDc(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactor) IssueInDc0(opts *bind.TransactOpts, to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "issueInDc0", to, licenseAlias)
}

// IssueInDc0 is a paid mutator transaction binding the contract method 0x6cca0c45.
//
// Solidity: function issueInDc(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoSession) IssueInDc0(to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDc0(&_Devlicensedimo.TransactOpts, to, licenseAlias)
}

// IssueInDc0 is a paid mutator transaction binding the contract method 0x6cca0c45.
//
// Solidity: function issueInDc(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactorSession) IssueInDc0(to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDc0(&_Devlicensedimo.TransactOpts, to, licenseAlias)
}

// IssueInDimo is a paid mutator transaction binding the contract method 0x5d9cbfb6.
//
// Solidity: function issueInDimo(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactor) IssueInDimo(opts *bind.TransactOpts, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "issueInDimo", licenseAlias)
}

// IssueInDimo is a paid mutator transaction binding the contract method 0x5d9cbfb6.
//
// Solidity: function issueInDimo(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoSession) IssueInDimo(licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDimo(&_Devlicensedimo.TransactOpts, licenseAlias)
}

// IssueInDimo is a paid mutator transaction binding the contract method 0x5d9cbfb6.
//
// Solidity: function issueInDimo(string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactorSession) IssueInDimo(licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDimo(&_Devlicensedimo.TransactOpts, licenseAlias)
}

// IssueInDimo0 is a paid mutator transaction binding the contract method 0x8efdce0a.
//
// Solidity: function issueInDimo(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactor) IssueInDimo0(opts *bind.TransactOpts, to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "issueInDimo0", to, licenseAlias)
}

// IssueInDimo0 is a paid mutator transaction binding the contract method 0x8efdce0a.
//
// Solidity: function issueInDimo(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoSession) IssueInDimo0(to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDimo0(&_Devlicensedimo.TransactOpts, to, licenseAlias)
}

// IssueInDimo0 is a paid mutator transaction binding the contract method 0x8efdce0a.
//
// Solidity: function issueInDimo(address to, string licenseAlias) returns(uint256 tokenId, address clientId)
func (_Devlicensedimo *DevlicensedimoTransactorSession) IssueInDimo0(to common.Address, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.IssueInDimo0(&_Devlicensedimo.TransactOpts, to, licenseAlias)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) Lock(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "lock", tokenId, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoSession) Lock(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Lock(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x1338736f.
//
// Solidity: function lock(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) Lock(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Lock(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Devlicensedimo *DevlicensedimoSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.RenounceRole(&_Devlicensedimo.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.RenounceRole(&_Devlicensedimo.TransactOpts, role, callerConfirmation)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 tokenId) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) Revoke(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "revoke", tokenId)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 tokenId) returns()
func (_Devlicensedimo *DevlicensedimoSession) Revoke(tokenId *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Revoke(&_Devlicensedimo.TransactOpts, tokenId)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(uint256 tokenId) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) Revoke(tokenId *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Revoke(&_Devlicensedimo.TransactOpts, tokenId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.RevokeRole(&_Devlicensedimo.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.RevokeRole(&_Devlicensedimo.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SafeTransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "safeTransferFrom", arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SafeTransferFrom(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SafeTransferFrom(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SafeTransferFrom0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "safeTransferFrom0", arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Devlicensedimo *DevlicensedimoSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SafeTransferFrom0(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SafeTransferFrom0(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetApprovalForAll(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setApprovalForAll", arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetApprovalForAll(&_Devlicensedimo.TransactOpts, arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetApprovalForAll(&_Devlicensedimo.TransactOpts, arg0, arg1)
}

// SetContractDescription is a paid mutator transaction binding the contract method 0xf754dd9b.
//
// Solidity: function setContractDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetContractDescription(opts *bind.TransactOpts, description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setContractDescription", description_)
}

// SetContractDescription is a paid mutator transaction binding the contract method 0xf754dd9b.
//
// Solidity: function setContractDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetContractDescription(description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetContractDescription(&_Devlicensedimo.TransactOpts, description_)
}

// SetContractDescription is a paid mutator transaction binding the contract method 0xf754dd9b.
//
// Solidity: function setContractDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetContractDescription(description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetContractDescription(&_Devlicensedimo.TransactOpts, description_)
}

// SetContractImage is a paid mutator transaction binding the contract method 0xdc42c255.
//
// Solidity: function setContractImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetContractImage(opts *bind.TransactOpts, image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setContractImage", image_)
}

// SetContractImage is a paid mutator transaction binding the contract method 0xdc42c255.
//
// Solidity: function setContractImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetContractImage(image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetContractImage(&_Devlicensedimo.TransactOpts, image_)
}

// SetContractImage is a paid mutator transaction binding the contract method 0xdc42c255.
//
// Solidity: function setContractImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetContractImage(image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetContractImage(&_Devlicensedimo.TransactOpts, image_)
}

// SetDimoCreditAddress is a paid mutator transaction binding the contract method 0xd866e391.
//
// Solidity: function setDimoCreditAddress(address dimoCreditAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetDimoCreditAddress(opts *bind.TransactOpts, dimoCreditAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setDimoCreditAddress", dimoCreditAddress_)
}

// SetDimoCreditAddress is a paid mutator transaction binding the contract method 0xd866e391.
//
// Solidity: function setDimoCreditAddress(address dimoCreditAddress_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetDimoCreditAddress(dimoCreditAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetDimoCreditAddress(&_Devlicensedimo.TransactOpts, dimoCreditAddress_)
}

// SetDimoCreditAddress is a paid mutator transaction binding the contract method 0xd866e391.
//
// Solidity: function setDimoCreditAddress(address dimoCreditAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetDimoCreditAddress(dimoCreditAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetDimoCreditAddress(&_Devlicensedimo.TransactOpts, dimoCreditAddress_)
}

// SetDimoTokenAddress is a paid mutator transaction binding the contract method 0x81c9bb13.
//
// Solidity: function setDimoTokenAddress(address dimoTokenAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetDimoTokenAddress(opts *bind.TransactOpts, dimoTokenAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setDimoTokenAddress", dimoTokenAddress_)
}

// SetDimoTokenAddress is a paid mutator transaction binding the contract method 0x81c9bb13.
//
// Solidity: function setDimoTokenAddress(address dimoTokenAddress_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetDimoTokenAddress(dimoTokenAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetDimoTokenAddress(&_Devlicensedimo.TransactOpts, dimoTokenAddress_)
}

// SetDimoTokenAddress is a paid mutator transaction binding the contract method 0x81c9bb13.
//
// Solidity: function setDimoTokenAddress(address dimoTokenAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetDimoTokenAddress(dimoTokenAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetDimoTokenAddress(&_Devlicensedimo.TransactOpts, dimoTokenAddress_)
}

// SetLicenseAlias is a paid mutator transaction binding the contract method 0x507c5a89.
//
// Solidity: function setLicenseAlias(uint256 tokenId, string licenseAlias) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetLicenseAlias(opts *bind.TransactOpts, tokenId *big.Int, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setLicenseAlias", tokenId, licenseAlias)
}

// SetLicenseAlias is a paid mutator transaction binding the contract method 0x507c5a89.
//
// Solidity: function setLicenseAlias(uint256 tokenId, string licenseAlias) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetLicenseAlias(tokenId *big.Int, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseAlias(&_Devlicensedimo.TransactOpts, tokenId, licenseAlias)
}

// SetLicenseAlias is a paid mutator transaction binding the contract method 0x507c5a89.
//
// Solidity: function setLicenseAlias(uint256 tokenId, string licenseAlias) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetLicenseAlias(tokenId *big.Int, licenseAlias string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseAlias(&_Devlicensedimo.TransactOpts, tokenId, licenseAlias)
}

// SetLicenseCost is a paid mutator transaction binding the contract method 0x1d6f8ef3.
//
// Solidity: function setLicenseCost(uint256 licenseCostInUsd1e18_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetLicenseCost(opts *bind.TransactOpts, licenseCostInUsd1e18_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setLicenseCost", licenseCostInUsd1e18_)
}

// SetLicenseCost is a paid mutator transaction binding the contract method 0x1d6f8ef3.
//
// Solidity: function setLicenseCost(uint256 licenseCostInUsd1e18_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetLicenseCost(licenseCostInUsd1e18_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseCost(&_Devlicensedimo.TransactOpts, licenseCostInUsd1e18_)
}

// SetLicenseCost is a paid mutator transaction binding the contract method 0x1d6f8ef3.
//
// Solidity: function setLicenseCost(uint256 licenseCostInUsd1e18_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetLicenseCost(licenseCostInUsd1e18_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseCost(&_Devlicensedimo.TransactOpts, licenseCostInUsd1e18_)
}

// SetLicenseFactoryAddress is a paid mutator transaction binding the contract method 0x5b10dd30.
//
// Solidity: function setLicenseFactoryAddress(address licenseAccountFactory_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetLicenseFactoryAddress(opts *bind.TransactOpts, licenseAccountFactory_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setLicenseFactoryAddress", licenseAccountFactory_)
}

// SetLicenseFactoryAddress is a paid mutator transaction binding the contract method 0x5b10dd30.
//
// Solidity: function setLicenseFactoryAddress(address licenseAccountFactory_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetLicenseFactoryAddress(licenseAccountFactory_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseFactoryAddress(&_Devlicensedimo.TransactOpts, licenseAccountFactory_)
}

// SetLicenseFactoryAddress is a paid mutator transaction binding the contract method 0x5b10dd30.
//
// Solidity: function setLicenseFactoryAddress(address licenseAccountFactory_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetLicenseFactoryAddress(licenseAccountFactory_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetLicenseFactoryAddress(&_Devlicensedimo.TransactOpts, licenseAccountFactory_)
}

// SetPeriodValidity is a paid mutator transaction binding the contract method 0xd0dde29b.
//
// Solidity: function setPeriodValidity(uint256 periodValidity_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetPeriodValidity(opts *bind.TransactOpts, periodValidity_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setPeriodValidity", periodValidity_)
}

// SetPeriodValidity is a paid mutator transaction binding the contract method 0xd0dde29b.
//
// Solidity: function setPeriodValidity(uint256 periodValidity_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetPeriodValidity(periodValidity_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetPeriodValidity(&_Devlicensedimo.TransactOpts, periodValidity_)
}

// SetPeriodValidity is a paid mutator transaction binding the contract method 0xd0dde29b.
//
// Solidity: function setPeriodValidity(uint256 periodValidity_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetPeriodValidity(periodValidity_ *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetPeriodValidity(&_Devlicensedimo.TransactOpts, periodValidity_)
}

// SetPriceProviderAddress is a paid mutator transaction binding the contract method 0x920e7190.
//
// Solidity: function setPriceProviderAddress(address providerAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetPriceProviderAddress(opts *bind.TransactOpts, providerAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setPriceProviderAddress", providerAddress_)
}

// SetPriceProviderAddress is a paid mutator transaction binding the contract method 0x920e7190.
//
// Solidity: function setPriceProviderAddress(address providerAddress_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetPriceProviderAddress(providerAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetPriceProviderAddress(&_Devlicensedimo.TransactOpts, providerAddress_)
}

// SetPriceProviderAddress is a paid mutator transaction binding the contract method 0x920e7190.
//
// Solidity: function setPriceProviderAddress(address providerAddress_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetPriceProviderAddress(providerAddress_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetPriceProviderAddress(&_Devlicensedimo.TransactOpts, providerAddress_)
}

// SetReceiverAddress is a paid mutator transaction binding the contract method 0x8279c7db.
//
// Solidity: function setReceiverAddress(address receiver_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetReceiverAddress(opts *bind.TransactOpts, receiver_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setReceiverAddress", receiver_)
}

// SetReceiverAddress is a paid mutator transaction binding the contract method 0x8279c7db.
//
// Solidity: function setReceiverAddress(address receiver_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetReceiverAddress(receiver_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetReceiverAddress(&_Devlicensedimo.TransactOpts, receiver_)
}

// SetReceiverAddress is a paid mutator transaction binding the contract method 0x8279c7db.
//
// Solidity: function setReceiverAddress(address receiver_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetReceiverAddress(receiver_ common.Address) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetReceiverAddress(&_Devlicensedimo.TransactOpts, receiver_)
}

// SetRedirectUri is a paid mutator transaction binding the contract method 0xba1bedfc.
//
// Solidity: function setRedirectUri(uint256 tokenId, bool enabled, string uri) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetRedirectUri(opts *bind.TransactOpts, tokenId *big.Int, enabled bool, uri string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setRedirectUri", tokenId, enabled, uri)
}

// SetRedirectUri is a paid mutator transaction binding the contract method 0xba1bedfc.
//
// Solidity: function setRedirectUri(uint256 tokenId, bool enabled, string uri) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetRedirectUri(tokenId *big.Int, enabled bool, uri string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetRedirectUri(&_Devlicensedimo.TransactOpts, tokenId, enabled, uri)
}

// SetRedirectUri is a paid mutator transaction binding the contract method 0xba1bedfc.
//
// Solidity: function setRedirectUri(uint256 tokenId, bool enabled, string uri) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetRedirectUri(tokenId *big.Int, enabled bool, uri string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetRedirectUri(&_Devlicensedimo.TransactOpts, tokenId, enabled, uri)
}

// SetTokenDescription is a paid mutator transaction binding the contract method 0x00d5da02.
//
// Solidity: function setTokenDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetTokenDescription(opts *bind.TransactOpts, description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setTokenDescription", description_)
}

// SetTokenDescription is a paid mutator transaction binding the contract method 0x00d5da02.
//
// Solidity: function setTokenDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetTokenDescription(description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetTokenDescription(&_Devlicensedimo.TransactOpts, description_)
}

// SetTokenDescription is a paid mutator transaction binding the contract method 0x00d5da02.
//
// Solidity: function setTokenDescription(string description_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetTokenDescription(description_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetTokenDescription(&_Devlicensedimo.TransactOpts, description_)
}

// SetTokenImage is a paid mutator transaction binding the contract method 0xf806147d.
//
// Solidity: function setTokenImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) SetTokenImage(opts *bind.TransactOpts, image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "setTokenImage", image_)
}

// SetTokenImage is a paid mutator transaction binding the contract method 0xf806147d.
//
// Solidity: function setTokenImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoSession) SetTokenImage(image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetTokenImage(&_Devlicensedimo.TransactOpts, image_)
}

// SetTokenImage is a paid mutator transaction binding the contract method 0xf806147d.
//
// Solidity: function setTokenImage(string image_) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) SetTokenImage(image_ string) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.SetTokenImage(&_Devlicensedimo.TransactOpts, image_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.TransferFrom(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.TransferFrom(&_Devlicensedimo.TransactOpts, arg0, arg1, arg2)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Devlicensedimo *DevlicensedimoTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Devlicensedimo *DevlicensedimoSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.UpgradeToAndCall(&_Devlicensedimo.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.UpgradeToAndCall(&_Devlicensedimo.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactor) Withdraw(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.contract.Transact(opts, "withdraw", tokenId, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoSession) Withdraw(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Withdraw(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 amount) returns()
func (_Devlicensedimo *DevlicensedimoTransactorSession) Withdraw(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Devlicensedimo.Contract.Withdraw(&_Devlicensedimo.TransactOpts, tokenId, amount)
}

// DevlicensedimoAssetForfeitIterator is returned from FilterAssetForfeit and is used to iterate over the raw logs and unpacked data for AssetForfeit events raised by the Devlicensedimo contract.
type DevlicensedimoAssetForfeitIterator struct {
	Event *DevlicensedimoAssetForfeit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoAssetForfeitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoAssetForfeit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoAssetForfeit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoAssetForfeitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoAssetForfeitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoAssetForfeit represents a AssetForfeit event raised by the Devlicensedimo contract.
type DevlicensedimoAssetForfeit struct {
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetForfeit is a free log retrieval operation binding the contract event 0x52d30ff0b778a7633bd7cbd2a5f340804ec47e42354939955b95c7647b5250b7.
//
// Solidity: event AssetForfeit(uint256 indexed tokenId, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterAssetForfeit(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoAssetForfeitIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "AssetForfeit", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoAssetForfeitIterator{contract: _Devlicensedimo.contract, event: "AssetForfeit", logs: logs, sub: sub}, nil
}

// WatchAssetForfeit is a free log subscription operation binding the contract event 0x52d30ff0b778a7633bd7cbd2a5f340804ec47e42354939955b95c7647b5250b7.
//
// Solidity: event AssetForfeit(uint256 indexed tokenId, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchAssetForfeit(opts *bind.WatchOpts, sink chan<- *DevlicensedimoAssetForfeit, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "AssetForfeit", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoAssetForfeit)
				if err := _Devlicensedimo.contract.UnpackLog(event, "AssetForfeit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetForfeit is a log parse operation binding the contract event 0x52d30ff0b778a7633bd7cbd2a5f340804ec47e42354939955b95c7647b5250b7.
//
// Solidity: event AssetForfeit(uint256 indexed tokenId, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseAssetForfeit(log types.Log) (*DevlicensedimoAssetForfeit, error) {
	event := new(DevlicensedimoAssetForfeit)
	if err := _Devlicensedimo.contract.UnpackLog(event, "AssetForfeit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoAssetFreezeUpdateIterator is returned from FilterAssetFreezeUpdate and is used to iterate over the raw logs and unpacked data for AssetFreezeUpdate events raised by the Devlicensedimo contract.
type DevlicensedimoAssetFreezeUpdateIterator struct {
	Event *DevlicensedimoAssetFreezeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoAssetFreezeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoAssetFreezeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoAssetFreezeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoAssetFreezeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoAssetFreezeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoAssetFreezeUpdate represents a AssetFreezeUpdate event raised by the Devlicensedimo contract.
type DevlicensedimoAssetFreezeUpdate struct {
	TokenId *big.Int
	Amount  *big.Int
	Frozen  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetFreezeUpdate is a free log retrieval operation binding the contract event 0x16ba6956fa2c7ee4362cc81ca5ab2cb6bbd122fe59d7ce140387b1cd4e044899.
//
// Solidity: event AssetFreezeUpdate(uint256 indexed tokenId, uint256 amount, bool frozen)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterAssetFreezeUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoAssetFreezeUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "AssetFreezeUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoAssetFreezeUpdateIterator{contract: _Devlicensedimo.contract, event: "AssetFreezeUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetFreezeUpdate is a free log subscription operation binding the contract event 0x16ba6956fa2c7ee4362cc81ca5ab2cb6bbd122fe59d7ce140387b1cd4e044899.
//
// Solidity: event AssetFreezeUpdate(uint256 indexed tokenId, uint256 amount, bool frozen)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchAssetFreezeUpdate(opts *bind.WatchOpts, sink chan<- *DevlicensedimoAssetFreezeUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "AssetFreezeUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoAssetFreezeUpdate)
				if err := _Devlicensedimo.contract.UnpackLog(event, "AssetFreezeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetFreezeUpdate is a log parse operation binding the contract event 0x16ba6956fa2c7ee4362cc81ca5ab2cb6bbd122fe59d7ce140387b1cd4e044899.
//
// Solidity: event AssetFreezeUpdate(uint256 indexed tokenId, uint256 amount, bool frozen)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseAssetFreezeUpdate(log types.Log) (*DevlicensedimoAssetFreezeUpdate, error) {
	event := new(DevlicensedimoAssetFreezeUpdate)
	if err := _Devlicensedimo.contract.UnpackLog(event, "AssetFreezeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Devlicensedimo contract.
type DevlicensedimoInitializedIterator struct {
	Event *DevlicensedimoInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoInitialized represents a Initialized event raised by the Devlicensedimo contract.
type DevlicensedimoInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterInitialized(opts *bind.FilterOpts) (*DevlicensedimoInitializedIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoInitializedIterator{contract: _Devlicensedimo.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DevlicensedimoInitialized) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoInitialized)
				if err := _Devlicensedimo.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseInitialized(log types.Log) (*DevlicensedimoInitialized, error) {
	event := new(DevlicensedimoInitialized)
	if err := _Devlicensedimo.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Devlicensedimo contract.
type DevlicensedimoIssuedIterator struct {
	Event *DevlicensedimoIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoIssued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoIssued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoIssued represents a Issued event raised by the Devlicensedimo contract.
type DevlicensedimoIssued struct {
	TokenId  *big.Int
	Owner    common.Address
	ClientId common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0x7533f62ec6601bf9c87f8d96bf756b4b495e2a0e26ec9284e4927926ed6b3afd.
//
// Solidity: event Issued(uint256 indexed tokenId, address indexed owner, address indexed clientId)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterIssued(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, clientId []common.Address) (*DevlicensedimoIssuedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "Issued", tokenIdRule, ownerRule, clientIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoIssuedIterator{contract: _Devlicensedimo.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0x7533f62ec6601bf9c87f8d96bf756b4b495e2a0e26ec9284e4927926ed6b3afd.
//
// Solidity: event Issued(uint256 indexed tokenId, address indexed owner, address indexed clientId)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *DevlicensedimoIssued, tokenId []*big.Int, owner []common.Address, clientId []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var clientIdRule []interface{}
	for _, clientIdItem := range clientId {
		clientIdRule = append(clientIdRule, clientIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "Issued", tokenIdRule, ownerRule, clientIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoIssued)
				if err := _Devlicensedimo.contract.UnpackLog(event, "Issued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssued is a log parse operation binding the contract event 0x7533f62ec6601bf9c87f8d96bf756b4b495e2a0e26ec9284e4927926ed6b3afd.
//
// Solidity: event Issued(uint256 indexed tokenId, address indexed owner, address indexed clientId)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseIssued(log types.Log) (*DevlicensedimoIssued, error) {
	event := new(DevlicensedimoIssued)
	if err := _Devlicensedimo.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoLicenseAliasSetIterator is returned from FilterLicenseAliasSet and is used to iterate over the raw logs and unpacked data for LicenseAliasSet events raised by the Devlicensedimo contract.
type DevlicensedimoLicenseAliasSetIterator struct {
	Event *DevlicensedimoLicenseAliasSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoLicenseAliasSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoLicenseAliasSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoLicenseAliasSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoLicenseAliasSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoLicenseAliasSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoLicenseAliasSet represents a LicenseAliasSet event raised by the Devlicensedimo contract.
type DevlicensedimoLicenseAliasSet struct {
	TokenId      *big.Int
	LicenseAlias string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLicenseAliasSet is a free log retrieval operation binding the contract event 0xcd1647e82913e163e110815fdac0265c513795a376da7a90f90090cea69905a7.
//
// Solidity: event LicenseAliasSet(uint256 indexed tokenId, string licenseAlias)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterLicenseAliasSet(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoLicenseAliasSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "LicenseAliasSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoLicenseAliasSetIterator{contract: _Devlicensedimo.contract, event: "LicenseAliasSet", logs: logs, sub: sub}, nil
}

// WatchLicenseAliasSet is a free log subscription operation binding the contract event 0xcd1647e82913e163e110815fdac0265c513795a376da7a90f90090cea69905a7.
//
// Solidity: event LicenseAliasSet(uint256 indexed tokenId, string licenseAlias)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchLicenseAliasSet(opts *bind.WatchOpts, sink chan<- *DevlicensedimoLicenseAliasSet, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "LicenseAliasSet", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoLicenseAliasSet)
				if err := _Devlicensedimo.contract.UnpackLog(event, "LicenseAliasSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLicenseAliasSet is a log parse operation binding the contract event 0xcd1647e82913e163e110815fdac0265c513795a376da7a90f90090cea69905a7.
//
// Solidity: event LicenseAliasSet(uint256 indexed tokenId, string licenseAlias)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseLicenseAliasSet(log types.Log) (*DevlicensedimoLicenseAliasSet, error) {
	event := new(DevlicensedimoLicenseAliasSet)
	if err := _Devlicensedimo.contract.UnpackLog(event, "LicenseAliasSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Devlicensedimo contract.
type DevlicensedimoLockedIterator struct {
	Event *DevlicensedimoLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoLocked represents a Locked event raised by the Devlicensedimo contract.
type DevlicensedimoLocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterLocked(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoLockedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "Locked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoLockedIterator{contract: _Devlicensedimo.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *DevlicensedimoLocked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "Locked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoLocked)
				if err := _Devlicensedimo.contract.UnpackLog(event, "Locked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLocked is a log parse operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseLocked(log types.Log) (*DevlicensedimoLocked, error) {
	event := new(DevlicensedimoLocked)
	if err := _Devlicensedimo.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoRedirectUriDisabledIterator is returned from FilterRedirectUriDisabled and is used to iterate over the raw logs and unpacked data for RedirectUriDisabled events raised by the Devlicensedimo contract.
type DevlicensedimoRedirectUriDisabledIterator struct {
	Event *DevlicensedimoRedirectUriDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoRedirectUriDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoRedirectUriDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoRedirectUriDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoRedirectUriDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoRedirectUriDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoRedirectUriDisabled represents a RedirectUriDisabled event raised by the Devlicensedimo contract.
type DevlicensedimoRedirectUriDisabled struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedirectUriDisabled is a free log retrieval operation binding the contract event 0xa5ad928ebfe9bbabb9b65e85335dc567892182dc87c0b7b1fef8d054db2383a3.
//
// Solidity: event RedirectUriDisabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterRedirectUriDisabled(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoRedirectUriDisabledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "RedirectUriDisabled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoRedirectUriDisabledIterator{contract: _Devlicensedimo.contract, event: "RedirectUriDisabled", logs: logs, sub: sub}, nil
}

// WatchRedirectUriDisabled is a free log subscription operation binding the contract event 0xa5ad928ebfe9bbabb9b65e85335dc567892182dc87c0b7b1fef8d054db2383a3.
//
// Solidity: event RedirectUriDisabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchRedirectUriDisabled(opts *bind.WatchOpts, sink chan<- *DevlicensedimoRedirectUriDisabled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "RedirectUriDisabled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoRedirectUriDisabled)
				if err := _Devlicensedimo.contract.UnpackLog(event, "RedirectUriDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedirectUriDisabled is a log parse operation binding the contract event 0xa5ad928ebfe9bbabb9b65e85335dc567892182dc87c0b7b1fef8d054db2383a3.
//
// Solidity: event RedirectUriDisabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseRedirectUriDisabled(log types.Log) (*DevlicensedimoRedirectUriDisabled, error) {
	event := new(DevlicensedimoRedirectUriDisabled)
	if err := _Devlicensedimo.contract.UnpackLog(event, "RedirectUriDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoRedirectUriEnabledIterator is returned from FilterRedirectUriEnabled and is used to iterate over the raw logs and unpacked data for RedirectUriEnabled events raised by the Devlicensedimo contract.
type DevlicensedimoRedirectUriEnabledIterator struct {
	Event *DevlicensedimoRedirectUriEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoRedirectUriEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoRedirectUriEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoRedirectUriEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoRedirectUriEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoRedirectUriEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoRedirectUriEnabled represents a RedirectUriEnabled event raised by the Devlicensedimo contract.
type DevlicensedimoRedirectUriEnabled struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRedirectUriEnabled is a free log retrieval operation binding the contract event 0x100a5934f50dcbea6ca17171581b081113167276cba43b6ca421a7bb1b9704f0.
//
// Solidity: event RedirectUriEnabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterRedirectUriEnabled(opts *bind.FilterOpts, tokenId []*big.Int) (*DevlicensedimoRedirectUriEnabledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "RedirectUriEnabled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoRedirectUriEnabledIterator{contract: _Devlicensedimo.contract, event: "RedirectUriEnabled", logs: logs, sub: sub}, nil
}

// WatchRedirectUriEnabled is a free log subscription operation binding the contract event 0x100a5934f50dcbea6ca17171581b081113167276cba43b6ca421a7bb1b9704f0.
//
// Solidity: event RedirectUriEnabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchRedirectUriEnabled(opts *bind.WatchOpts, sink chan<- *DevlicensedimoRedirectUriEnabled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "RedirectUriEnabled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoRedirectUriEnabled)
				if err := _Devlicensedimo.contract.UnpackLog(event, "RedirectUriEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedirectUriEnabled is a log parse operation binding the contract event 0x100a5934f50dcbea6ca17171581b081113167276cba43b6ca421a7bb1b9704f0.
//
// Solidity: event RedirectUriEnabled(uint256 indexed tokenId, string uri)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseRedirectUriEnabled(log types.Log) (*DevlicensedimoRedirectUriEnabled, error) {
	event := new(DevlicensedimoRedirectUriEnabled)
	if err := _Devlicensedimo.contract.UnpackLog(event, "RedirectUriEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Devlicensedimo contract.
type DevlicensedimoRoleAdminChangedIterator struct {
	Event *DevlicensedimoRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoRoleAdminChanged represents a RoleAdminChanged event raised by the Devlicensedimo contract.
type DevlicensedimoRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DevlicensedimoRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoRoleAdminChangedIterator{contract: _Devlicensedimo.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DevlicensedimoRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoRoleAdminChanged)
				if err := _Devlicensedimo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseRoleAdminChanged(log types.Log) (*DevlicensedimoRoleAdminChanged, error) {
	event := new(DevlicensedimoRoleAdminChanged)
	if err := _Devlicensedimo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Devlicensedimo contract.
type DevlicensedimoRoleGrantedIterator struct {
	Event *DevlicensedimoRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoRoleGranted represents a RoleGranted event raised by the Devlicensedimo contract.
type DevlicensedimoRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DevlicensedimoRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoRoleGrantedIterator{contract: _Devlicensedimo.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DevlicensedimoRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoRoleGranted)
				if err := _Devlicensedimo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseRoleGranted(log types.Log) (*DevlicensedimoRoleGranted, error) {
	event := new(DevlicensedimoRoleGranted)
	if err := _Devlicensedimo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Devlicensedimo contract.
type DevlicensedimoRoleRevokedIterator struct {
	Event *DevlicensedimoRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoRoleRevoked represents a RoleRevoked event raised by the Devlicensedimo contract.
type DevlicensedimoRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DevlicensedimoRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoRoleRevokedIterator{contract: _Devlicensedimo.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DevlicensedimoRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoRoleRevoked)
				if err := _Devlicensedimo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseRoleRevoked(log types.Log) (*DevlicensedimoRoleRevoked, error) {
	event := new(DevlicensedimoRoleRevoked)
	if err := _Devlicensedimo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoSignerDisabledIterator is returned from FilterSignerDisabled and is used to iterate over the raw logs and unpacked data for SignerDisabled events raised by the Devlicensedimo contract.
type DevlicensedimoSignerDisabledIterator struct {
	Event *DevlicensedimoSignerDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoSignerDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoSignerDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoSignerDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoSignerDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoSignerDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoSignerDisabled represents a SignerDisabled event raised by the Devlicensedimo contract.
type DevlicensedimoSignerDisabled struct {
	TokenId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerDisabled is a free log retrieval operation binding the contract event 0xdbd8d9c504c83e3bec8383041e811568d4eb37e94697dbfce9a50e8ae572017a.
//
// Solidity: event SignerDisabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterSignerDisabled(opts *bind.FilterOpts, tokenId []*big.Int, signer []common.Address) (*DevlicensedimoSignerDisabledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "SignerDisabled", tokenIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoSignerDisabledIterator{contract: _Devlicensedimo.contract, event: "SignerDisabled", logs: logs, sub: sub}, nil
}

// WatchSignerDisabled is a free log subscription operation binding the contract event 0xdbd8d9c504c83e3bec8383041e811568d4eb37e94697dbfce9a50e8ae572017a.
//
// Solidity: event SignerDisabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchSignerDisabled(opts *bind.WatchOpts, sink chan<- *DevlicensedimoSignerDisabled, tokenId []*big.Int, signer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "SignerDisabled", tokenIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoSignerDisabled)
				if err := _Devlicensedimo.contract.UnpackLog(event, "SignerDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerDisabled is a log parse operation binding the contract event 0xdbd8d9c504c83e3bec8383041e811568d4eb37e94697dbfce9a50e8ae572017a.
//
// Solidity: event SignerDisabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseSignerDisabled(log types.Log) (*DevlicensedimoSignerDisabled, error) {
	event := new(DevlicensedimoSignerDisabled)
	if err := _Devlicensedimo.contract.UnpackLog(event, "SignerDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoSignerEnabledIterator is returned from FilterSignerEnabled and is used to iterate over the raw logs and unpacked data for SignerEnabled events raised by the Devlicensedimo contract.
type DevlicensedimoSignerEnabledIterator struct {
	Event *DevlicensedimoSignerEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoSignerEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoSignerEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoSignerEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoSignerEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoSignerEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoSignerEnabled represents a SignerEnabled event raised by the Devlicensedimo contract.
type DevlicensedimoSignerEnabled struct {
	TokenId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignerEnabled is a free log retrieval operation binding the contract event 0x0ea45d41f8d2f0fd7e53e2a0821276268b7cf41414e2c2eac7bf8cd69241386b.
//
// Solidity: event SignerEnabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterSignerEnabled(opts *bind.FilterOpts, tokenId []*big.Int, signer []common.Address) (*DevlicensedimoSignerEnabledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "SignerEnabled", tokenIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoSignerEnabledIterator{contract: _Devlicensedimo.contract, event: "SignerEnabled", logs: logs, sub: sub}, nil
}

// WatchSignerEnabled is a free log subscription operation binding the contract event 0x0ea45d41f8d2f0fd7e53e2a0821276268b7cf41414e2c2eac7bf8cd69241386b.
//
// Solidity: event SignerEnabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchSignerEnabled(opts *bind.WatchOpts, sink chan<- *DevlicensedimoSignerEnabled, tokenId []*big.Int, signer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "SignerEnabled", tokenIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoSignerEnabled)
				if err := _Devlicensedimo.contract.UnpackLog(event, "SignerEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerEnabled is a log parse operation binding the contract event 0x0ea45d41f8d2f0fd7e53e2a0821276268b7cf41414e2c2eac7bf8cd69241386b.
//
// Solidity: event SignerEnabled(uint256 indexed tokenId, address indexed signer)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseSignerEnabled(log types.Log) (*DevlicensedimoSignerEnabled, error) {
	event := new(DevlicensedimoSignerEnabled)
	if err := _Devlicensedimo.contract.UnpackLog(event, "SignerEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoStakeDepositIterator is returned from FilterStakeDeposit and is used to iterate over the raw logs and unpacked data for StakeDeposit events raised by the Devlicensedimo contract.
type DevlicensedimoStakeDepositIterator struct {
	Event *DevlicensedimoStakeDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoStakeDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoStakeDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoStakeDeposit represents a StakeDeposit event raised by the Devlicensedimo contract.
type DevlicensedimoStakeDeposit struct {
	TokenId *big.Int
	User    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposit is a free log retrieval operation binding the contract event 0x1cd3a07c960f34280cc48968d5b5315035ebdde3a1e371aa539bf0a415743a3b.
//
// Solidity: event StakeDeposit(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterStakeDeposit(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*DevlicensedimoStakeDepositIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "StakeDeposit", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoStakeDepositIterator{contract: _Devlicensedimo.contract, event: "StakeDeposit", logs: logs, sub: sub}, nil
}

// WatchStakeDeposit is a free log subscription operation binding the contract event 0x1cd3a07c960f34280cc48968d5b5315035ebdde3a1e371aa539bf0a415743a3b.
//
// Solidity: event StakeDeposit(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchStakeDeposit(opts *bind.WatchOpts, sink chan<- *DevlicensedimoStakeDeposit, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "StakeDeposit", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoStakeDeposit)
				if err := _Devlicensedimo.contract.UnpackLog(event, "StakeDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeDeposit is a log parse operation binding the contract event 0x1cd3a07c960f34280cc48968d5b5315035ebdde3a1e371aa539bf0a415743a3b.
//
// Solidity: event StakeDeposit(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseStakeDeposit(log types.Log) (*DevlicensedimoStakeDeposit, error) {
	event := new(DevlicensedimoStakeDeposit)
	if err := _Devlicensedimo.contract.UnpackLog(event, "StakeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoStakeWithdrawIterator is returned from FilterStakeWithdraw and is used to iterate over the raw logs and unpacked data for StakeWithdraw events raised by the Devlicensedimo contract.
type DevlicensedimoStakeWithdrawIterator struct {
	Event *DevlicensedimoStakeWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoStakeWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoStakeWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoStakeWithdraw represents a StakeWithdraw event raised by the Devlicensedimo contract.
type DevlicensedimoStakeWithdraw struct {
	TokenId *big.Int
	User    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdraw is a free log retrieval operation binding the contract event 0xb632f81a2da981f3daa6d867e773bbab8f2931ca133d2cca3595eb61af045c3a.
//
// Solidity: event StakeWithdraw(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterStakeWithdraw(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*DevlicensedimoStakeWithdrawIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "StakeWithdraw", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoStakeWithdrawIterator{contract: _Devlicensedimo.contract, event: "StakeWithdraw", logs: logs, sub: sub}, nil
}

// WatchStakeWithdraw is a free log subscription operation binding the contract event 0xb632f81a2da981f3daa6d867e773bbab8f2931ca133d2cca3595eb61af045c3a.
//
// Solidity: event StakeWithdraw(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchStakeWithdraw(opts *bind.WatchOpts, sink chan<- *DevlicensedimoStakeWithdraw, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "StakeWithdraw", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoStakeWithdraw)
				if err := _Devlicensedimo.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdraw is a log parse operation binding the contract event 0xb632f81a2da981f3daa6d867e773bbab8f2931ca133d2cca3595eb61af045c3a.
//
// Solidity: event StakeWithdraw(uint256 indexed tokenId, address indexed user, uint256 amount)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseStakeWithdraw(log types.Log) (*DevlicensedimoStakeWithdraw, error) {
	event := new(DevlicensedimoStakeWithdraw)
	if err := _Devlicensedimo.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Devlicensedimo contract.
type DevlicensedimoTransferIterator struct {
	Event *DevlicensedimoTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoTransfer represents a Transfer event raised by the Devlicensedimo contract.
type DevlicensedimoTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DevlicensedimoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoTransferIterator{contract: _Devlicensedimo.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DevlicensedimoTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoTransfer)
				if err := _Devlicensedimo.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseTransfer(log types.Log) (*DevlicensedimoTransfer, error) {
	event := new(DevlicensedimoTransfer)
	if err := _Devlicensedimo.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdateDimoCreditAddressIterator is returned from FilterUpdateDimoCreditAddress and is used to iterate over the raw logs and unpacked data for UpdateDimoCreditAddress events raised by the Devlicensedimo contract.
type DevlicensedimoUpdateDimoCreditAddressIterator struct {
	Event *DevlicensedimoUpdateDimoCreditAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdateDimoCreditAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdateDimoCreditAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdateDimoCreditAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdateDimoCreditAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdateDimoCreditAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdateDimoCreditAddress represents a UpdateDimoCreditAddress event raised by the Devlicensedimo contract.
type DevlicensedimoUpdateDimoCreditAddress struct {
	DimoCredit common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDimoCreditAddress is a free log retrieval operation binding the contract event 0x695a889a1604567698a4ed7ebaa247b3baee08ed9030d7bceb2a8e66201823fe.
//
// Solidity: event UpdateDimoCreditAddress(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdateDimoCreditAddress(opts *bind.FilterOpts) (*DevlicensedimoUpdateDimoCreditAddressIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdateDimoCreditAddress")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdateDimoCreditAddressIterator{contract: _Devlicensedimo.contract, event: "UpdateDimoCreditAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateDimoCreditAddress is a free log subscription operation binding the contract event 0x695a889a1604567698a4ed7ebaa247b3baee08ed9030d7bceb2a8e66201823fe.
//
// Solidity: event UpdateDimoCreditAddress(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdateDimoCreditAddress(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdateDimoCreditAddress) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdateDimoCreditAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdateDimoCreditAddress)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateDimoCreditAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateDimoCreditAddress is a log parse operation binding the contract event 0x695a889a1604567698a4ed7ebaa247b3baee08ed9030d7bceb2a8e66201823fe.
//
// Solidity: event UpdateDimoCreditAddress(address dimoCredit_)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdateDimoCreditAddress(log types.Log) (*DevlicensedimoUpdateDimoCreditAddress, error) {
	event := new(DevlicensedimoUpdateDimoCreditAddress)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateDimoCreditAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdateDimoTokenAddressIterator is returned from FilterUpdateDimoTokenAddress and is used to iterate over the raw logs and unpacked data for UpdateDimoTokenAddress events raised by the Devlicensedimo contract.
type DevlicensedimoUpdateDimoTokenAddressIterator struct {
	Event *DevlicensedimoUpdateDimoTokenAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdateDimoTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdateDimoTokenAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdateDimoTokenAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdateDimoTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdateDimoTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdateDimoTokenAddress represents a UpdateDimoTokenAddress event raised by the Devlicensedimo contract.
type DevlicensedimoUpdateDimoTokenAddress struct {
	DimoToken common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateDimoTokenAddress is a free log retrieval operation binding the contract event 0xbf31bafe39dc11dc61bafdb5ef79888f1b0ec64e94705fd2da62dc8d8baf1f67.
//
// Solidity: event UpdateDimoTokenAddress(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdateDimoTokenAddress(opts *bind.FilterOpts) (*DevlicensedimoUpdateDimoTokenAddressIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdateDimoTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdateDimoTokenAddressIterator{contract: _Devlicensedimo.contract, event: "UpdateDimoTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateDimoTokenAddress is a free log subscription operation binding the contract event 0xbf31bafe39dc11dc61bafdb5ef79888f1b0ec64e94705fd2da62dc8d8baf1f67.
//
// Solidity: event UpdateDimoTokenAddress(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdateDimoTokenAddress(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdateDimoTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdateDimoTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdateDimoTokenAddress)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateDimoTokenAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateDimoTokenAddress is a log parse operation binding the contract event 0xbf31bafe39dc11dc61bafdb5ef79888f1b0ec64e94705fd2da62dc8d8baf1f67.
//
// Solidity: event UpdateDimoTokenAddress(address dimoToken_)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdateDimoTokenAddress(log types.Log) (*DevlicensedimoUpdateDimoTokenAddress, error) {
	event := new(DevlicensedimoUpdateDimoTokenAddress)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateDimoTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdateLicenseAccountFactoryAddressIterator is returned from FilterUpdateLicenseAccountFactoryAddress and is used to iterate over the raw logs and unpacked data for UpdateLicenseAccountFactoryAddress events raised by the Devlicensedimo contract.
type DevlicensedimoUpdateLicenseAccountFactoryAddressIterator struct {
	Event *DevlicensedimoUpdateLicenseAccountFactoryAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdateLicenseAccountFactoryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdateLicenseAccountFactoryAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdateLicenseAccountFactoryAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdateLicenseAccountFactoryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdateLicenseAccountFactoryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdateLicenseAccountFactoryAddress represents a UpdateLicenseAccountFactoryAddress event raised by the Devlicensedimo contract.
type DevlicensedimoUpdateLicenseAccountFactoryAddress struct {
	LicenseAccountFactory common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateLicenseAccountFactoryAddress is a free log retrieval operation binding the contract event 0xa5898a6a6186a7e784bff87e4255e15d37f8329e62341db8c79a10e117b23f38.
//
// Solidity: event UpdateLicenseAccountFactoryAddress(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdateLicenseAccountFactoryAddress(opts *bind.FilterOpts) (*DevlicensedimoUpdateLicenseAccountFactoryAddressIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdateLicenseAccountFactoryAddress")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdateLicenseAccountFactoryAddressIterator{contract: _Devlicensedimo.contract, event: "UpdateLicenseAccountFactoryAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateLicenseAccountFactoryAddress is a free log subscription operation binding the contract event 0xa5898a6a6186a7e784bff87e4255e15d37f8329e62341db8c79a10e117b23f38.
//
// Solidity: event UpdateLicenseAccountFactoryAddress(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdateLicenseAccountFactoryAddress(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdateLicenseAccountFactoryAddress) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdateLicenseAccountFactoryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdateLicenseAccountFactoryAddress)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateLicenseAccountFactoryAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateLicenseAccountFactoryAddress is a log parse operation binding the contract event 0xa5898a6a6186a7e784bff87e4255e15d37f8329e62341db8c79a10e117b23f38.
//
// Solidity: event UpdateLicenseAccountFactoryAddress(address licenseAccountFactory_)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdateLicenseAccountFactoryAddress(log types.Log) (*DevlicensedimoUpdateLicenseAccountFactoryAddress, error) {
	event := new(DevlicensedimoUpdateLicenseAccountFactoryAddress)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateLicenseAccountFactoryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdateLicenseCostIterator is returned from FilterUpdateLicenseCost and is used to iterate over the raw logs and unpacked data for UpdateLicenseCost events raised by the Devlicensedimo contract.
type DevlicensedimoUpdateLicenseCostIterator struct {
	Event *DevlicensedimoUpdateLicenseCost // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdateLicenseCostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdateLicenseCost)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdateLicenseCost)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdateLicenseCostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdateLicenseCostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdateLicenseCost represents a UpdateLicenseCost event raised by the Devlicensedimo contract.
type DevlicensedimoUpdateLicenseCost struct {
	LicenseCost *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateLicenseCost is a free log retrieval operation binding the contract event 0x0aca62d366c4fbb1ef303879453cb890975077ec0dcc742f25c0736c2f563652.
//
// Solidity: event UpdateLicenseCost(uint256 licenseCost)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdateLicenseCost(opts *bind.FilterOpts) (*DevlicensedimoUpdateLicenseCostIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdateLicenseCost")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdateLicenseCostIterator{contract: _Devlicensedimo.contract, event: "UpdateLicenseCost", logs: logs, sub: sub}, nil
}

// WatchUpdateLicenseCost is a free log subscription operation binding the contract event 0x0aca62d366c4fbb1ef303879453cb890975077ec0dcc742f25c0736c2f563652.
//
// Solidity: event UpdateLicenseCost(uint256 licenseCost)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdateLicenseCost(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdateLicenseCost) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdateLicenseCost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdateLicenseCost)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateLicenseCost", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateLicenseCost is a log parse operation binding the contract event 0x0aca62d366c4fbb1ef303879453cb890975077ec0dcc742f25c0736c2f563652.
//
// Solidity: event UpdateLicenseCost(uint256 licenseCost)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdateLicenseCost(log types.Log) (*DevlicensedimoUpdateLicenseCost, error) {
	event := new(DevlicensedimoUpdateLicenseCost)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateLicenseCost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdatePeriodValidityIterator is returned from FilterUpdatePeriodValidity and is used to iterate over the raw logs and unpacked data for UpdatePeriodValidity events raised by the Devlicensedimo contract.
type DevlicensedimoUpdatePeriodValidityIterator struct {
	Event *DevlicensedimoUpdatePeriodValidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdatePeriodValidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdatePeriodValidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdatePeriodValidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdatePeriodValidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdatePeriodValidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdatePeriodValidity represents a UpdatePeriodValidity event raised by the Devlicensedimo contract.
type DevlicensedimoUpdatePeriodValidity struct {
	PeriodValidity *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatePeriodValidity is a free log retrieval operation binding the contract event 0x9a360497eca32e7df7472fb414b7188a59505b079b391e7f3a12d59cc4c3fabd.
//
// Solidity: event UpdatePeriodValidity(uint256 periodValidity)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdatePeriodValidity(opts *bind.FilterOpts) (*DevlicensedimoUpdatePeriodValidityIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdatePeriodValidity")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdatePeriodValidityIterator{contract: _Devlicensedimo.contract, event: "UpdatePeriodValidity", logs: logs, sub: sub}, nil
}

// WatchUpdatePeriodValidity is a free log subscription operation binding the contract event 0x9a360497eca32e7df7472fb414b7188a59505b079b391e7f3a12d59cc4c3fabd.
//
// Solidity: event UpdatePeriodValidity(uint256 periodValidity)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdatePeriodValidity(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdatePeriodValidity) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdatePeriodValidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdatePeriodValidity)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdatePeriodValidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePeriodValidity is a log parse operation binding the contract event 0x9a360497eca32e7df7472fb414b7188a59505b079b391e7f3a12d59cc4c3fabd.
//
// Solidity: event UpdatePeriodValidity(uint256 periodValidity)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdatePeriodValidity(log types.Log) (*DevlicensedimoUpdatePeriodValidity, error) {
	event := new(DevlicensedimoUpdatePeriodValidity)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdatePeriodValidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdatePriceProviderAddressIterator is returned from FilterUpdatePriceProviderAddress and is used to iterate over the raw logs and unpacked data for UpdatePriceProviderAddress events raised by the Devlicensedimo contract.
type DevlicensedimoUpdatePriceProviderAddressIterator struct {
	Event *DevlicensedimoUpdatePriceProviderAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdatePriceProviderAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdatePriceProviderAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdatePriceProviderAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdatePriceProviderAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdatePriceProviderAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdatePriceProviderAddress represents a UpdatePriceProviderAddress event raised by the Devlicensedimo contract.
type DevlicensedimoUpdatePriceProviderAddress struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatePriceProviderAddress is a free log retrieval operation binding the contract event 0x47641d988225801ecd2b57b094546634e12b60dbb6082cd3bb56e40a134ee643.
//
// Solidity: event UpdatePriceProviderAddress(address provider)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdatePriceProviderAddress(opts *bind.FilterOpts) (*DevlicensedimoUpdatePriceProviderAddressIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdatePriceProviderAddress")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdatePriceProviderAddressIterator{contract: _Devlicensedimo.contract, event: "UpdatePriceProviderAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatePriceProviderAddress is a free log subscription operation binding the contract event 0x47641d988225801ecd2b57b094546634e12b60dbb6082cd3bb56e40a134ee643.
//
// Solidity: event UpdatePriceProviderAddress(address provider)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdatePriceProviderAddress(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdatePriceProviderAddress) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdatePriceProviderAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdatePriceProviderAddress)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdatePriceProviderAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePriceProviderAddress is a log parse operation binding the contract event 0x47641d988225801ecd2b57b094546634e12b60dbb6082cd3bb56e40a134ee643.
//
// Solidity: event UpdatePriceProviderAddress(address provider)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdatePriceProviderAddress(log types.Log) (*DevlicensedimoUpdatePriceProviderAddress, error) {
	event := new(DevlicensedimoUpdatePriceProviderAddress)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdatePriceProviderAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpdateReceiverAddressIterator is returned from FilterUpdateReceiverAddress and is used to iterate over the raw logs and unpacked data for UpdateReceiverAddress events raised by the Devlicensedimo contract.
type DevlicensedimoUpdateReceiverAddressIterator struct {
	Event *DevlicensedimoUpdateReceiverAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpdateReceiverAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpdateReceiverAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpdateReceiverAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpdateReceiverAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpdateReceiverAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpdateReceiverAddress represents a UpdateReceiverAddress event raised by the Devlicensedimo contract.
type DevlicensedimoUpdateReceiverAddress struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateReceiverAddress is a free log retrieval operation binding the contract event 0x61dcb70b81338b32ccb30ca03da4d7aa2719e47931bfa009e8b83c0043234e4a.
//
// Solidity: event UpdateReceiverAddress(address receiver_)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpdateReceiverAddress(opts *bind.FilterOpts) (*DevlicensedimoUpdateReceiverAddressIterator, error) {

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "UpdateReceiverAddress")
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpdateReceiverAddressIterator{contract: _Devlicensedimo.contract, event: "UpdateReceiverAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateReceiverAddress is a free log subscription operation binding the contract event 0x61dcb70b81338b32ccb30ca03da4d7aa2719e47931bfa009e8b83c0043234e4a.
//
// Solidity: event UpdateReceiverAddress(address receiver_)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpdateReceiverAddress(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpdateReceiverAddress) (event.Subscription, error) {

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "UpdateReceiverAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpdateReceiverAddress)
				if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateReceiverAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateReceiverAddress is a log parse operation binding the contract event 0x61dcb70b81338b32ccb30ca03da4d7aa2719e47931bfa009e8b83c0043234e4a.
//
// Solidity: event UpdateReceiverAddress(address receiver_)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpdateReceiverAddress(log types.Log) (*DevlicensedimoUpdateReceiverAddress, error) {
	event := new(DevlicensedimoUpdateReceiverAddress)
	if err := _Devlicensedimo.contract.UnpackLog(event, "UpdateReceiverAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DevlicensedimoUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Devlicensedimo contract.
type DevlicensedimoUpgradedIterator struct {
	Event *DevlicensedimoUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DevlicensedimoUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DevlicensedimoUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DevlicensedimoUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DevlicensedimoUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DevlicensedimoUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DevlicensedimoUpgraded represents a Upgraded event raised by the Devlicensedimo contract.
type DevlicensedimoUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Devlicensedimo *DevlicensedimoFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DevlicensedimoUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Devlicensedimo.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DevlicensedimoUpgradedIterator{contract: _Devlicensedimo.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Devlicensedimo *DevlicensedimoFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DevlicensedimoUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Devlicensedimo.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DevlicensedimoUpgraded)
				if err := _Devlicensedimo.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Devlicensedimo *DevlicensedimoFilterer) ParseUpgraded(log types.Log) (*DevlicensedimoUpgraded, error) {
	event := new(DevlicensedimoUpgraded)
	if err := _Devlicensedimo.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
