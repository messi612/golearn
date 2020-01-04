package main

import (
	mysort "github.com/messi612/golearn/msort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error Happend:Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return values, nil
}

func main() {
	var infile = flag.String("i", "unsorted.dat", "File contains values for sorting")
	var outfile = flag.String("o", "console", "Output file contains sorted values")
	var algorithm = flag.String("s", "qsort", "Sort algorithm")

	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("values for sorting data:")
	for _, value := range values {
		fmt.Print(value, " ")
	}
	fmt.Println()
	switch *algorithm {
	case "qsort":
		mysort.QuickSort(values)
	case "isort":
		mysort.InsertionSort(values)
	default:
		mysort.BubbleSort(values)
	}
	if *outfile == "console" {
		fmt.Println("values for sorted data:")
		for _, value := range values {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
	writeValues(values, *outfile)
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create output file", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
