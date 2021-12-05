package main
import (
    y "fmt"
    sc "strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var a, b, c int
    y.Scanf("%d %d\n", &a, &b)
    var ab[50] int
    var l string
    for i:= 0; i<a;i++ {
    y.Scan(&l)
    tc, err := sc.Atoi(l)
    if err == nil {
    ab[i] = tc
    }
    }
    ik := 0
    for i := 0; i < a-1; i++ {
    for j:= i+1;j<a; j++ {
    if c=ab[i]+ab[j]; c % b == 0 {
    ik++
    }
    }
    }
    y.Println(ik)
    }
    
