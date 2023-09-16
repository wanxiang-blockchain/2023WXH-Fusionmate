package common

import "errors"

var MsgFlags = map[int]string{
	SUCCESS:                    "success",
	INVALID_PARAM:              "invalid params",
	CANNOT_ACCESS_TO_RESOURCES: "cannot access to this resources",
	BAD_CREDENTIALS:            "bad credentials for user's signature",
	AUTH_TOKEN_ERROR:           "auth token error",
	INTERNAL_ERROR:             "internal server error",
	EXCEED_FREQUENCY_LIMITS:    "api request exceed frequency limits",
	API_NOT_IN_SUPPORT_NOW:     "this api is not in support now",
}

// GetMsg get error information based on Code
func GetMsg(code int) (string, error) {
	msg, ok := MsgFlags[code]
	if ok {
		return msg, nil
	}
	return msg, errors.New("no such code found")
}
