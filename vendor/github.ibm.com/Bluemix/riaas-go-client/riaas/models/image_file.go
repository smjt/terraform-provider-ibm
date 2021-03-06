// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ImageFile ImageFile
//
// The checksum and download location of the image file
// swagger:model imageFile
type ImageFile struct {

	// A SHA-3 checksum of the image file
	Checksum string `json:"checksum,omitempty"`

	// The location of the image file
	Href string `json:"href,omitempty"`

	// The size of the stored image file rounded up to the next gigabyte
	Size int64 `json:"size,omitempty"`
}

// Validate validates this image file
func (m *ImageFile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageFile) UnmarshalBinary(b []byte) error {
	var res ImageFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
