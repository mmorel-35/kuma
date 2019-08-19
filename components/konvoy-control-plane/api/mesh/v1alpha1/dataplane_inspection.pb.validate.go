// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mesh/v1alpha1/dataplane_inspection.proto

package v1alpha1

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on DataplaneInspection with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DataplaneInspection) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetDataplane()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DataplaneInspectionValidationError{
					field:  "Dataplane",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetDataplaneInsight()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DataplaneInspectionValidationError{
					field:  "DataplaneInsight",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// DataplaneInspectionValidationError is the validation error returned by
// DataplaneInspection.Validate if the designated constraints aren't met.
type DataplaneInspectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataplaneInspectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataplaneInspectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataplaneInspectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataplaneInspectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataplaneInspectionValidationError) ErrorName() string {
	return "DataplaneInspectionValidationError"
}

// Error satisfies the builtin error interface
func (e DataplaneInspectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataplaneInspection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataplaneInspectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataplaneInspectionValidationError{}
