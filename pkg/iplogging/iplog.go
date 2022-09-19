package iplogging

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInformation struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

func Iplog() {
	var ip string
	fmt.Println("\n \n Please enter the IP address: ")
	fmt.Scanln(&ip)
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	var data IPInformation
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		print(err)
	}
	fmt.Printf("Query: %s \n", data.Query)
	fmt.Printf("Status: %s \n", data.Status)
	fmt.Printf("Country: %s \n", data.Country)
	fmt.Printf("CountryCode: %s \n", data.CountryCode)
	fmt.Printf("Region: %s \n", data.Region)
	fmt.Printf("RegionName: %s \n", data.RegionName)
	fmt.Printf("City: %s \n", data.City)
	fmt.Printf("Zip: %s \n", data.Zip)
	fmt.Printf("Lat: %f \n", data.Lat)
	fmt.Printf("Lon: %f \n", data.Lon)
	fmt.Printf("Timezone: %s \n", data.Timezone)
	fmt.Printf("Isp: %s \n", data.Isp)
	fmt.Printf("Org: %s \n", data.Org)
	fmt.Printf("As: %s \n", data.As)
	fmt.Printf("\nView location at: https://www.google.com/maps/@%f,%f,15z \n \n", data.Lat, data.Lon)
}
