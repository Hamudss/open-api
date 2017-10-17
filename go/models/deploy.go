package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Deploy deploy

swagger:model deploy
*/
type Deploy struct {

	/* admin url
	 */
	AdminURL string `json:"admin_url,omitempty"`

	/* branch
	 */
	Branch string `json:"branch,omitempty"`

	/* build id
	 */
	BuildID string `json:"build_id,omitempty"`

	/* commit ref
	 */
	CommitRef string `json:"commit_ref,omitempty"`

	/* commit url
	 */
	CommitURL string `json:"commit_url,omitempty"`

	/* created at
	 */
	CreatedAt string `json:"created_at,omitempty"`

	/* deploy ssl url
	 */
	DeploySslURL string `json:"deploy_ssl_url,omitempty"`

	/* deploy url
	 */
	DeployURL string `json:"deploy_url,omitempty"`

	/* draft
	 */
	Draft bool `json:"draft,omitempty"`

	/* error message
	 */
	ErrorMessage string `json:"error_message,omitempty"`

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* required
	 */
	Required []string `json:"required,omitempty"`

	/* required functions
	 */
	RequiredFunctions []string `json:"required_functions,omitempty"`

	/* review id
	 */
	ReviewID float64 `json:"review_id,omitempty"`

	/* screenshot url
	 */
	ScreenshotURL string `json:"screenshot_url,omitempty"`

	/* site id
	 */
	SiteID string `json:"site_id,omitempty"`

	/* skipped
	 */
	Skipped bool `json:"skipped,omitempty"`

	/* state
	 */
	State string `json:"state,omitempty"`

	/* title
	 */
	Title string `json:"title,omitempty"`

	/* updated at
	 */
	UpdatedAt string `json:"updated_at,omitempty"`

	/* url
	 */
	URL string `json:"url,omitempty"`

	/* user id
	 */
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this deploy
func (m *Deploy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequired(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequiredFunctions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deploy) validateRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.Required) { // not required
		return nil
	}

	return nil
}

func (m *Deploy) validateRequiredFunctions(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredFunctions) { // not required
		return nil
	}

	return nil
}
