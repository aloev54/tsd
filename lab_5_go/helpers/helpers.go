package helpers

import "math"

func checkPriority(element string) int {
	if element == "(" || element == ")" {
		return 5
	} else if element == "cos" || element == "sin" || element == "tg" {
		return 4
	} else if element == "^" {
		return 3
	} else if element == "*" || element == "/" {
		return 2
	} else if element == "+" || element == "-" {
		return 1
	} else {
		return 0
	}

}

func calculationBinary(operand string, nums ...int) int {
	if operand == "+" {
		return nums[0] + nums[1]
	} else if operand == "-" {
		return nums[0] - nums[1]
	} else if operand == "*" {
		return nums[0] * nums[1]
	} else if operand == "/" {
		return nums[0] / nums[1]
	} else if operand == "^" {
		return int(math.Pow(float64(nums[0]), float64(nums[1])))
	} else if operand == "sin" {
		return int(math.Sin(float64(nums[0])))
	} else if operand == "cos" {
		return int(math.Cos(float64(nums[0])))
	} else if operand == "tg" {
		return int(math.Tan(float64(nums[0])))
	}
}
