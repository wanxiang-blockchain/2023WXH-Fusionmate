package service

import (
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/util"
	"time"
)

type loginAddress struct {
	Token     string
	Address   string
	AuthUntil time.Time
}

func Web3Login(address, message, timestamp, signature string) (string, int, error) {
	ok, err := LoginVerify(address, message, timestamp, signature)
	if err != nil {
		return "", common.BAD_CREDENTIALS, err
	}
	if !ok {
		return "", common.BAD_CREDENTIALS, nil
	}
	token, err := GenerateTokenForAddress(address)
	if err != nil {
		return "", common.INTERNAL_ERROR, err
	}
	return token, common.SUCCESS, nil
}

func GenerateTokenForAddress(address string) (string, error) {
	token, err := util.GenerateToken(address)
	if err != nil {
		return "", err
	}
	l := loginAddress{
		Token:     token,
		Address:   address,
		AuthUntil: time.Now().Add(12 * time.Hour),
	}
	loginAddresses[token] = l
	return token, nil
}

func TokenCheck(token string) (string, bool) {
	l, ok := loginAddresses[token]
	if !ok {
		return "", false
	}
	_, err := util.ParseToken(token)
	if err != nil {
		return "", false
	}
	if isTimeOut := util.IsTimeOut(time.Now(), l.AuthUntil); isTimeOut {
		return "", false
	}
	setExpireTimeByKey(token)
	return l.Address, true
}

func setExpireTimeByKey(token string) {
	l := loginAddresses[token]
	l.AuthUntil = time.Now().Add(12 * time.Hour)
}
