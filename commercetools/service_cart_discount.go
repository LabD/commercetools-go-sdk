// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// CartDiscountURLPath is the commercetools API path.
const CartDiscountURLPath = "cart-discounts"

// CartDiscountCreate creates a new instance of type CartDiscount
func (client *Client) CartDiscountCreate(ctx context.Context, draft *CartDiscountDraft) (result *CartDiscount, err error) {
	err = client.Create(ctx, CartDiscountURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountQuery allows querying for type CartDiscount
func (client *Client) CartDiscountQuery(ctx context.Context, input *QueryInput) (result *CartDiscountPagedQueryResponse, err error) {
	err = client.Query(ctx, CartDiscountURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountDeleteWithKey for type CartDiscount
func (client *Client) CartDiscountDeleteWithKey(ctx context.Context, key string, version int) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("cart-discounts/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountGetWithKey for type CartDiscount
func (client *Client) CartDiscountGetWithKey(ctx context.Context, key string) (result *CartDiscount, err error) {
	err = client.Get(ctx, strings.Replace("cart-discounts/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountUpdateWithKeyInput is input for function CartDiscountUpdateWithKey
type CartDiscountUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CartDiscountUpdateAction
}

// CartDiscountUpdateWithKey for type CartDiscount
func (client *Client) CartDiscountUpdateWithKey(ctx context.Context, input *CartDiscountUpdateWithKeyInput) (result *CartDiscount, err error) {
	err = client.Update(ctx, strings.Replace("cart-discounts/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountDeleteWithID for type CartDiscount
func (client *Client) CartDiscountDeleteWithID(ctx context.Context, ID string, version int) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("cart-discounts/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountGetWithID for type CartDiscount
func (client *Client) CartDiscountGetWithID(ctx context.Context, ID string) (result *CartDiscount, err error) {
	err = client.Get(ctx, strings.Replace("cart-discounts/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountUpdateWithIDInput is input for function CartDiscountUpdateWithID
type CartDiscountUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartDiscountUpdateAction
}

// CartDiscountUpdateWithID for type CartDiscount
func (client *Client) CartDiscountUpdateWithID(ctx context.Context, input *CartDiscountUpdateWithIDInput) (result *CartDiscount, err error) {
	err = client.Update(ctx, strings.Replace("cart-discounts/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
