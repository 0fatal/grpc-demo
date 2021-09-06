package client

import "log"

var MyClient Client

func init() {
	err := MyClient.Init()
	if err != nil {
		log.Fatal(err)
	}
}