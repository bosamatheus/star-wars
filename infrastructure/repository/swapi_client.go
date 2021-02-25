package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type StarWarsClient struct {
	baseUrl string
}

func NewStarWarsClient(baseUrl string) *StarWarsClient {
	return &StarWarsClient{
		baseUrl: baseUrl,
	}
}

func (r *StarWarsClient) Search(name string) (int, error) {
	resp, err := http.Get(r.baseUrl + "?search=" + name)
	log.Printf("SWAPI Status: %s\n", resp.Status)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
		return 0, err
	}
	result := data["results"].([]interface{})[0]
	films := result.(map[string]interface{})["films"]
	return len(films.([]interface{})), nil
}
