// Code generated by go-swagger; DO NOT EDIT.

package peers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochain/models"
)

// GetChainOKCode is the HTTP code returned for type GetChainOK
const GetChainOKCode int = 200

/*GetChainOK success

swagger:response getChainOK
*/
type GetChainOK struct {

	/*
	  In: Body
	*/
	Payload *models.Chain `json:"body,omitempty"`
}

// NewGetChainOK creates GetChainOK with default headers values
func NewGetChainOK() *GetChainOK {

	return &GetChainOK{}
}

// WithPayload adds the payload to the get chain o k response
func (o *GetChainOK) WithPayload(payload *models.Chain) *GetChainOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chain o k response
func (o *GetChainOK) SetPayload(payload *models.Chain) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChainOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
