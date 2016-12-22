package models




import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// EventsListResponse List of events.
// swagger:model EventsListResponse
type EventsListResponse struct {

	// The actual list of events in reverse chronological order.
	Events []*Event `json:"events"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#eventsListResponse".
	Kind *string `json:"kind,omitempty"`

	// Token for the next page of events.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// The total number of events for the query. The number of items in a response may be smaller due to paging.
	TotalResults int32 `json:"totalResults,omitempty"`
}

// Validate validates this events list response
func (m *EventsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventsListResponse) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {

		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {

			if err := m.Events[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}