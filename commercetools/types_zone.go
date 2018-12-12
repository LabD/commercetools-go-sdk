// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ZoneUpdateAction uses action as discriminator attribute
type ZoneUpdateAction interface{}

func mapDiscriminatorZoneUpdateAction(input ZoneUpdateAction) ZoneUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addLocation":
		new := ZoneAddLocationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ZoneChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeLocation":
		new := ZoneRemoveLocationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := ZoneSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ZoneSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Location is a standalone struct
type Location struct {
	State   string      `json:"state,omitempty"`
	Country CountryCode `json:"country"`
}

// Zone is of type Resource
type Zone struct {
	Version        int        `json:"version"`
	LastModifiedAt time.Time  `json:"lastModifiedAt"`
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	Name           string     `json:"name"`
	Locations      []Location `json:"locations"`
	Key            string     `json:"key,omitempty"`
	Description    string     `json:"description,omitempty"`
}

// ZoneAddLocationAction implements the interface ZoneUpdateAction
type ZoneAddLocationAction struct {
	Location *Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneAddLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneAddLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocation", Alias: (*Alias)(&obj)})
}

// ZoneChangeNameAction implements the interface ZoneUpdateAction
type ZoneChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ZoneDraft is a standalone struct
type ZoneDraft struct {
	Name        string     `json:"name"`
	Locations   []Location `json:"locations"`
	Key         string     `json:"key,omitempty"`
	Description string     `json:"description,omitempty"`
}

// ZonePagedQueryResponse is of type PagedQueryResponse
type ZonePagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Results []Zone `json:"results"`
}

// ZoneReference implements the interface Reference
type ZoneReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Zone  `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneReference) MarshalJSON() ([]byte, error) {
	type Alias ZoneReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "zone", Alias: (*Alias)(&obj)})
}

// ZoneRemoveLocationAction implements the interface ZoneUpdateAction
type ZoneRemoveLocationAction struct {
	Location *Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneRemoveLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneRemoveLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLocation", Alias: (*Alias)(&obj)})
}

// ZoneSetDescriptionAction implements the interface ZoneUpdateAction
type ZoneSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ZoneSetKeyAction implements the interface ZoneUpdateAction
type ZoneSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ZoneUpdate is of type Update
type ZoneUpdate struct {
	Version int                `json:"version"`
	Actions []ZoneUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ZoneUpdate) UnmarshalJSON(data []byte) error {
	type Alias ZoneUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorZoneUpdateAction(obj.Actions[i])
	}

	return nil
}
