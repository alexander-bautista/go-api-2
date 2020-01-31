package comics

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	publicKey  = "ea9048e277273d185472a588dd1f1537"
	privateKey = "bbb4d6492f11babda02d52e71d3d0af9a9f20d8d"
)

// ComicsResponse : represents a single database item
type ComicsResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status,omitempty"`
	Data   ComicDataReponse `json:"data"`
}

type ComicDataReponse struct {
	Total   int     `json:"total"`
	Count   int     `json:"count"`
	Results []Comic `json:"results"`
}

// GetComics : get comics by filter
func GetComics(titleStartsWith, dateRange string) (comicResponse ComicsResponse) {

	ts, hash := md5Hash()

	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics?ts=%s&apikey=%s&hash=%s&titleStartsWith=%s&dateRange=%s", ts, publicKey, hash, titleStartsWith, dateRange)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	json.Unmarshal([]byte(body), &comicResponse)

	return
}

// GetOneComicById : get one commic by id
func GetOneComicById(id string) (comicResponse ComicsResponse) {
	ts, hash := md5Hash()

	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/comics/%s?ts=%s&apikey=%s&hash=%s", id, ts, publicKey, hash)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	json.Unmarshal([]byte(body), &comicResponse)

	return
}

func md5Hash() (string, string) {
	t := strconv.Itoa(rand.Intn(1000))
	return t, fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}
