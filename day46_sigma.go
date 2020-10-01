//https://www.hackerrank.com/challenges/saying-hi/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    var DD = "d"[0]

    for ; n > 0 && scanner.Scan(); n-- {
        inp := scanner.Text()
        lw := strings.ToLower(inp)

        if strings.HasPrefix(lw, "hi ") {
            if len(lw) >= 4 && lw[3] != DD {
                fmt.Println(inp)
            }
        }
    }
}