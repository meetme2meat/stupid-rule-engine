package main

import "fmt"

type LessThan struct {
	leftValue  Valuer
	rightValue Valuer
}

func ltBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("lessThan comparsion require 2 operand")
	}

	lt := new(LessThan)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	lt.leftValue = buildComparator(leftOperand)
	lt.rightValue = buildComparator(rightOperand)
	return lt
}

func (lt *LessThan) Evaluate(data map[string]interface{}) bool {

	lType := lt.leftValue.GetType()
	rType := lt.rightValue.GetType()
	if lType != rType {
		panic("lessThan type comparsion failed")
	}
	lValue := lt.leftValue.GetValue(data)
	rValue := lt.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a < b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a < b
	default:
		panic(fmt.Sprintf("lessThan not supported for given type %s", lType))
	}
}
