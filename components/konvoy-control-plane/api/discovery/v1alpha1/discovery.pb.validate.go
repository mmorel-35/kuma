// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: discovery/v1alpha1/discovery.proto

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

// Validate checks the field values on Inventory with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Inventory) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return InventoryValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// InventoryValidationError is the validation error returned by
// Inventory.Validate if the designated constraints aren't met.
type InventoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InventoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InventoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InventoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InventoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InventoryValidationError) ErrorName() string { return "InventoryValidationError" }

// Error satisfies the builtin error interface
func (e InventoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInventory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InventoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InventoryValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetId()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ServiceValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ServiceValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Endpoint with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Endpoint) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetWorkload()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return EndpointValidationError{
					field:  "Workload",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Address

	// no validation rules for Port

	return nil
}

// EndpointValidationError is the validation error returned by
// Endpoint.Validate if the designated constraints aren't met.
type EndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointValidationError) ErrorName() string { return "EndpointValidationError" }

// Error satisfies the builtin error interface
func (e EndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointValidationError{}

// Validate checks the field values on Workload with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Workload) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetId()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WorkloadValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMeta()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WorkloadValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLocality()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WorkloadValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// WorkloadValidationError is the validation error returned by
// Workload.Validate if the designated constraints aren't met.
type WorkloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkloadValidationError) ErrorName() string { return "WorkloadValidationError" }

// Error satisfies the builtin error interface
func (e WorkloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkloadValidationError{}

// Validate checks the field values on Id with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Id) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Namespace

	// no validation rules for Name

	return nil
}

// IdValidationError is the validation error returned by Id.Validate if the
// designated constraints aren't met.
type IdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdValidationError) ErrorName() string { return "IdValidationError" }

// Error satisfies the builtin error interface
func (e IdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdValidationError{}

// Validate checks the field values on Meta with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Meta) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Labels

	return nil
}

// MetaValidationError is the validation error returned by Meta.Validate if the
// designated constraints aren't met.
type MetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaValidationError) ErrorName() string { return "MetaValidationError" }

// Error satisfies the builtin error interface
func (e MetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaValidationError{}

// Validate checks the field values on Locality with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Locality) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Region

	// no validation rules for Zone

	return nil
}

// LocalityValidationError is the validation error returned by
// Locality.Validate if the designated constraints aren't met.
type LocalityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityValidationError) ErrorName() string { return "LocalityValidationError" }

// Error satisfies the builtin error interface
func (e LocalityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocality.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityValidationError{}

// Validate checks the field values on Inventory_Item with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Inventory_Item) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ItemType.(type) {

	case *Inventory_Item_Service:

		{
			tmp := m.GetService()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Inventory_ItemValidationError{
						field:  "Service",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *Inventory_Item_Workload:

		{
			tmp := m.GetWorkload()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Inventory_ItemValidationError{
						field:  "Workload",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// Inventory_ItemValidationError is the validation error returned by
// Inventory_Item.Validate if the designated constraints aren't met.
type Inventory_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Inventory_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Inventory_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Inventory_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Inventory_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Inventory_ItemValidationError) ErrorName() string { return "Inventory_ItemValidationError" }

// Error satisfies the builtin error interface
func (e Inventory_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInventory_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Inventory_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Inventory_ItemValidationError{}
