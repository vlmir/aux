package main

import "testing"
import "sync"

type t1 struct {
	arg1 []float64
	arg2 []float64
	val float64
}
var t1s = []t1{
	{ []float64{0.9, 1.1}, []float64{1.1, 0.9}, -1 },
	{ []float64{1.1, 0.9}, []float64{1.1, 0.9}, 1 },
}
func TestCorrelation(t *testing.T) {
	for _, t1 := range t1s {
		v, _ := Correlation(t1.arg1, t1.arg2)
		if v != t1.val {
			t.Error(
			"For", t1.arg1, t1.arg2,
			"expected", t1.val,
			"got", v,
			)
		}
	}
}

type t2 struct {
	arg []string
	val float64
}
var t2s = []t2{
	{ []string{"1.0", "-1.0"}, 0.0 },
	{ []string{"1.0", "1.0"}, 2.0 },
}
func TestXfloats(t *testing.T) {
	for _, t2 := range t2s {
		vec, _ := Xfloats(t2.arg)
		v := 0.0
		for _, val := range vec {
			v += val
		}
		if v != t2.val {
			t.Error(
			"For", t2.arg,
			"expected", t2.val,
			"got", v,
			)
		}
	}
}

type t3 struct {
	arg1 []string
	arg2 int
}
func Tln2ch(t *testing.T) {
	var t3s = []t3{
		{[]string{"lbl1", "1.0", "-1.0"}, 1},
		{[]string{"lbl2", "2.0", "-2.0"}, 2},
	}
	wg := new(sync.WaitGroup) // 1. initiation , pointer
	ch := make(chan Dataln)
	cnt := 0
	for _, t3 := range t3s {
		ln2ch(t3.arg1, t3.arg2, wg, ch)
		cnt += 1
	}
	v := len(ch)
	if v != cnt {
		t.Error(
		"expected", 2,
		"got", v,
		)
	}
}

type t4 struct {
	arg1 Dataln
	arg2 Dataln
}
func Tcr2ch(t *testing.T) {
	var ln1 = Dataln{1, "lbl1", []float64{0.1, -0.1}}
	var ln2 = Dataln{2, "lbl2", []float64{0.2, -0.2}}
	var t4s = []t4{
		{ ln1, ln2},
		{ ln2, ln1 },
	}
	wg := new(sync.WaitGroup) // 1. initiation , pointer
	ch := make(chan string)
	abs := false
	aP := &abs
	cnt := 0
	for _, t4 := range t4s {
		cr2ch(t4.arg1, t4.arg2, wg, ch, aP)
		cnt += 1
	}
	v := len(ch)
	if v != cnt {
		t.Error(
		"expected", 2,
		"got", v,
		)
	}
}

