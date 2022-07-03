// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/nft/v1/comic.proto

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

// Validate checks the field values on ComicCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ComicCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ComicCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ComicCreateRequestMultiError, or nil if none found.
func (m *ComicCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ComicCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetOriginNftContractAddress()) != 42 {
		err := ComicCreateRequestValidationError{
			field:  "OriginNftContractAddress",
			reason: "value length must be 42 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetOriginNftTokenId()) < 1 {
		err := ComicCreateRequestValidationError{
			field:  "OriginNftTokenId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMintLimit() < 1 {
		err := ComicCreateRequestValidationError{
			field:  "MintLimit",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMintPrice() < 0 {
		err := ComicCreateRequestValidationError{
			field:  "MintPrice",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ComicCreateRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for MetadataJson

	if utf8.RuneCountInString(m.GetMinterAddress()) != 42 {
		err := ComicCreateRequestValidationError{
			field:  "MinterAddress",
			reason: "value length must be 42 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return ComicCreateRequestMultiError(errors)
	}
	return nil
}

// ComicCreateRequestMultiError is an error wrapping multiple validation errors
// returned by ComicCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type ComicCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ComicCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ComicCreateRequestMultiError) AllErrors() []error { return m }

// ComicCreateRequestValidationError is the validation error returned by
// ComicCreateRequest.Validate if the designated constraints aren't met.
type ComicCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComicCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComicCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComicCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComicCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComicCreateRequestValidationError) ErrorName() string {
	return "ComicCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ComicCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComicCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComicCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComicCreateRequestValidationError{}

// Validate checks the field values on ComicCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ComicCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ComicCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ComicCreateResponseMultiError, or nil if none found.
func (m *ComicCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ComicCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ComicCreateResponseMultiError(errors)
	}
	return nil
}

// ComicCreateResponseMultiError is an error wrapping multiple validation
// errors returned by ComicCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type ComicCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ComicCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ComicCreateResponseMultiError) AllErrors() []error { return m }

// ComicCreateResponseValidationError is the validation error returned by
// ComicCreateResponse.Validate if the designated constraints aren't met.
type ComicCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComicCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComicCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComicCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComicCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComicCreateResponseValidationError) ErrorName() string {
	return "ComicCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ComicCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComicCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComicCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComicCreateResponseValidationError{}

// Validate checks the field values on ListComicWorkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListComicWorkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListComicWorkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListComicWorkRequestMultiError, or nil if none found.
func (m *ListComicWorkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListComicWorkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBaseListRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListComicWorkRequestValidationError{
					field:  "BaseListRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListComicWorkRequestValidationError{
					field:  "BaseListRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBaseListRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListComicWorkRequestValidationError{
				field:  "BaseListRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if len(errors) > 0 {
		return ListComicWorkRequestMultiError(errors)
	}
	return nil
}

// ListComicWorkRequestMultiError is an error wrapping multiple validation
// errors returned by ListComicWorkRequest.ValidateAll() if the designated
// constraints aren't met.
type ListComicWorkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListComicWorkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListComicWorkRequestMultiError) AllErrors() []error { return m }

// ListComicWorkRequestValidationError is the validation error returned by
// ListComicWorkRequest.Validate if the designated constraints aren't met.
type ListComicWorkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListComicWorkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListComicWorkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListComicWorkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListComicWorkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListComicWorkRequestValidationError) ErrorName() string {
	return "ListComicWorkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListComicWorkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListComicWorkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListComicWorkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListComicWorkRequestValidationError{}

// Validate checks the field values on ListComicWorkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListComicWorkResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListComicWorkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListComicWorkResponseMultiError, or nil if none found.
func (m *ListComicWorkResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListComicWorkResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBaseListResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListComicWorkResponseValidationError{
					field:  "BaseListResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListComicWorkResponseValidationError{
					field:  "BaseListResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBaseListResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListComicWorkResponseValidationError{
				field:  "BaseListResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetComicWorks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListComicWorkResponseValidationError{
						field:  fmt.Sprintf("ComicWorks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListComicWorkResponseValidationError{
						field:  fmt.Sprintf("ComicWorks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListComicWorkResponseValidationError{
					field:  fmt.Sprintf("ComicWorks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListComicWorkResponseMultiError(errors)
	}
	return nil
}

// ListComicWorkResponseMultiError is an error wrapping multiple validation
// errors returned by ListComicWorkResponse.ValidateAll() if the designated
// constraints aren't met.
type ListComicWorkResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListComicWorkResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListComicWorkResponseMultiError) AllErrors() []error { return m }

// ListComicWorkResponseValidationError is the validation error returned by
// ListComicWorkResponse.Validate if the designated constraints aren't met.
type ListComicWorkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListComicWorkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListComicWorkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListComicWorkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListComicWorkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListComicWorkResponseValidationError) ErrorName() string {
	return "ListComicWorkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListComicWorkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListComicWorkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListComicWorkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListComicWorkResponseValidationError{}
