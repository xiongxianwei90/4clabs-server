// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/nft/v1/typs.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on ComicWork with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ComicWork) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOriginNft()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComicWorkValidationError{
				field:  "OriginNft",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ComicId

	// no validation rules for MintLimit

	// no validation rules for MintPrice

	// no validation rules for Name

	// no validation rules for CreatedAtTimestamp

	// no validation rules for MinterAddress

	return nil
}

// ComicWorkValidationError is the validation error returned by
// ComicWork.Validate if the designated constraints aren't met.
type ComicWorkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComicWorkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComicWorkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComicWorkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComicWorkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComicWorkValidationError) ErrorName() string { return "ComicWorkValidationError" }

// Error satisfies the builtin error interface
func (e ComicWorkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComicWork.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComicWorkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComicWorkValidationError{}

// Validate checks the field values on ComicNft with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ComicNft) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TokenId

	// no validation rules for Owner

	if v, ok := interface{}(m.GetSummary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComicNftValidationError{
				field:  "Summary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ComicNftValidationError is the validation error returned by
// ComicNft.Validate if the designated constraints aren't met.
type ComicNftValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComicNftValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComicNftValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComicNftValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComicNftValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComicNftValidationError) ErrorName() string { return "ComicNftValidationError" }

// Error satisfies the builtin error interface
func (e ComicNftValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComicNft.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComicNftValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComicNftValidationError{}

// Validate checks the field values on Trait with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trait) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for Value

	return nil
}

// TraitValidationError is the validation error returned by Trait.Validate if
// the designated constraints aren't met.
type TraitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraitValidationError) ErrorName() string { return "TraitValidationError" }

// Error satisfies the builtin error interface
func (e TraitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrait.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraitValidationError{}

// Validate checks the field values on Rarity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Rarity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Score

	// no validation rules for Rank

	// no validation rules for Total

	return nil
}

// RarityValidationError is the validation error returned by Rarity.Validate if
// the designated constraints aren't met.
type RarityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RarityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RarityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RarityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RarityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RarityValidationError) ErrorName() string { return "RarityValidationError" }

// Error satisfies the builtin error interface
func (e RarityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRarity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RarityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RarityValidationError{}

// Validate checks the field values on Detail with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Detail) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSummary()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DetailValidationError{
				field:  "Summary",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStat()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DetailValidationError{
				field:  "Stat",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DetailValidationError is the validation error returned by Detail.Validate if
// the designated constraints aren't met.
type DetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailValidationError) ErrorName() string { return "DetailValidationError" }

// Error satisfies the builtin error interface
func (e DetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailValidationError{}

// Validate checks the field values on Summary with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Summary) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Blockchain

	// no validation rules for CollectionName

	// no validation rules for CollectionSlug

	// no validation rules for ContractAddress

	// no validation rules for TokenId

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Image

	// no validation rules for AnimationUrl

	for idx, item := range m.GetTraits() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SummaryValidationError{
					field:  fmt.Sprintf("Traits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetRarity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SummaryValidationError{
				field:  "Rarity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Registered

	return nil
}

// SummaryValidationError is the validation error returned by Summary.Validate
// if the designated constraints aren't met.
type SummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SummaryValidationError) ErrorName() string { return "SummaryValidationError" }

// Error satisfies the builtin error interface
func (e SummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SummaryValidationError{}

// Validate checks the field values on Rank with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Rank) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Score

	// no validation rules for Rank

	// no validation rules for Total

	// no validation rules for LastUpdated

	return nil
}

// RankValidationError is the validation error returned by Rank.Validate if the
// designated constraints aren't met.
type RankValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RankValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RankValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RankValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RankValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RankValidationError) ErrorName() string { return "RankValidationError" }

// Error satisfies the builtin error interface
func (e RankValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRank.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RankValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RankValidationError{}

// Validate checks the field values on Stat with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Stat) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for LastUpdated

	// no validation rules for SaleNum_7D

	// no validation rules for SaleNumAll

	if v, ok := interface{}(m.GetMaxPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatValidationError{
				field:  "MaxPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMinPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatValidationError{
				field:  "MinPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLastPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatValidationError{
				field:  "LastPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPastOwners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StatValidationError{
					field:  fmt.Sprintf("PastOwners[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreateTime

	// no validation rules for StartHoldingTime

	// no validation rules for LongestHoldingTime

	return nil
}

// StatValidationError is the validation error returned by Stat.Validate if the
// designated constraints aren't met.
type StatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatValidationError) ErrorName() string { return "StatValidationError" }

// Error satisfies the builtin error interface
func (e StatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatValidationError{}

// Validate checks the field values on Sort with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Sort) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ByPrice

	return nil
}

// SortValidationError is the validation error returned by Sort.Validate if the
// designated constraints aren't met.
type SortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SortValidationError) ErrorName() string { return "SortValidationError" }

// Error satisfies the builtin error interface
func (e SortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SortValidationError{}

// Validate checks the field values on Stat_PriceInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Stat_PriceInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TxHash

	// no validation rules for PriceToken

	// no validation rules for TokenSymbol

	// no validation rules for TokenContractAddress

	// no validation rules for PriceUsd

	// no validation rules for Time

	return nil
}

// Stat_PriceInfoValidationError is the validation error returned by
// Stat_PriceInfo.Validate if the designated constraints aren't met.
type Stat_PriceInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Stat_PriceInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Stat_PriceInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Stat_PriceInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Stat_PriceInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Stat_PriceInfoValidationError) ErrorName() string { return "Stat_PriceInfoValidationError" }

// Error satisfies the builtin error interface
func (e Stat_PriceInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStat_PriceInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Stat_PriceInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Stat_PriceInfoValidationError{}

// Validate checks the field values on Stat_OwnerStat with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Stat_OwnerStat) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	// no validation rules for HoldingTime

	return nil
}

// Stat_OwnerStatValidationError is the validation error returned by
// Stat_OwnerStat.Validate if the designated constraints aren't met.
type Stat_OwnerStatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Stat_OwnerStatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Stat_OwnerStatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Stat_OwnerStatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Stat_OwnerStatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Stat_OwnerStatValidationError) ErrorName() string { return "Stat_OwnerStatValidationError" }

// Error satisfies the builtin error interface
func (e Stat_OwnerStatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStat_OwnerStat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Stat_OwnerStatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Stat_OwnerStatValidationError{}
