package main

import "fmt"

/*
- 문제 설명
첫 번째 분수의 분자와 분모를 뜻하는 numer1, denom1, 두 번째 분수의 분자와 분모를 뜻하는 numer2, denom2가 매개변수로 주어집니다.
두 분수를 더한 값을 기약 분수로 나타냈을 때 분자와 분모를 순서대로 담은 배열을 return 하도록 solution 함수를 완성해보세요.

- 제한사항
0 <numer1, denom1, numer2, denom2 < 1,000

- 입출력 예
numer1	denom1	numer2	denom2	result
1	2	3	4	[5, 4]
9	2	1	3	[29, 6]
*/

//func addFraction(numer1 int, denom1 int, numer2 int, denom2 int) []int {
//
//}

func main() {
	num := float64(1)/float64(2) + float64(3)/float64(4)
	fmt.Println(num / 100)
}
