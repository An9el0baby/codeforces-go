// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/757/E
// https://codeforces.com/problemset/status/757/problem/E
func TestCF757E(t *testing.T) {
	testCases := [][2]string{
		{
			`5
0 30
1 25
3 65
2 5
4 48`,
			`8
5
25
4
630`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF757E)
}