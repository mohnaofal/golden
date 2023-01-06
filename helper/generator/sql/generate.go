package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// GenerateSql
func GenerateSql(name string) {
	f, err := os.Create("./misc/sql/" + name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Naming file invalid")
	}

	// prefix
	prefix := strconv.Itoa(int(time.Now().UnixMilli()))
	// filename
	fileName := strings.Join(strings.Split(args[1], " "), "_")

	// filename down
	fileNameDown := prefix + "_" + fileName + ".down.sql"
	// filename up
	fileNameUp := prefix + "_" + fileName + ".up.sql"

	// generate file down sql
	GenerateSql(fileNameDown)
	// generate file up sql
	GenerateSql(fileNameUp)
}
