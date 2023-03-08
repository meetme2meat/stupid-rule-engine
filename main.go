package main

import (
	"encoding/json"
	"fmt"
)

type falsify bool                                      // true if
func (n falsify) Evaluate(map[string]interface{}) bool { return false }

var Noop falsify

func unaryOperator(key string) ValueBuilder {
	builder := map[string]ValueBuilder{
		"plus":  plusbuilder,
		"minus": minusBuilder,
		"div":   divBuilder,
		"mod":   modBuilder,
		"not":   notBuilder,
	}
	return builder[key]
}

func buildComparator(rule map[string]interface{}) Valuer {
	for key, value := range rule {
		operator := unaryOperator(key)
		if operator == nil {
			break
		}
		return operator(value)
	}

	return buildValue(rule)
}

func main() {
	// "gt":
	// name := "rule 1"
	// rules := `{"or": [{"gte": [{"plus": [{"field": "age", "type": "int64"}, {"value": 9, "type": "int64"}]}, {"plus": [{"field": "newAge", "type": "int64"}, {"value": 0, "type": "int64"}]}]}, {"and": [{"gte": [{"field": "pincode","type": "int64"}, {"value": 400000, "type": "int64"}]},{"neq": [{"field": "name", "type":"string"},{"value": "John Doe", "type": "string"}]}]}]}`
	// check if John does is male or female
	name := "john does is female"
	rules := `{"and": [{"eq": [{"not": [{"field": "male", "type": "boolean"}]}, {"value": true, "type": "boolean"}]}]}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(rules), &data)
	if err != nil {
		panic(err)
	}

	r := Init(name, data)
	r.Evaluate(map[string]interface{}{"name": "John Doe", "male": false})
	// r.Evaluate(map[string]interface{}{"name": "John Doe1", "pincode": 400000, "age": 0, "newAge": 10})
}

func Init(name string, rule map[string]interface{}) *Rule {
	r := &Rule{Name: name}

	if value, ok := rule["or"]; ok {
		newBuilder := getBuilder("or")(value)
		r.condition = newBuilder
	}

	if value, ok := rule["and"]; ok {
		newBuilder := getBuilder("and")(value)
		r.condition = newBuilder
	}

	fmt.Println("rule build")
	return r
}
