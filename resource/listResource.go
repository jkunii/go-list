package resource

import (
	"net/http"

	"github.com/labstack/echo"
)

func (tt PropertyResource) Get(c echo.Context) error {

	return c.JSON(http.StatusOK, "works")
}

type PropertyResource struct {
}

type (
	property struct {
		ID   string    `json:"id"`
		ListingType string `json:"listingType"`
		ListingStatus string `json:"listingStatus"`
		
		ParkingSpaces string `json:"parkingSpaces"`
		Owner string `json:"owner"`
		images []string `json:"images"`
		Address address `json:"address"`
		Bathrooms int `json:"bathrooms"`
		Bedrooms int `json:"bedrooms"`
		CreatedAt string `json:"createdAt"`
		updatedAt string `json:"updatedAt"`

		TotalCount int `json:"totalCount"`
		PageNumber int `json:"pageNumber"`
		PageSize int `json:"pageSize"`
	}

	address struct {
		City string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		GeoLocation geoLocation `json:"geoLocation"`
	}

	geoLocation struct {
		Precision string `json:"precision"`
		Location location `json:"location"`
	}

	location struct {
		Lon string `json:"lon"`
		Lat string `json:"lat"`
	}

	pricingInfos struct {
		YearlyIptu	string `json:"yearlyIptu"`
		Price	string `json:"price"`
		BusinessType	string `json:"businessType"`
		MonthlyCondoFee	string `json:"monthlyCondoFee"`
	}


)
