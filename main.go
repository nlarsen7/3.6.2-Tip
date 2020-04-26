package main

import "fmt"
func cost(total float64, tip float64)float64{
return total*tip
}
func main() {
 total:=754.0
 tip:=0.20
 fmt.Println(cost(total, tip),"is the amount for your tip")
}