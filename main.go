package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("path/to/db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	data, err := db.Get([]byte("key"), nil)
	log.Println(err)
	log.Println(data)
	err = db.Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err = db.Get([]byte("key"), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	err = db.Delete([]byte("key"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err = db.Get([]byte("key"), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(err)
}
