package schema

type GovernmentForm struct {
	GovernmentType    string `json:"gov_type"`
	NumberOfCountries int    `json:"country_number"`
}
type CountCountryByContinent struct {
	Continent         string `json:"continent"`
	NumberOfCountries int    `json:"country_number"`
}
