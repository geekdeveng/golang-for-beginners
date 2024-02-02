package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	formula := "{(10 +2)/ 2}* 3"

	formula = removeSpace(formula)

	if !checkFormula(formula) {
		fmt.Println("Math Formula Error~!!!")
		return
	} else {
		fmt.Println("Math Formula normality~!!!")
	}

	fmt.Println(eval(formula))
}

func eval(formula string) int {
	var symbols []byte
	var nums []int

	for i := 0; i < len(formula); {
		switch formula[i] {
		case ' ':
			i++
		case '(', '{', '[':
			symbols = append(symbols, formula[i])
			i++
		case '+', '-':
			for len(symbols) > 0 && symbols[len(symbols)-1] != '(' && symbols[len(symbols)-1] != '{' && symbols[len(symbols)-1] != '[' {
				nums, symbols = calc(nums, symbols)
			}
			symbols = append(symbols, formula[i])
			i++
		case '*', '/', '%':
			for len(symbols) > 0 && (symbols[len(symbols)-1] == '*' || symbols[len(symbols)-1] == '/' ||
				symbols[len(symbols)-1] == '%') {
				nums, symbols = calc(nums, symbols)
			}
			symbols = append(symbols, formula[i])
			i++
		case ')', '}', ']':
			for len(symbols) > 0 && symbols[len(symbols)-1] != '(' && symbols[len(symbols)-1] != '{' && symbols[len(symbols)-1] != '[' {
				nums, symbols = calc(nums, symbols)
			}
			if len(symbols) > 0 {
				symbols = symbols[:len(symbols)-1]
			}
			i++
		default:
			j := i
			for ; j < len(formula) && formula[j] >= '0' && formula[j] <= '9'; j++ {
			}
			num, _ := strconv.Atoi(formula[i:j])
			nums = append(nums, num)
			i = j
		}
	}

	for len(symbols) > 0 {
		nums, symbols = calc(nums, symbols)
	}

	return nums[0]
}

func calc(nums []int, symbols []byte) ([]int, []byte) {
	symbol := symbols[len(symbols)-1]
	symbols = symbols[:len(symbols)-1]
	b := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	a := nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	switch symbol {
	case '+':
		nums = append(nums, a+b)
	case '-':
		nums = append(nums, a-b)
	case '*':
		nums = append(nums, a*b)
	case '/':
		nums = append(nums, a/b)
	case '%':
		nums = append(nums, a%b)
	default:
		fmt.Println(fmt.Sprintf("unsupported operator: %q", symbol))
	}

	return nums, symbols
}

func checkFormula(formula string) bool {
	openCnt := findChar(formula, "\\(") + findChar(formula, "\\{") + findChar(formula, "\\[")
	closeCnt := findChar(formula, "\\)") + findChar(formula, "\\}") + findChar(formula, "\\]")
	if (openCnt+closeCnt)%2 != 0 {
		return false
	}

	signCnt := findChar(formula, "\\+") + findChar(formula, "\\-") +
		findChar(formula, "\\*") + findChar(formula, "\\/") + findChar(formula, "\\%")

	sep := "()+-*/% {}[]"
	fields := strings.FieldsFunc(formula, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
	numCnt := len(fields)

	if (numCnt - 1) != signCnt {
		return false
	}

	return true
}

func removeSpace(formula string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(formula, "")
}

func findChar(formula string, find string) int {
	re := regexp.MustCompile(find)
	return len(re.FindAllString(formula, -1))
}
