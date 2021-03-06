package main

// 某DP, 再帰でやったらTLEになる原因がメモ化再帰しか浮かばないんだけどループでやるDPってやったことないので練習のため来ました
// もう初心にかえってナップザックをやる
// 前回メモ化再帰でやったので今回はループでやる

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
var n, w int
var values, weights []int

var dp [101][10101]int

func main() {
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!coution!!!! you must use Fprint(out, ) not Print()
	/* --- code ---*/
	n, w = nextInt(), nextInt()
	values, weights = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		vi, wi := nextInt(), nextInt()
		values[i] = vi
		weights[i] = wi
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ { // 未満じゃなくて以上なことに注意
			if j < weights[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = int(math.Max(float64(dp[i][j]), float64(dp[i][j-weights[i]]+values[i])))
			}
		}
	}
	fmt.Fprintln(out, dp[n][w])
}

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
