package youtube

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SetYoutubeConfig() {
	file, err := os.Open("ytkey.txt")
	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadAll(file)
	fmt.Printf("Key read: %s", key)

	// TODO: authenticate youtube API
}
