package main

import (
	"os"
)

func DbName(DbName string) string {
	os.Setenv(DbName, "mongodb+srv://nexta:foobar@cluster0.h9grc.mongodb.net/zuriChat?retryWrites=true&w=majority")

	// return the env variable using os package
	return os.Getenv(DbName)
}
