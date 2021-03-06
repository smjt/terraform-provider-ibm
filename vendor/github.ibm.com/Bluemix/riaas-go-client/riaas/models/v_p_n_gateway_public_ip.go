// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VPNGatewayPublicIP v p n gateway public Ip
// swagger:model vPNGatewayPublicIp
type VPNGatewayPublicIP struct {

	// The IP address assigned to this VPN gateway
	Address string `json:"address,omitempty"`
}

// Validate validates this v p n gateway public Ip
func (m *VPNGatewayPublicIP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewayPublicIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewayPublicIP) UnmarshalBinary(b []byte) error {
	var res VPNGatewayPublicIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
