package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"fmt"
	"time"
	"testing-assigment/src"
)



func main() {
	flag.Parse()
	inputfile := flag.Arg(0)
	outputfile := "./example_output.csv"

	fmt.Println(inputfile)

	fileReader, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(fileReader)

	data, err := src.ReadData(csvReader)

	err = fileReader.Close()

	t := time.Now()
	fmt.Println(t)
	start := t.AddDate(-1, 0, 0)

	data, err = src.DateFilter(data, start)

	outputData := src.ProcessData(data)

	outfile, err := os.Create(outputfile)
	if err != nil {
		log.Fatal(err)
	}

	src.WriteOutputData(outfile, outputData)
	err = outfile.Close()

}
