package main

import (
	"fmt"
	"log"

	"github.com/lavishag4193/csv2jsonx/pkg/csv2jsonx"
)

func main() {
	result, err := csv2jsonx.ConvertFile("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}
