package cloudinary

import (
	"flag"
	"log"
	"os"
	"sync"
)

type singleton struct {
}

var instance *Service
var once sync.Once

var accountKey = os.Getenv("CLOUDINARY_API_KEY")
var secretKey = os.Getenv("CLOUDINARY_API_SECRET")
var cloudName = os.Getenv("CLOUDINARY_CLOUD_NAME")

func GetService() *Service {
	flag.Parse()
	endpoint := "cloudinary://" + accountKey + ":" + secretKey + "@" + cloudName
	once.Do(func() {
		var err error
		instance, err = Dial(endpoint)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
