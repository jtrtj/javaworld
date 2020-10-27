package services

import 	(
	"net/url"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func Call() PlacesResponse {
	base, err := url.Parse("https://maps.googleapis.com/maps/api/place/")

	if err != nil {
		log.Fatalln(err)
	}	

	// Path params
	base.Path += "nearbysearch/json"

	// Query params
	params := url.Values{}
	params.Add("key", os.Getenv("GOOGLE_KEY"))
	params.Add("location", "39.7392,-104.9903" )
	params.Add("radius", "33000" )
	params.Add("type", "cafe" )
	base.RawQuery = params.Encode()

	resp, err := http.Get(base.String())

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var placesStruct PlacesResponse
	json.Unmarshal(bodyBytes, &placesStruct)
	return placesStruct
}
