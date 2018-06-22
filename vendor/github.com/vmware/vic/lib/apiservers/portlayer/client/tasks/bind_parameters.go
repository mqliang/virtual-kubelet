package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewBindParams creates a new BindParams object
// with the default values initialized.
func NewBindParams() *BindParams {
	var ()
	return &BindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBindParamsWithTimeout creates a new BindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBindParamsWithTimeout(timeout time.Duration) *BindParams {
	var ()
	return &BindParams{

		timeout: timeout,
	}
}

// NewBindParamsWithContext creates a new BindParams object
// with the default values initialized, and the ability to set a context for a request
func NewBindParamsWithContext(ctx context.Context) *BindParams {
	var ()
	return &BindParams{

		Context: ctx,
	}
}

// NewBindParamsWithHTTPClient creates a new BindParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBindParamsWithHTTPClient(client *http.Client) *BindParams {
	var ()
	return &BindParams{
		HTTPClient: client,
	}
}

/*BindParams contains all the parameters to send to the API endpoint
for the bind operation typically these are written to a http.Request
*/
type BindParams struct {

	/*Config*/
	Config *models.TaskBindConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bind params
func (o *BindParams) WithTimeout(timeout time.Duration) *BindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind params
func (o *BindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind params
func (o *BindParams) WithContext(ctx context.Context) *BindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind params
func (o *BindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind params
func (o *BindParams) WithHTTPClient(client *http.Client) *BindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind params
func (o *BindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the bind params
func (o *BindParams) WithConfig(config *models.TaskBindConfig) *BindParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the bind params
func (o *BindParams) SetConfig(config *models.TaskBindConfig) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *BindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Config == nil {
		o.Config = new(models.TaskBindConfig)
	}

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}