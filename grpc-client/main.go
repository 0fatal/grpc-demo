package main

import "grpc-client/router"

func main() {
	r := router.RouterInit()
	err := r.Run(":8899")
	if err != nil {
		return
	}
}