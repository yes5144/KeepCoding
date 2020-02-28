package models

import (
	"log"
)

// User ...
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// SelectUserAll ...
func SelectUserAll() ([]User, error) {
	sqlStr := "select * from user "
	users := make([]User, 0)
	err := DB.Select(&users, sqlStr)
	if err != nil {
		log.Printf("select 查询失败, Failed code: %#v", err)
	}
	return users, err
}

// SelectUserOne ...
func SelectUserOne(id int64) (User, error) {
	sqlStr := "select * from user where id=? limit 1"
	u := User{}
	err := DB.Get(&u, sqlStr, id)
	if err != nil {
		log.Printf("select 查询失败, Failed code: %#v", err)
	}
	return u, err
}

// InsertUser ...
func InsertUser(mod *User) error {
	_, err := DB.Exec("insert into user (name,age) values(?,?)", mod.Name, mod.Age)
	if err != nil {
		log.Printf("insert User error, Failed code: %#v", err)
	}
	return err
}

// UpdateUser ...
func UpdateUser(mod *User) error {
	_, err := DB.Exec("update user set name=?,age=?", mod.Name, mod.Age)
	if err != nil {
		log.Printf("update user error, Failed code: %#v", err)
	}
	return err
}
