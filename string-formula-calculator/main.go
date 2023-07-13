package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	formula := "( 10 +2)/ 2* 3"

	formula = removeSpace(formula)

	if !checkFormula(formula) {
		fmt.Println("Math Formula Error~!!!")
		return
	} else {
		fmt.Println("Math Formula normality~!!!")
	}

	fmt.Println(eval(formula))
}

// desc : 연산 기호를 후위 방식으로 stack 에 저장
// arg formula : 계산해야 할 문자열 수식
// return : 모든 계산이 끝난 결과 값
func eval(formula string) int {
	var symbols []byte
	var nums []int

	for i := 0; i < len(formula); {
		switch formula[i] {
		case ' ':
			i++
		case '(':
			symbols = append(symbols, formula[i])
			i++
		case '+', '-':
			for len(symbols) > 0 && symbols[len(symbols)-1] != '(' {
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
		case ')':
			for len(symbols) > 0 && symbols[len(symbols)-1] != '(' {
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

// desc : eval() 에서 후위 처리된 stack 을 기반으로 실제 계산을 수행
// arg nums : 계산해야 할 숫자
// arg symbols : 계산해야 할 수식 기호 +-*/%, ()
// return : nums == numbers, symbols == 수식 연산 기호들
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

// desc : 수식 오류 체크
// arg formula : 수식
// return : true == normal, false == error
func checkFormula(formula string) bool {
	// 괄호 열기, 닫기 체크
	openCnt := findChar(formula, "\\(")
	closeCnt := findChar(formula, "\\)")
	if (openCnt+closeCnt)%2 != 0 {
		return false
	}

	// 4칙 연산 체크
	signCnt := findChar(formula, "\\+") + findChar(formula, "\\-") +
		findChar(formula, "\\*") + findChar(formula, "\\/") + findChar(formula, "\\%")

	// 숫자(metric value) 갯수 체크
	sep := "()+-*/% "
	fields := strings.FieldsFunc(formula, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
	numCnt := len(fields)

	// 연산 대상 숫자 갯수보다 항상 4칙연산 기호 수가 하나 적어야 맞는 수식임.
	if (numCnt - 1) != signCnt {
		return false
	}

	return true
}

// desc : 공백 제거
// arg formula : 문자열 수식
// return : 공백 제거한 문자열 수식
func removeSpace(formula string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(formula, "")
}

// desc : 수식 문자열에서 특정 문자에 전체 갯수 확인
// arg formula : 문자열 수식
// arg find : 찾을 특정 문자, 문자열로 찾기가 편해 string type
func findChar(formula string, find string) int {
	re := regexp.MustCompile(find)
	return len(re.FindAllString(formula, -1))
}
