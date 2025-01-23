package rest

import (
	rk7client "github.com/alexandr-andreyev/rk7-xml-apiclient"
	"github.com/gofiber/fiber/v2"
)

type RestaurantResult struct {
	ID                   string              `json:"id"`
	RKeeperCode          string              `json:"r_keeper_code"`
	FullRestaurantCode   string              `json:"full_restaurant_code"`
	Name                 string              `json:"name"`
	Guid                 string              `json:"guid"`
	Address              string              `json:"address"`
	NationalCurrencyCode string              `json:"national_currency_code"`
	Concept              RestaurantConcept   `json:"concept"`
	Franchise            RestaurantFranchise `json:"franchise"`
	LocationLat          string              `json:"location_lat"`
	LocationLong         string              `json:"location_long"`
	Region               RestaurantRegion    `json:"region"`
	Image                string              `json:"image"`
	WorkHours            string              `json:"work_hours"`
	OperationHours       string              `json:"operation_hours"`
}

type RestaurantConcept struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Guid string `json:"guid"`
}

type RestaurantFranchise struct {
	ID string `json:"id"`
}

type RestaurantRegion struct {
	ID         string `json:"id"`
	Сode       string `json:"code"`
	Name       string `json:"name"`
	Guid       string `json:"guid"`
	UTCOffsets string `json:"utc_offsets"`
}

func (h Handler) GetRestaurants(c *fiber.Ctx) error {
	response, err := h.services.RKeeperService.GetRestaurants()
	if err != nil {
		return err
	}

	restaurantResult := &[]RestaurantResult{}

	concepts := make(map[string]RestaurantConcept)
	regions := make(map[string]RestaurantRegion)

	for _, result := range response.CommandResult[0].Data {
		for _, item := range result.(rk7client.RK7Reference).Items {
			concept := RestaurantConcept{
				ID:   item.Ident,
				Code: item.Code,
				Name: item.Name,
				Guid: item.GuidString,
			}

			concepts[item.Ident] = concept
		}
	}

	for _, result := range response.CommandResult[1].Data {
		for _, item := range result.(rk7client.RK7Reference).Items {
			region := RestaurantRegion{
				ID:         item.Ident,
				Сode:       item.Code,
				Name:       item.Name,
				Guid:       item.GuidString,
				UTCOffsets: item.Attrs["UTCOffsets"],
			}

			regions[item.Ident] = region
		}
	}

	for _, result := range response.CommandResult[2].Data {
		for _, item := range result.(rk7client.RK7Reference).Items {
			restaurant := RestaurantResult{
				ID:                   item.Ident,
				RKeeperCode:          item.Code,
				FullRestaurantCode:   item.Attrs["FullRestaurantCode"],
				Name:                 item.Name,
				Guid:                 item.GuidString,
				Address:              item.Attrs["Address"],
				NationalCurrencyCode: item.Attrs["NationalCurrency"],
				Concept:              concepts[item.Attrs["Concept"]],
				Franchise:            RestaurantFranchise{ID: item.Attrs["Franchise"]},
				LocationLat:          item.Attrs["LocationLat"],
				LocationLong:         item.Attrs["LocationLong"],
				Region:               regions[item.Attrs["Region"]],
				Image:                item.Attrs["genPhotoForDelivery"],
				WorkHours:            item.Attrs["genworkPeriod"],
				OperationHours:       item.Attrs["OperationalHours"],
			}

			*restaurantResult = append(*restaurantResult, restaurant)
		}
	}

	return c.JSON(restaurantResult)
}
