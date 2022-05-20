package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"math/rand"
	"time"
)

// TODO: 定数ではなくデッキから減らしていけるようにするとか。
func getStack() []int {
	return []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 14, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
}

func getNumber() []int {
	stack := getStack()
	rand.Seed(time.Now().UnixNano())

	res := []int{}
	for i := 0; i < 5; i++ {
		num := rand.Intn(len(stack))
		res = append(res, stack[num])
		// 取ったカードはデッキから削除
		stack = append(stack[:num], stack[num+1:]...)
	}

	return res
}

func eval(e string) (string, error) {
	expr, err := parser.ParseExpr(e)
	if err != nil {
		return "", err
	}

	v, err := evalExpr(expr)
	if err != nil {
		return "", err
	}

	return v.String(), nil
}

func evalExpr(expr ast.Expr) (v constant.Value, rerr error) {
	defer func() {
		if r := recover(); r != nil {
			v, rerr = constant.MakeUnknown(), fmt.Errorf("%v", r)
		}
	}()

	switch e := expr.(type) {
	case *ast.ParenExpr:
		return evalExpr(e.X)
	case *ast.BinaryExpr:
		return evalBinaryExpr(e)
	case *ast.UnaryExpr:
		return evalUnaryExpr(e)
	case *ast.BasicLit:
		return constant.MakeFromLiteral(e.Value, e.Kind, 0), nil
	}

	return constant.MakeUnknown(), errors.New("unkown formula")
}

func evalBinaryExpr(expr *ast.BinaryExpr) (constant.Value, error) {
	x, err := evalExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), err
	}

	y, err := evalExpr(expr.Y)
	if err != nil {
		return constant.MakeUnknown(), err
	}

	switch expr.Op {
	case token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
		return constant.MakeBool(constant.Compare(x, expr.Op, y)), nil
	}

	return constant.BinaryOp(x, expr.Op, y), nil
}

func evalUnaryExpr(expr *ast.UnaryExpr) (constant.Value, error) {
	x, err := evalExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), err
	}

	return constant.UnaryOp(expr.Op, x, 0), nil
}

func main() {
	fmt.Println("Pick up 5 card! Please make 15 with four arithmetic operations.")
	fmt.Println("")
	for _, v := range getNumber() {
		fmt.Print("[")
		fmt.Print(v)
		fmt.Print("]")
	}
	fmt.Println("")

	fmt.Println("Please enter the formula.")
	var l string
	fmt.Scan(&l)
	r, err := eval(l)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("[Calculation Results]: " + r)
	if r == "15" {
		fmt.Printf("\x1b[32m%s\x1b[0m", "Congratulations!!!!!")
		fmt.Println("")
		return
	}
	fmt.Printf("\x1b[31m%s\x1b[0m", "The result was not 15")
	fmt.Println("")
}
