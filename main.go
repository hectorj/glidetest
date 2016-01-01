package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	ldb, err := leveldb.OpenFile("test.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	ldb.Close()
}
