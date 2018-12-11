package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// DeleteInput provides the data required to delete a shipping zone.
type ShippingZoneDeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a shipping zone.
type ShippingZoneUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []ZoneUpdateAction
}

// Service contains client information and bundles all actions.
type ShippingZoneService struct {
	client *Client
}

// GetByID will return a shipping zone matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (svc *ShippingZoneService) GetByID(id string) (result *Zone, err error) {
	err = svc.client.Get(fmt.Sprintf("zones/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new shipping zone from a draft, and return the newly
// created shipping zone. OAuth2 Scopes: manage_products:{projectKey}
func (svc *ShippingZoneService) Create(draft *ZoneDraft) (result *Zone, err error) {
	err = svc.client.Create("zones", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a shipping zone matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (svc *ShippingZoneService) Update(input *ShippingZoneUpdateInput) (result *Zone, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("zones/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a shipping zone matching the provided ID. OAuth2
// Scopes: manage_products:{projectKey}
func (svc *ShippingZoneService) DeleteByID(id string, version int) (result *Zone, err error) {
	endpoint := fmt.Sprintf("zones/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
