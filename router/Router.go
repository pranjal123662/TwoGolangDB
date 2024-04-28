package router

import (
	"TwoDB/helper"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// var wg sync.WaitGroup
	router := mux.NewRouter()
	router.HandleFunc("/mergeTwoDB", helper.MergeTwoDataBase)
	return router
}
