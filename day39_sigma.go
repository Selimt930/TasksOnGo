//https://www.hackerrank.com/challenges/big-sorting/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {
	var res []string
	sort.Slice(unsorted, func(i, k int) bool {
		a, b := unsorted[i], unsorted[k]
		if diff := len(a) - len(b); diff != 0 {
			return diff < 0
		}

		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return true
			}
			if a[i] > b[i] {
				return false
			}
		}
		return false
	})

	for _, a := range unsorted {
		res = append(res, a)
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
