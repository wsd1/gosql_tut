package main

import (
	"database/sql"
	"fmt"
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

	// Execute the query
	rows, err := db.Query("SELECT word,content,compression,encryption,created,modified,visited,readonly FROM wikiwordcontent LIMIT 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	/*
	   db.Query()表示向数据库发送一个query，defer rows.Close()非常重要，遍历rows使用rows.Next()，
	   把遍历到的数据存入变量使用rows.Scan(), 遍历完成后检查error。有几点需要注意：

	   检查遍历是否有error
	   结果集(rows)未关闭前，底层的连接处于繁忙状态。当遍历读到最后一条记录时，会发生一个内部EOF错误，自动
	   调用rows.Close()，但是如果提前退出循环，rows不会关闭，连接不会回到连接池中，连接也不会关闭。所以手
	   动关闭非常重要。rows.Close()可以多次调用，是无害操作。

	*/

	word := new(Wikiwordcontent)

	for rows.Next() {
		err := rows.Scan(&word.Word, &word.Content, &word.Compression, &word.Encryption, &word.Created, &word.Modified, &word.Visited, &word.Readonly)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("---word.Word---")
		fmt.Println(word.Word)
		fmt.Println("---word.Content---")
		fmt.Println(string(word.Content))
		fmt.Println("---word.Compression---")
		fmt.Println(word.Compression)
		fmt.Println("---word.Encryption---")
		fmt.Println(word.Encryption)

		fmt.Println("---word.Created---")
		fmt.Println(time.Unix(int64(word.Created), 0).Format("2006-01-02 03:04:05 PM"))

		fmt.Println("---word.Modified---")
		fmt.Println(time.Unix(int64(word.Modified), 0).Format("2006-01-02 03:04:05 PM"))

		fmt.Println("---word.Visited---")
		fmt.Println(word.Visited)

		fmt.Println("---word.Readonly---")
		fmt.Println(word.Readonly)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//single row query

	row := db.QueryRow("SELECT word,content,compression,encryption,created,modified,visited,readonly FROM wikiwordcontent WHERE word=?", "ExMemory")
	err = row.Scan(&word.Word, &word.Content, &word.Compression, &word.Encryption, &word.Created, &word.Modified, &word.Visited, &word.Readonly)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---word.Word---")
	fmt.Println(word.Word)
	fmt.Println("---word.Content---")
	fmt.Println(string(word.Content))
	fmt.Println("---word.Compression---")
	fmt.Println(word.Compression)
	fmt.Println("---word.Encryption---")
	fmt.Println(word.Encryption)

	fmt.Println("---word.Created---")
	fmt.Println(time.Unix(int64(word.Created), 0).Format("2006-01-02 03:04:05 PM"))

	fmt.Println("---word.Modified---")
	fmt.Println(time.Unix(int64(word.Modified), 0).Format("2006-01-02 03:04:05 PM"))

	fmt.Println("---word.Visited---")
	fmt.Println(word.Visited)

	fmt.Println("---word.Readonly---")
	fmt.Println(word.Readonly)

}
