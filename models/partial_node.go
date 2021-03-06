// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartialNode Post a node into RackHD
// swagger:model PartialNode

type PartialNode struct {

	// auto discover
	AutoDiscover string `json:"autoDiscover,omitempty"`

	// boot settings
	BootSettings interface{} `json:"bootSettings,omitempty"`

	// identifiers
	Identifiers []string `json:"identifiers"`

	// Name of the node
	Name string `json:"name,omitempty"`

	// obms
	Obms PartialNodeObms `json:"obms"`

	// relations
	Relations PartialNodeRelations `json:"relations"`

	// tags
	Tags []string `json:"tags"`

	// Type of node
	Type string `json:"type,omitempty"`
}

/* polymorph PartialNode autoDiscover false */

/* polymorph PartialNode bootSettings false */

/* polymorph PartialNode identifiers false */

/* polymorph PartialNode name false */

/* polymorph PartialNode obms false */

/* polymorph PartialNode relations false */

/* polymorph PartialNode tags false */

/* polymorph PartialNode type false */

// Validate validates this partial node
func (m *PartialNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifiers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartialNode) validateIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {

	}

	return nil
}

func (m *PartialNode) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

	}

	return nil
}

var partialNodeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["compute","compute-container","switch","dae","pdu","mgmt","enclosure","rack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partialNodeTypeTypePropEnum = append(partialNodeTypeTypePropEnum, v)
	}
}

const (
	// PartialNodeTypeCompute captures enum value "compute"
	PartialNodeTypeCompute string = "compute"
	// PartialNodeTypeComputeContainer captures enum value "compute-container"
	PartialNodeTypeComputeContainer string = "compute-container"
	// PartialNodeTypeSwitch captures enum value "switch"
	PartialNodeTypeSwitch string = "switch"
	// PartialNodeTypeDae captures enum value "dae"
	PartialNodeTypeDae string = "dae"
	// PartialNodeTypePdu captures enum value "pdu"
	PartialNodeTypePdu string = "pdu"
	// PartialNodeTypeMgmt captures enum value "mgmt"
	PartialNodeTypeMgmt string = "mgmt"
	// PartialNodeTypeEnclosure captures enum value "enclosure"
	PartialNodeTypeEnclosure string = "enclosure"
	// PartialNodeTypeRack captures enum value "rack"
	PartialNodeTypeRack string = "rack"
)

// prop value enum
func (m *PartialNode) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partialNodeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartialNode) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartialNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartialNode) UnmarshalBinary(b []byte) error {
	var res PartialNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
