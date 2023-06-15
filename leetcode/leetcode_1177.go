package leetcode

import "math/bits"

//2023-06-15
//关键知识点：前缀和，bits包二进制操作
//https://leetcode.cn/problems/can-make-palindrome-from-substring/submissions/

func canMakePaliQueries1(s string, queries [][]int) []bool {
	sum := make([][26]int, len(s)+1)
	for i, c := range s {
		sum[i+1] = sum[i]
		sum[i+1][c-'a']++
	}
	ans := []bool{}
	for _, q := range queries {
		left, right, k, m := q[0], q[1], q[2], 0
		for j := 0; j < 26; j++ {
			m = m + (sum[right+1][j]-sum[left][j])%2
		}
		ans = append(ans, m/2 <= k)
	}
	return ans
}

func canMakePaliQueries2(s string, queries [][]int) []bool {
	sum := make([]uint32, len(s)+1)
	for i, c := range s {
		sum[i+1] = sum[i]
		sum[i+1] = sum[i+1] ^ (uint32(1) << (c - 'a'))
	}
	ans := []bool{}
	for _, q := range queries {
		left, right, k := q[0], q[1], q[2]
		m := bits.OnesCount32(sum[right+1] ^ sum[left])
		ans = append(ans, m/2 <= k)
	}
	return ans
}
