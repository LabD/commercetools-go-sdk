// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"
)

// StoreUpdateAction uses action as discriminator attribute
type StoreUpdateAction interface{}

func mapDiscriminatorStoreUpdateAction(input interface{}) (StoreUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "setLanguages":
		new := StoreSetLanguagesAction{}
		err := DecodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := StoreSetNameAction{}
		err := DecodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addDistributionChannel":
		new := StoresAddDistributionChannelsAction{}
		err := DecodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDistributionChannel":
		new := StoresRemoveDistributionChannelsAction{}
		err := DecodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDistributionChannels":
		new := StoresSetDistributionChannelsAction{}
		err := DecodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Store is of type BaseResource
type Store struct {
	Version              int                `json:"version"`
	Name                 *LocalizedString   `json:"name,omitempty"`
	LastModifiedBy       *LastModifiedBy    `json:"lastModifiedBy,omitempty"`
	LastModifiedAt       time.Time          `json:"lastModifiedAt"`
	Languages            []string           `json:"languages,omitempty"`
	Key                  string             `json:"key"`
	ID                   string             `json:"id"`
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	CreatedBy            *CreatedBy         `json:"createdBy,omitempty"`
	CreatedAt            time.Time          `json:"createdAt"`
}

// StoreDraft is a standalone struct
type StoreDraft struct {
	Name                 *LocalizedString            `json:"name"`
	Languages            []string                    `json:"languages,omitempty"`
	Key                  string                      `json:"key"`
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
}

// StoreKeyReference implements the interface KeyReference
type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StorePagedQueryResponse is a standalone struct
type StorePagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Results []Store `json:"results"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
}

// StoreReference implements the interface Reference
type StoreReference struct {
	ID  string `json:"id"`
	Obj *Store `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreReference) MarshalJSON() ([]byte, error) {
	type Alias StoreReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StoreResourceIdentifier implements the interface ResourceIdentifier
type StoreResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StoreResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "store", Alias: (*Alias)(&obj)})
}

// StoreSetLanguagesAction implements the interface StoreUpdateAction
type StoreSetLanguagesAction struct {
	Languages []string `json:"languages,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLanguages", Alias: (*Alias)(&obj)})
}

// StoreSetNameAction implements the interface StoreUpdateAction
type StoreSetNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

// StoreUpdate is a standalone struct
type StoreUpdate struct {
	Version int                 `json:"version"`
	Actions []StoreUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreUpdate) UnmarshalJSON(data []byte) error {
	type Alias StoreUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStoreUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// StoresAddDistributionChannelsAction implements the interface StoreUpdateAction
type StoresAddDistributionChannelsAction struct {
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoresAddDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoresAddDistributionChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDistributionChannel", Alias: (*Alias)(&obj)})
}

// StoresRemoveDistributionChannelsAction implements the interface StoreUpdateAction
type StoresRemoveDistributionChannelsAction struct {
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoresRemoveDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoresRemoveDistributionChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDistributionChannel", Alias: (*Alias)(&obj)})
}

// StoresSetDistributionChannelsAction implements the interface StoreUpdateAction
type StoresSetDistributionChannelsAction struct {
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoresSetDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoresSetDistributionChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDistributionChannels", Alias: (*Alias)(&obj)})
}
