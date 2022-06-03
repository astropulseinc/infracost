package aws

import (
	"fmt"
	"github.com/infracost/infracost/internal/resources"
	"github.com/infracost/infracost/internal/schema"
	"github.com/infracost/infracost/internal/usage"
	"github.com/shopspring/decimal"
	"reflect"
	"regexp"
	"strings"
)

// GlobalAccelerator struct represents AWS Global Accelerator service
//
// Resource information: https://aws.amazon.com/global-accelerator
// Pricing information: https://aws.amazon.com/global-accelerator/pricing/
type GlobalAccelerator struct {
	Name          string
	IPAddressType string
	Enabled       bool

	MonthlyInboundDataTransferGB  *globalAcceleratorRegionDataTransferUsage `infracost_usage:"monthly_inbound_data_transfer_gb"`
	MonthlyOutboundDataTransferGB *globalAcceleratorRegionDataTransferUsage `infracost_usage:"monthly_outbound_data_transfer_gb"`
}

type globalAcceleratorRegionDataTransferUsage struct {
	FromAsiaPacificToAsiaPacific   *float64 `infracost_usage:"from_asia_pacific_to_asia_pacific"`
	FromAsiaPacificToAustralia     *float64 `infracost_usage:"from_asia_pacific_to_australia"`
	FromAsiaPacificToEurope        *float64 `infracost_usage:"from_asia_pacific_to_europe"`
	FromAsiaPacificToIndia         *float64 `infracost_usage:"from_asia_pacific_to_india"`
	FromAsiaPacificToSouthKorea    *float64 `infracost_usage:"from_asia_pacific_to_south_korea"`
	FromAsiaPacificToMiddleEast    *float64 `infracost_usage:"from_asia_pacific_to_middle_east"`
	FromAsiaPacificToNorthAmerica  *float64 `infracost_usage:"from_asia_pacific_to_north_america"`
	FromAsiaPacificToSouthAmerica  *float64 `infracost_usage:"from_asia_pacific_to_south_america"`
	FromAsiaPacificToSouthAfrica   *float64 `infracost_usage:"from_asia_pacific_to_south_africa"`
	FromAustraliaToAsiaPacific     *float64 `infracost_usage:"from_australia_to_asia_pacific"`
	FromAustraliaToAustralia       *float64 `infracost_usage:"from_australia_to_australia"`
	FromAustraliaToEurope          *float64 `infracost_usage:"from_australia_to_europe"`
	FromAustraliaToIndia           *float64 `infracost_usage:"from_australia_to_india"`
	FromAustraliaToSouthKorea      *float64 `infracost_usage:"from_australia_to_south_korea"`
	FromAustraliaToMiddleEast      *float64 `infracost_usage:"from_australia_to_middle_east"`
	FromAustraliaToNorthAmerica    *float64 `infracost_usage:"from_australia_to_north_america"`
	FromAustraliaToSouthAmerica    *float64 `infracost_usage:"from_australia_to_south_america"`
	FromAustraliaToSouthAfrica     *float64 `infracost_usage:"from_australia_to_south_africa"`
	FromEuropeToAsiaPacific        *float64 `infracost_usage:"from_europe_to_asia_pacific"`
	FromEuropeToAustralia          *float64 `infracost_usage:"from_europe_to_australia"`
	FromEuropeToEurope             *float64 `infracost_usage:"from_europe_to_europe"`
	FromEuropeToIndia              *float64 `infracost_usage:"from_europe_to_india"`
	FromEuropeToSouthKorea         *float64 `infracost_usage:"from_europe_to_south_korea"`
	FromEuropeToMiddleEast         *float64 `infracost_usage:"from_europe_to_middle_east"`
	FromEuropeToNorthAmerica       *float64 `infracost_usage:"from_europe_to_north_america"`
	FromEuropeToSouthAmerica       *float64 `infracost_usage:"from_europe_to_south_america"`
	FromEuropeToSouthAfrica        *float64 `infracost_usage:"from_europe_to_south_africa"`
	FromIndiaToAsiaPacific         *float64 `infracost_usage:"from_india_to_asia_pacific"`
	FromIndiaToAustralia           *float64 `infracost_usage:"from_india_to_australia"`
	FromIndiaToEurope              *float64 `infracost_usage:"from_india_to_europe"`
	FromIndiaToIndia               *float64 `infracost_usage:"from_india_to_india"`
	FromIndiaToSouthKorea          *float64 `infracost_usage:"from_india_to_south_korea"`
	FromIndiaToMiddleEast          *float64 `infracost_usage:"from_india_to_middle_east"`
	FromIndiaToNorthAmerica        *float64 `infracost_usage:"from_india_to_north_america"`
	FromIndiaToSouthAmerica        *float64 `infracost_usage:"from_india_to_south_america"`
	FromIndiaToSouthAfrica         *float64 `infracost_usage:"from_india_to_south_africa"`
	FromSouthKoreaToAsiaPacific    *float64 `infracost_usage:"from_south_korea_to_asia_pacific"`
	FromSouthKoreaToAustralia      *float64 `infracost_usage:"from_south_korea_to_australia"`
	FromSouthKoreaToEurope         *float64 `infracost_usage:"from_south_korea_to_europe"`
	FromSouthKoreaToIndia          *float64 `infracost_usage:"from_south_korea_to_india"`
	FromSouthKoreaToSouthKorea     *float64 `infracost_usage:"from_south_korea_to_south_korea"`
	FromSouthKoreaToMiddleEast     *float64 `infracost_usage:"from_south_korea_to_middle_east"`
	FromSouthKoreaToNorthAmerica   *float64 `infracost_usage:"from_south_korea_to_north_america"`
	FromSouthKoreaToSouthAmerica   *float64 `infracost_usage:"from_south_korea_to_south_america"`
	FromSouthKoreaToSouthAfrica    *float64 `infracost_usage:"from_south_korea_to_south_africa"`
	FromMiddleEastToAsiaPacific    *float64 `infracost_usage:"from_middle_east_to_asia_pacific"`
	FromMiddleEastToAustralia      *float64 `infracost_usage:"from_middle_east_to_australia"`
	FromMiddleEastToEurope         *float64 `infracost_usage:"from_middle_east_to_europe"`
	FromMiddleEastToIndia          *float64 `infracost_usage:"from_middle_east_to_india"`
	FromMiddleEastToSouthKorea     *float64 `infracost_usage:"from_middle_east_to_south_korea"`
	FromMiddleEastToMiddleEast     *float64 `infracost_usage:"from_middle_east_to_middle_east"`
	FromMiddleEastToNorthAmerica   *float64 `infracost_usage:"from_middle_east_to_north_america"`
	FromMiddleEastToSouthAmerica   *float64 `infracost_usage:"from_middle_east_to_south_america"`
	FromMiddleEastToSouthAfrica    *float64 `infracost_usage:"from_middle_east_to_south_africa"`
	FromNorthAmericaToAsiaPacific  *float64 `infracost_usage:"from_north_america_to_asia_pacific"`
	FromNorthAmericaToAustralia    *float64 `infracost_usage:"from_north_america_to_australia"`
	FromNorthAmericaToEurope       *float64 `infracost_usage:"from_north_america_to_europe"`
	FromNorthAmericaToIndia        *float64 `infracost_usage:"from_north_america_to_india"`
	FromNorthAmericaToSouthKorea   *float64 `infracost_usage:"from_north_america_to_south_korea"`
	FromNorthAmericaToMiddleEast   *float64 `infracost_usage:"from_north_america_to_middle_east"`
	FromNorthAmericaToNorthAmerica *float64 `infracost_usage:"from_north_america_to_north_america"`
	FromNorthAmericaToSouthAmerica *float64 `infracost_usage:"from_north_america_to_south_america"`
	FromNorthAmericaToSouthAfrica  *float64 `infracost_usage:"from_north_america_to_south_africa"`
	FromSouthAmericaToAsiaPacific  *float64 `infracost_usage:"from_south_america_to_asia_pacific"`
	FromSouthAmericaToAustralia    *float64 `infracost_usage:"from_south_america_to_australia"`
	FromSouthAmericaToEurope       *float64 `infracost_usage:"from_south_america_to_europe"`
	FromSouthAmericaToIndia        *float64 `infracost_usage:"from_south_america_to_india"`
	FromSouthAmericaToSouthKorea   *float64 `infracost_usage:"from_south_america_to_south_korea"`
	FromSouthAmericaToMiddleEast   *float64 `infracost_usage:"from_south_america_to_middle_east"`
	FromSouthAmericaToNorthAmerica *float64 `infracost_usage:"from_south_america_to_north_america"`
	FromSouthAmericaToSouthAmerica *float64 `infracost_usage:"from_south_america_to_south_america"`
	FromSouthAmericaToSouthAfrica  *float64 `infracost_usage:"from_south_america_to_south_africa"`
	FromSouthAfricaToAsiaPacific   *float64 `infracost_usage:"from_south_africa_to_asia_pacific"`
	FromSouthAfricaToAustralia     *float64 `infracost_usage:"from_south_africa_to_australia"`
	FromSouthAfricaToEurope        *float64 `infracost_usage:"from_south_africa_to_europe"`
	FromSouthAfricaToIndia         *float64 `infracost_usage:"from_south_africa_to_india"`
	FromSouthAfricaToSouthKorea    *float64 `infracost_usage:"from_south_africa_to_south_korea"`
	FromSouthAfricaToMiddleEast    *float64 `infracost_usage:"from_south_africa_to_middle_east"`
	FromSouthAfricaToNorthAmerica  *float64 `infracost_usage:"from_south_africa_to_north_america"`
	FromSouthAfricaToSouthAmerica  *float64 `infracost_usage:"from_south_africa_to_south_america"`
	FromSouthAfricaToSouthAfrica   *float64 `infracost_usage:"from_south_africa_to_south_africa"`
}

type dataTransferElement struct {
	from             string
	to               string
	trafficDirection string
	quantity         float64
}

var (
	GlobalAcceleratorUsageSchema = []*schema.UsageItem{
		{
			Key:          "monthly_inbound_data_transfer_gb",
			DefaultValue: &usage.ResourceUsage{Name: "monthly_inbound_data_transfer_gb", Items: globalAcceleratorRegionDataTransferUsageSchema},
			ValueType:    schema.SubResourceUsage,
		},
		{
			Key:          "monthly_outbound_data_transfer_gb",
			DefaultValue: &usage.ResourceUsage{Name: "monthly_outbound_data_transfer_gb", Items: globalAcceleratorRegionDataTransferUsageSchema},
			ValueType:    schema.SubResourceUsage,
		},
	}
	globalAcceleratorRegionDataTransferUsageSchema = []*schema.UsageItem{
		{Key: "from_asia_pacific_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_asia_pacific_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_europe_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_india_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_korea_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_middle_east_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_north_america_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_america_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_asia_pacific", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_australia", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_europe", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_india", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_australia_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_south_korea", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_middle_east", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_north_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_south_america", DefaultValue: 0, ValueType: schema.Float64},
		{Key: "from_south_africa_to_south_africa", DefaultValue: 0, ValueType: schema.Float64},
	}
	regionToCodeMap = map[string]string{
		"asia_pacific":  "AP",
		"australia":     "AU",
		"europe":        "EU",
		"india":         "IN",
		"south_korea":   "KR",
		"middle_east":   "ME",
		"north_america": "NA",
		"south_america": "SA",
		"south_africa":  "ZA",
	}
	fromToUsageRegex = regexp.MustCompile(`from_(?P<From>[a-z_]+)_to_(?P<To>[a-z_]+)`)
)

func (r *GlobalAccelerator) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *GlobalAccelerator) BuildResource() *schema.Resource {
	var (
		inboundDataTransferUsage  float64          = 0
		outboundDataTransferUsage float64          = 0
		resource                  *schema.Resource = &schema.Resource{
			Name:           r.Name,
			UsageSchema:    GlobalAcceleratorUsageSchema,
			CostComponents: nil,
		}
	)

	if !r.Enabled {
		return resource
	}

	if r.MonthlyInboundDataTransferGB != nil {
		inboundDataTransferUsage = calculateDataTransferUsage(r.MonthlyInboundDataTransferGB)
	}
	if r.MonthlyOutboundDataTransferGB != nil {
		outboundDataTransferUsage = calculateDataTransferUsage(r.MonthlyOutboundDataTransferGB)
	}

	costComponents := []*schema.CostComponent{
		r.fixedCostComponent(),
	}

	if inboundDataTransferUsage > 0 || outboundDataTransferUsage > 0 {
		direction := "In"
		dominantDirectionUsage := r.MonthlyInboundDataTransferGB
		// There is no info in the AWS docs about the very remote corner case inboundDataTransferUsage == outboundDataTransferUsage
		if outboundDataTransferUsage > inboundDataTransferUsage {
			direction = "Out"
			dominantDirectionUsage = r.MonthlyOutboundDataTransferGB
		}
		costComponents = append(costComponents, r.dataTransferCostComponents(direction, dominantDirectionUsage)...)
	}
	resource.CostComponents = costComponents
	return resource
}

func (r *GlobalAccelerator) fixedCostComponent() *schema.CostComponent {
	c := &schema.CostComponent{
		Name:           fmt.Sprintf("AWS Global Accelerator %s Fixed Fee", r.Name),
		Unit:           "hours",
		UnitMultiplier: decimal.NewFromInt(1),
		HourlyQuantity: decimalPtr(decimal.NewFromInt(1)),
		ProductFilter: &schema.ProductFilter{
			VendorName: strPtr("aws"),
			Service:    strPtr("AWSGlobalAccelerator"),
		},
	}
	// AWS Global Accelerator has a fixed fee of 0.025$ per hour.
	// This price unfortunately is not mapped actually in AWS Pricing API
	// More: AWS_DEFAULT_REGION=us-east-1 aws pricing get-products --service-code AWSGlobalAccelerator | jq -r '.PriceList[] | fromjson | .product'
	c.SetCustomPrice(decimalPtr(decimal.NewFromFloat(0.025)))
	return c
}

func (r *GlobalAccelerator) dataTransferCostComponents(direction string, usage *globalAcceleratorRegionDataTransferUsage) []*schema.CostComponent {
	f := reflect.TypeOf(*usage)
	v := reflect.ValueOf(*usage)
	dataTransferElements := []*dataTransferElement{}
	costComponents := []*schema.CostComponent{}
	for i := 0; i < f.NumField(); i++ {
		value := reflect.Indirect(v.Field(i))
		if value.Kind() == 0 {
			continue
		}
		quantity := reflect.Indirect(v.Field(i)).Float()
		tag := f.Field(i).Tag.Get("infracost_usage")
		regexRes := fromToUsageRegex.FindStringSubmatch(tag)
		dataTransferElements = append(dataTransferElements, &dataTransferElement{
			from:             regionToCodeMap[regexRes[1]],
			to:               regionToCodeMap[regexRes[2]],
			trafficDirection: direction,
			quantity:         quantity,
		})
	}
	for _, d := range dataTransferElements {
		name := fmt.Sprintf("AWS Global Accelerator %s DT-Premium Usage %s from %s to %s", r.Name, strings.ToUpper(d.trafficDirection), d.from, d.to)
		// Even if there are multiple price record entries the price for -Bytes-Internet and -Bytes-AWS for the same regions are equal
		// So one of these two can be fixed to avoid multiple prices found
		usageType := fmt.Sprintf("%s-%s-%s-Bytes-Internet", strings.ToUpper(d.from), strings.ToUpper(d.to), strings.ToUpper(d.trafficDirection))
		costComponents = append(costComponents, &schema.CostComponent{
			Name:            name,
			Unit:            "GB",
			UnitMultiplier:  decimal.NewFromInt(1),
			MonthlyQuantity: decimalPtr(decimal.NewFromFloat(d.quantity)),
			ProductFilter: &schema.ProductFilter{
				VendorName: strPtr("aws"),
				Service:    strPtr("AWSGlobalAccelerator"),
				AttributeFilters: []*schema.AttributeFilter{
					{Key: "trafficDirection", Value: strPtr(d.trafficDirection)},
					{Key: "fromLocation", Value: strPtr(d.from)},
					{Key: "toLocation", Value: strPtr(d.to)},
					{Key: "operation", Value: strPtr("Dominant")},
					{Key: "usagetype", Value: strPtr(usageType)},
				},
			},
		})
	}
	return costComponents
}

func calculateDataTransferUsage(usage *globalAcceleratorRegionDataTransferUsage) float64 {
	var (
		sum float64 = 0
	)
	v := reflect.ValueOf(*usage)
	for i := 0; i < v.NumField(); i++ {
		value := reflect.Indirect(v.Field(i))
		if value.Kind() != 0 {
			sum += reflect.Indirect(v.Field(i)).Float()
		}
	}
	return sum
}
