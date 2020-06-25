// Code generated by go-swagger; DO NOT EDIT.

package foods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"food_app/food_app/models"
)

// AddFoodCreatedCode is the HTTP code returned for type AddFoodCreated
const AddFoodCreatedCode int = 201

/*AddFoodCreated create a food

swagger:response addFoodCreated
*/
type AddFoodCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Food `json:"body,omitempty"`
}

// NewAddFoodCreated creates AddFoodCreated with default headers values
func NewAddFoodCreated() *AddFoodCreated {

	return &AddFoodCreated{}
}

// WithPayload adds the payload to the add food created response
func (o *AddFoodCreated) WithPayload(payload *models.Food) *AddFoodCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add food created response
func (o *AddFoodCreated) SetPayload(payload *models.Food) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFoodCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
