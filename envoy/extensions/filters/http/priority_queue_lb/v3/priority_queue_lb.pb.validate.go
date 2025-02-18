//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/priority_queue_lb/v3/priority_queue_lb.proto

package priority_queue_lbv3

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

// Validate checks the field values on PriorityQueueLBGradientControllerConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *PriorityQueueLBGradientControllerConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// PriorityQueueLBGradientControllerConfig with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// PriorityQueueLBGradientControllerConfigMultiError, or nil if none found.
func (m *PriorityQueueLBGradientControllerConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PriorityQueueLBGradientControllerConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMaxLatency()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "MaxLatency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "MaxLatency",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxLatency()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "MaxLatency",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHighLatencyTriggerThreshold()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HighLatencyTriggerThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HighLatencyTriggerThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHighLatencyTriggerThreshold()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "HighLatencyTriggerThreshold",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHappinessAlpha()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HappinessAlpha",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HappinessAlpha",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHappinessAlpha()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "HappinessAlpha",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHappinessBeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HappinessBeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "HappinessBeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHappinessBeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "HappinessBeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStepDownPercentage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "StepDownPercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "StepDownPercentage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStepDownPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "StepDownPercentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExperimentModeDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentModeDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentModeDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExperimentModeDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "ExperimentModeDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExperimentModeRecheckDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentModeRecheckDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentModeRecheckDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExperimentModeRecheckDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "ExperimentModeRecheckDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNormalModeDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "NormalModeDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "NormalModeDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNormalModeDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "NormalModeDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExperimentMinSampleSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentMinSampleSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentMinSampleSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExperimentMinSampleSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "ExperimentMinSampleSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExperimentMaxDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentMaxDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBGradientControllerConfigValidationError{
					field:  "ExperimentMaxDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExperimentMaxDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBGradientControllerConfigValidationError{
				field:  "ExperimentMaxDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PriorityQueueLBGradientControllerConfigMultiError(errors)
	}

	return nil
}

// PriorityQueueLBGradientControllerConfigMultiError is an error wrapping
// multiple validation errors returned by
// PriorityQueueLBGradientControllerConfig.ValidateAll() if the designated
// constraints aren't met.
type PriorityQueueLBGradientControllerConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PriorityQueueLBGradientControllerConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PriorityQueueLBGradientControllerConfigMultiError) AllErrors() []error { return m }

// PriorityQueueLBGradientControllerConfigValidationError is the validation
// error returned by PriorityQueueLBGradientControllerConfig.Validate if the
// designated constraints aren't met.
type PriorityQueueLBGradientControllerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriorityQueueLBGradientControllerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriorityQueueLBGradientControllerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriorityQueueLBGradientControllerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriorityQueueLBGradientControllerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriorityQueueLBGradientControllerConfigValidationError) ErrorName() string {
	return "PriorityQueueLBGradientControllerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e PriorityQueueLBGradientControllerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPriorityQueueLBGradientControllerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriorityQueueLBGradientControllerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriorityQueueLBGradientControllerConfigValidationError{}

// Validate checks the field values on PriorityQueueLB with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PriorityQueueLB) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PriorityQueueLB with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PriorityQueueLBMultiError, or nil if none found.
func (m *PriorityQueueLB) ValidateAll() error {
	return m.validate(true)
}

func (m *PriorityQueueLB) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PriorityQueueLBValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PriorityQueueLBValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PriorityQueueLBValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	oneofConcurrencyControllerConfigPresent := false
	switch v := m.ConcurrencyControllerConfig.(type) {
	case *PriorityQueueLB_GradientControllerConfig:
		if v == nil {
			err := PriorityQueueLBValidationError{
				field:  "ConcurrencyControllerConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConcurrencyControllerConfigPresent = true

		if m.GetGradientControllerConfig() == nil {
			err := PriorityQueueLBValidationError{
				field:  "GradientControllerConfig",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetGradientControllerConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PriorityQueueLBValidationError{
						field:  "GradientControllerConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PriorityQueueLBValidationError{
						field:  "GradientControllerConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGradientControllerConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PriorityQueueLBValidationError{
					field:  "GradientControllerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConcurrencyControllerConfigPresent {
		err := PriorityQueueLBValidationError{
			field:  "ConcurrencyControllerConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PriorityQueueLBMultiError(errors)
	}

	return nil
}

// PriorityQueueLBMultiError is an error wrapping multiple validation errors
// returned by PriorityQueueLB.ValidateAll() if the designated constraints
// aren't met.
type PriorityQueueLBMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PriorityQueueLBMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PriorityQueueLBMultiError) AllErrors() []error { return m }

// PriorityQueueLBValidationError is the validation error returned by
// PriorityQueueLB.Validate if the designated constraints aren't met.
type PriorityQueueLBValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriorityQueueLBValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriorityQueueLBValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriorityQueueLBValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriorityQueueLBValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriorityQueueLBValidationError) ErrorName() string { return "PriorityQueueLBValidationError" }

// Error satisfies the builtin error interface
func (e PriorityQueueLBValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPriorityQueueLB.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriorityQueueLBValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriorityQueueLBValidationError{}
