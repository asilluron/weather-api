package models

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type WeatherResponse struct {
	Coordinates        Coordinates `json:"coordinates"`
	ShortForecast      string      `json:"shortForecast"`
	TemperatureSummary string      `json:"temperatureSummary"`
}

type Point struct {
	Properties struct {
		GridId   string `json:"gridId"`
		GridX    int    `json:"gridX"`
		GridY    int    `json:"gridY"`
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type Forecast struct {
	Properties struct {
		Periods []Period `json:"periods"`
	} `json:"properties"`
}

type Period struct {
	Number        int    `json:"number"`
	Name          string `json:"name"`
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}
