package main

import (
	bshttp "blockchain_services/http"
	bsdb "blockchain_services/postgres"
	"blockchain_services/services"
)

func main() {
	err := bsdb.InitPgConn()
	if err != nil {
		panic(err)
	}

	go services.Sync()

	bshttp.StartHttp()
}
