package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()
	/* --- code --- */
	// これワ―シャルフロイドして間に合うんか？
	// 毎回 時間と料金をそれぞれコストとして別にワ―シャルフロイドを作成する
	INF := pow(10, 16)
	for {
		n, m := nextInt(), nextInt()
		if n == 0 || m == 0 {
			return
		}
		cost_table, time_table := make([][]int, m), make([][]int, m)
		for i := 0; i < m; i++ {
			cost_table[i], time_table[i] = make([]int, m), make([]int, m)
			for j := 0; j < m; j++ {
				cost_table[i][j], time_table[i][j] = INF, INF
			}
		}
		for i := 0; i < n; i++ {
			a, b, c, t := nextInt()-1, nextInt()-1, nextInt(), nextInt()
			cost_table[a][b], time_table[a][b] = c, t
			cost_table[b][a], time_table[b][a] = c, t
		}

		// ワ―シャルフロイド
		for k := 0; k < m; k++ {
			for i := 0; i < m; i++ {
				if cost_table[i][k] == INF {
					continue
				}
				for j := 0; j < m; j++ {
					if cost_table[k][j] == INF {
						continue
					}
					cost_table[i][j] = min(cost_table[i][j], cost_table[i][k]+cost_table[k][j])
				}
			}
		}

		for k := 0; k < m; k++ {
			for i := 0; i < m; i++ {
				if time_table[i][k] == INF {
					continue
				}
				for j := 0; j < m; j++ {
					if time_table[k][j] == INF {
						continue
					}
					time_table[i][j] = min(time_table[i][j], time_table[i][k]+time_table[k][j])
				}
			}
		}

		k := nextInt()
		for i := 0; i < k; i++ {
			p, q, r := nextInt()-1, nextInt()-1, nextInt()
			if r == 0 {
				fmt.Fprintln(out, cost_table[p][q])
			} else {
				fmt.Fprintln(out, time_table[p][q])
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// -*-*-*-*-*-*-*-*-
// * I/O utilities *
// -*-*-*-*-*-*-*-*-

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func nextFloat64() float64 {
	a, _ := strconv.ParseFloat(next(), 64)
	return a
}

func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}
func nextFloats(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = nextFloat64()
	}
	return ret
}
func nextStrings(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = next()
	}
	return ret
}

func split(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

func PrintOut(src interface{}, joinner string) {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)
		for idx := 0; idx < s.Len(); idx++ {
			fmt.Fprintf(out, "%v", s.Index(idx))
			if idx < s.Len()-1 {
				fmt.Fprintf(out, "%s", joinner)
			} else {
				fmt.Fprintln(out)
			}
		}
	default:
		fmt.Fprintln(out, "fuck")
	}
}

// -*-*-*-*-*-*-*-*-
// * tool snippets *
// -*-*-*-*-*-*-*-*-
func duplicate2Int(base [][]int) (ret [][]int) {
	ret = make([][]int, len(base))
	for i, v := range base {
		ret[i] = append([]int{}, v...)
	}
	return
}

// -*-*-*-*-*-*-*-*-*-*-*-*-*-
// * Algorithms Utility Zone *
// -*-*-*-*-*-*-*-*-*-*-*-*-*-

// -*-*-*-*-*-*-*-
// * 1. nibutan  *
// -*-*-*-*-*-*-*-
func lower_bound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func upper_bound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if target < arr[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
// * math flavor typical theories *
// *-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
func gcd(a, b int) int {
	if a > b {
		return gcd(b, a)
	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func pow(a, b int) (ret int) {
	ret = a
	// 10^2 = 100ってことは10に10を1回掛けることだね
	// なので初期値を含めると上限b-1未満
	for i := 0; i < b-1; i++ {
		ret *= a
	}
	return
}

func powMod(n, m, mod int) (ret int) {
	ret = 1
	for m > 0 {
		if m&1 == 1 {
			ret *= n
			ret %= mod
		}
		n *= n
		n %= mod
		m >>= 1
	}
	return ret
}

func next_permutation(arr []int) func() []int {
	/*
		how to use it:
			this is a generator, so should be invoked such as below example.

			"""code"""
			np := next_permutation(arr)
			for{
				lis := np()
				if len(lis) == 0{
					break
				}
				fmt.Println(lis)
			}
			"""code end"""
	*/
	first := true
	ret := append([]int{}, arr...)
	_next_permutation := func() []int {
		if first {
			first = false
			return arr
		}
		n := len(ret)
		for i := n - 2; i >= 0; i-- {
			if ret[i] < ret[i+1] {
				j := n
				for {
					j -= 1
					if ret[i] < ret[j] {
						break
					}
				}
				ret[i], ret[j] = ret[j], ret[i]
				for k := n - 1; i < k; i, k = i+1, k-1 {
					ret[i+1], ret[k] = ret[k], ret[i+1]
				}
				return ret
			}
		}
		return []int{}
	}
	return _next_permutation
}
