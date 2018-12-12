// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// CustomerGroupUpdateAction uses action as discriminator attribute
type CustomerGroupUpdateAction interface{}

func mapDiscriminatorCustomerGroupUpdateAction(input interface{}) CustomerGroupUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeName":
		new := CustomerGroupChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := CustomerGroupSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CustomerGroupSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := CustomerGroupSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// CustomerGroup is of type Resource
type CustomerGroup struct {
	Version        int           `json:"version"`
	LastModifiedAt time.Time     `json:"lastModifiedAt"`
	ID             string        `json:"id"`
	CreatedAt      time.Time     `json:"createdAt"`
	Name           string        `json:"name"`
	Key            string        `json:"key,omitempty"`
	Custom         *CustomFields `json:"custom,omitempty"`
}

// CustomerGroupChangeNameAction implements the interface CustomerGroupUpdateAction
type CustomerGroupChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// CustomerGroupDraft is a standalone struct
type CustomerGroupDraft struct {
	Key       string        `json:"key,omitempty"`
	GroupName string        `json:"groupName"`
	Custom    *CustomFields `json:"custom,omitempty"`
}

// CustomerGroupPagedQueryResponse is of type PagedQueryResponse
type CustomerGroupPagedQueryResponse struct {
	Total   int             `json:"total,omitempty"`
	Offset  int             `json:"offset"`
	Count   int             `json:"count"`
	Results []CustomerGroup `json:"results"`
}

// CustomerGroupReference implements the interface Reference
type CustomerGroupReference struct {
	Key string         `json:"key,omitempty"`
	ID  string         `json:"id,omitempty"`
	Obj *CustomerGroup `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer-group", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetCustomFieldAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetCustomTypeAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetKeyAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CustomerGroupUpdate is of type Update
type CustomerGroupUpdate struct {
	Version int                         `json:"version"`
	Actions []CustomerGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorCustomerGroupUpdateAction(obj.Actions[i])
	}

	return nil
}
