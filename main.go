package main

import (
	"fmt"
	"strings"
)

func main() {
  a:=[]string{"tomas beyne", "sara byene","joseph beyne"}
  
   
  for i:=0; i<len(a); i++{
	fmt.Println(strings.Split(a[i]," ")[0])
	  }
}