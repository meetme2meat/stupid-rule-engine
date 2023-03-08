package main

import "fmt"

func notBuilder(rule interface{}) Valuer {
	unaryRules, ok := rule.([]interface{})
	if !ok {
		panic("not a valid minus rule")
	}

	if len(unaryRules) != 1 {
		panic("unaryNot rule require 1 operand")
	}

	not := new(notOperator)
	not.Value = buildValue(unaryRules[0])
	return not
}

type notOperator struct {
	Value Valuer
}

func (p *notOperator) GetValue(data map[string]interface{}) interface{} {
	value := p.Value.GetValue(data)
	typeOf := p.Value.GetType()

	switch typeOf {
	case "boolean":
		a := value.(bool)
		return !a
	default: //
		panic(fmt.Sprintf("minus does not support given type %s", typeOf))
	}
}

func (p *notOperator) GetType() string {
	return "boolean"
}
