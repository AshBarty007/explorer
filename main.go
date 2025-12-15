package main

import (
	"blockchain_services/config"
	bshttp "blockchain_services/http"
	bsdb "blockchain_services/postgres"
	"blockchain_services/services"
)

func main() {
	err := bsdb.InitPgConn()
	if err != nil {
		panic(err)
	}

	go services.Sync(config.TestUrl)

	bshttp.StartHttp()
}
