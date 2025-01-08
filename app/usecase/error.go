package usecase

import (
	"errors"
	"fmt"
	"strings"
)

// ErrResourceNotFound はリソースが見つからなかったことを示す一般エラー
var ErrResourceNotFound = errors.New("usecase: resource not found")

// ErrBearAlreadyUpdated はプッシュ通知更新の空中衝突が発生したことを表すエラー
var ErrBearAlreadyUpdated = errors.New("usecase: Bear already updated")

// InvalidInputError 入力値が不正であることを示すエラー
type InvalidInputError struct {
	cause  error
	fields []string
}

// Error implements error interface
func (e *InvalidInputError) Error() string {
	if len(e.fields) == 0 {
		return fmt.Sprintf("usecase: invalid input, cause: %s", e.cause.Error())
	}

	return fmt.Sprintf("usecase: invalid %s, cause: %s", strings.Join(e.fields, ", "), e.cause.Error())
}

// NewInvalidInputError InvalidInputError を生成する
func NewInvalidInputError(cause error, fields ...string) *InvalidInputError {
	return &InvalidInputError{
		cause:  cause,
		fields: fields,
	}
}
