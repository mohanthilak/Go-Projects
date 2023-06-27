package Controllers

import (
	"encoding/json"
	"example/todo-list/Infra/DB"
	"fmt"
	"log"
	"net/http"
)

type ControllerStruct struct {
	DB *DB.LinkMongoWorker
}

func (CS *ControllerStruct) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var link LinkStruct
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("link:", link.Link)
	isThere := CS.DB.Findlink(link.Link)
	if isThere {
		fmt.Fprintf(w, "already there")
	} else {
		link.ShortendLink = fmt.Sprintf("http://localhost:8000/%d", len(linkList)+1)
		CS.DB.AddRecordToURLCol(link.Link, link.ShortendLink)
		log.Println(link)
		linkList = append(linkList, link)
		fmt.Fprintf(w, "%+v", link)

	}
}
