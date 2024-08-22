// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/716/B
// https://codeforces.com/problemset/status/716/problem/B
func Test_cf716B(t *testing.T) {
	testCases := [][2]string{
		{
			`ABC??FGHIJK???OPQR?TUVWXY?`,
			`ABCDEFGHIJKLMNOPQRZTUVWXYS`,
		},
		{
			`WELCOMETOCODEFORCESROUNDTHREEHUNDREDANDSEVENTYTWO`,
			`-1`,
		},
		{
			`??????????????????????????`,
			`MNBVCXZLKJHGFDSAQPWOEIRUYT`,
		},
		{
			`AABCDEFGHIJKLMNOPQRSTUVW??M`,
			`-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf716B)
}