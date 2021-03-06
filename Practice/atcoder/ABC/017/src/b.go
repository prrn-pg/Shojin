// 判定するのがめんどくさいので、逆に考えて、最後が空文字になればいい

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
	defer out.Flush() // !!!!coution!!!! you must use Fprint(out, ) not Print()
	/* --- code ---*/
	x := next()
	// golang, []runeとstringでなんとかできちゃう感じ？
	// atcoderのgolangのバージョンが古いせい？？？REになるんだけど

	for len([]rune(x)) > 0 {
		size := len([]rune(x))
		if len([]rune(x)) >= 2 && string(x[size-2:]) == "ch" {
			x = string(x[:size-2])
		} else if string(x[size-1]) == "o" || string(x[size-1]) == "k" || string(x[size-1]) == "u" {
			x = string(x[:size-1])
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
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
