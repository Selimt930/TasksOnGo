//https://www.hackerrank.com/challenges/hackerrank-language/problem?h_r=internal-search
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const ALL = "C:CPP:JAVA:PYTHON:PERL:PHP:RUBY:CSHARP:HASKELL:CLOJURE:BASH:SCALA:ERLANG:CLISP:LUA:BRAINFUCK:JAVASCRIPT:GO:D:OCAML:R:PASCAL:SBCL:DART:GROOVY:OBJECTIVEC"

func main() {
    Ml := make(map[string]bool)

    parts := strings.Split(ALL, ":")
    for _, s := range parts {
        Ml[s] = true
    }

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    for ; n > 0 && scanner.Scan(); n-- {
        line := scanner.Text()
        p := strings.Split(line, " ")

        if len(p) >= 2 {
            if _, contains := Ml[p[1]]; contains {
                fmt.Println("VALID")
            } else {
                fmt.Println("INVALID")
            }
        } else {
            fmt.Println("INVALID")
        }
    }
}