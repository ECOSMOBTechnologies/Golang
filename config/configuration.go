package config

//Configuration file
type Configuration struct {
	Dev struct {
		PORT string `json:"PORT"`
	} `json:"dev"`
	Production struct {
		PORT string `json:"PORT"`
	} `json:"production"`
}
