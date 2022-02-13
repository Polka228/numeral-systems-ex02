package main

import (
	"testing"
)

func testCaseError(t *testing.T) {
	t.Error("forced error")
}

func testCaseOK(t *testing.T) {
	for _, tc := range []struct{
		name string
		num string
		fromBase int
		want int
	}
	{"simple",num: "123", fromBase: 10, want: 123},
	{"to 2",num: "111101", fromBase: 2, want: 61},
	{"letters",num: "ABCD", fromBase: 16, want: 43981},
	{"empty",num: "123", fromBase: 10, want: 123},
	{"negative",num: "-123", fromBase: 10, want: },
	{"zero",num: "0", fromBase: 10, want: 0},
	{
		t.Run(name, func(t *testing.T) {
			gotL, gotC := SomeFuncReturningLaureatesAndCounters(tc.num,tc.fromBase)
			sort.Strings(gotL)
			if reflect.DeepEqual(gotL, tc.want){
				t.Errorf("got = (%v, %v), want = (%v, %v)", gotL, gotC, tc.want)
			}
		})
	}
}