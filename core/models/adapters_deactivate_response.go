package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AdaptersDeactivateResponse adapters deactivate response
// swagger:model AdaptersDeactivateResponse
type AdaptersDeactivateResponse struct {

	// The number of devices that were deleted.
	DeletedCount int32 `json:"deletedCount,omitempty"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#adaptersDeactivateResponse".
	Kind *string `json:"kind,omitempty"`
}

// Validate validates this adapters deactivate response
func (m *AdaptersDeactivateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}