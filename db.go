package main

import (
	"fmt"
	"os"
	// "regexp"
	// "time"
	
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id      int64
	Name    string
	// Created time.Time `xorm:"created"`
	// Updated time.Time `xorm:"updated"`
}

func main() {
	f := "conversion.db"
	os.Remove(f)

	orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL(true)

	err = orm.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = orm.Insert(&User{Id: 1, Name: "xlw"});
	_, err = orm.Insert(&User{Id: 0, Name: "Mercury"});
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []User
	// err = orm.Table("user").Cols("name").Where("name = ?", "Mercury").Find(&users);
	// fmt.Println(users);
	
	res, err := orm.Query("select name from user where(name = 'Mercury'))");
	fmt.Println(res);
	
	err = orm.Find(&users);
	fmt.Println(users);
	// users := make([]User, 0)
	// err = orm.Find(&users)
	// 
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(users)
}
