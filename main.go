package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func main(){
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=agartala&appid=b61be9af89958211991024056a1291da")
	if err != nil{
		fmt.Printf("err: %v\n", err)
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("responseBody: %s\n", responseBody)

}