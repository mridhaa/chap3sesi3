package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestBody struct{
	Water int  `json:"water"`
	Wind int `json:"wind"`
}

func main(){

	apiUrl :="https://jsonplaceholder.typicode.com/posts"


	// rand.Seed(time.Now().UnixNano())

	for {
		data:= RequestBody{
			Water: rand.Intn(16),
			Wind: rand.Intn(16),
		}

		bs, err := json.Marshal(data)

		if err != nil{
			log.Panicf("error while converting struct to json : %s", err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bs))

		if err != nil{
			log.Panicf("error while defining  the request instance : %s", err.Error())
		}		

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		response, err := client.Do(request)

		if err != nil{
			log.Panicf("error while sending the api request : %s", err.Error())
		}	

		defer response.Body.Close()

		responseBody, err := ioutil.ReadAll(response.Body)

		fmt.Println(string(responseBody))

		if (data.Water < 5) {
			fmt.Println("Status water: aman")	
		}else if(data.Water >= 6 && data.Water<=8){
			fmt.Println("Status water: siaga")	
		}else if(data.Water>8){
			fmt.Println("Status water: bahaya")	
		}


		if (data.Wind < 6) {
			fmt.Println("Status wind: aman")	
		}else if(data.Wind >= 7 && data.Wind<=15){
			fmt.Println("Status wind: siaga")	
		}else if(data.Wind>15){
			fmt.Println("Status wind: bahaya")	
		}

		time.Sleep(time.Second *2)
	}


}