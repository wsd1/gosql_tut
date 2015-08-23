package main

import (
	"tuts_gosql/app"
	"tuts_gosql/app/model"

	"fmt"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func main() {

	app.Init()
	model.Init()

	word := model.WikiM.GetWikiwordByWord("ooxx")

	if nil != word {

		spew.Dump(word)
		fmt.Println("Content:", string(word.Content))
		fmt.Println("Created:", time.Unix(int64(word.Created), 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Println("Modified:", time.Unix(int64(word.Modified), 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Println("Visited:", time.Unix(int64(word.Visited), 0).Format("2006-01-02 03:04:05 PM"))

	} else {
		fmt.Println("Find nothing")
	}
}
