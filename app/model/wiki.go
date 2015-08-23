package model

import (
	"log"
	"tuts_gosql/app"
)

//sql tutorial can be found at:
//	http://segmentfault.com/a/1190000003036452
//	http://go-database-sql.org/retrieving.html

type Wikiwordstruct struct {
	Word        string
	Content     []byte
	Compression bool
	Encryption  bool

	Created  float64
	Modified float64
	Visited  float64

	Readonly bool
	//	Metadataprocessed     int
	//	Presentationdatablock string
}

type WikiwordModel struct {
	wikiwordsCache map[string]*Wikiwordstruct
}

func (this *WikiwordModel) cacheWikiword(words ...*Wikiwordstruct) {
	for _, w := range words {
		if w == nil {
			return
		}
		this.wikiwordsCache[w.Word] = w
	}
}

// get one user by id.
// if no cached, query from db and cache it.
func (this *WikiwordModel) GetWikiwordByWord(w string) *Wikiwordstruct {

	//cache check
	if this.wikiwordsCache[w] != nil {
		return this.wikiwordsCache[w]
	}

	//sql retrive
	word := new(Wikiwordstruct)

	row := app.SqlDb.QueryRow("SELECT word,content,compression,encryption,created,modified,visited,readonly FROM wikiwordcontent WHERE word=?", w)
	err := row.Scan(&word.Word, &word.Content, &word.Compression, &word.Encryption, &word.Created, &word.Modified, &word.Visited, &word.Readonly)
	if err != nil {
		log.Println("Try query: "+w, "get err:"+err)
	}

	// cache it
	if word.Word == w {
		this.cacheWikiword(word)
	} else {
		word = nil
	}

	return word
}

func (this *WikiwordModel) Reset() {
	this.wikiwordsCache = make(map[string]*Wikiwordstruct)

	//	this.GetAllUser()
}

func NewWikiModel() *WikiwordModel {
	wikiM := new(WikiwordModel)
	wikiM.Reset()
	return wikiM
}
