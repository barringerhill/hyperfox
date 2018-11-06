package main;

import (
	"encoding/json"
	"strconv"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// db part
type Tx struct{
        Number   uint64 `gorm: "not null"`
        Hash     string `gorm: "not null"`
        Data     string `gorm: "not null; unique; index;"`
}

func (tx *Tx) to_json() string {
	data, err := json.Marshal(tx);
	assert(err);

	return string(data[:]);
}

type Allblue struct {}

func (a *Allblue) read(page int) []string {
	db, err := gorm.Open("postgres", "dbname=allblue sslmode=disable");
	assert(err); defer db.Close();
	
	var tx []Tx;	
	db.Raw("SELECT * FROM txes order by number offset " + strconv.Itoa(page * 10) + " limit 10").Scan(&tx);

	var res []string
	for _, i := range(tx) {
		res = append(res, i.to_json());
	}
	
	return res;
}

func (a *Allblue) search(ctx string, page int) []string {
	db, err := gorm.Open("postgres", "dbname=allblue sslmode=disable");
	assert(err); defer db.Close();
	
	var tx []Tx;
	db.Raw("SELECT * FROM txes where data ~* '" + ctx + "' " + strconv.Itoa(page * 10) + " limit 10").Scan(&tx);

	var res []string
	for _, i := range(tx) {
		res = append(res, i.Hash);
		res = append(res, i.Data);
		res = append(res, strconv.FormatUint(i.Number, 10))
	}
	
	return res
}
