package main

type Value struct {
	Field *string
	Value interface{}
	Type  string
}

func buildValue(rule interface{}) Valuer {
	operand, ok := rule.(map[string]interface{})
	if !ok {
		panic("not a valid value to build")
	}

	_, ok = operand["field"]

	if _, ok1 := operand["value"]; ok || ok1 {
		return newValue(operand)
	}

	return buildValue(operand)
}

type ValueBuilder func(interface{}) Valuer

type Valuer interface {
	GetValue(map[string]interface{}) interface{}
	GetType() string
}

func (v *Value) GetType() string {
	return v.Type
}

func (v *Value) GetValue(d map[string]interface{}) interface{} {
	if v.Field == nil {
		return v.CastValue()
	}

	field := *v.Field
	val, ok := d[field]
	if !ok {
		panic("value must contain field")
	}

	newValue := &Value{Value: val, Type: v.Type}
	return newValue.CastValue()
}

func (v *Value) CastValue() interface{} {
	switch v.Type {
	case "int64":
		return toInt(v.Value)
	case "float64":
		return toFloat(v.Value)
	case "string":
		return toString(v.Value)
	case "boolean":
		return toBool(v.Value)
	}

	return nil
}
