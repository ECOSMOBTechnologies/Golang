package controler

import (
	"net/http"
	"practical/db"
	"practical/utils"
)

//BooksWordsRepeating show get repetition words
func BooksWordsRepeating(w http.ResponseWriter, r *http.Request) {

	bookstr := r.FormValue(utils.BOOK_STR)

	if bookstr == "" {
		db.RespondJSON(r, w, utils.BOOKSTRINGINVALID, utils.STATUFAILS, nil)
		return
	}

	words := utils.GetRepetitionWords(bookstr)
	db.RespondJSON(r, w, utils.SUCCESS, utils.STATUSSUCESS, words)
}
