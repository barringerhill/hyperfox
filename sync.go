
package main

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// SyncUser2 describes a user
type SyncUser2 struct {
	Id      int64
	Name    string `xorm:"unique"`
	Age     int    `xorm:"index"`
	Title   string
	Address string
	Genre   string
	Area    string
	Date    int
}

func sync(engine *xorm.Engine) error {
	return engine.Sync(&SyncLoginInfo2{}, &SyncUser2{})
}

func sqliteEngine() (*xorm.Engine, error) {
	f := "sync.db"
	//os.Remove(f)

	return xorm.NewEngine("sqlite3", f)
}

func postgresEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", "dbname=xorm_test sslmode=disable")
}

type engineFunc func() (*xorm.Engine, error)

lfunc main() {
	
	engines := []engineFunc{postgresEngine, sqliteEngine}
	for _, enginefunc := range engines {
		Orm, err := enginefunc()
		fmt.Println("--------", Orm.DriverName(), "----------")
		if err != nil {
			fmt.Println(err)
			return
		}
		Orm.ShowSQL(true)
		err = sync(Orm)
		if err != nil {
			fmt.Println(err)
		}

		_, err = Orm.Where("id > 0").Delete(&SyncUser2{})
		if err != nil {
			fmt.Println(err)
		}

		user := &SyncUser2{
			Name:    "testsdf",
			Age:     15,
			Title:   "newsfds",
			Address: "fasfdsafdsaf",
			Genre:   "fsafd",
			Area:    "fafdsafd",
			Date:    1000,
		}
		_, err = Orm.Insert(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		isexist, err := Orm.IsTableExist("sync_user2")
		if err != nil {
			fmt.Println(err)
			return
		}
		if !isexist {
			fmt.Println("sync_user2 is not exist")
			return
		}
	}
}
