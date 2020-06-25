package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	r := strings.Fields(s)
	for _, v := range r {
		_,ok:= m[v]
		if ok {
		   m[v] = m[v] +1
		}else{
		   m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

