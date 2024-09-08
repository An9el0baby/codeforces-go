// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1864/D
// https://codeforces.com/problemset/status/1864/problem/D
func Test_cf1864D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
5
00100
01110
11111
11111
11111
3
100
110
110
6
010101
111101
011110
000000
111010
001110`,
			`1
2
15`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1864D)
}