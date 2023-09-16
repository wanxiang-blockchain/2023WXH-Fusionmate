package common

import "errors"

var (
	ErrDuplicatedBotCreation error = errors.New("duplicated, bot have been created by same collection id")
	ErrPoeRejectBotCreation  error = errors.New("bot creation request be rejected by poe server")
	ErrBotLimitReaches       error = errors.New("bot limit reaches")
)
