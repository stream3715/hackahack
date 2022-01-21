package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type ApiResponse struct {
	URL               string `json:"url"`
	ArchivedSnapshots struct {
		Closest struct {
			Status    string `json:"status"`
			Available bool   `json:"available"`
			URL       string `json:"url"`
			Timestamp string `json:"timestamp"`
		} `json:"closest"`
	} `json:"archived_snapshots"`
}

func RequestArchiveUrl(path string) (string, int, error) {
	url := "https://archive.org/wayback/available"
	fakeDomain := os.Getenv("FAKE_DOMAIN")
	if fakeDomain == "" {
		log.Fatal("Fake domain not specified")
		return "", 500, errors.New("FAKE_DOMAIN_NULL")
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return "", 500, err
	}

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("url", "https://"+fakeDomain+path)
	request.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", response.StatusCode, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "", 500, err
	}
	data := new(ApiResponse)

	if err := json.Unmarshal(body, data); err != nil {
		return "", 500, err
	} else if !(data.ArchivedSnapshots.Closest.Available) {
		return "", 404, errors.New("ARCHIVE_NOT_FOUND")
	}

	addIf_ := regexp.MustCompile("/https://" + fakeDomain)
	archiveUrl := addIf_.ReplaceAllString(data.ArchivedSnapshots.Closest.URL, "if_/https://"+fakeDomain)
	return archiveUrl, response.StatusCode, nil
}
