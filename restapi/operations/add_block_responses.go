// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochain/models"
)

// AddBlockOKCode is the HTTP code returned for type AddBlockOK
const AddBlockOKCode int = 200

/*AddBlockOK success

swagger:response addBlockOK
*/
type AddBlockOK struct {
}

// NewAddBlockOK creates AddBlockOK with default headers values
func NewAddBlockOK() *AddBlockOK {

	return &AddBlockOK{}
}

// WriteResponse to the client
func (o *AddBlockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddBlockBadRequestCode is the HTTP code returned for type AddBlockBadRequest
const AddBlockBadRequestCode int = 400

/*AddBlockBadRequest bad request

swagger:response addBlockBadRequest
*/
type AddBlockBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddBlockBadRequest creates AddBlockBadRequest with default headers values
func NewAddBlockBadRequest() *AddBlockBadRequest {

	return &AddBlockBadRequest{}
}

// WithPayload adds the payload to the add block bad request response
func (o *AddBlockBadRequest) WithPayload(payload *models.Error) *AddBlockBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add block bad request response
func (o *AddBlockBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBlockBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
