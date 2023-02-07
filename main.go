package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type report struct {
	CityName    string `json:"name"`
	Main m `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`

	}`json:"wind"`
}
type m struct{
	TempActual float32 `json:"temp"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}

func main(){
	var city string
	fmt.Println("Enter any city name")
	fmt.Scanf("%s", &city)
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=b61be9af89958211991024056a1291da")
	if err != nil{
		fmt.Printf("err: %v\n", err)
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	// fmt.Printf("responseBody: %s\n", responseBody)

	var responseObject report
    json.Unmarshal(responseBody, &responseObject)
	// fmt.Printf("responseObject: %v\n", responseObject)
    fmt.Println("City name",responseObject.CityName)
	fmt.Println("Temparature ", responseObject.Main.TempActual - 273.15)
	fmt.Println("Maximum Temparature ", responseObject.Main.TempMax - 273.15)
	fmt.Println("Minimum Temparature ", responseObject.Main.TempMin - 273.15)
    fmt.Println("Wind details",responseObject.Wind.Speed)
	
	
}
