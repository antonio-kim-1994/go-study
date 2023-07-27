// 문제 설명
// 정수 num1과 num2가 매개변수로 주어질 때, num1을 num2로 나눈 값에 1,000을 곱한 후 정수 부분을 return 하도록 soltuion 함수를 완성해주세요.
//
// 제한사항
// 0 < num1 ≤ 100
// 0 < num2 ≤ 100
// 입출력 예
// num1	num2	result
// 3	2	1500
// 7	3	2333
// 1	16	62

package main

func twoNumberDivide(num1 int, num2 int) int {
	var result float64 //
	//func main() {
	//	fmt.Println(twoNumberDivide(1, 16))
	//}

	result = float64(num1) / float64(num2)
	return int(result * 1000)
}
