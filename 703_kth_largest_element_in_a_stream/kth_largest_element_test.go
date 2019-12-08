package kth_largest_element

import (
	"log"
	"testing"
)

func TestKth(t *testing.T) {

	kth := Constructor(3, []int{4, 5, 8, 3})
	log.Printf("kth=%+v", kth)

	kth.Add(3)
	log.Printf("kth=%+v", kth)

	kth.Add(5)
	log.Printf("kth=%+v", kth)

}

func TestKth2(t *testing.T) {

	kth := Constructor(1, []int{})
	log.Printf("kth=%+v", kth)

	got := kth.Add(-3)
	log.Printf("kth=%+v, got=%d", kth, got)

	got = kth.Add(-2)
	log.Printf("kth=%+v, got=%d", kth, got)

}
