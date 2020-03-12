package models

import "log"

// Person ...
type Person struct {
	ID        int64  `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

func InsertPerson(first, last string) error {
	rs, err := DB.Exec("insert into person(first_name, last_name) value(?,?)", first, last)
	if err != nil {
		log.Printf("insert error, Failed code: %#v", err)
		return err
	}
	log.Printf("Type: %T, value: %#v", rs, rs)
	return nil
}

func QueryPersonOne(id int) (Person, error) {
	var person Person
	// DB get 和 select 区别？
	err := DB.Get(&person, "select * from person where id=?", id)
	if err != nil {
		log.Printf("select one err, Failed code: %#v", err)
	}
	return person, err
}

func QueryPersonAll() ([]Person, error) {
	persons := make([]Person, 0)
	err := DB.Select(&persons, "select * from person")
	if err != nil {
		log.Printf("select err, Failed code: %#v", err)
	}
	return persons, err
}

// 写成 结构体的方法的方式也不错
// https://gist.github.com/rsj217/26492af115a083876570f003c64df118
/*
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(40) NOT NULL DEFAULT '',
  `last_name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/
