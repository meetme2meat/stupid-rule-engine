package main

import "fmt"

type Equal struct {
	leftValue  Valuer
	rightValue Valuer
}

func eqBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("equal comparsion require 2 operand")
	}

	eq := new(Equal)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	eq.leftValue = buildComparator(leftOperand)
	eq.rightValue = buildComparator(rightOperand)
	return eq
}

func (eq *Equal) Evaluate(data map[string]interface{}) bool {

	lType := eq.leftValue.GetType()
	rType := eq.rightValue.GetType()
	if lType != rType {
		panic("equal type comparsion failed")
	}
	lValue := eq.leftValue.GetValue(data)
	rValue := eq.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a == b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a == b
	case "string":
		a := lValue.(string)
		b := rValue.(string)
		return a == b
	case "boolean":
		a := lValue.(bool)
		b := rValue.(bool)
		return a == b
	default:
		panic(fmt.Sprintf("equal not supported for given type %s", lType))
	}
}
