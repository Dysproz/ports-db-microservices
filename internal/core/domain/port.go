package domain

type Port struct {
	Name        string    `json:"name,omitempty"`
	City        string    `json:"city,omitempty"`
	Country     string    `json:"country,omitempty"`
	Alias       []string  `json:"alias,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Coordinates []float32 `json:"coordinates,omitempty"`
	Province    string    `json:"province,omitempty"`
	Timezone    string    `json:"timezone,omitempty"`
	Unlocs      []string  `json:"unlocs,omitempty"`
	Code        string    `json:"code,omitempty"`
}
