package listingsapi

import (
	"context"
	"log"

	"listings.com/db"
	listings "listings.com/gen/listings"
	"listings.com/models"
)

// listings service example implementation.
// The example methods log the requests and return zero values.
type listingssrvc struct {
	logger *log.Logger
}

// NewListings returns the listings service implementation.
func NewListings(logger *log.Logger) listings.Service {
	db.InitDB()
	return &listingssrvc{logger}
}

func (s *listingssrvc) List(ctx context.Context) (res []*listings.Listing, err error) {
    dbListings, err := models.GetAllListings()
    if err != nil {
        s.logger.Printf("error fetching listings: %v", err)
        return nil, err
    }
    for _, dbListing := range dbListings {
        res = append(res, &listings.Listing{
            ID:      &dbListing.ID,
            Slug:    dbListing.Slug,
            Address: dbListing.Address,
            Price:   dbListing.Price,
        })
    }
    return res, nil
}

// Create a listing
func (s *listingssrvc) Add(ctx context.Context, p *listings.Listing) (*listings.Listing, error) {
	newListing := models.Listing{
        Slug:    p.Slug,
        Address: p.Address,
        Price:   p.Price,
    }

    err := newListing.Save()
    if err != nil {
        s.logger.Printf("error creating listing: %v", err)
        return nil, listings.MakeInternalError(err)
    }

    return &listings.Listing{
        ID:      &newListing.ID,
        Slug:    newListing.Slug,
        Address: newListing.Address,
        Price:   newListing.Price,
    }, nil
}

// Retrieve a single listing
func (s *listingssrvc) Get(ctx context.Context, p *listings.GetPayload) (res *listings.Listing, err error) {
	dbListing, err := models.GetListingBySlug(p.Slug)
    if err != nil {
        s.logger.Printf("error retrieving listing: %v", err)
        return nil, listings.MakeNotFoundError(err)
    }

    return &listings.Listing{
        ID:      &dbListing.ID,
        Slug:    dbListing.Slug,
        Address: dbListing.Address,
        Price:   dbListing.Price,
    }, nil
}

// Update a listing
func (s *listingssrvc) Update(ctx context.Context, p *listings.UpdatePayload) error {
	dbListing, err := models.GetListingBySlug(p.Slug)
    if err != nil {
        s.logger.Printf("error fetching listing to update: %v", err)
        return listings.MakeNotFoundError(err)
    }

    dbListing.Slug = p.Listing.Slug
    dbListing.Address = p.Listing.Address
    dbListing.Price = p.Listing.Price

    err = dbListing.Update()
    if err != nil {
        s.logger.Printf("error updating listing: %v", err)
        return listings.MakeInternalError(err)
    }

    return nil
}

// Delete a listing
func (s *listingssrvc) Delete(ctx context.Context, p *listings.DeletePayload) error {
	dbListing, err := models.GetListingBySlug(p.Slug)
    if err != nil {
        s.logger.Printf("error fetching listing to delete: %v", err)
        return listings.MakeNotFoundError(err)
    }

    err = dbListing.Delete()
    if err != nil {
        s.logger.Printf("error deleting listing: %v", err)
        return listings.MakeInternalError(err)
    }

    return nil
}
