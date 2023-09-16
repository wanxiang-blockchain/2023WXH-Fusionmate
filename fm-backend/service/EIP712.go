package service

import (
	"encoding/hex"
	"fmt"
	"github.com/FusionMate/fm-backend/conf"
	"github.com/FusionMate/fm-backend/util"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

func newLoginSignerData(address, message, timestamp string) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"FusionMateChallenge": []apitypes.Type{
				{Name: "address", Type: "address"},
				{Name: "message", Type: "string"},
				{Name: "timestamp", Type: "string"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "FusionMateChallenge",
		Domain: apitypes.TypedDataDomain{
			Name:    "FusionMate",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(conf.GConfig.GetInt64("contract.networkId")),
		},
		Message: apitypes.TypedDataMessage{
			"timestamp": timestamp,
			"address":   address,
			"message":   message,
		},
	}
}

func newNFTContractCreateSignerData(name, symbol, baseURI, makerAddress string, collectionId, maxSupply, mintPrice *math.HexOrDecimal256, factoryContract string) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"FusionMateNFTContractCreation": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "symbol", Type: "string"},
				{Name: "baseURI", Type: "string"},
				{Name: "makerAddress", Type: "address"},
				{Name: "collectionId", Type: "uint256"},
				{Name: "maxSupply", Type: "uint256"},
				{Name: "mintPrice", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "FusionMateNFTContractCreation",
		Domain: apitypes.TypedDataDomain{
			Name:              "FusionMate",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(conf.GConfig.GetInt64("contract.networkId")),
			VerifyingContract: factoryContract,
		},
		Message: apitypes.TypedDataMessage{
			"name":         name,
			"symbol":       symbol,
			"baseURI":      baseURI,
			"makerAddress": makerAddress,
			"collectionId": collectionId,
			"maxSupply":    maxSupply,
			"mintPrice":    mintPrice,
		},
	}
}

func newHarvestSignerData(harvested, collectionId, tokenId *math.HexOrDecimal256, factoryContract string) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"FusionMateHarvest": []apitypes.Type{
				{Name: "amount", Type: "uint256"},
				{Name: "collectionId", Type: "uint256"},
				{Name: "tokenId", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "FusionMateHarvest",
		Domain: apitypes.TypedDataDomain{
			Name:              "FusionMate",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(conf.GConfig.GetInt64("contract.networkId")),
			VerifyingContract: factoryContract,
		},
		Message: apitypes.TypedDataMessage{
			"amount":       harvested,
			"collectionId": collectionId,
			"tokenId":      tokenId,
		},
	}
}

func LoginSign(address, message, timestamp, privateKey string) (string, error) {
	signerData := newLoginSignerData(address, message, timestamp)
	return EIP712SignatureSign(signerData, privateKey)
}

func LoginVerify(address, message, timestamp, signature string) (bool, error) {
	signerData := newLoginSignerData(address, message, timestamp)
	return EIP712SignatureVerify(address, signature, signerData)
}

// NFTContractCreateSign use for signing NFT contract(AssistantCollection) creations
// Params:
//   - name: name of Assistant Collection, param from frontend
//   - symbol: symbol of Assistant Collection, usually as a shortened name, created by backend
//   - baseURI: created by backend, tokenURI = baseURI+"/:tokenId", address like https://api.fusionmate.xyz/api/v1/metadata/:CollectionId/:TokenId
//   - makerAddress: the creator of this Assistant Collection, can be fetched from the access token
//   - privateKey: the backend EOA private Key, can be imported from ENV params by:
//     ```privateKey := os.Getenv("PRIVATE_KEY")```
//   - _collectionId: created by backend
//   - _maxSupply: max supply of Assistant Collection, param from frontend
//   - _mintPrice: mint price of single Assistant NFT, param from frontend
//   - factoryContract: the address of NFT factory contract(verifying Contract), can be imported from config by:
//     ```factoryContract := conf.GConfig.GetString("contract.assistantFactoryContractAddress")```
//
// Returns:
// - signature(string): use for contract verification
// - err(error): the signature is valid when err is nil
func NFTContractCreateSign(name, symbol, baseURI, makerAddress, privateKey string, _collectionId int64, _maxSupply uint, _mintPrice string, factoryContract string) (string, error) {
	maxSupply := math.NewHexOrDecimal256(int64(_maxSupply))
	mintPriceBigInt, ok := math.ParseBig256(_mintPrice)
	if !ok {
		return "", fmt.Errorf("fail to parse mint price to big int")
	}
	mintPrice := math.HexOrDecimal256(*mintPriceBigInt)

	collectionId := math.NewHexOrDecimal256(_collectionId)
	signerData := newNFTContractCreateSignerData(name, symbol, baseURI, makerAddress, collectionId, maxSupply, &mintPrice, factoryContract)
	return EIP712SignatureSign(signerData, privateKey)
}

func NFTContractCreateVerify(name, symbol, baseURI, makerAddress, address, signature string, _collectionId int64, _maxSupply uint, _mintPrice string, factoryContract string) (bool, error) {
	maxSupply := math.NewHexOrDecimal256(int64(_maxSupply))
	mintPriceBigInt, ok := math.ParseBig256(_mintPrice)
	if !ok {
		return false, fmt.Errorf("fail to parse mint price to big int")
	}
	mintPrice := math.HexOrDecimal256(*mintPriceBigInt)

	collectionId := math.NewHexOrDecimal256(_collectionId)
	signerData := newNFTContractCreateSignerData(name, symbol, baseURI, makerAddress, collectionId, maxSupply, &mintPrice, factoryContract)
	return EIP712SignatureVerify(address, signature, signerData)
}

// HarvestSign use for signing FM token harvests
// Params:
//   - _harvested: amount of harvested FM token, calculated by backend
//   - _collectionId: created by backend
//   - _tokenId: tokenId of the assistant NFT, the first token of user's token list
//   - factoryContract: the address of NFT factory contract(verifying Contract), can be imported from config by:
//     ```factoryContract := conf.GConfig.GetString("contract.assistantFactoryContractAddress")```
//   - privateKey: the backend EOA private Key, can be imported from ENV params by:
//     ```privateKey := os.Getenv("PRIVATE_KEY")```
//
// Returns:
// - signature(string): use for contract verification
// - err(error): the signature is valid when err is nil
func HarvestSign(_collectionId int64, _harvested *big.Int, _tokenId uint, factoryContract, privateKey string) (string, error) {
	harvest := math.HexOrDecimal256(*_harvested)
	collectionId := math.NewHexOrDecimal256(_collectionId)
	tokenId := math.NewHexOrDecimal256(int64(_tokenId))
	signerData := newHarvestSignerData(&harvest, collectionId, tokenId, factoryContract)
	return EIP712SignatureSign(signerData, privateKey)
}

func HarvestVerify(_collectionId int64, address, signature, factoryContract string, _harvested, _tokenId uint) (bool, error) {
	harvested := math.NewHexOrDecimal256(int64(_harvested))
	collectionId := math.NewHexOrDecimal256(_collectionId)
	tokenId := math.NewHexOrDecimal256(int64(_tokenId))
	signerData := newHarvestSignerData(harvested, collectionId, tokenId, factoryContract)
	return EIP712SignatureVerify(address, signature, signerData)
}

func EIP712SignatureSign(signerData apitypes.TypedData, privateKey string) (string, error) {
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return "", fmt.Errorf("hash struct message error: %s", err.Error())
	}
	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return "", fmt.Errorf("hash struct EIP712 domain error: %s", err.Error())
	}
	// format them into an EIP-712 compliant byte string.
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md#specification
	rawData := append([]byte("\x19\x01"), append(domainSeparator, typedDataHash...)...)
	challengeHash := crypto.Keccak256Hash(rawData)
	key, err := crypto.HexToECDSA(util.Strip0x(privateKey))
	if err != nil {
		return "", err
	}
	if sig, err := crypto.Sign(challengeHash.Bytes(), key); err != nil {
		return "", err
	} else {
		//https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
		if sig[64] == 0 || sig[64] == 1 {
			sig[64] += 27
		}
		return "0x" + hex.EncodeToString(sig), nil
	}
}

func EIP712SignatureVerify(address, signature string, signerData apitypes.TypedData) (bool, error) {
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return false, fmt.Errorf("hash struct message error: %s", err.Error())
	}
	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return false, fmt.Errorf("hash struct EIP712 domain error: %s", err.Error())
	}
	// format them into an EIP-712 compliant byte string.
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md#specification
	rawData := append([]byte("\x19\x01"), append(domainSeparator, typedDataHash...)...)
	challengeHash := crypto.Keccak256Hash(rawData)
	signatureBytes, ok, err := util.PreCheckSignature(signature)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	pubKey, err := util.EcRecoverForPubKey(signatureBytes, challengeHash.Bytes())
	if err != nil {
		return false, err
	}
	ok, err = util.RecoveryAddressCheck(pubKey, address)
	if err != nil {
		return false, err
	}
	return ok, nil
}
