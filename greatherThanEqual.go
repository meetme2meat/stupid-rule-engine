package main

import "fmt"

type GreaterThanEqual struct {
	leftValue  Valuer
	rightValue Valuer
}

func gteBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("greaterThanEqual comparsion require 2 operand")
	}

	gt := new(GreaterThanEqual)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	gt.leftValue = buildComparator(leftOperand)
	gt.rightValue = buildComparator(rightOperand)
	return gt
}

func (gte *GreaterThanEqual) Evaluate(data map[string]interface{}) bool {

	lType := gte.leftValue.GetType()
	rType := gte.rightValue.GetType()
	if lType != rType {
		panic("greaterThanEqual type comparsion failed")
	}
	lValue := gte.leftValue.GetValue(data)
	rValue := gte.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a >= b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a >= b
	default:
		panic(fmt.Sprintf("greaterThanEqual not supported for given type %s", lType))
	}
}
