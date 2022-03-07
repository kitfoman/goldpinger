// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CheckAllPodResult check all pod result
//
// swagger:model CheckAllPodResult
type CheckAllPodResult struct {

	// host IP
	// Format: ipv4
	HostIP strfmt.IPv4 `json:"HostIP,omitempty"`

	// hostname
	Hostname string `json:"Hostname,omitempty"`

	// o k
	OK *bool `json:"OK,omitempty"`

	// pod IP
	// Format: ipv4
	PodIP strfmt.IPv4 `json:"PodIP,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// response
	Response *CheckResults `json:"response,omitempty"`

	// status code
	StatusCode int32 `json:"status-code,omitempty"`
}

// Validate validates this check all pod result
func (m *CheckAllPodResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAllPodResult) validateHostIP(formats strfmt.Registry) error {
	if swag.IsZero(m.HostIP) { // not required
		return nil
	}

	if err := validate.FormatOf("HostIP", "body", "ipv4", m.HostIP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CheckAllPodResult) validatePodIP(formats strfmt.Registry) error {
	if swag.IsZero(m.PodIP) { // not required
		return nil
	}

	if err := validate.FormatOf("PodIP", "body", "ipv4", m.PodIP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CheckAllPodResult) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check all pod result based on the context it is used
func (m *CheckAllPodResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAllPodResult) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckAllPodResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckAllPodResult) UnmarshalBinary(b []byte) error {
	var res CheckAllPodResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
