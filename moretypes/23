package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
 	m := make(map[string]int)
 	array_strings:=strings.Fields(s)
 	for _,v:=range array_strings {
 		if _,ok:=m[v]; ok {
			m[v]++
 		} else {
 			m[v]=1
 		} 
 	}
	return m
}

func main() {
	wc.Test(WordCount)
}
