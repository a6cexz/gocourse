package calculator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"
	"strings"
)

// Func info
type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

var funcMap map[string]Func

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0])
		},
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0])
		},
	}
	funcMap["tan"] = Func{
		Name: "tan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Tan(args[0])
		},
	}
	funcMap["arcsin"] = Func{
		Name: "arcsin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Asin(args[0])
		},
	}
	funcMap["arccos"] = Func{
		Name: "arccos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Acos(args[0])
		},
	}
	funcMap["arctan"] = Func{
		Name: "arctan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Atan(args[0])
		},
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
	funcMap["help"] = Func{
		Name: "help",
		Args: 0,
		Func: func(args ...float64) float64 {
			fmt.Println("Calculator help")
			fmt.Println("Enter expression and press Enter to evaluate it.")
			fmt.Println("E.g. 10 + 20")
			fmt.Println("To quit type exit and press Enter")
			return 0
		},
	}
}

func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	case *ast.BinaryExpr:
		return binary(expr.(*ast.BinaryExpr))
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Unknown operator")
	}
}

func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)
		if err != nil {
			return -1, err
		}
		return float64(i), nil

	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)
		if err != nil {
			return -1, err
		}
		return i, nil

	default:
		return -1, errors.New("Unknown argument")
	}
}

func binary(expr *ast.BinaryExpr) (ret float64, err error) {
	x, err1 := eval(expr.X)
	y, err2 := eval(expr.Y)
	ret = -1
	if (err1 == nil) && (err2 == nil) {
		switch expr.Op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = float64(int64(x) % int64(y))
		case token.AND:
			ret = float64(int64(x) & int64(y))
		case token.OR:
			ret = float64(int64(x) | int64(y))
		case token.XOR:
			ret = math.Pow(x, y)
		default:
			err = errors.New("Unknown binary operator")
		}
	} else {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}
	return
}

func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil
	case "phi":
		return math.Phi, nil
	default:
		return -1, errors.New("Unknown const " + n)
	}
}

func call(c *ast.CallExpr) (float64, error) {
	switch t := c.Fun.(type) {
	case *ast.Ident:
	default:
		_ = t
		return -1, errors.New("Unknown func")
	}

	ident := c.Fun.(*ast.Ident)
	args := make([]float64, len(c.Args))
	for i, expr := range c.Args {
		var err error
		args[i], err = eval(expr)
		if err != nil {
			return -1, err
		}
	}

	name := strings.ToLower(ident.Name)
	if val, ok := funcMap[name]; ok {
		if len(args) == val.Args {
			return val.Func(args...), nil
		}
		return -1, errors.New("Too many arguments for  " + name)
	}
	return -1, errors.New("Unknown func " + name)
}

// Calculate expr
func Calculate(expr string) (float64, error) {
	root, err := parser.ParseExpr(expr)
	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}

// Run functions shows sort demo
func Run() {
	fmt.Println()
	fmt.Println("Calculator Demo")
	fmt.Println()

	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := Calculate(input); err == nil {
			fmt.Printf("Result: %v\n", res)
		} else {
			fmt.Println("Can not calculate")
		}
	}
}
