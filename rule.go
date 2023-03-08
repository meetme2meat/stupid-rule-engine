package main

import "fmt"

type Rule struct {
	Name      string
	condition Evaluator
}

func (r *Rule) Evaluate(fact map[string]interface{}) {
	fmt.Println(r.Name, "=>", r.condition.Evaluate(fact))
}
