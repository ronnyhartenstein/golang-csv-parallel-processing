package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	//"io"
	"os"
	"strconv"
	"strings"
)

var filename = flag.String("f", "REQUIRED", "target CSV file (tuncated if exists)")
var num = flag.Int("n", 1000, "rows to generate")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), "\n"))
	if *filename == "REQUIRED" {
		return
	}

	csvfile, err := os.Create(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	//fmt.Println(writer)
	for i := 0; i < *num; i++ {

		single := []string{strconv.Itoa(i), "bla", "fasel"}
		er := writer.Write(single)
		if er != nil {
			fmt.Println("error", er)
			return
		}
		//fmt.Print(".")
		if i%1000 == 0 {
			fmt.Printf("\r%d", i)
		}
		writer.Flush()
	}
}
