package DB

type LinkStruct struct {
	Link         string `bson: "link"`
	ShortendLink string `bson: "shortendLink"`
}
