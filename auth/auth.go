package auth

import (
	"encoding/json"
	"log"
	"os"
)

//Configuration struct reperesnts.
type Configuration struct {
	AWSAccessKeyID     string
	AWSSecretKey       string
	GoogleClientID     string
	GoogleClientSecret string
}

var Config Configuration

func LoadConfig() {

	var home string = os.Getenv("HOME")
	file, err := os.Open(home + "/gocloudconfig.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {

		Config.AWSAccessKeyID = os.Getenv("AWSAccessKeyID")
		Config.AWSSecretKey = os.Getenv("AWSSecretKey")

		if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {
			log.Fatalln("Cannot Get access key and secert key")
		}
	}
}
