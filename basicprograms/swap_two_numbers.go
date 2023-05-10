package basicprograms

import "fmt"

func SwapTwoNumbers(num1, num2 int) {
	temp := num1
	num1 = num2
	num2 = temp
	fmt.Println("after swapping", num1, num2)
}

func SwapTwoNumbersWithoutThirdVar(num1, num2 int) {
	num1 = num1 - num2
	num2 = num1 + num2
	num1 = num2 - num1
	fmt.Println("after swapping without third variable", num1, num2)
}

func SimpleMethodToSwapTwoNumbers(num1, num2 int) {
	num1, num2 = num2, num1
	fmt.Println("simple after swapping", num1, num2)
}
