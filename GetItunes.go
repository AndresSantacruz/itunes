package itunes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	ResultCount int
	Results []map[string]string
}

func GetItunes(query string)  {
	res, err := http.Get("https://itunes.apple.com/search?term=" + query)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)

	var response Response
	json.Unmarshal(data, &response)

	encjson, _ := json.Marshal(response.Results)

	fmt.Println(string(encjson))
}