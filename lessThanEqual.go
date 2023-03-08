package main

import "fmt"

type LessThanEqual struct {
	leftValue  Valuer
	rightValue Valuer
}

func lteBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("lessThanEqual comparsion require 2 operand")
	}

	lte := new(LessThanEqual)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	lte.leftValue = buildComparator(leftOperand)
	lte.rightValue = buildComparator(rightOperand)
	return lte
}

func (lte *LessThanEqual) Evaluate(data map[string]interface{}) bool {

	lType := lte.leftValue.GetType()
	rType := lte.rightValue.GetType()
	if lType != rType {
		panic("lessThanEqual type comparsion failed")
	}
	lValue := lte.leftValue.GetValue(data)
	rValue := lte.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a <= b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a <= b
	default:
		panic(fmt.Sprintf("lessThanEqual not supported for given type %s", lType))
	}
}
