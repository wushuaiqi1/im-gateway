package tests

import (
	"log"
	"testing"
)

func TestMapAdd(t *testing.T) {
	amap := map[int]int{}
	amap[1] = 2
	amap[2] = 1
	log.Println(amap)
	// panic: assignment to entry in nil map [recovered]
	//	panic: assignment to entry in nil map
	var aamap map[int]map[int]string = map[int]map[int]string{}
	log.Println(aamap == nil)
	log.Println(aamap[1])
	log.Println(aamap[2] == nil)
	//aamap[2] = "1"
}
