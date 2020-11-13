package main

import (
	"database/sql"
	"log"
	"time"

	"gopkg.in/gorp.v1"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	// db tag lets you specify the colum name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`
	Body    string `db:"article_body,size:1024"`
}

func newPost(title, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	// delete
	err := dbmap.TruncateTables()
	checkErr(err, "Truncatetable failed")

	// create two posts
	p1 := newPost("Go 1.1 released", "Lorem ipsum 1.1")
	p2 := newPost("Go 1.2 released", "Lorem ipsum 2.2")

	// insert rows
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "Insert failed")

	// use convenience selectint
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("row after inserting ", count)

}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "post_db.db")
	checkErr(err, "sql.Open failed")
	//
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("err, Failed code: %#v", err)
	}
}

// https://github.com/go-gorp/gorp
// install gcc
// https://jmeubank.github.io/tdm-gcc/
