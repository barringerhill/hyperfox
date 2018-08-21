// 
// Sync insert.
//

package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func sync(engine *xorm.Engine) error {
	return engine.Sync(&Block{}, &Tx{});
}

func sqliteEngine() (*xorm.Engine, error) {
	f := "sync.db"
	return xorm.NewEngine("sqlite3", f)
}

func postgresEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", "dbname=xorm_test sslmode=disable")
}

type engineFunc func() (*xorm.Engine, error)

func read() {
	engine,err := sqliteEngine();
	var blocks []Block
	err = engine.Find(&blocks);
	if err != nil {
		fmt.Printf("{}", err);
	}
	fmt.Printf("{}", blocks)
}

func insert(block Block) {
	
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

		// _, err = Orm.Where("id > 0").Delete(&Block{})
		// if err != nil {
		// 	fmt.Println(err)
		// }		
		
		_, err = Orm.Insert(block)
		if err != nil {
			fmt.Println(err)
			return
		}

		isexist, err := Orm.IsTableExist("block")
		if err != nil {
			fmt.Println(err)
			return
		}
		if !isexist {
			fmt.Println("block is not exist")
			return
		}
	}
}

