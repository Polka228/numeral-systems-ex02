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
		num int
		toBase int
		want string
	}
	{"simple",num: 123, toBase: 10, want: "123"},
	{"",num: 61, toBase: 2, want: "111101"},
	{"to letters",num: 43981, toBase: 16, want: "ABCD"},
	{"",num: 123, toBase: 10, want: "123"},
	{"empty",num: , toBase: 10, want: ""},
	{"zero",num: 0, toBase: 10, want: "0"},
	{"negative",num: -123, toBase: 10, want: ""},
	{
		t.Run(name, func(t *testing.T) {
			gotL, gotC := SomeFuncReturningLaureatesAndCounters(tc.num,tc.toBase)
			sort.Strings(gotL)
			if reflect.DeepEqual(gotL, tc.wantL) || gotC != tc.wantC {
				t.Errorf("got = (%v, %v), want = (%v, %v)", gotL, gotC, tc.wantL, tc.wantC)
			}
		})
	}
}