package Controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type LinkStruct struct {
// 	Link         string
// 	ShortendLink string
// }

// var linkList []LinkStruct

// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	var link LinkStruct
// 	err := json.NewDecoder(r.Body).Decode(&link)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("link:", link.Link)
// 	for _, el := range linkList {
// 		if el.Link == link.Link {
// 			fmt.Fprintf(w, "Link is %s", el.ShortendLink)
// 			return
// 		}
// 	}

// 	link.ShortendLink = fmt.Sprintf("http://localhost:8000/%d", len(linkList)+1)
// 	log.Println(link)
// 	linkList = append(linkList, link)
// 	fmt.Fprintf(w, "%+v", link)
// 	return
// }