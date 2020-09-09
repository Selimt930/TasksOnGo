//https://www.hackerrank.com/challenges/mini-max-sum/problem
package main

    import (
        "fmt"
        "sort"
    )

    func main(){

        array := make([]int, 5)
        fmt.Scan(&array[0],&array[1],&array[2],&array[3],&array[4]) 
        sort.Ints(array) 
        min := array[0]+array[1]+array[2]+array[3]
        max := array[1]+array[2]+array[3]+array[4]
        fmt.Println(min,max)
    }