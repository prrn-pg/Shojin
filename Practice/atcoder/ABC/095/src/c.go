package main
import (
	"fmt"
)

func main(){
	// a+b<cならcを選んだほうが得になる?
	// a/2+b/2=cなので駄目だったわ えっとね
	// a+b = 2c
	// 出力例1はa+b < 2cなのでcをx,y=3,2の小さい方に合わせて ()+（x,yのうち小さい方*2*c） = x+4c = 1*1500 + 4*1600 が最小。a,b,c = 1500, 2000, 1600
	// 出力例2はa+b > 2cなのでcをx,y=3,2の大きい方に合わせて3x+2y = 3*1500 + 2*2000 が最小。a,b,c = 1500, 2000, 1900
	// 出力例3はa+b < 2cなのでcをx,y=90000, 100000の大きい方に合わせてみると、a,b,c = 1500, 2000, 500なので、c=100000*500=5千万, 2c = 一億

	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)
	if a+b<2*c{
		min := x
		if x>y{
			min = y
		}

	}
}