package main

import (
	"fmt"
	"strconv"
)

func toInt(v interface{}) interface{} {
	switch vType := v.(type) {
	case int:
		val := v.(int)
		return int64(val)
	case int64:
		return v
	case float64:
		val := v.(float64)
		return int64(val)
	case string:
		val := v.(string)
		newVal, _ := strconv.ParseInt(val, 10, 64)
		return newVal
	default:
		panic(fmt.Sprintf("cannot convert int from type %s", vType))
	}
}

func toFloat(v interface{}) interface{} {
	switch vType := v.(type) {
	case int:
		val := v.(int)
		return float64(val)
	case int64:
		val := v.(int64)
		return float64(val)
	case float64:
		return v
	case string:
		val := v.(string)
		newVal, _ := strconv.ParseFloat(val, 64)
		return newVal
	default:
		panic(fmt.Sprintf("cannot convert float64 from type %s", vType))
	}

}

func toString(v interface{}) interface{} {
	switch vType := v.(type) {
	case int, int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%d", v)
	case string:
		return v
	default:
		panic(fmt.Sprintf("cannot convert string from type %s", vType))
	}
}

func toBool(v interface{}) interface{} {
	switch vType := v.(type) {
	case bool:
		return v.(bool)
	default:
		panic(fmt.Sprintf("cannot convert string from type %s", vType))
	}
}

func newValue(operand map[string]interface{}) Valuer {
	v := &Value{}
	field, ok := operand["field"]
	if ok {
		f := field.(string)
		v.Field = &f
	}

	value, ok1 := operand["value"]
	if ok1 {
		v.Value = value
	}

	v.Type = operand["type"].(string)
	return v
}
