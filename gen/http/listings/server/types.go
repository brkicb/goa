// Code generated by goa v3.16.1, DO NOT EDIT.
//
// listings HTTP server types
//
// Command:
// $ goa gen listings.com/design

package server

import (
	goa "goa.design/goa/v3/pkg"
	listings "listings.com/gen/listings"
)

// AddRequestBody is the type of the "listings" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Unique identifier of the listing
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the listing
	Slug *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
	// Address of the house
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Price of the house
	Price *int64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// UpdateRequestBody is the type of the "listings" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Listing *ListingRequestBody `form:"listing,omitempty" json:"listing,omitempty" xml:"listing,omitempty"`
}

// ListResponseBody is the type of the "listings" service "list" endpoint HTTP
// response body.
type ListResponseBody []*ListingResponse

// GetResponseBody is the type of the "listings" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	// Unique identifier of the listing
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the listing
	Slug string `form:"slug" json:"slug" xml:"slug"`
	// Address of the house
	Address string `form:"address" json:"address" xml:"address"`
	// Price of the house
	Price int64 `form:"price" json:"price" xml:"price"`
}

// ListingResponse is used to define fields on response body types.
type ListingResponse struct {
	// Unique identifier of the listing
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the listing
	Slug string `form:"slug" json:"slug" xml:"slug"`
	// Address of the house
	Address string `form:"address" json:"address" xml:"address"`
	// Price of the house
	Price int64 `form:"price" json:"price" xml:"price"`
}

// ListingRequestBody is used to define fields on request body types.
type ListingRequestBody struct {
	// Unique identifier of the listing
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the listing
	Slug *string `form:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
	// Address of the house
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Price of the house
	Price *int64 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "listings" service.
func NewListResponseBody(res []*listings.Listing) ListResponseBody {
	body := make([]*ListingResponse, len(res))
	for i, val := range res {
		body[i] = marshalListingsListingToListingResponse(val)
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "listings" service.
func NewGetResponseBody(res *listings.Listing) *GetResponseBody {
	body := &GetResponseBody{
		ID:      res.ID,
		Slug:    res.Slug,
		Address: res.Address,
		Price:   res.Price,
	}
	return body
}

// NewAddListing builds a listings service add endpoint payload.
func NewAddListing(body *AddRequestBody) *listings.Listing {
	v := &listings.Listing{
		ID:      body.ID,
		Slug:    *body.Slug,
		Address: *body.Address,
		Price:   *body.Price,
	}

	return v
}

// NewGetPayload builds a listings service get endpoint payload.
func NewGetPayload(slug string) *listings.GetPayload {
	v := &listings.GetPayload{}
	v.Slug = slug

	return v
}

// NewUpdatePayload builds a listings service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, slug string) *listings.UpdatePayload {
	v := &listings.UpdatePayload{}
	v.Listing = unmarshalListingRequestBodyToListingsListing(body.Listing)
	v.Slug = slug

	return v
}

// NewDeletePayload builds a listings service delete endpoint payload.
func NewDeletePayload(slug string) *listings.DeletePayload {
	v := &listings.DeletePayload{}
	v.Slug = slug

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Listing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("listing", "body"))
	}
	if body.Listing != nil {
		if err2 := ValidateListingRequestBody(body.Listing); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateListingRequestBody runs the validations defined on ListingRequestBody
func ValidateListingRequestBody(body *ListingRequestBody) (err error) {
	if body.Slug == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("slug", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	return
}
