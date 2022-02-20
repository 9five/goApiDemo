package controller

import (
	"net/http"

	"portfolio/services"

	"github.com/gorilla/mux"
)

type (
	Todo struct {
		Id   int64
		Item string
	}

	ApiResponse struct {
		ResultCode    string
		ResultMessage interface{}
	}
)

var TodoList []Todo

func PostDemo(w http.ResponseWriter, r *http.Request) {
	/*	body process
		defer r.Body.Close()
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
		if err != nil {
			fmt.Println(err.Error())
		}
	*/

	/*	form data process
		r.Form["username"]
		r.Form["password"]
	*/
}

func GetDemo(w http.ResponseWriter, r *http.Request) {
	//	get query variable
	vars := mux.Vars(r)
	queryV1 := vars["nickname"]

	response := ApiResponse{"200", "nickname:" + queryV1}

	services.ResponseWithJson(w, http.StatusOK, response)
}

func PutDemo(w http.ResponseWriter, r *http.Request) {

}

func DeleteDemo(w http.ResponseWriter, r *http.Request) {

}
