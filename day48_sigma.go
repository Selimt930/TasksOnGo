//https://www.hackerrank.com/challenges/split-number/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

var reg = regexp.MustCompile(`([0-9]+)[\-\ ]([0-9]+)[\-\ ]([0-9]+)`)
var n int

func main() {
    fmt.Scanf("%d", &n)
    reader := bufio.NewScanner(os.Stdin)

    for ; n > 0; n-- {
        reader.Scan()

        group := reg.FindStringSubmatch(reader.Text())

        fmt.Printf("CountryCode=%s,LocalAreaCode=%s,Number=%s\n", group[1], group[2],group[3])
    }
}