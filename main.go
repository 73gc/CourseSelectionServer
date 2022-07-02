package main

import (
	server0 "courseselection/kitex_gen/Server/adminservice"
	"log"
)

func main() {
	svr := server0.NewServer(new(AdminServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
