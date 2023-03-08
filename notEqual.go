package main

import "fmt"

type NotEqual struct {
	leftValue  Valuer
	rightValue Valuer
}

func neqBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("notEqual comparsion require 2 operand")
	}

	neq := new(NotEqual)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	neq.leftValue = buildComparator(leftOperand)
	neq.rightValue = buildComparator(rightOperand)
	return neq
}

func (neq *NotEqual) Evaluate(data map[string]interface{}) bool {

	lType := neq.leftValue.GetType()
	rType := neq.rightValue.GetType()
	if lType != rType {
		panic("notEqual type comparsion failed")
	}
	lValue := neq.leftValue.GetValue(data)
	rValue := neq.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a != b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a != b
	case "string":
		a := lValue.(string)
		b := rValue.(string)
		return a != b
	default:
		panic(fmt.Sprintf("notEqual does not supported for given type %s", lType))
	}
}
