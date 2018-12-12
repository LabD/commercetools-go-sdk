// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// TaxCategoryUpdateAction uses action as discriminator attribute
type TaxCategoryUpdateAction interface{}

func mapDiscriminatorTaxCategoryUpdateAction(input TaxCategoryUpdateAction) TaxCategoryUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addTaxRate":
		new := TaxCategoryAddTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := TaxCategoryChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeTaxRate":
		new := TaxCategoryRemoveTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "replaceTaxRate":
		new := TaxCategoryReplaceTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := TaxCategorySetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := TaxCategorySetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// SubRate is a standalone struct
type SubRate struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

// TaxCategory is of type Resource
type TaxCategory struct {
	Version        int       `json:"version"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Rates          []TaxRate `json:"rates"`
	Name           string    `json:"name"`
	Key            string    `json:"key,omitempty"`
	Description    string    `json:"description,omitempty"`
}

// TaxCategoryAddTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryAddTaxRateAction struct {
	TaxRate *TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryAddTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryAddTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategoryChangeNameAction implements the interface TaxCategoryUpdateAction
type TaxCategoryChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// TaxCategoryDraft is a standalone struct
type TaxCategoryDraft struct {
	Rates       []TaxRateDraft `json:"rates"`
	Name        string         `json:"name"`
	Key         string         `json:"key,omitempty"`
	Description string         `json:"description,omitempty"`
}

// TaxCategoryPagedQueryResponse is of type PagedQueryResponse
type TaxCategoryPagedQueryResponse struct {
	Total   int           `json:"total,omitempty"`
	Offset  int           `json:"offset"`
	Count   int           `json:"count"`
	Results []TaxCategory `json:"results"`
}

// TaxCategoryReference implements the interface Reference
type TaxCategoryReference struct {
	Key string       `json:"key,omitempty"`
	ID  string       `json:"id,omitempty"`
	Obj *TaxCategory `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "tax-category", Alias: (*Alias)(&obj)})
}

// TaxCategoryRemoveTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryRemoveTaxRateAction struct {
	TaxRateID string `json:"taxRateId"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategoryReplaceTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryReplaceTaxRateAction struct {
	TaxRateID string        `json:"taxRateId"`
	TaxRate   *TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategorySetDescriptionAction implements the interface TaxCategoryUpdateAction
type TaxCategorySetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// TaxCategorySetKeyAction implements the interface TaxCategoryUpdateAction
type TaxCategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// TaxCategoryUpdate is of type Update
type TaxCategoryUpdate struct {
	Version int                       `json:"version"`
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorTaxCategoryUpdateAction(obj.Actions[i])
	}

	return nil
}

// TaxRate is a standalone struct
type TaxRate struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	ID              string      `json:"id,omitempty"`
	Country         CountryCode `json:"country"`
	Amount          float64     `json:"amount"`
}

// TaxRateDraft is a standalone struct
type TaxRateDraft struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	Country         CountryCode `json:"country"`
	Amount          float64     `json:"amount,omitempty"`
}
