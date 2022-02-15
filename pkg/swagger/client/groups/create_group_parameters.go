// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/v2/pkg/swagger/models"
)

// NewCreateGroupParams creates a new CreateGroupParams object
// with the default values initialized.
func NewCreateGroupParams() *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGroupParamsWithTimeout creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGroupParamsWithTimeout(timeout time.Duration) *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		timeout: timeout,
	}
}

// NewCreateGroupParamsWithContext creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGroupParamsWithContext(ctx context.Context) *CreateGroupParams {
	var ()
	return &CreateGroupParams{

		Context: ctx,
	}
}

// NewCreateGroupParamsWithHTTPClient creates a new CreateGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGroupParamsWithHTTPClient(client *http.Client) *CreateGroupParams {
	var ()
	return &CreateGroupParams{
		HTTPClient: client,
	}
}

/*CreateGroupParams contains all the parameters to send to the API endpoint
for the create group operation typically these are written to a http.Request
*/
type CreateGroupParams struct {

	/*Group
	  Configuration options for the new group.

	*/
	Group *models.CreateGroupInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create group params
func (o *CreateGroupParams) WithTimeout(timeout time.Duration) *CreateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create group params
func (o *CreateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create group params
func (o *CreateGroupParams) WithContext(ctx context.Context) *CreateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create group params
func (o *CreateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) WithHTTPClient(client *http.Client) *CreateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create group params
func (o *CreateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the create group params
func (o *CreateGroupParams) WithGroup(group *models.CreateGroupInput) *CreateGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the create group params
func (o *CreateGroupParams) SetGroup(group *models.CreateGroupInput) {
	o.Group = group
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Group != nil {
		if err := r.SetBodyParam(o.Group); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
