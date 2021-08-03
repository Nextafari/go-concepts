package main

import (
	"os"
)

func api() {
	os.Setenv("Db_name", "mongodb+srv://nexta:foobar@cluster0.h9grc.mongodb.net/zuriChat?retryWrites=true&w=majority")
}
