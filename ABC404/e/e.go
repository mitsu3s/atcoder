// This is sample code.
// I did not answer the question.
// Please note that this is a sample code.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	// c[i] := 茶碗 i のラベル（1 ≤ i ≤ n-1）, c[0] は使わない
	c := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	// 初期豆がある茶碗のリスト
	var initial []int
	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		if a > 0 {
			initial = append(initial, i)
		}
	}
	sort.Ints(initial)

	// s_c[x] == true なら「x は既に操作リストに含まれ根側に繋がっている」
	s_c := make([]bool, n)
	s_c[0] = true
	// ops[x] == true なら「茶碗 x で操作を行う」
	ops := make([]bool, n)

	var connect func(x int)
	connect = func(x int) {
		if s_c[x] {
			return
		}
		// x から左へ c[x] 範囲を調べる
		l := x - c[x]
		if l < 0 {
			l = 0
		}
		// その区間にすでに s_c にあるノードがあれば，この x は繋がっている
		for y := l; y < x; y++ {
			if s_c[y] {
				return
			}
		}
		// まだ繋がっていない ⇒ 区間内から「より根側まで届く」ノード p を選ぶ
		// 評価基準は p - c[p] が最小（根側の端点が最も小さい）．
		// 同値なら p を大きい方を選ぶ（直前の端点を大きく保つことで，次の繋ぎやすさを確保）
		bestP := -1
		bestLP := math.MaxInt32
		if l < 1 {
			l = 1
		}
		for p := l; p < x; p++ {
			lp := p - c[p]
			if lp < bestLP || (lp == bestLP && p > bestP) {
				bestLP = lp
				bestP = p
			}
		}
		// 再帰的に p を繋いでから，自分も繋ぐ
		connect(bestP)
		s_c[bestP] = true
		ops[bestP] = true
	}

	// 各初期豆のある i について connect し，i 自身でも操作
	for _, i := range initial {
		connect(i)
		s_c[i] = true
		ops[i] = true
	}

	// ops[] の true の数が答え
	ans := 0
	for i := 1; i < n; i++ {
		if ops[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
