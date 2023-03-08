package main

func orBuilder(rule interface{}) Evaluator {
	if rule == nil {
		return Noop
	}

	newRules, ok := rule.([]interface{})
	if !ok {
		return Noop
	}

	or := new(Or)
	for _, newRule := range newRules {
		if rule, ok := newRule.(map[string]interface{}); ok {
			for key, value := range rule {
				newBuilder := getBuilder(key)
				or.evaluators = append(or.evaluators, newBuilder(value))
			}
		}
	}

	return or
}

type Or struct {
	evaluators []Evaluator
}

func (o *Or) Evaluate(data map[string]interface{}) bool {
	for _, evaluate := range o.evaluators {
		if evaluate.Evaluate(data) {
			return true
		}
	}

	return false
}
