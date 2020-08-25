//https://www.hackerrank.com/challenges/extra-long-factorials/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)

// Complete the extraLongFactorials function below.
func factorial(x *big.Int) *big.Int {
    n := big.NewInt(1)
    if x.Cmp(big.NewInt(0)) == 0 {
        return n
    }
    return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := big.NewInt(nTemp)

    
    fmt.Print(factorial(n))
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
