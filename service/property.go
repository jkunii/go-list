package service

import (
	"net/http"
	"sort"

	"encoding/json"

	"github.com/jkunii/go-list/global"
)

var (
	data            []Property
	url             = "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
	limitPerPage    = 50       // default
	offSet          = 0        // default
	sale_property   []Property // sale
	rental_property []Property // rental
)

//USED | NEW
func SortByType(property []Property, types string) ([]Property, error) {

	return property, nil
}

// SALE | RENTAL
func SliceBybusinessType(property []Property, groupType string) ([]Property, error) {
	for _, element := range property {
		if element.PricingInfos.BusinessType == "SALE" {
			sale_property = append(sale_property, element)
		}
		if element.PricingInfos.BusinessType == "RENTAL" {
			rental_property = append(rental_property, element)
		}
	}

	if groupType == "SALE" {
		return sale_property, nil
	} else {
		return rental_property, nil
	}

}

// Price by asc
func SortByPriceAsc(property []Property) ([]Property, error) {
	sort.Slice(property, func(i, j int) bool {
		return property[i].PricingInfos.Price < property[j].PricingInfos.Price
	})
	return property, nil
}

// Price by desc
func SortByPriceDesc(property []Property) ([]Property, error) {
	sort.Slice(property, func(i, j int) bool {
		return property[i].PricingInfos.Price > property[j].PricingInfos.Price
	})
	return property, nil
}

func GetPropertiesByPagination(property []Property, offSet, limitPerPage int) ([]Property, error) {
	var sm []Property
	for i := 0; i < limitPerPage; i++ {
		// log.Println("########### offSet ###########")
		// log.Println(offSet)
		// log.Println(property[offSet].ID)
		sm = append(sm, property[offSet])
		offSet++
	}
	return sm, nil
}

func GetProperties(groupType, orderByPrice string, offSet, limitPerPage int) ([]Property, error) {
	var smp []Property

	if data == nil {
		res, err := http.Get(url)
		if err != nil {
			global.Error(err)
		}
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&data)
	}

	propertyByBusinessType, _ := SliceBybusinessType(data, groupType)
	smp, _ = GetPropertiesByPagination(propertyByBusinessType, offSet, limitPerPage)
	ppp, _ := SortByPriceAsc(smp)
	if orderByPrice == "desc" {
		ppp, _ = SortByPriceDesc(smp)
	}
	return ppp, nil
}

type (
	Property struct {
		ID            string       `json:"id"`
		ListingType   string       `json:"listingType"`
		ListingStatus string       `json:"listingStatus"`
		UsableAreas   string       `json:"usableAreas"`
		ParkingSpaces int          `json:"parkingSpaces"`
		Owner         string       `json:"owner"`
		Images        []string     `json:"images"`
		Address       Address      `json:"address"`
		Bathrooms     int          `json:"bathrooms"`
		Bedrooms      int          `json:"bedrooms"`
		PricingInfos  PricingInfos `json:"pricingInfos"`
		CreatedAt     string       `json:"createdAt"`
		UpdatedAt     string       `json:"updatedAt"`
	}

	Address struct {
		City         string      `json:"city"`
		Neighborhood string      `json:"neighborhood"`
		GeoLocation  GeoLocation `json:"geoLocation"`
	}

	GeoLocation struct {
		Precision string   `json:"precision"`
		Location  Location `json:"location"`
	}

	Location struct {
		Lon string `json:"lon"`
		Lat string `json:"lat"`
	}

	PricingInfos struct {
		YearlyIptu      string `json:"yearlyIptu"`
		Price           string `json:"price"`
		BusinessType    string `json:"businessType"`
		MonthlyCondoFee string `json:"monthlyCondoFee"`
	}
)
