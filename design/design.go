package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("listings", func() {
    Title("House Listings Service")
    Description("Service for managing house listings")
    Server("listings", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
            URI("grpc://localhost:8080")
        })
    })
})

var Listing = Type("Listing", func() {
    Description("A Listing describes a house available for sale")
    Attribute("id", Int64, "Unique identifier of the listing")
    Attribute("slug", String, "Unique identifier of the listing")
    Attribute("address", String, "Address of the house")
    Attribute("price", Int64, "Price of the house")
    Required("slug", "address", "price")
})

var _ = Service("listings", func() {
    Description("The listings service provides operations on managing house listings.")

    HTTP(func() {
        Path("/api/listings")
    })

    Method("list", func() {
        Description("Retrieve a list of all listings")
        Result(ArrayOf(Listing))
        HTTP(func() {
            GET("/")
            Response(StatusOK)
        })
    })

    Method("add", func() {
        Description("Create a listing")
        Payload(Listing)
        HTTP(func() {
            POST("/")
            Response(StatusCreated)
        })
    })

    Method("get", func() {
        Description("Retrieve a single listing")
        Payload(func() {
            Attribute("slug", String, "Unique identifier of the listing")
            Required("slug")
        })
        Result(Listing)
        HTTP(func() {
            GET("/{slug}")
            Response(StatusOK)
        })
    })

    Method("update", func() {
        Description("Update a listing")
        Payload(func() {
            Attribute("slug", String, "Unique identifier of the listing")
            Attribute("listing", Listing)
            Required("slug", "listing")
        })
        HTTP(func() {
            PUT("/{slug}")
            Response(StatusNoContent)
        })
    })

    Method("delete", func() {
        Description("Delete a listing")
        Payload(func() {
            Attribute("slug", String, "Unique identifier of the listing")
            Required("slug")
        })
        HTTP(func() {
            DELETE("/{slug}")
            Response(StatusNoContent)
        })
    })

    Error("not_found", ErrorResult, "Listing not found")
    Error("internal_error", ErrorResult, "Internal server error")

    HTTP(func() {
        Response("not_found", StatusNotFound)
        Response("internal_error", StatusInternalServerError)
    })
})
