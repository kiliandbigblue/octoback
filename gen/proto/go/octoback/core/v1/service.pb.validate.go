// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: octoback/core/v1/service.proto

package corev1

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

// Validate checks the field values on GetGroceryListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetGroceryListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroceryListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroceryListRequestMultiError, or nil if none found.
func (m *GetGroceryListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroceryListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetGroceryListRequestMultiError(errors)
	}

	return nil
}

// GetGroceryListRequestMultiError is an error wrapping multiple validation
// errors returned by GetGroceryListRequest.ValidateAll() if the designated
// constraints aren't met.
type GetGroceryListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroceryListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroceryListRequestMultiError) AllErrors() []error { return m }

// GetGroceryListRequestValidationError is the validation error returned by
// GetGroceryListRequest.Validate if the designated constraints aren't met.
type GetGroceryListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroceryListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroceryListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroceryListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroceryListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroceryListRequestValidationError) ErrorName() string {
	return "GetGroceryListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetGroceryListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroceryListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroceryListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroceryListRequestValidationError{}

// Validate checks the field values on GetGroceryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetGroceryListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetGroceryListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetGroceryListResponseMultiError, or nil if none found.
func (m *GetGroceryListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetGroceryListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetGroceryList()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetGroceryListResponseValidationError{
					field:  "GroceryList",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetGroceryListResponseValidationError{
					field:  "GroceryList",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroceryList()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGroceryListResponseValidationError{
				field:  "GroceryList",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetGroceryListResponseMultiError(errors)
	}

	return nil
}

// GetGroceryListResponseMultiError is an error wrapping multiple validation
// errors returned by GetGroceryListResponse.ValidateAll() if the designated
// constraints aren't met.
type GetGroceryListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetGroceryListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetGroceryListResponseMultiError) AllErrors() []error { return m }

// GetGroceryListResponseValidationError is the validation error returned by
// GetGroceryListResponse.Validate if the designated constraints aren't met.
type GetGroceryListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroceryListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroceryListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroceryListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroceryListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroceryListResponseValidationError) ErrorName() string {
	return "GetGroceryListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetGroceryListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroceryListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroceryListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroceryListResponseValidationError{}
