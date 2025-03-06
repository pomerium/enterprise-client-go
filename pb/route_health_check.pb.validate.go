// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: route_health_check.proto

package pb

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

// Validate checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthCheckMultiError, or
// nil if none found.
func (m *HealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTimeout() == nil {
		err := HealthCheckValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = HealthCheckValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := HealthCheckValidationError{
					field:  "Timeout",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if m.GetInterval() == nil {
		err := HealthCheckValidationError{
			field:  "Interval",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = HealthCheckValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := HealthCheckValidationError{
					field:  "Interval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if all {
		switch v := interface{}(m.GetInitialJitter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "InitialJitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "InitialJitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInitialJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "InitialJitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIntervalJitter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "IntervalJitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "IntervalJitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIntervalJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "IntervalJitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IntervalJitterPercent

	if m.GetUnhealthyThreshold() < 1 {
		err := HealthCheckValidationError{
			field:  "UnhealthyThreshold",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetHealthyThreshold() < 1 {
		err := HealthCheckValidationError{
			field:  "HealthyThreshold",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofHealthCheckerPresent := false
	switch v := m.HealthChecker.(type) {
	case *HealthCheck_HttpHealthCheck_:
		if v == nil {
			err := HealthCheckValidationError{
				field:  "HealthChecker",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofHealthCheckerPresent = true

		if all {
			switch v := interface{}(m.GetHttpHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "HttpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "HttpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHttpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "HttpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_TcpHealthCheck_:
		if v == nil {
			err := HealthCheckValidationError{
				field:  "HealthChecker",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofHealthCheckerPresent = true

		if all {
			switch v := interface{}(m.GetTcpHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "TcpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "TcpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTcpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "TcpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_GrpcHealthCheck_:
		if v == nil {
			err := HealthCheckValidationError{
				field:  "HealthChecker",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofHealthCheckerPresent = true

		if all {
			switch v := interface{}(m.GetGrpcHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "GrpcHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "GrpcHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGrpcHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "GrpcHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofHealthCheckerPresent {
		err := HealthCheckValidationError{
			field:  "HealthChecker",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HealthCheckMultiError(errors)
	}

	return nil
}

// HealthCheckMultiError is an error wrapping multiple validation errors
// returned by HealthCheck.ValidateAll() if the designated constraints aren't met.
type HealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckMultiError) AllErrors() []error { return m }

// HealthCheckValidationError is the validation error returned by
// HealthCheck.Validate if the designated constraints aren't met.
type HealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckValidationError) ErrorName() string { return "HealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckValidationError{}

// Validate checks the field values on HealthCheck_Payload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheck_Payload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck_Payload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheck_PayloadMultiError, or nil if none found.
func (m *HealthCheck_Payload) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck_Payload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofPayloadPresent := false
	switch v := m.Payload.(type) {
	case *HealthCheck_Payload_Text:
		if v == nil {
			err := HealthCheck_PayloadValidationError{
				field:  "Payload",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPayloadPresent = true

		if utf8.RuneCountInString(m.GetText()) < 1 {
			err := HealthCheck_PayloadValidationError{
				field:  "Text",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *HealthCheck_Payload_Binary:
		if v == nil {
			err := HealthCheck_PayloadValidationError{
				field:  "Payload",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPayloadPresent = true
		// no validation rules for Binary
	default:
		_ = v // ensures v is used
	}
	if !oneofPayloadPresent {
		err := HealthCheck_PayloadValidationError{
			field:  "Payload",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HealthCheck_PayloadMultiError(errors)
	}

	return nil
}

// HealthCheck_PayloadMultiError is an error wrapping multiple validation
// errors returned by HealthCheck_Payload.ValidateAll() if the designated
// constraints aren't met.
type HealthCheck_PayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheck_PayloadMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheck_PayloadMultiError) AllErrors() []error { return m }

// HealthCheck_PayloadValidationError is the validation error returned by
// HealthCheck_Payload.Validate if the designated constraints aren't met.
type HealthCheck_PayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_PayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_PayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_PayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_PayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_PayloadValidationError) ErrorName() string {
	return "HealthCheck_PayloadValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_PayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_Payload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_PayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_PayloadValidationError{}

// Validate checks the field values on HealthCheck_HttpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheck_HttpHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck_HttpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheck_HttpHealthCheckMultiError, or nil if none found.
func (m *HealthCheck_HttpHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck_HttpHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_HealthCheck_HttpHealthCheck_Host_Pattern.MatchString(m.GetHost()) {
		err := HealthCheck_HttpHealthCheckValidationError{
			field:  "Host",
			reason: "value does not match regex pattern \"^[^\\x00-\\b\\n-\\x1f\\x7f]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := HealthCheck_HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_HealthCheck_HttpHealthCheck_Path_Pattern.MatchString(m.GetPath()) {
		err := HealthCheck_HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value does not match regex pattern \"^[^\\x00-\\b\\n-\\x1f\\x7f]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetExpectedStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("ExpectedStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("ExpectedStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  fmt.Sprintf("ExpectedStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRetriableStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("RetriableStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("RetriableStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  fmt.Sprintf("RetriableStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if _, ok := CodecClientType_name[int32(m.GetCodecClientType())]; !ok {
		err := HealthCheck_HttpHealthCheckValidationError{
			field:  "CodecClientType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HealthCheck_HttpHealthCheckMultiError(errors)
	}

	return nil
}

// HealthCheck_HttpHealthCheckMultiError is an error wrapping multiple
// validation errors returned by HealthCheck_HttpHealthCheck.ValidateAll() if
// the designated constraints aren't met.
type HealthCheck_HttpHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheck_HttpHealthCheckMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheck_HttpHealthCheckMultiError) AllErrors() []error { return m }

// HealthCheck_HttpHealthCheckValidationError is the validation error returned
// by HealthCheck_HttpHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_HttpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_HttpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_HttpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_HttpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_HttpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_HttpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_HttpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_HttpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_HttpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_HttpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_HttpHealthCheckValidationError{}

var _HealthCheck_HttpHealthCheck_Host_Pattern = regexp.MustCompile("^[^\x00-\b\n-\x1f\x7f]*$")

var _HealthCheck_HttpHealthCheck_Path_Pattern = regexp.MustCompile("^[^\x00-\b\n-\x1f\x7f]*$")

// Validate checks the field values on HealthCheck_TcpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheck_TcpHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck_TcpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheck_TcpHealthCheckMultiError, or nil if none found.
func (m *HealthCheck_TcpHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck_TcpHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSend()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheck_TcpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheck_TcpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheck_TcpHealthCheckValidationError{
				field:  "Send",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetReceive() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheck_TcpHealthCheckValidationError{
						field:  fmt.Sprintf("Receive[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheck_TcpHealthCheckValidationError{
						field:  fmt.Sprintf("Receive[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheck_TcpHealthCheckValidationError{
					field:  fmt.Sprintf("Receive[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return HealthCheck_TcpHealthCheckMultiError(errors)
	}

	return nil
}

// HealthCheck_TcpHealthCheckMultiError is an error wrapping multiple
// validation errors returned by HealthCheck_TcpHealthCheck.ValidateAll() if
// the designated constraints aren't met.
type HealthCheck_TcpHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheck_TcpHealthCheckMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheck_TcpHealthCheckMultiError) AllErrors() []error { return m }

// HealthCheck_TcpHealthCheckValidationError is the validation error returned
// by HealthCheck_TcpHealthCheck.Validate if the designated constraints aren't met.
type HealthCheck_TcpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_TcpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_TcpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_TcpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_TcpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_TcpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_TcpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_TcpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_TcpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_TcpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_TcpHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_GrpcHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheck_GrpcHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck_GrpcHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheck_GrpcHealthCheckMultiError, or nil if none found.
func (m *HealthCheck_GrpcHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck_GrpcHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceName

	if !_HealthCheck_GrpcHealthCheck_Authority_Pattern.MatchString(m.GetAuthority()) {
		err := HealthCheck_GrpcHealthCheckValidationError{
			field:  "Authority",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HealthCheck_GrpcHealthCheckMultiError(errors)
	}

	return nil
}

// HealthCheck_GrpcHealthCheckMultiError is an error wrapping multiple
// validation errors returned by HealthCheck_GrpcHealthCheck.ValidateAll() if
// the designated constraints aren't met.
type HealthCheck_GrpcHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheck_GrpcHealthCheckMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheck_GrpcHealthCheckMultiError) AllErrors() []error { return m }

// HealthCheck_GrpcHealthCheckValidationError is the validation error returned
// by HealthCheck_GrpcHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_GrpcHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_GrpcHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_GrpcHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_GrpcHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_GrpcHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_GrpcHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_GrpcHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_GrpcHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_GrpcHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_GrpcHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_GrpcHealthCheckValidationError{}

var _HealthCheck_GrpcHealthCheck_Authority_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")
