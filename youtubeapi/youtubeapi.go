package youtubeapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SetConfig() {
	file, err := os.Open("./youtubeapi/ytkey.txt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadAll(file)
	fmt.Printf("Key read: %s", key)

	// TODO: authenticate youtube API
}
