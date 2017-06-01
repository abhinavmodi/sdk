package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlacementNetwork placement network
// swagger:model PlacementNetwork
type PlacementNetwork struct {

	//  It is a reference to an object of type Network.
	// Required: true
	NetworkRef *string `json:"network_ref"`

	// Placeholder for description of property subnet of obj type PlacementNetwork field type str  type object
	Subnet *IPAddrPrefix `json:"subnet,omitempty"`
}

// Validate validates this placement network
func (m *PlacementNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkRef(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementNetwork) validateNetworkRef(formats strfmt.Registry) error {

	if err := validate.Required("network_ref", "body", m.NetworkRef); err != nil {
		return err
	}

	return nil
}

func (m *PlacementNetwork) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {

		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}