package bpn

import (
	"bpn/pkg/structs"
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	MathFormulaError = errors.New("math formula error")
	BPNFormulaError  = errors.New("bpn formula error")

	signsPriority = map[string]int{
		"(": 0,
		")": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
)

// ToRPN первод математической формулы в обратную польскую аннотацию
func ToRPN(fml string) (string, error) {
	if len(fml) == 0 {
		return "", MathFormulaError
	}

	var (
		signs structs.Queue[string]
		res   strings.Builder
	)

	fields := strings.Fields(fml)
	for i := 0; i < len(fields); i++ {
		field := fields[i]

		switch field {
		case "+", "-", "*", "/", "^":
			if i == len(fields)-1 {
				return "", MathFormulaError
			}
			priority, priorityLastSign := signsPriority[field], 0
			if signs.Len() > 0 {
				priorityLastSign = signsPriority[signs.PeekLast()]
			}
			if priority >= priorityLastSign {
				signs.PushFirst(field)
				continue
			}
			for signs.Len() > 0 {
				res.WriteString(signs.PopLast() + " ")
			}
			signs.PushLast(field)
		case "(":
			var (
				iterRes  strings.Builder
				brackets int = 1
			)

			i++
			for ; i < len(fields); i++ {
				if fields[i] == "(" {
					brackets++
				} else if fields[i] == ")" {
					brackets--
					if brackets == 0 {
						break
					}
				}
				iterRes.WriteString(fields[i] + " ")
			}
			if brackets != 0 {
				return "", MathFormulaError
			}

			sib, err := ToRPN(strings.TrimSpace(iterRes.String()))
			if err != nil {
				return "", MathFormulaError
			}
			res.WriteString(sib + " ")
		default:
			if _, err := strconv.Atoi(field); err != nil {
				return "", MathFormulaError
			}
			res.WriteString(field + " ")
		}
	}

	for signs.Len() > 0 {
		res.WriteString(signs.PopFirst())
		if signs.Len() > 0 {
			res.WriteString(" ")
		}
	}

	return res.String(), nil
}

// CalcRPN вычисление обратной польской аннотации
func CalcRPN(bpnFormula string) (float64, error) {
	if len(bpnFormula) == 0 {
		return 0, BPNFormulaError
	}

	var (
		res  float64
		nums structs.Stack[float64]
	)

	bpn := strings.Fields(bpnFormula)
	for i := 0; i < len(bpn); i++ {
		switch bpn[i] {
		case "*", "/", "+", "-", "^":
			n2 := nums.Pop()
			n1 := nums.Pop()
			nums.Push(calc(n1, n2, bpn[i]))
		default:
			n, err := strconv.ParseFloat(bpn[i], 64)
			if err != nil {
				return res, BPNFormulaError
			}
			nums.Push(n)
		}
	}

	return nums.Pop(), nil
}

func calc(n, m float64, op string) float64 {
	var res float64

	switch op {
	case "+":
		res = n + m
	case "-":
		res = n - m
	case "*":
		res = n * m
	case "/":
		res = n / m
	case "^":
		res = math.Pow(n, m)
	}
	return res
}
