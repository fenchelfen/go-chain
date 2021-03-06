// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Block block
//
// swagger:model Block
type Block struct {

	// index
	// Required: true
	Index *int64 `json:"index"`

	// sha256 + nonce digest
	PrevBlockHash string `json:"prevBlockHash,omitempty"`

	// sha256 digest
	// Required: true
	ProofOfWork *string `json:"proofOfWork"`

	// transactions
	// Required: true
	Transactions []*Transaction `json:"transactions"`
}

// Validate validates this block
func (m *Block) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProofOfWork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Block) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *Block) validateProofOfWork(formats strfmt.Registry) error {

	if err := validate.Required("proofOfWork", "body", m.ProofOfWork); err != nil {
		return err
	}

	return nil
}

func (m *Block) validateTransactions(formats strfmt.Registry) error {

	if err := validate.Required("transactions", "body", m.Transactions); err != nil {
		return err
	}

	for i := 0; i < len(m.Transactions); i++ {
		if swag.IsZero(m.Transactions[i]) { // not required
			continue
		}

		if m.Transactions[i] != nil {
			if err := m.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Block) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Block) UnmarshalBinary(b []byte) error {
	var res Block
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
