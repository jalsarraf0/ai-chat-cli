package extra

func Sum(a, b int) int { return a + b }

func Multiply(a, b int) int { return a * b }

func Subtract(a, b int) int { return a - b }

func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
