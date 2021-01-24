package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ReadURL(url string) {
	if xmlBytes, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &AllData)
	}
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
