package routing

import (
	"practical/controler"
	"practical/utils"

	"github.com/gorilla/mux"
)

//BookRouting user routing
func BookWordsRouting(r *mux.Router) {
	r.HandleFunc(utils.BOOKAPI, controler.BooksWordsRepeating).Methods("POST")
}
