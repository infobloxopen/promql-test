package promock

import (
	"fmt"
	"math"

	"github.com/Knetic/govaluate"
)

func convertToFloat64(r interface{}) (float64, error) {

	switch v := r.(type) {
	case float64:
		return r.(float64), nil
	case float32:
		return float64(r.(float32)), nil
	case int:
		return float64(r.(int)), nil
	case int64:
		return float64(r.(int64)), nil
	case int32:
		return float64(r.(int32)), nil
	case int16:
		return float64(r.(int16)), nil
	case int8:
		return float64(r.(int8)), nil
	case uint:
		return float64(r.(uint)), nil
	case uint64:
		return float64(r.(uint64)), nil
	case uint32:
		return float64(r.(uint32)), nil
	case uint16:
		return float64(r.(uint16)), nil
	case uint8:
		return float64(r.(uint8)), nil
	default:
		return nan, fmt.Errorf("unsupported result type %T", v)
	}
}

func fnFloat(arguments ...interface{}) (interface{}, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("float() takes exactly one argument")
	}
	return convertToFloat64(arguments[0])
}

func fnMOD(arguments ...interface{}) (interface{}, error) {
	if len(arguments) != 2 {
		return nil, fmt.Errorf("mod(x,y) takes exactly 2 arguments")
	}
	x, ok := arguments[0].(float64)
	if !ok {
		return nil, fmt.Errorf("argument x to mod(x,y) must be float64")
	}
	y, ok := arguments[1].(float64)
	if !ok {
		return nil, fmt.Errorf("argument y to mod(x,y) must be float64")
	}
	return math.Mod(x, y), nil
}

func fnSIN(arguments ...interface{}) (interface{}, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("sin takes exactly one argument")
	}
	a, ok := arguments[0].(float64)
	if !ok {
		return nil, fmt.Errorf("sin argument must be float64")
	}
	return math.Sin(a), nil
}

func fnToUnix(arguments ...interface{}) (interface{}, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("to_unix() takes exactly one argument")
	}
	ms, err := convertToFloat64(arguments[0])
	if err != nil {
		return nan, err
	}
	return ms / 1000.0, nil
}

var commonFuncs = map[string]govaluate.ExpressionFunction{
	"float":   fnFloat,
	"mod":     fnMOD,
	"sin":     fnSIN,
	"to_unix": fnToUnix,
}
