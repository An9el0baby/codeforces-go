// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, closetTarget, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-325/problems/shortest-distance-to-target-string-in-a-circular-array/