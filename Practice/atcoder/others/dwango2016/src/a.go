// なんじゃこりゃ
// 2525...を&&して愚直で割った商の和
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	ans := 0
	ans += N / 25
	// よくよくかんがえたら以下いらねぇ
	// for i := 2525; i <= N; i += 2525 {
	// 	if N%i == 0 && N%25 != 0 {
	// 		ans += 1
	// 	}
	// }
	// for i := 252525; i <= N; i += 252525 {
	// 	if N%i == 0 && N%25 != 0 && N%2525 != 0 {
	// 		ans += 1
	// 	}
	// }

	// for i := 25252525; i <= N; i += 25252525 {
	// 	if N%i == 0 && N%25 != 0 && N%2525 != 0 && N%252525 != 0 {
	// 		ans += 1
	// 	}
	// }

	// for i := 2525252525; i <= N; i += 2525252525 {
	// 	if N%i == 0 && N%25 != 0 && N%2525 != 0 && N%252525 != 0 && N%25252525 != 0 {
	// 		ans += 1
	// 	}
	// 	fmt.Println(ans)
	// }

	fmt.Println(ans)
}
