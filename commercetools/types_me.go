// Automatically generated, do not edit

package commercetools

type MyCartDraft struct {
	TaxMode                         TaxMode                  `json:"taxMode,omitempty"`
	ShippingMethod                  *ShippingMethodReference `json:"shippingMethod,omitempty"`
	ShippingAddress                 *Address                 `json:"shippingAddress,omitempty"`
	Locale                          string                   `json:"locale,omitempty"`
	LineItems                       []MyLineItemDraft        `json:"lineItems,omitempty"`
	ItemShippingAddresses           []Address                `json:"itemShippingAddresses,omitempty"`
	InventoryMode                   InventoryMode            `json:"inventoryMode,omitempty"`
	DeleteDaysAfterLastModification int                      `json:"deleteDaysAfterLastModification,omitempty"`
	CustomerEmail                   string                   `json:"customerEmail,omitempty"`
	Custom                          *CustomFieldsDraft       `json:"custom,omitempty"`
	Currency                        CurrencyCode             `json:"currency"`
	Country                         string                   `json:"country,omitempty"`
	BillingAddress                  *Address                 `json:"billingAddress,omitempty"`
}

type MyCustomerDraft struct {
	VatID                  string        `json:"vatId,omitempty"`
	Title                  string        `json:"title,omitempty"`
	Password               string        `json:"password"`
	MiddleName             string        `json:"middleName,omitempty"`
	Locale                 string        `json:"locale,omitempty"`
	LastName               string        `json:"lastName,omitempty"`
	FirstName              string        `json:"firstName,omitempty"`
	Email                  string        `json:"email"`
	DefaultShippingAddress int           `json:"defaultShippingAddress,omitempty"`
	DefaultBillingAddress  int           `json:"defaultBillingAddress,omitempty"`
	DateOfBirth            interface{}   `json:"dateOfBirth,omitempty"`
	Custom                 *CustomFields `json:"custom,omitempty"`
	CompanyName            string        `json:"companyName,omitempty"`
	Addresses              []Address     `json:"addresses,omitempty"`
}

type MyLineItemDraft struct {
	VariantID           int                       `json:"variantId"`
	SupplyChannel       *ChannelReference         `json:"supplyChannel,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	Quantity            float64                   `json:"quantity"`
	ProductID           string                    `json:"productId"`
	DistributionChannel *ChannelReference         `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft        `json:"custom,omitempty"`
}

type MyOrderFromCartDraft struct {
	Version int    `json:"version"`
	ID      string `json:"id"`
}
