// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/nft/v1/nft.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetNftDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNftDetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNftDetailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNftDetailRequestMultiError, or nil if none found.
func (m *GetNftDetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNftDetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContractAddress

	// no validation rules for TokenId

	if len(errors) > 0 {
		return GetNftDetailRequestMultiError(errors)
	}
	return nil
}

// GetNftDetailRequestMultiError is an error wrapping multiple validation
// errors returned by GetNftDetailRequest.ValidateAll() if the designated
// constraints aren't met.
type GetNftDetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNftDetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNftDetailRequestMultiError) AllErrors() []error { return m }

// GetNftDetailRequestValidationError is the validation error returned by
// GetNftDetailRequest.Validate if the designated constraints aren't met.
type GetNftDetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNftDetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNftDetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNftDetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNftDetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNftDetailRequestValidationError) ErrorName() string {
	return "GetNftDetailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNftDetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNftDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNftDetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNftDetailRequestValidationError{}

// Validate checks the field values on GetNftDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetNftDetailResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNftDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNftDetailResponseMultiError, or nil if none found.
func (m *GetNftDetailResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNftDetailResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNftDetailResponseValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNftDetailResponseValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNftDetailResponseValidationError{
				field:  "Detail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNftDetailResponseMultiError(errors)
	}
	return nil
}

// GetNftDetailResponseMultiError is an error wrapping multiple validation
// errors returned by GetNftDetailResponse.ValidateAll() if the designated
// constraints aren't met.
type GetNftDetailResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNftDetailResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNftDetailResponseMultiError) AllErrors() []error { return m }

// GetNftDetailResponseValidationError is the validation error returned by
// GetNftDetailResponse.Validate if the designated constraints aren't met.
type GetNftDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNftDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNftDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNftDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNftDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNftDetailResponseValidationError) ErrorName() string {
	return "GetNftDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetNftDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNftDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNftDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNftDetailResponseValidationError{}

// Validate checks the field values on GetAddressNftsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAddressNftsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAddressNftsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAddressNftsRequestMultiError, or nil if none found.
func (m *GetAddressNftsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAddressNftsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBaseListRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAddressNftsRequestValidationError{
					field:  "BaseListRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAddressNftsRequestValidationError{
					field:  "BaseListRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBaseListRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAddressNftsRequestValidationError{
				field:  "BaseListRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if len(errors) > 0 {
		return GetAddressNftsRequestMultiError(errors)
	}
	return nil
}

// GetAddressNftsRequestMultiError is an error wrapping multiple validation
// errors returned by GetAddressNftsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAddressNftsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAddressNftsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAddressNftsRequestMultiError) AllErrors() []error { return m }

// GetAddressNftsRequestValidationError is the validation error returned by
// GetAddressNftsRequest.Validate if the designated constraints aren't met.
type GetAddressNftsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAddressNftsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAddressNftsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAddressNftsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAddressNftsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAddressNftsRequestValidationError) ErrorName() string {
	return "GetAddressNftsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAddressNftsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAddressNftsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAddressNftsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAddressNftsRequestValidationError{}

// Validate checks the field values on GetAddressNftResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAddressNftResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAddressNftResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAddressNftResponseMultiError, or nil if none found.
func (m *GetAddressNftResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAddressNftResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBaseListResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAddressNftResponseValidationError{
					field:  "BaseListResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAddressNftResponseValidationError{
					field:  "BaseListResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBaseListResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAddressNftResponseValidationError{
				field:  "BaseListResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSummaries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAddressNftResponseValidationError{
						field:  fmt.Sprintf("Summaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAddressNftResponseValidationError{
						field:  fmt.Sprintf("Summaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAddressNftResponseValidationError{
					field:  fmt.Sprintf("Summaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAddressNftResponseMultiError(errors)
	}
	return nil
}

// GetAddressNftResponseMultiError is an error wrapping multiple validation
// errors returned by GetAddressNftResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAddressNftResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAddressNftResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAddressNftResponseMultiError) AllErrors() []error { return m }

// GetAddressNftResponseValidationError is the validation error returned by
// GetAddressNftResponse.Validate if the designated constraints aren't met.
type GetAddressNftResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAddressNftResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAddressNftResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAddressNftResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAddressNftResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAddressNftResponseValidationError) ErrorName() string {
	return "GetAddressNftResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAddressNftResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAddressNftResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAddressNftResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAddressNftResponseValidationError{}
