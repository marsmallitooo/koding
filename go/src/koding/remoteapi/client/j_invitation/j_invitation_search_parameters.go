package j_invitation

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

	"koding/remoteapi/models"
)

// NewJInvitationSearchParams creates a new JInvitationSearchParams object
// with the default values initialized.
func NewJInvitationSearchParams() *JInvitationSearchParams {
	var ()
	return &JInvitationSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJInvitationSearchParamsWithTimeout creates a new JInvitationSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJInvitationSearchParamsWithTimeout(timeout time.Duration) *JInvitationSearchParams {
	var ()
	return &JInvitationSearchParams{

		timeout: timeout,
	}
}

// NewJInvitationSearchParamsWithContext creates a new JInvitationSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewJInvitationSearchParamsWithContext(ctx context.Context) *JInvitationSearchParams {
	var ()
	return &JInvitationSearchParams{

		Context: ctx,
	}
}

/*JInvitationSearchParams contains all the parameters to send to the API endpoint
for the j invitation search operation typically these are written to a http.Request
*/
type JInvitationSearchParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j invitation search params
func (o *JInvitationSearchParams) WithTimeout(timeout time.Duration) *JInvitationSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j invitation search params
func (o *JInvitationSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j invitation search params
func (o *JInvitationSearchParams) WithContext(ctx context.Context) *JInvitationSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j invitation search params
func (o *JInvitationSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j invitation search params
func (o *JInvitationSearchParams) WithBody(body models.DefaultSelector) *JInvitationSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j invitation search params
func (o *JInvitationSearchParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JInvitationSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
