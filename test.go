package main

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Wikiwordcontent struct {
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

//http://go-database-sql.org/retrieving.html

func main() {
	db, err := sql.Open("sqlite3", "wiki.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(30)

	defer db.Close()

	word := new(Wikiwordcontent)

	word.Word = "ooxx"

	fmt.Println("------ ------ delete ------ ------")

	{

		sql_del := "DELETE FROM wikiwordcontent WHERE word = ?"
		res_del, err_del := db.Exec(sql_del, word.Word)
		if err_del != nil {
			log.Fatal(err_del)
		}

		lastId, err_lastInsertId := res_del.LastInsertId()
		if err_lastInsertId != nil {
			log.Fatal(err_lastInsertId)
		}
		rowCnt, err_affected := res_del.RowsAffected()
		if err_affected != nil {
			log.Fatal(err_affected)
		}
		fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	}

	fmt.Println("------ ------ inserting ------ ------")

	word.Content = []byte("Go String与Byte切片之间的转换")
	word.Compression = false
	word.Encryption = false
	word.Created = float64(time.Now().Unix())
	word.Modified = float64(time.Now().Unix())
	word.Visited = float64(time.Now().Unix())
	word.Readonly = false
	sql_insert := "INSERT INTO wikiwordcontent(word,content,compression,encryption,created,modified,visited,readonly) "
	sql_insert += "VALUES (?,?,?,?,?,?,?,?)"
	res_insert, err_insert := db.Exec(sql_insert, word.Word, word.Content, word.Compression, word.Encryption, word.Created, word.Modified, word.Visited, word.Readonly)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	lastId, err_lastInsertId := res_insert.LastInsertId()
	if err_lastInsertId != nil {
		log.Fatal(err_lastInsertId)
	}
	rowCnt, err_affected := res_insert.RowsAffected()
	if err_affected != nil {
		log.Fatal(err_affected)
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	fmt.Println("------ ------ Query ------ ------")
	//single row query
	row := db.QueryRow("SELECT word,content,compression,encryption,created,modified,visited,readonly FROM wikiwordcontent WHERE word=?", word.Word)
	err = row.Scan(&word.Word, &word.Content, &word.Compression, &word.Encryption, &word.Created, &word.Modified, &word.Visited, &word.Readonly)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(word)

	fmt.Println("Content:%s", string(word.Content))
	fmt.Println("Created:%s", time.Unix(int64(word.Created), 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Println("Modified:%s", time.Unix(int64(word.Modified), 0).Format("2006-01-02 03:04:05 PM"))
	fmt.Println("Visited:%s", time.Unix(int64(word.Visited), 0).Format("2006-01-02 03:04:05 PM"))

	fmt.Println("------ ------ delete ------ ------")

	{

		sql_del := "DELETE FROM wikiwordcontent WHERE word = ?"
		res_del, err_del := db.Exec(sql_del, word.Word)
		if err_del != nil {
			log.Fatal(err_del)
		}

		lastId, err_lastInsertId := res_del.LastInsertId()
		if err_lastInsertId != nil {
			log.Fatal(err_lastInsertId)
		}
		rowCnt, err_affected := res_del.RowsAffected()
		if err_affected != nil {
			log.Fatal(err_affected)
		}
		fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	}

}
