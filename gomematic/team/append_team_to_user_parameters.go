// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gomematic/gomematic-go/models"
)

// NewAppendTeamToUserParams creates a new AppendTeamToUserParams object
// with the default values initialized.
func NewAppendTeamToUserParams() *AppendTeamToUserParams {
	var ()
	return &AppendTeamToUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppendTeamToUserParamsWithTimeout creates a new AppendTeamToUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppendTeamToUserParamsWithTimeout(timeout time.Duration) *AppendTeamToUserParams {
	var ()
	return &AppendTeamToUserParams{

		timeout: timeout,
	}
}

// NewAppendTeamToUserParamsWithContext creates a new AppendTeamToUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppendTeamToUserParamsWithContext(ctx context.Context) *AppendTeamToUserParams {
	var ()
	return &AppendTeamToUserParams{

		Context: ctx,
	}
}

// NewAppendTeamToUserParamsWithHTTPClient creates a new AppendTeamToUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppendTeamToUserParamsWithHTTPClient(client *http.Client) *AppendTeamToUserParams {
	var ()
	return &AppendTeamToUserParams{
		HTTPClient: client,
	}
}

/*AppendTeamToUserParams contains all the parameters to send to the API endpoint
for the append team to user operation typically these are written to a http.Request
*/
type AppendTeamToUserParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string
	/*TeamUser
	  The team user data to assign

	*/
	TeamUser *models.TeamUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the append team to user params
func (o *AppendTeamToUserParams) WithTimeout(timeout time.Duration) *AppendTeamToUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the append team to user params
func (o *AppendTeamToUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the append team to user params
func (o *AppendTeamToUserParams) WithContext(ctx context.Context) *AppendTeamToUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the append team to user params
func (o *AppendTeamToUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the append team to user params
func (o *AppendTeamToUserParams) WithHTTPClient(client *http.Client) *AppendTeamToUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the append team to user params
func (o *AppendTeamToUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the append team to user params
func (o *AppendTeamToUserParams) WithTeamID(teamID string) *AppendTeamToUserParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the append team to user params
func (o *AppendTeamToUserParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamUser adds the teamUser to the append team to user params
func (o *AppendTeamToUserParams) WithTeamUser(teamUser *models.TeamUserParams) *AppendTeamToUserParams {
	o.SetTeamUser(teamUser)
	return o
}

// SetTeamUser adds the teamUser to the append team to user params
func (o *AppendTeamToUserParams) SetTeamUser(teamUser *models.TeamUserParams) {
	o.TeamUser = teamUser
}

// WriteToRequest writes these params to a swagger request
func (o *AppendTeamToUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if o.TeamUser != nil {
		if err := r.SetBodyParam(o.TeamUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
