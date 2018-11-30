package service

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/jkunii/go-list/global"
)

var (
	data []Property
	url  = "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
)

func GetProperties() ([]Property, error) {
	if data != nil {
		log.Println("##########")

		log.Println("Announcement already populated.")

		return data, nil

	} else {
		log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$")

		log.Println("Announcement not populated.")
		res, err := http.Get(url)
		if err != nil {
			global.Error(err)
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)

		err = decoder.Decode(&data)
	}
	return data, nil
}

type (
	Property struct {
		ID            string   `json:"id"`
		ListingType   string   `json:"listingType"`
		ListingStatus string   `json:"listingStatus"`
		UsableAreas   string   `json:"usableAreas"`
		ParkingSpaces int      `json:"parkingSpaces"`
		Owner         string   `json:"owner"`
		Images        []string `json:"images"`
		Address       Address  `json:"address"`
		Bathrooms     int      `json:"bathrooms"`
		Bedrooms      int      `json:"bedrooms"`
		CreatedAt     string   `json:"createdAt"`
		UpdatedAt     string   `json:"updatedAt"`
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
