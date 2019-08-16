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

// TriplesRequest triples request
// swagger:model TriplesRequest
type TriplesRequest struct {

	// dcids
	Dcids []string `json:"dcids"`

	// limit
	// Maximum: 500
	Limit int64 `json:"limit,omitempty"`
}

// Validate validates this triples request
func (m *TriplesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriplesRequest) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.Limit) { // not required
		return nil
	}

	if err := validate.MaximumInt("limit", "body", int64(m.Limit), 500, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriplesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriplesRequest) UnmarshalBinary(b []byte) error {
	var res TriplesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
