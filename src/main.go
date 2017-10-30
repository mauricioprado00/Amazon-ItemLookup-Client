package main

import (
	"fmt"
	"net/url"
	"github.com/dominicphillips/amazing"
	"encoding/json"
	"flag"
	"reflect"
	"os"
)

var Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s :\n", os.Args[0])
        flag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\nRequired Environment Variables:\n")
        fmt.Fprintf(os.Stderr, "  -AMAZING_ASSOCIATE_TAG\n")
        fmt.Fprintf(os.Stderr, "  -AMAZING_ACCESS_KEY\n")
        fmt.Fprintf(os.Stderr, "  -AMAZING_SECRET_KEY\n")
        fmt.Fprintf(os.Stderr, "\n see http://docs.aws.amazon.com/AWSECommerceService/latest/DG/ItemLookup.html\n")
        fmt.Fprintf(os.Stderr, " see https://app.teampassword.com/dashboard#account/104478\n")
}

func main() {
	var ItemId string
	var IdType string
	var ResponseGroup string
	var TruncateReviewsAt string
	var SearchIndex string
	var RelationshipType string
	var RelatedItemPage string
	var IncludeReviewsSummary string
	var Condition string

	client, _ := amazing.NewAmazingFromEnv("DE")

	flag.Usage = Usage;
	
	flag.StringVar(&ItemId, "ItemId", "", "One or more (up to ten) positive integers that uniquely identify an item.\n" + 
		"\tThe meaning of the number is specified by IdType. That is, if IdType is ASIN, the ItemId value is an ASIN.\n" +
		"\tIf ItemIdis an ASIN, a search index cannot be specified in the request.\n\n" + 
		"\tType: String\n\n" + 
		"\tDefault: None\n\n" + 
		"\tConstraints: Must be a valid item ID. For more than one ID, use a comma-separated list of up to ten IDs.\n")

	flag.StringVar(&IdType, "IdType", "", "Type of item identifier used to look up an item. All IdTypes except ASINx require a SearchIndex to be specified.\n\n" + 
		"\tType: String\n\n" + 
		"\tDefault: ASIN\n\n" + 
		"\tValid Values: SKU | UPC | EAN | ISBN (US only, when search index is Books). UPC is not valid in the CA locale.\n")

	flag.StringVar(&ResponseGroup, "ResponseGroup", "", "Specifies the types of values to return.\n" + 
		"\tYou can specify multiple response groups in one request by separating them with commas.\n\n" +
		"\tType: String\n\n" +
		"\tDefault: Small\n\n" +
		"\tValid Values: Accessories | BrowseNodes | EditorialReview | Images | ItemAttributes | \n" +
		"\t\tItemIds | Large | Medium | OfferFull | Offers | PromotionSummary | OfferSummary| \n" +
		"\t\tRelatedItems | Reviews | SalesRank | Similarities | Small | Tracks | VariationImages | \n" +
		"\t\tVariations (US only) | VariationSummary\n")

	flag.StringVar(&TruncateReviewsAt, "TruncateReviewsAt", "", "By default, reviews are truncated to 1000 characters within the Reviews iframe. \n" +
		"\tTo specify a different length, enter the value.\n" + 
		"\tTo return complete reviews, specify 0.\n\n" + 
		"\tType: Integer\n\n" + 
		"\tDefault: 1000\n\n" + 
		"\tConstraints: Must be a positive integer or 0 (returns entire review)")

	flag.StringVar(&SearchIndex, "SearchIndex", "", "The product category to search.\n\n" + 
		"\tType: String\n\n" + 
		"\tDefault: None\n\n" + 
		"\tValid Values: A search index, for example, Apparel, Beauty, Blended, Books, and so forth. \n" +
		"\tFor a complete of search indices, see Locale Reference for the Product Advertising API.\n\n" + 
		"\tConstraint: If ItemIdis an ASIN, a search index cannot be specified in the request. Required for non-ASIN ItemIds.\n" +
		"\thttp://docs.aws.amazon.com/AWSECommerceService/latest/DG/localevalues.html\n" +
		"\thttp://docs.aws.amazon.com/AWSECommerceService/latest/DG/LocaleDE.html\n")

	flag.StringVar(&RelationshipType, "RelationshipType", "", "This parameter is required when the RelatedItems response group is used.\n" + 
		"\tThe type of related item returned is specified by the RelationshipType parameter.\n" + 
		"\tSample values include Episode, Season, and Tracks.\n" + 
		"\tFor a complete list of types, see Relationship Types.\n\n" + 
		"\tRequired when RelatedItems response group is used.\n" +
		"\thttp://docs.aws.amazon.com/AWSECommerceService/latest/DG/Motivating_RelatedItems.html#RelationshipTypes\n")

	flag.StringVar(&RelatedItemPage, "RelatedItemPage", "", "This optional parameter is only valid when the RelatedItems response group is used.\n" + 
		"\tEach ItemLookup request can return, at most, ten related items.\n" + 
		"\tThe RelatedItemPage value specifies the set of ten related items to return.\n" + 
		"\tA value of 2, for example, returns the second set of ten related items\n")

	flag.StringVar(&IncludeReviewsSummary, "IncludeReviewsSummary", "", "When set to true, returns the reviews summary within the Reviews iframe.\n\n" + 
		"\tType: Boolean\n\n" + 
		"\tDefault: True\n\n" + 
		"\tValid Values: True | False\n")

	flag.StringVar(&Condition, "Condition", "", "Specifies an item's condition.\n" + 
		"\tIf Condition is set to \"All\", a separate set of responses is returned for each valid value of Condition.\n" + 
		"\tThe default value is \"New\" (not \"All\").\n" + 
		"\tSo, if your request does not return results, consider setting the value to \"All\".\n" + 
		"\tWhen the value is \"New\", the ItemLookup availability parameter cannot be set to \"Available\".\n" + 
		"\tAmazon only sells items that are \"New\".\n\n" + 
		"\tType: String\n\n" + 
		"\tDefault: New\n\n" + 
		"\tValid Values: Used | Collectible | Refurbished, All\n")

	Verbose := flag.Bool("v", false, "increase verbosity")

	flag.Parse()

	params := url.Values{
	  "Operation":     []string{"ItemLookup"},
	}

	if len(ItemId) > 0 {
		params.Set("ItemId", ItemId)
	}

	if len(IdType) > 0 {
		params.Set("IdType", IdType)
	}

	if len(ResponseGroup) > 0 {
		params.Set("ResponseGroup", ResponseGroup)
	}

	if len(TruncateReviewsAt) > 0 {
		params.Set("TruncateReviewsAt", TruncateReviewsAt)
	}

	if len(SearchIndex) > 0 {
		params.Set("SearchIndex", SearchIndex)
	}

	if len(RelationshipType) > 0 {
		params.Set("RelationshipType", RelationshipType)
	}

	if len(RelatedItemPage) > 0 {
		params.Set("RelatedItemPage", RelatedItemPage)
	}

	if len(IncludeReviewsSummary) > 0 {
		params.Set("IncludeReviewsSummary", IncludeReviewsSummary)
	}

	if len(Condition) > 0 {
		params.Set("Condition", Condition)
	}

	if *Verbose {
		fmt.Printf("ItemId: %s\n", ItemId)
		fmt.Printf("params: %+v\n", params)
		fmt.Println(reflect.TypeOf(params))
	}

	result, _ := client.ItemLookup(params)

	b, _ := json.Marshal(result)

	fmt.Printf("%s\n", b)
}