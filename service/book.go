package service

import (
	"database/sql"
	"log"
	"gogo/models"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"reflect"
	"fmt"
	"strings"
)

func GetBook() []models.Book {
	db := getInstance()
	defer db.Close()
	books := make([]models.Book, 0)
	rows, err := db.Query("select * from books ")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next(){
		var book models.Book
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price,&book.Create_time, &book.Update_time)
		books = append(books, book)
	}

	return books
}

func UpdateBookById(id int, data models.Book ) interface{} {
	db := getInstance()
	defer db.Close()
	data.Update_time = time.Now()
	stmt , err := db.Prepare("update books set name = ? ,price = ? where id = ?")
	ret, err := stmt.Exec(data.Name, data.Price, id)
	row, _ := ret.RowsAffected()
	if err != nil{
		log.Fatalln(err)
	}

	return row
}

func DeleteBook(book_id int) (interface{}, error) {

	if book_id <= 0{
		return 0, nil
	}

	db := getInstance()
	defer db.Close()
	ret, err := db.Exec("DELETE from books where id = ?", book_id)
	if err != nil{
		log.Fatalln(err)
	}
	row, _:= ret.RowsAffected()
	return row, nil
}

func InsertBook(table string, data interface{}) (int64, error) {
	sql, binds := InsertSQL(table, data)
	dbInstance := getInstance()

	ret, err := dbInstance.Exec(sql, binds...)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	insert_id, _ := ret.LastInsertId()
	return  insert_id, nil
}

func InsertSQL(table string, data interface{}) (string, []interface{}){
	v := reflect.Indirect(reflect.ValueOf(data))
	sqls := ""
	binds := []interface{}{}
	syna_kind := v.Kind()
	switch syna_kind {
	case reflect.Struct :
		 sqls, binds = singleInsertWithStruct(table, v)
	default:
		//
	}
	return sqls, binds
}

func singleInsertWithStruct(table string, v reflect.Value)(string, []interface{}) {
	fieldNum := v.NumField()
	columns := make([]string, 0, fieldNum)
	placeholoder := make([]string, 0, fieldNum)
	binds := make([]interface{}, 0, fieldNum)
	t := v.Type()
	for i := 0; i < fieldNum ; i ++ {
		column := t.Field(i).Tag.Get("db")
		if column == "-" {
			continue
		}

		columns = append(columns, column)
		placeholoder = append(placeholoder, "?")
		binds = append(binds, v.Field(i).Interface())
	}
	sqls := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholoder, ", "))
	return sqls, binds
}


func getInstance() (*sql.DB){
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gogo?")
	if err != nil {
		log.Fatalln(err)
	}

	if ping :=db.Ping(); ping != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	return db
}

