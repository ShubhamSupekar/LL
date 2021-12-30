package main

import (
	"bufio"
	"fmt"
	"os"
)
type node struct{
	data string
	next *node
} 
var head *node 
var last *node 
func main(){

	//myname := []int{10,20,30,40,50}
	reader := bufio.NewReader(os.Stdin)
	for counter:=0;;counter++{
		 fmt.Print("Enter a number: ")
		 myname,_ := reader.ReadString('\n')
		for {
			n := node{data: myname,}
			if head==nil {
				head = &n
				last = head
				break	
			}
		last.next = &n
		last = last.next
		output()
		}
	}
}
func output(){
	for head !=nil{
		fmt.Println(head.data,head.next)
		head = head.next
	}
}



