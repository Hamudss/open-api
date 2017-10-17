package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSiteDeployParams creates a new GetSiteDeployParams object
// with the default values initialized.
func NewGetSiteDeployParams() *GetSiteDeployParams {
	var ()
	return &GetSiteDeployParams{}
}

/*GetSiteDeployParams contains all the parameters to send to the API endpoint
for the get site deploy operation typically these are written to a http.Request
*/
type GetSiteDeployParams struct {

	/*DeployID*/
	DeployID string
	/*SiteID*/
	SiteID string
}

// WithDeployID adds the deployId to the get site deploy params
func (o *GetSiteDeployParams) WithDeployID(DeployID string) *GetSiteDeployParams {
	o.DeployID = DeployID
	return o
}

// WithSiteID adds the siteId to the get site deploy params
func (o *GetSiteDeployParams) WithSiteID(SiteID string) *GetSiteDeployParams {
	o.SiteID = SiteID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
