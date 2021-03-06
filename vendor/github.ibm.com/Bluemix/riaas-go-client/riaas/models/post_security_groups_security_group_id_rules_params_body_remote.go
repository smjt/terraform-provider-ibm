// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote SecurityGroupIdentityByName
//
// Uniquely identifies a security group using any one of ID, CRN, or name.
// swagger:model postSecurityGroupsSecurityGroupIdRulesParamsBodyRemote
type PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote struct {

	// A single IPv4 or IPv6 address.
	Address string `json:"address,omitempty"`

	// A range of IPv4 or IPv6 addresses, in CIDR format.
	CidrBlock string `json:"cidr_block,omitempty"`

	// The security group's CRN
	Crn string `json:"crn,omitempty"`

	// The security group's unique identifier.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post security groups security group Id rules params body remote
func (m *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) UnmarshalBinary(b []byte) error {
	var res PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
