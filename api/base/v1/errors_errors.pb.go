// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsNoAuthHeader(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NO_AUTH_HEADER.String() && e.Code == 401
}

func ErrorNoAuthHeader(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NO_AUTH_HEADER.String(), fmt.Sprintf(format, args...))
}

func IsAuthTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AUTH_TOKEN_EXPIRED.String() && e.Code == 401
}

func ErrorAuthTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_AUTH_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

func IsSignVerifyFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SIGN_VERIFY_FAILED.String() && e.Code == 400
}

func ErrorSignVerifyFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_SIGN_VERIFY_FAILED.String(), fmt.Sprintf(format, args...))
}