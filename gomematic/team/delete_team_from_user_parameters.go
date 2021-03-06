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

// NewDeleteTeamFromUserParams creates a new DeleteTeamFromUserParams object
// with the default values initialized.
func NewDeleteTeamFromUserParams() *DeleteTeamFromUserParams {
	var ()
	return &DeleteTeamFromUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamFromUserParamsWithTimeout creates a new DeleteTeamFromUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTeamFromUserParamsWithTimeout(timeout time.Duration) *DeleteTeamFromUserParams {
	var ()
	return &DeleteTeamFromUserParams{

		timeout: timeout,
	}
}

// NewDeleteTeamFromUserParamsWithContext creates a new DeleteTeamFromUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTeamFromUserParamsWithContext(ctx context.Context) *DeleteTeamFromUserParams {
	var ()
	return &DeleteTeamFromUserParams{

		Context: ctx,
	}
}

// NewDeleteTeamFromUserParamsWithHTTPClient creates a new DeleteTeamFromUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTeamFromUserParamsWithHTTPClient(client *http.Client) *DeleteTeamFromUserParams {
	var ()
	return &DeleteTeamFromUserParams{
		HTTPClient: client,
	}
}

/*DeleteTeamFromUserParams contains all the parameters to send to the API endpoint
for the delete team from user operation typically these are written to a http.Request
*/
type DeleteTeamFromUserParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string
	/*TeamUser
	  The team user data to delete

	*/
	TeamUser *models.TeamUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete team from user params
func (o *DeleteTeamFromUserParams) WithTimeout(timeout time.Duration) *DeleteTeamFromUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team from user params
func (o *DeleteTeamFromUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team from user params
func (o *DeleteTeamFromUserParams) WithContext(ctx context.Context) *DeleteTeamFromUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team from user params
func (o *DeleteTeamFromUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team from user params
func (o *DeleteTeamFromUserParams) WithHTTPClient(client *http.Client) *DeleteTeamFromUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team from user params
func (o *DeleteTeamFromUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the delete team from user params
func (o *DeleteTeamFromUserParams) WithTeamID(teamID string) *DeleteTeamFromUserParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete team from user params
func (o *DeleteTeamFromUserParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamUser adds the teamUser to the delete team from user params
func (o *DeleteTeamFromUserParams) WithTeamUser(teamUser *models.TeamUserParams) *DeleteTeamFromUserParams {
	o.SetTeamUser(teamUser)
	return o
}

// SetTeamUser adds the teamUser to the delete team from user params
func (o *DeleteTeamFromUserParams) SetTeamUser(teamUser *models.TeamUserParams) {
	o.TeamUser = teamUser
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamFromUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
