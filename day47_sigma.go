//https://www.hackerrank.com/challenges/utopian-identification-number/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

var rg = regexp.MustCompile(`^[a-z]{0,3}[0-9]{2,8}[A-Z]{3,}`)

func main() {
    var n int
    fmt.Scanf("%d", &n)
    scanner := bufio.NewScanner(os.Stdin)

    for ; n > 0 && scanner.Scan(); n-- {
        if rg.Match(scanner.Bytes()) {
            fmt.Println("VALID")
        } else {
            fmt.Println("INVALID")
        }
    }

}