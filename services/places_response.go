package services

type PlacesResponse struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results          []struct {
		BusinessStatus string `json:"business_status"`
		Geometry       struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Icon         string `json:"icon"`
		Name         string `json:"name"`
		OpeningHours struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours,omitempty"`
		Photos []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		PriceLevel        int      `json:"price_level,omitempty"`
		Rating            float64  `json:"rating,omitempty"`
		Reference         string   `json:"reference"`
		Scope             string   `json:"scope"`
		Types             []string `json:"types"`
		UserRatingsTotal  int      `json:"user_ratings_total,omitempty"`
		Vicinity          string   `json:"vicinity"`
		PermanentlyClosed bool     `json:"permanently_closed,omitempty"`
	} `json:"results"`
	Status string `json:"status"`
}