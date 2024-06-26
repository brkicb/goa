// Code generated by goa v3.16.1, DO NOT EDIT.
//
// listings HTTP client CLI support package
//
// Command:
// $ goa gen listings.com/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	listingsc "listings.com/gen/http/listings/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `listings (list|add|get|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` listings list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		listingsFlags = flag.NewFlagSet("listings", flag.ContinueOnError)

		listingsListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		listingsAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		listingsAddBodyFlag = listingsAddFlags.String("body", "REQUIRED", "")

		listingsGetFlags    = flag.NewFlagSet("get", flag.ExitOnError)
		listingsGetSlugFlag = listingsGetFlags.String("slug", "REQUIRED", "Unique identifier of the listing")

		listingsUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		listingsUpdateBodyFlag = listingsUpdateFlags.String("body", "REQUIRED", "")
		listingsUpdateSlugFlag = listingsUpdateFlags.String("slug", "REQUIRED", "Unique identifier of the listing")

		listingsDeleteFlags    = flag.NewFlagSet("delete", flag.ExitOnError)
		listingsDeleteSlugFlag = listingsDeleteFlags.String("slug", "REQUIRED", "Unique identifier of the listing")
	)
	listingsFlags.Usage = listingsUsage
	listingsListFlags.Usage = listingsListUsage
	listingsAddFlags.Usage = listingsAddUsage
	listingsGetFlags.Usage = listingsGetUsage
	listingsUpdateFlags.Usage = listingsUpdateUsage
	listingsDeleteFlags.Usage = listingsDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "listings":
			svcf = listingsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "listings":
			switch epn {
			case "list":
				epf = listingsListFlags

			case "add":
				epf = listingsAddFlags

			case "get":
				epf = listingsGetFlags

			case "update":
				epf = listingsUpdateFlags

			case "delete":
				epf = listingsDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "listings":
			c := listingsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "add":
				endpoint = c.Add()
				data, err = listingsc.BuildAddPayload(*listingsAddBodyFlag)
			case "get":
				endpoint = c.Get()
				data, err = listingsc.BuildGetPayload(*listingsGetSlugFlag)
			case "update":
				endpoint = c.Update()
				data, err = listingsc.BuildUpdatePayload(*listingsUpdateBodyFlag, *listingsUpdateSlugFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = listingsc.BuildDeletePayload(*listingsDeleteSlugFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// listingsUsage displays the usage of the listings command and its subcommands.
func listingsUsage() {
	fmt.Fprintf(os.Stderr, `The listings service provides operations on managing house listings.
Usage:
    %[1]s [globalflags] listings COMMAND [flags]

COMMAND:
    list: Retrieve a list of all listings
    add: Create a listing
    get: Retrieve a single listing
    update: Update a listing
    delete: Delete a listing

Additional help:
    %[1]s listings COMMAND --help
`, os.Args[0])
}
func listingsListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] listings list

Retrieve a list of all listings

Example:
    %[1]s listings list
`, os.Args[0])
}

func listingsAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] listings add -body JSON

Create a listing
    -body JSON: 

Example:
    %[1]s listings add --body '{
      "address": "Fugiat aut voluptatem non exercitationem sed omnis.",
      "id": 7084724792209511641,
      "price": 3845457252673368,
      "slug": "Accusamus et assumenda corrupti."
   }'
`, os.Args[0])
}

func listingsGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] listings get -slug STRING

Retrieve a single listing
    -slug STRING: Unique identifier of the listing

Example:
    %[1]s listings get --slug "Repudiandae voluptatibus."
`, os.Args[0])
}

func listingsUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] listings update -body JSON -slug STRING

Update a listing
    -body JSON: 
    -slug STRING: Unique identifier of the listing

Example:
    %[1]s listings update --body '{
      "listing": {
         "address": "Numquam ut sint possimus velit eos nostrum.",
         "id": 7626958432408663309,
         "price": 8607624620613692515,
         "slug": "Ducimus neque natus blanditiis quam iusto."
      }
   }' --slug "Voluptas alias harum."
`, os.Args[0])
}

func listingsDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] listings delete -slug STRING

Delete a listing
    -slug STRING: Unique identifier of the listing

Example:
    %[1]s listings delete --slug "Totam ea odit tempora nulla aliquid voluptates."
`, os.Args[0])
}
