package helper

import (
	"TwoDB/controller"
	"TwoDB/model"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var response model.ResponseData

func MergeTwoDataBase(w http.ResponseWriter, r *http.Request) {
	var result model.UserData
	_ = json.NewDecoder(r.Body).Decode(&result)
	fmt.Println(result)
	if controller.FetchFromLoginDB(result.Number) {
		wg.Add(2)
		fetchofResult1 := make(chan bool, 1)
		go func() {
			defer wg.Done()
			result2 := controller.FetchFromLoginDB(result.Number)
			fetchofResult1 <- result2

		}()
		fetch1 := <-fetchofResult1
		fetchofResult2 := make(chan bool, 1)
		go func() {
			defer wg.Done()
			result1 := controller.FetchFromUserDataDB(result.Name)
			fmt.Println(result1)
			fetchofResult2 <- result1
		}()
		fetch2 := <-fetchofResult2
		wg.Wait()
		if fetch1 && fetch2 {
			response = model.ResponseData{Code: "200", DataCookie: &model.UserData{Name: result.Name, Number: result.Number}}
		}
	} else {
		insertmakechannel1 := make(chan bool)
		insertmakechannel2 := make(chan bool)
		start := time.Now()
		go func() {
			result1 := controller.InsertIntoLoginBucket(result.Number)
			insertmakechannel1 <- result1
		}()
		insert1 := <-insertmakechannel1
		go func() {
			result2 := controller.InsertIntoUserDataBucket(result.Name)
			insertmakechannel2 <- result2
		}()
		insert2 := <-insertmakechannel2
		// start := time.Now()
		// insert1 := controller.InsertIntoLoginBucket(result.Number)
		// insert2 := controller.InsertIntoUserDataBucket(result.Name)

		// fmt.Println(time.Now())
		// fmt.Println(insert1, insert2)
		if insert1 && insert2 {
			response = model.ResponseData{Code: "200", DataCookie: &model.UserData{Name: result.Name, Number: result.Number}}
			end := time.Now()
			fmt.Println("duration", end.Sub(start))
		}
	}

	json.NewEncoder(w).Encode(response)
}
