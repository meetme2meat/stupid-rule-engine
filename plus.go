package main

import "fmt"

func plusbuilder(rule interface{}) Valuer {
	unaryRules, ok := rule.([]interface{})
	if !ok {
		panic("not a valid plus rule")
	}

	if len(unaryRules) != 2 {
		panic("plus must be 2 operand")
	}

	plus := new(plusOperator)
	plus.leftValue = buildValue(unaryRules[0])
	plus.rightValue = buildValue(unaryRules[1])
	return plus
}

type plusOperator struct {
	leftValue  Valuer
	rightValue Valuer
}

func (p *plusOperator) GetValue(data map[string]interface{}) interface{} {
	leftValue := p.leftValue.GetValue(data)
	rightValue := p.rightValue.GetValue(data)
	typeOf := p.leftValue.GetType()

	switch typeOf {
	case "int64":
		a := leftValue.(int64)
		b := rightValue.(int64)
		return a + b
	case "float64":
		a := leftValue.(float64)
		b := rightValue.(float64)
		return a + b
	case "string":
		a := leftValue.(string)
		b := rightValue.(string)
		return a + b
	default:
		panic(fmt.Sprintf("plus does not support given type %s", typeOf))
	}
}

func (p *plusOperator) GetType() string {
	return p.leftValue.GetType()
}
