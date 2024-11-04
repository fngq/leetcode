package main 

import (
	"testing"
	"log"
)

func Test_t(t *testing.T){
	n1:=[]int{1,3}
	n2 := []int{2}
	n := findMedianSortedArrays(n1,n2)
	log.Println("expect",2,"get",n)
	log.Println()
	n1 = []int{1,2}
	n2 = []int{3,4}
	n = findMedianSortedArrays(n1,n2)
	log.Println("expect",2,"get",n)
}