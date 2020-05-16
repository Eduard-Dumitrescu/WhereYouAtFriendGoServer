package models

// AccountCreationData model struct
type AccountCreationData struct {
	PostalCode        string `json:"postalCode" binding:"required"`
	City              string `json:"city" binding:"required"`
	IsLocationFromAPI bool   `json:"isLocationFromAPI" binding:"required"`
}

// ZonesStatusPlaceAndCount model struct
type ZonesStatusPlaceAndCount struct {
	PostalCode    string `json:"postalCode" binding:"required"`
	City          string `json:"city" binding:"required"`
	IsInsideCount int    `json:"isInsideCount" binding:"required"`
	TotalCount    int    `json:"totalCount" binding:"required"`
}

// ZonesStatusCount model struct
type ZonesStatusCount struct {
	IsInsideCount int `json:"isInsideCount" binding:"required"`
	TotalCount    int `json:"totalCount" binding:"required"`
}
