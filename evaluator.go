package main

type Builder func(interface{}) Evaluator

type Evaluator interface {
	Evaluate(map[string]interface{}) bool
}

func getBuilder(key string) Builder {
	builder := map[string]Builder{
		"or":  orBuilder,
		"and": andBuilder,
		"gt":  gtBuilder,
		"gte": gteBuilder,
		"lt":  ltBuilder,
		"lte": lteBuilder,
		"eq":  eqBuilder,
		"neq": neqBuilder,
	}
	return builder[key]
}
