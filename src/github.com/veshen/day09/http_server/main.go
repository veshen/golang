package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"net/http"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

type Book struct {
	id     int64
	title  string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User `pg:"rel:has-one"`
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}


func fn1(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello world"))
}

func main()  {

	fmt.Println("开始连接数据库")
	db := pg.Connect(&pg.Options{
		User: "postgres",
		Password: "test123456",
		Addr: "localhost:5432",
	})
	_, err := db.ExecContext()
	if err != nil {
		panic(err)
	}
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	book := new(Book)
	err = db.Model(book).Where("title = ?", "测试").Select()
	if err != nil {
       fmt.Println(err)
    }
	fmt.Println(book)
	//var book1 = &Book{
	//	id:    2,
	//	title: "22222",
	//}
	//_, err := db.Model(book1).Insert()
	//if err != nil {
	//	fmt.Println(err)
	//}

	http.HandleFunc("/api/1",fn1)
	err = http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		return 
	}
}
