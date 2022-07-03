package main

import (
	server0 "courseselection/kitex_gen/Server/service"
	"courseselection/sqlcontroller"
	"log"
)

func main() {
	err := sqlcontroller.Init()

	if err != nil {
		log.Println(err.Error())
	}

	svr := server0.NewServer(new(ServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
