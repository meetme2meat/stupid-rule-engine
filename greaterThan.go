package main

import "fmt"

type GreaterThan struct {
	leftValue  Valuer
	rightValue Valuer
}

func gtBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	fields, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	if len(fields) != 2 {
		panic("greaterThan comparsion require 2 operand")
	}

	gt := new(GreaterThan)
	leftOperand := fields[0].(map[string]interface{})
	rightOperand := fields[1].(map[string]interface{})

	gt.leftValue = buildComparator(leftOperand)
	gt.rightValue = buildComparator(rightOperand)
	return gt
}

func (gt *GreaterThan) Evaluate(data map[string]interface{}) bool {

	lType := gt.leftValue.GetType()
	rType := gt.rightValue.GetType()
	if lType != rType {
		panic("greaterThan type comparsion failed")
	}
	lValue := gt.leftValue.GetValue(data)
	rValue := gt.rightValue.GetValue(data)
	switch lType {
	case "int64":
		a := lValue.(int64)
		b := rValue.(int64)
		return a > b
	case "float64":
		a := lValue.(float64)
		b := rValue.(float64)
		return a > b
	default:
		panic(fmt.Sprintf("greaterThan not supported for given type %s", lType))
	}
}
