package main

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Id int
	Name string
	CreatedAt time.Time
}

func init() {
	orm.RegisterDataBase("default","mysql","root:channel@tcp(192.168.204.222:3306)/beegodemo?charset=utf8",30)
	// create new table
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",true,true)
}

func main(){
	orm.Debug=true
	o:=orm.NewOrm()
	user:=User{Name: "wh",CreatedAt: time.Now()}

	// insert
	id,err:=o.Insert(&user)
	if err != nil {
		log.Fatalf("beego orm insert error, Failed code %#v",err)
	}
	logs.Info("hhh",id)

	// update
	user.Name="astaxie"
	num,err:=o.Update(&user)
	fmt.Printf("num: %d, Err: %#v",num,err)

	// read one
	u:=User{Id: user.Id}
	err =o.Read(&u)
	log.Printf("value: %#v type: %T\n", err,err)
	log.Println(u)

	// delete
	// num,err = o.Delete(&u)
	// log.Printf("value: %#v type: %T\n", num,err)
	

}
