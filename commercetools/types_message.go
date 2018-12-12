// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// MessagePayload uses type as discriminator attribute
type MessagePayload interface{}

func mapDiscriminatorMessagePayload(input interface{}) MessagePayload {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "CategoryCreated":
		new := CategoryCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CategorySlugChanged":
		new := CategorySlugChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerCreated":
		new := CustomerCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerGroupSet":
		new := CustomerGroupSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryAdded":
		new := DeliveryAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryRemoved":
		new := DeliveryRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCreated":
		new := OrderCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomerSet":
		new := OrderCustomerSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDeleted":
		new := OrderDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderEditApplied":
		new := OrderEditAppliedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderImported":
		new := OrderImportedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderPaymentStateChanged":
		new := OrderPaymentChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.OldShippingRateInput != nil {
			new.OldShippingRateInput = mapDiscriminatorShippingRateInput(new.OldShippingRateInput)
		}
		if new.ShippingRateInput != nil {
			new.ShippingRateInput = mapDiscriminatorShippingRateInput(new.ShippingRateInput)
		}

		return new
	case "OrderStateChanged":
		new := OrderStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderStateTransition":
		new := OrderStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentCreated":
		new := PaymentCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductCreated":
		new := ProductCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductDeleted":
		new := ProductDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductImageAdded":
		new := ProductImageAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductPublished":
		new := ProductPublishedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductSlugChanged":
		new := ProductSlugChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductStateTransition":
		new := ProductStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductUnpublished":
		new := ProductUnpublishedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReviewCreated":
		new := ReviewCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReviewRatingSet":
		new := ReviewRatingSetMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = mapDiscriminatorReference(new.Target)
		}

		return new
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = mapDiscriminatorReference(new.Target)
		}

		return new
	}
	return nil
}

// CategoryCreatedMessage is of type Message
type CategoryCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Category        *Category `json:"category"`
}

// CategoryCreatedMessagePayload implements the interface MessagePayload
type CategoryCreatedMessagePayload struct {
	Category *Category `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategoryCreated", Alias: (*Alias)(&obj)})
}

// CategorySlugChangedMessage is of type Message
type CategorySlugChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	Slug            *LocalizedString `json:"slug"`
}

// CategorySlugChangedMessagePayload implements the interface MessagePayload
type CategorySlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

// CustomLineItemStateTransitionMessage is of type Message
type CustomLineItemStateTransitionMessage struct {
	Version          int             `json:"version"`
	LastModifiedAt   time.Time       `json:"lastModifiedAt"`
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"createdAt"`
	Type             string          `json:"type"`
	SequenceNumber   int             `json:"sequenceNumber"`
	ResourceVersion  int             `json:"resourceVersion"`
	Resource         Reference       `json:"resource"`
	TransitionDate   time.Time       `json:"transitionDate"`
	ToState          *StateReference `json:"toState"`
	Quantity         int             `json:"quantity"`
	FromState        *StateReference `json:"fromState"`
	CustomLineItemID string          `json:"customLineItemId"`
}

// CustomLineItemStateTransitionMessagePayload implements the interface MessagePayload
type CustomLineItemStateTransitionMessagePayload struct {
	TransitionDate   time.Time       `json:"transitionDate"`
	ToState          *StateReference `json:"toState"`
	Quantity         int             `json:"quantity"`
	FromState        *StateReference `json:"fromState"`
	CustomLineItemID string          `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

// CustomerAddressAddedMessage is of type Message
type CustomerAddressAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

// CustomerAddressAddedMessagePayload implements the interface MessagePayload
type CustomerAddressAddedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

// CustomerAddressChangedMessage is of type Message
type CustomerAddressChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

// CustomerAddressChangedMessagePayload implements the interface MessagePayload
type CustomerAddressChangedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

// CustomerAddressRemovedMessage is of type Message
type CustomerAddressRemovedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

// CustomerAddressRemovedMessagePayload implements the interface MessagePayload
type CustomerAddressRemovedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

// CustomerCompanyNameSetMessage is of type Message
type CustomerCompanyNameSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	CompanyName     string    `json:"companyName"`
}

// CustomerCompanyNameSetMessagePayload implements the interface MessagePayload
type CustomerCompanyNameSetMessagePayload struct {
	CompanyName string `json:"companyName"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCompanyNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

// CustomerCreatedMessage is of type Message
type CustomerCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Customer        *Customer `json:"customer"`
}

// CustomerCreatedMessagePayload implements the interface MessagePayload
type CustomerCreatedMessagePayload struct {
	Customer *Customer `json:"customer"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCreated", Alias: (*Alias)(&obj)})
}

// CustomerDateOfBirthSetMessage is of type Message
type CustomerDateOfBirthSetMessage struct {
	Version         int         `json:"version"`
	LastModifiedAt  time.Time   `json:"lastModifiedAt"`
	ID              string      `json:"id"`
	CreatedAt       time.Time   `json:"createdAt"`
	Type            string      `json:"type"`
	SequenceNumber  int         `json:"sequenceNumber"`
	ResourceVersion int         `json:"resourceVersion"`
	Resource        Reference   `json:"resource"`
	DateOfBirth     interface{} `json:"dateOfBirth"`
}

// CustomerDateOfBirthSetMessagePayload implements the interface MessagePayload
type CustomerDateOfBirthSetMessagePayload struct {
	DateOfBirth interface{} `json:"dateOfBirth"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerDateOfBirthSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

// CustomerEmailChangedMessage is of type Message
type CustomerEmailChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Email           string    `json:"email"`
}

// CustomerEmailChangedMessagePayload implements the interface MessagePayload
type CustomerEmailChangedMessagePayload struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

// CustomerEmailVerifiedMessage is of type Message
type CustomerEmailVerifiedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

// CustomerEmailVerifiedMessagePayload implements the interface MessagePayload
type CustomerEmailVerifiedMessagePayload struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailVerifiedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetMessage is of type Message
type CustomerGroupSetMessage struct {
	Version         int                     `json:"version"`
	LastModifiedAt  time.Time               `json:"lastModifiedAt"`
	ID              string                  `json:"id"`
	CreatedAt       time.Time               `json:"createdAt"`
	Type            string                  `json:"type"`
	SequenceNumber  int                     `json:"sequenceNumber"`
	ResourceVersion int                     `json:"resourceVersion"`
	Resource        Reference               `json:"resource"`
	CustomerGroup   *CustomerGroupReference `json:"customerGroup"`
}

// CustomerGroupSetMessagePayload implements the interface MessagePayload
type CustomerGroupSetMessagePayload struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

// DeliveryAddedMessage is of type Message
type DeliveryAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Delivery        *Delivery `json:"delivery"`
}

// DeliveryAddedMessagePayload implements the interface MessagePayload
type DeliveryAddedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

// DeliveryAddressSetMessage is of type Message
type DeliveryAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	DeliveryID      string    `json:"deliveryId"`
	Address         *Address  `json:"address,omitempty"`
}

// DeliveryAddressSetMessagePayload implements the interface MessagePayload
type DeliveryAddressSetMessagePayload struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

// DeliveryItemsUpdatedMessage is of type Message
type DeliveryItemsUpdatedMessage struct {
	Version         int            `json:"version"`
	LastModifiedAt  time.Time      `json:"lastModifiedAt"`
	ID              string         `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	Type            string         `json:"type"`
	SequenceNumber  int            `json:"sequenceNumber"`
	ResourceVersion int            `json:"resourceVersion"`
	Resource        Reference      `json:"resource"`
	Items           []DeliveryItem `json:"items"`
	DeliveryID      string         `json:"deliveryId"`
}

// DeliveryItemsUpdatedMessagePayload implements the interface MessagePayload
type DeliveryItemsUpdatedMessagePayload struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

// DeliveryRemovedMessage is of type Message
type DeliveryRemovedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Delivery        *Delivery `json:"delivery"`
}

// DeliveryRemovedMessagePayload implements the interface MessagePayload
type DeliveryRemovedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

// InventoryEntryDeletedMessage is of type Message
type InventoryEntryDeletedMessage struct {
	Version         int               `json:"version"`
	LastModifiedAt  time.Time         `json:"lastModifiedAt"`
	ID              string            `json:"id"`
	CreatedAt       time.Time         `json:"createdAt"`
	Type            string            `json:"type"`
	SequenceNumber  int               `json:"sequenceNumber"`
	ResourceVersion int               `json:"resourceVersion"`
	Resource        Reference         `json:"resource"`
	SupplyChannel   *ChannelReference `json:"supplyChannel"`
	SKU             string            `json:"sku"`
}

// InventoryEntryDeletedMessagePayload implements the interface MessagePayload
type InventoryEntryDeletedMessagePayload struct {
	SupplyChannel *ChannelReference `json:"supplyChannel"`
	SKU           string            `json:"sku"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

// LineItemStateTransitionMessage is of type Message
type LineItemStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	TransitionDate  time.Time       `json:"transitionDate"`
	ToState         *StateReference `json:"toState"`
	Quantity        int             `json:"quantity"`
	LineItemID      string          `json:"lineItemId"`
	FromState       *StateReference `json:"fromState"`
}

// LineItemStateTransitionMessagePayload implements the interface MessagePayload
type LineItemStateTransitionMessagePayload struct {
	TransitionDate time.Time       `json:"transitionDate"`
	ToState        *StateReference `json:"toState"`
	Quantity       int             `json:"quantity"`
	LineItemID     string          `json:"lineItemId"`
	FromState      *StateReference `json:"fromState"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

// Message is of type Resource
type Message struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		obj.Resource = mapDiscriminatorReference(obj.Resource)
	}

	return nil
}

// MessageConfiguration is a standalone struct
type MessageConfiguration struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation,omitempty"`
}

// MessageConfigurationDraft is a standalone struct
type MessageConfigurationDraft struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation"`
}

// MessagePagedQueryResponse is of type PagedQueryResponse
type MessagePagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Message `json:"results"`
}

// OrderBillingAddressSetMessage is of type Message
type OrderBillingAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

// OrderBillingAddressSetMessagePayload implements the interface MessagePayload
type OrderBillingAddressSetMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderBillingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderCreatedMessage is of type Message
type OrderCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

// OrderCreatedMessagePayload implements the interface MessagePayload
type OrderCreatedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCreated", Alias: (*Alias)(&obj)})
}

// OrderCustomLineItemDiscountSetMessage is of type Message
type OrderCustomLineItemDiscountSetMessage struct {
	Version                    int                                  `json:"version"`
	LastModifiedAt             time.Time                            `json:"lastModifiedAt"`
	ID                         string                               `json:"id"`
	CreatedAt                  time.Time                            `json:"createdAt"`
	Type                       string                               `json:"type"`
	SequenceNumber             int                                  `json:"sequenceNumber"`
	ResourceVersion            int                                  `json:"resourceVersion"`
	Resource                   Reference                            `json:"resource"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID           string                               `json:"customLineItemId"`
}

// OrderCustomLineItemDiscountSetMessagePayload implements the interface MessagePayload
type OrderCustomLineItemDiscountSetMessagePayload struct {
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID           string                               `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerEmailSetMessage is of type Message
type OrderCustomerEmailSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Email           string    `json:"email"`
}

// OrderCustomerEmailSetMessagePayload implements the interface MessagePayload
type OrderCustomerEmailSetMessagePayload struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerSetMessage is of type Message
type OrderCustomerSetMessage struct {
	Version          int                     `json:"version"`
	LastModifiedAt   time.Time               `json:"lastModifiedAt"`
	ID               string                  `json:"id"`
	CreatedAt        time.Time               `json:"createdAt"`
	Type             string                  `json:"type"`
	SequenceNumber   int                     `json:"sequenceNumber"`
	ResourceVersion  int                     `json:"resourceVersion"`
	Resource         Reference               `json:"resource"`
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	OldCustomer      *CustomerReference      `json:"oldCustomer,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
	Customer         *CustomerReference      `json:"customer,omitempty"`
}

// OrderCustomerSetMessagePayload implements the interface MessagePayload
type OrderCustomerSetMessagePayload struct {
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	OldCustomer      *CustomerReference      `json:"oldCustomer,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
	Customer         *CustomerReference      `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

// OrderDeletedMessage is of type Message
type OrderDeletedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

// OrderDeletedMessagePayload implements the interface MessagePayload
type OrderDeletedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDeleted", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeAddedMessage is of type Message
type OrderDiscountCodeAddedMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

// OrderDiscountCodeAddedMessagePayload implements the interface MessagePayload
type OrderDiscountCodeAddedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeRemovedMessage is of type Message
type OrderDiscountCodeRemovedMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

// OrderDiscountCodeRemovedMessagePayload implements the interface MessagePayload
type OrderDiscountCodeRemovedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeStateSetMessage is of type Message
type OrderDiscountCodeStateSetMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	State           DiscountCodeState      `json:"state"`
	OldState        DiscountCodeState      `json:"oldState,omitempty"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

// OrderDiscountCodeStateSetMessagePayload implements the interface MessagePayload
type OrderDiscountCodeStateSetMessagePayload struct {
	State        DiscountCodeState      `json:"state"`
	OldState     DiscountCodeState      `json:"oldState,omitempty"`
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeStateSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

// OrderEditAppliedMessage is of type Message
type OrderEditAppliedMessage struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	Type            string              `json:"type"`
	SequenceNumber  int                 `json:"sequenceNumber"`
	ResourceVersion int                 `json:"resourceVersion"`
	Resource        Reference           `json:"resource"`
	Result          *OrderEditApplied   `json:"result"`
	Edit            *OrderEditReference `json:"edit"`
}

// OrderEditAppliedMessagePayload implements the interface MessagePayload
type OrderEditAppliedMessagePayload struct {
	Result *OrderEditApplied   `json:"result"`
	Edit   *OrderEditReference `json:"edit"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAppliedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

// OrderImportedMessage is of type Message
type OrderImportedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

// OrderImportedMessagePayload implements the interface MessagePayload
type OrderImportedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderImported", Alias: (*Alias)(&obj)})
}

// OrderLineItemDiscountSetMessage is of type Message
type OrderLineItemDiscountSetMessage struct {
	Version                    int                                  `json:"version"`
	LastModifiedAt             time.Time                            `json:"lastModifiedAt"`
	ID                         string                               `json:"id"`
	CreatedAt                  time.Time                            `json:"createdAt"`
	Type                       string                               `json:"type"`
	SequenceNumber             int                                  `json:"sequenceNumber"`
	ResourceVersion            int                                  `json:"resourceVersion"`
	Resource                   Reference                            `json:"resource"`
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

// OrderLineItemDiscountSetMessagePayload implements the interface MessagePayload
type OrderLineItemDiscountSetMessagePayload struct {
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderPaymentChangedMessage is of type Message
type OrderPaymentChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	PaymentState    string    `json:"paymentState"`
}

// OrderPaymentChangedMessagePayload implements the interface MessagePayload
type OrderPaymentChangedMessagePayload struct {
	PaymentState string `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderReturnInfoAddedMessage is of type Message
type OrderReturnInfoAddedMessage struct {
	Version         int         `json:"version"`
	LastModifiedAt  time.Time   `json:"lastModifiedAt"`
	ID              string      `json:"id"`
	CreatedAt       time.Time   `json:"createdAt"`
	Type            string      `json:"type"`
	SequenceNumber  int         `json:"sequenceNumber"`
	ResourceVersion int         `json:"resourceVersion"`
	Resource        Reference   `json:"resource"`
	ReturnInfo      *ReturnInfo `json:"returnInfo"`
}

// OrderReturnInfoAddedMessagePayload implements the interface MessagePayload
type OrderReturnInfoAddedMessagePayload struct {
	ReturnInfo *ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

// OrderReturnShipmentStateChangedMessage is of type Message
type OrderReturnShipmentStateChangedMessage struct {
	Version             int                 `json:"version"`
	LastModifiedAt      time.Time           `json:"lastModifiedAt"`
	ID                  string              `json:"id"`
	CreatedAt           time.Time           `json:"createdAt"`
	Type                string              `json:"type"`
	SequenceNumber      int                 `json:"sequenceNumber"`
	ResourceVersion     int                 `json:"resourceVersion"`
	Resource            Reference           `json:"resource"`
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
	ReturnItemID        string              `json:"returnItemId"`
}

// OrderReturnShipmentStateChangedMessagePayload implements the interface MessagePayload
type OrderReturnShipmentStateChangedMessagePayload struct {
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
	ReturnItemID        string              `json:"returnItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShipmentStateChangedMessage is of type Message
type OrderShipmentStateChangedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	ShipmentState   ShipmentState `json:"shipmentState"`
}

// OrderShipmentStateChangedMessagePayload implements the interface MessagePayload
type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState ShipmentState `json:"shipmentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShippingAddressSetMessage is of type Message
type OrderShippingAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

// OrderShippingAddressSetMessagePayload implements the interface MessagePayload
type OrderShippingAddressSetMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderShippingInfoSetMessage is of type Message
type OrderShippingInfoSetMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

// OrderShippingInfoSetMessagePayload implements the interface MessagePayload
type OrderShippingInfoSetMessagePayload struct {
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

// OrderShippingRateInputSetMessage is of type Message
type OrderShippingRateInputSetMessage struct {
	Version              int               `json:"version"`
	LastModifiedAt       time.Time         `json:"lastModifiedAt"`
	ID                   string            `json:"id"`
	CreatedAt            time.Time         `json:"createdAt"`
	Type                 string            `json:"type"`
	SequenceNumber       int               `json:"sequenceNumber"`
	ResourceVersion      int               `json:"resourceVersion"`
	Resource             Reference         `json:"resource"`
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		obj.OldShippingRateInput = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
	}

	return nil
}

// OrderShippingRateInputSetMessagePayload implements the interface MessagePayload
type OrderShippingRateInputSetMessagePayload struct {
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingRateInputSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		obj.OldShippingRateInput = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
	}

	return nil
}

// OrderStateChangedMessage is of type Message
type OrderStateChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	OrderState      string    `json:"orderState"`
}

// OrderStateChangedMessagePayload implements the interface MessagePayload
type OrderStateChangedMessagePayload struct {
	OrderState string `json:"orderState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

// OrderStateTransitionMessage is of type Message
type OrderStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

// OrderStateTransitionMessagePayload implements the interface MessagePayload
type OrderStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

// ParcelAddedToDeliveryMessage is of type Message
type ParcelAddedToDeliveryMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Parcel          *Parcel   `json:"parcel"`
	Delivery        *Delivery `json:"delivery"`
}

// ParcelAddedToDeliveryMessagePayload implements the interface MessagePayload
type ParcelAddedToDeliveryMessagePayload struct {
	Parcel   *Parcel   `json:"parcel"`
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelAddedToDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

// ParcelItemsUpdatedMessage is of type Message
type ParcelItemsUpdatedMessage struct {
	Version         int            `json:"version"`
	LastModifiedAt  time.Time      `json:"lastModifiedAt"`
	ID              string         `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	Type            string         `json:"type"`
	SequenceNumber  int            `json:"sequenceNumber"`
	ResourceVersion int            `json:"resourceVersion"`
	Resource        Reference      `json:"resource"`
	ParcelID        string         `json:"parcelId"`
	Items           []DeliveryItem `json:"items"`
	DeliveryID      string         `json:"deliveryId,omitempty"`
}

// ParcelItemsUpdatedMessagePayload implements the interface MessagePayload
type ParcelItemsUpdatedMessagePayload struct {
	ParcelID   string         `json:"parcelId"`
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelMeasurementsUpdatedMessage is of type Message
type ParcelMeasurementsUpdatedMessage struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	Type            string              `json:"type"`
	SequenceNumber  int                 `json:"sequenceNumber"`
	ResourceVersion int                 `json:"resourceVersion"`
	Resource        Reference           `json:"resource"`
	ParcelID        string              `json:"parcelId"`
	Measurements    *ParcelMeasurements `json:"measurements,omitempty"`
	DeliveryID      string              `json:"deliveryId"`
}

// ParcelMeasurementsUpdatedMessagePayload implements the interface MessagePayload
type ParcelMeasurementsUpdatedMessagePayload struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelMeasurementsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelRemovedFromDeliveryMessage is of type Message
type ParcelRemovedFromDeliveryMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Parcel          *Parcel   `json:"parcel"`
	DeliveryID      string    `json:"deliveryId"`
}

// ParcelRemovedFromDeliveryMessagePayload implements the interface MessagePayload
type ParcelRemovedFromDeliveryMessagePayload struct {
	Parcel     *Parcel `json:"parcel"`
	DeliveryID string  `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelRemovedFromDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

// ParcelTrackingDataUpdatedMessage is of type Message
type ParcelTrackingDataUpdatedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	TrackingData    *TrackingData `json:"trackingData,omitempty"`
	ParcelID        string        `json:"parcelId"`
	DeliveryID      string        `json:"deliveryId"`
}

// ParcelTrackingDataUpdatedMessagePayload implements the interface MessagePayload
type ParcelTrackingDataUpdatedMessagePayload struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
	DeliveryID   string        `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelTrackingDataUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

// PaymentCreatedMessage is of type Message
type PaymentCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Payment         *Payment  `json:"payment"`
}

// PaymentCreatedMessagePayload implements the interface MessagePayload
type PaymentCreatedMessagePayload struct {
	Payment *Payment `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentCreated", Alias: (*Alias)(&obj)})
}

// PaymentInteractionAddedMessage is of type Message
type PaymentInteractionAddedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	Interaction     *CustomFields `json:"interaction"`
}

// PaymentInteractionAddedMessagePayload implements the interface MessagePayload
type PaymentInteractionAddedMessagePayload struct {
	Interaction *CustomFields `json:"interaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentInteractionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

// PaymentStatusInterfaceCodeSetMessage is of type Message
type PaymentStatusInterfaceCodeSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	PaymentID       string    `json:"paymentId"`
	InterfaceCode   string    `json:"interfaceCode"`
}

// PaymentStatusInterfaceCodeSetMessagePayload implements the interface MessagePayload
type PaymentStatusInterfaceCodeSetMessagePayload struct {
	PaymentID     string `json:"paymentId"`
	InterfaceCode string `json:"interfaceCode"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusInterfaceCodeSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

// PaymentStatusStateTransitionMessage is of type Message
type PaymentStatusStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

// PaymentStatusStateTransitionMessagePayload implements the interface MessagePayload
type PaymentStatusStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

// PaymentTransactionAddedMessage is of type Message
type PaymentTransactionAddedMessage struct {
	Version         int          `json:"version"`
	LastModifiedAt  time.Time    `json:"lastModifiedAt"`
	ID              string       `json:"id"`
	CreatedAt       time.Time    `json:"createdAt"`
	Type            string       `json:"type"`
	SequenceNumber  int          `json:"sequenceNumber"`
	ResourceVersion int          `json:"resourceVersion"`
	Resource        Reference    `json:"resource"`
	Transaction     *Transaction `json:"transaction"`
}

// PaymentTransactionAddedMessagePayload implements the interface MessagePayload
type PaymentTransactionAddedMessagePayload struct {
	Transaction *Transaction `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

// PaymentTransactionStateChangedMessage is of type Message
type PaymentTransactionStateChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	TransactionID   string           `json:"transactionId"`
	State           TransactionState `json:"state"`
}

// PaymentTransactionStateChangedMessagePayload implements the interface MessagePayload
type PaymentTransactionStateChangedMessagePayload struct {
	TransactionID string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

// ProductCreatedMessage is of type Message
type ProductCreatedMessage struct {
	Version           int                `json:"version"`
	LastModifiedAt    time.Time          `json:"lastModifiedAt"`
	ID                string             `json:"id"`
	CreatedAt         time.Time          `json:"createdAt"`
	Type              string             `json:"type"`
	SequenceNumber    int                `json:"sequenceNumber"`
	ResourceVersion   int                `json:"resourceVersion"`
	Resource          Reference          `json:"resource"`
	ProductProjection *ProductProjection `json:"productProjection"`
}

// ProductCreatedMessagePayload implements the interface MessagePayload
type ProductCreatedMessagePayload struct {
	ProductProjection *ProductProjection `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductCreated", Alias: (*Alias)(&obj)})
}

// ProductDeletedMessage is of type Message
type ProductDeletedMessage struct {
	Version           int                `json:"version"`
	LastModifiedAt    time.Time          `json:"lastModifiedAt"`
	ID                string             `json:"id"`
	CreatedAt         time.Time          `json:"createdAt"`
	Type              string             `json:"type"`
	SequenceNumber    int                `json:"sequenceNumber"`
	ResourceVersion   int                `json:"resourceVersion"`
	Resource          Reference          `json:"resource"`
	RemovedImageUrls  []interface{}      `json:"removedImageUrls"`
	CurrentProjection *ProductProjection `json:"currentProjection"`
}

// ProductDeletedMessagePayload implements the interface MessagePayload
type ProductDeletedMessagePayload struct {
	RemovedImageUrls  []interface{}      `json:"removedImageUrls"`
	CurrentProjection *ProductProjection `json:"currentProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductDeleted", Alias: (*Alias)(&obj)})
}

// ProductImageAddedMessage is of type Message
type ProductImageAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	VariantID       int       `json:"variantId"`
	Staged          bool      `json:"staged"`
	Image           *Image    `json:"image"`
}

// ProductImageAddedMessagePayload implements the interface MessagePayload
type ProductImageAddedMessagePayload struct {
	VariantID int    `json:"variantId"`
	Staged    bool   `json:"staged"`
	Image     *Image `json:"image"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductImageAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

// ProductPublishedMessage is of type Message
type ProductPublishedMessage struct {
	Version           int                 `json:"version"`
	LastModifiedAt    time.Time           `json:"lastModifiedAt"`
	ID                string              `json:"id"`
	CreatedAt         time.Time           `json:"createdAt"`
	Type              string              `json:"type"`
	SequenceNumber    int                 `json:"sequenceNumber"`
	ResourceVersion   int                 `json:"resourceVersion"`
	Resource          Reference           `json:"resource"`
	Scope             ProductPublishScope `json:"scope"`
	RemovedImageUrls  []interface{}       `json:"removedImageUrls"`
	ProductProjection *ProductProjection  `json:"productProjection"`
}

// ProductPublishedMessagePayload implements the interface MessagePayload
type ProductPublishedMessagePayload struct {
	Scope             ProductPublishScope `json:"scope"`
	RemovedImageUrls  []interface{}       `json:"removedImageUrls"`
	ProductProjection *ProductProjection  `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPublished", Alias: (*Alias)(&obj)})
}

// ProductRevertedStagedChangesMessage is of type Message
type ProductRevertedStagedChangesMessage struct {
	Version          int           `json:"version"`
	LastModifiedAt   time.Time     `json:"lastModifiedAt"`
	ID               string        `json:"id"`
	CreatedAt        time.Time     `json:"createdAt"`
	Type             string        `json:"type"`
	SequenceNumber   int           `json:"sequenceNumber"`
	ResourceVersion  int           `json:"resourceVersion"`
	Resource         Reference     `json:"resource"`
	RemovedImageUrls []interface{} `json:"removedImageUrls"`
}

// ProductRevertedStagedChangesMessagePayload implements the interface MessagePayload
type ProductRevertedStagedChangesMessagePayload struct {
	RemovedImageUrls []interface{} `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

// ProductSlugChangedMessage is of type Message
type ProductSlugChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	Slug            *LocalizedString `json:"slug"`
}

// ProductSlugChangedMessagePayload implements the interface MessagePayload
type ProductSlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

// ProductStateTransitionMessage is of type Message
type ProductStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

// ProductStateTransitionMessagePayload implements the interface MessagePayload
type ProductStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

// ProductUnpublishedMessage is of type Message
type ProductUnpublishedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

// ProductUnpublishedMessagePayload implements the interface MessagePayload
type ProductUnpublishedMessagePayload struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

// ProductVariantDeletedMessage is of type Message
type ProductVariantDeletedMessage struct {
	Version          int             `json:"version"`
	LastModifiedAt   time.Time       `json:"lastModifiedAt"`
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"createdAt"`
	Type             string          `json:"type"`
	SequenceNumber   int             `json:"sequenceNumber"`
	ResourceVersion  int             `json:"resourceVersion"`
	Resource         Reference       `json:"resource"`
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []interface{}   `json:"removedImageUrls"`
}

// ProductVariantDeletedMessagePayload implements the interface MessagePayload
type ProductVariantDeletedMessagePayload struct {
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []interface{}   `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

// ReviewCreatedMessage is of type Message
type ReviewCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Review          *Review   `json:"review"`
}

// ReviewCreatedMessagePayload implements the interface MessagePayload
type ReviewCreatedMessagePayload struct {
	Review *Review `json:"review"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewCreated", Alias: (*Alias)(&obj)})
}

// ReviewRatingSetMessage is of type Message
type ReviewRatingSetMessage struct {
	Version              int       `json:"version"`
	LastModifiedAt       time.Time `json:"lastModifiedAt"`
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"createdAt"`
	Type                 string    `json:"type"`
	SequenceNumber       int       `json:"sequenceNumber"`
	ResourceVersion      int       `json:"resourceVersion"`
	Resource             Reference `json:"resource"`
	Target               Reference `json:"target"`
	OldRating            float64   `json:"oldRating"`
	NewRating            float64   `json:"newRating"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorReference(obj.Target)
	}

	return nil
}

// ReviewRatingSetMessagePayload implements the interface MessagePayload
type ReviewRatingSetMessagePayload struct {
	Target               Reference `json:"target"`
	OldRating            float64   `json:"oldRating"`
	NewRating            float64   `json:"newRating"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewRatingSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorReference(obj.Target)
	}

	return nil
}

// ReviewStateTransitionMessage is of type Message
type ReviewStateTransitionMessage struct {
	Version                 int             `json:"version"`
	LastModifiedAt          time.Time       `json:"lastModifiedAt"`
	ID                      string          `json:"id"`
	CreatedAt               time.Time       `json:"createdAt"`
	Type                    string          `json:"type"`
	SequenceNumber          int             `json:"sequenceNumber"`
	ResourceVersion         int             `json:"resourceVersion"`
	Resource                Reference       `json:"resource"`
	Target                  Reference       `json:"target"`
	OldState                *StateReference `json:"oldState"`
	OldIncludedInStatistics bool            `json:"oldIncludedInStatistics"`
	NewState                *StateReference `json:"newState"`
	NewIncludedInStatistics bool            `json:"newIncludedInStatistics"`
	Force                   bool            `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorReference(obj.Target)
	}

	return nil
}

// ReviewStateTransitionMessagePayload implements the interface MessagePayload
type ReviewStateTransitionMessagePayload struct {
	Target                  Reference       `json:"target"`
	OldState                *StateReference `json:"oldState"`
	OldIncludedInStatistics bool            `json:"oldIncludedInStatistics"`
	NewState                *StateReference `json:"newState"`
	NewIncludedInStatistics bool            `json:"newIncludedInStatistics"`
	Force                   bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorReference(obj.Target)
	}

	return nil
}
