package dynamic_method

import "reflect"

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}
