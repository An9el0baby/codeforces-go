// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, gcdValues, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-418/problems/sorted-gcd-pair-queries/
// https://leetcode.cn/problems/sorted-gcd-pair-queries/