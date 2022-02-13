package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound = errors.New("your requested item is not found")
	ErrConflict = errors.New("your item already exist")
	ErrBadParamInput = errors.New("given param is not valid")
	ErrUnauthorized = errors.New("your request need authentication")
	ErrBalanceNotEnough = errors.New("your balance is not enough, top up first")
	ErrStockNotEnough = errors.New("sorry our stock is not enough")
)