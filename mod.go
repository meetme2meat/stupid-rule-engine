package main

import "fmt"

func modBuilder(rule interface{}) Valuer {
	unaryRules, ok := rule.([]interface{})
	if !ok {
		panic("not a valid mod rule")
	}

	if len(unaryRules) != 2 {
		panic("mod must have 2 operand")
	}

	plus := new(modOperator)
	plus.leftValue = buildValue(unaryRules[0])
	plus.rightValue = buildValue(unaryRules[1])
	return plus
}

type modOperator struct {
	leftValue  Valuer
	rightValue Valuer
}

func (p *modOperator) GetValue(data map[string]interface{}) interface{} {
	leftValue := p.leftValue.GetValue(data)
	rightValue := p.rightValue.GetValue(data)
	typeOf := p.leftValue.GetType()

	switch typeOf {
	case "int64":
		a := leftValue.(int64)
		b := rightValue.(int64)
		return a % b
	default:
		panic(fmt.Sprintf("mod does not support given type %s", typeOf))
	}
}

func (p *modOperator) GetType() string {
	return p.leftValue.GetType()
}
