package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var filename = flag.String("f", "REQUIRED", "source CSV file")
var numChannels = flag.Int("c", 4, "num of parallel channels")

//var bufferedChannels = flag.Bool("b", false, "enable buffered channels")

func main() {
	start := time.Now()
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), "\n"))
	if *filename == "REQUIRED" {
		return
	}

	csvfile, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	i := 0
	for {
		i++
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		processData(record)
		fmt.Printf("\r%d %s  ..", i, record)
	}

	fmt.Printf("\n%2fs", time.Since(start).Seconds())

}

func processData([]string) {
	time.Sleep(10 * time.Millisecond)
}
