package main

func andBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	newRules, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	and := new(And)
	for _, newRule := range newRules {
		if rule, ok := newRule.(map[string]interface{}); ok {
			for key, value := range rule {
				newBuilder := getBuilder(key)
				and.evaluators = append(and.evaluators, newBuilder(value))
			}
		}
	}

	return and
}

type And struct {
	evaluators []Evaluator
}

func (a *And) Evaluate(data map[string]interface{}) bool {
	for _, evaluate := range a.evaluators {
		if !evaluate.Evaluate(data) {
			return false
		}
	}

	return true
}
