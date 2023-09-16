package util

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var jwtSecret = []byte("Nftfanatic")

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(address string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * time.Hour)
	claims := Claims{
		address,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "Nftfanatic",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func PreCheckSignature(_signature string) ([]byte, bool, error) {
	_signature = Strip0x(_signature)
	signature, err := hex.DecodeString(_signature)
	if err != nil {
		return []byte{}, false, err
	}
	if len(signature) != 65 {
		return []byte{}, false, fmt.Errorf("invalid signature length: %d", len(signature))
	}
	// making sure the signature's recovery ID (the last byte) is set to 27 or 28
	if signature[64] != 27 && signature[64] != 28 {
		return []byte{}, false, fmt.Errorf("invalid recovery id: %d", signature[64])
	}
	// subtract 27 from the recovery ID to convert it to a 0 or 1 , for Ecrecover function.
	signature[64] -= 27
	return signature, true, nil
}

// EcRecoverForPubKey derive a public key from the provided signature
func EcRecoverForPubKey(signature, challengeHash []byte) (*ecdsa.PublicKey, error) {
	pubKeyRaw, err := crypto.Ecrecover(challengeHash, signature)
	if err != nil {
		return &ecdsa.PublicKey{}, fmt.Errorf("invalid signature: %s", err.Error())
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return &ecdsa.PublicKey{}, err
	}
	return pubKey, nil
}

func RecoveryAddressCheck(pubkey *ecdsa.PublicKey, signatureAddress string) (bool, error) {
	ok := common.IsHexAddress(signatureAddress)
	if !ok {
		return false, fmt.Errorf("invalid input hex address: %s", signatureAddress)
	}
	signatureAddress_ := common.HexToAddress(signatureAddress)
	recoveredAddress := crypto.PubkeyToAddress(*pubkey)
	return bytes.Equal(signatureAddress_.Bytes(), recoveredAddress.Bytes()), nil
}

func IsTimeOut(now, checkTime time.Time) bool {
	if now.After(checkTime) {
		return true
	} else {
		return false
	}
}

func Strip0x(s string) string {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	return s
}
