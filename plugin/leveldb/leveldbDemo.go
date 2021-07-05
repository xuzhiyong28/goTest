package leveldb

import "github.com/syndtr/goleveldb/leveldb"

func Demo1(){
	db, err := leveldb.OpenFile("D:\\leveldb", nil)
	if err == nil {
		db.Put([]byte("key"), []byte("value"), nil)
	}
	defer db.Close()
}