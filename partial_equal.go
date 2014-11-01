package structmommy

import (
    "errors"
    "fmt"
    "reflect"
)

// Checks values of subset of fields in a given struct.
// Accepts both object or it's refference.
func PartialEqual(obj interface{}, keyValue ...interface{}) (err error) {
    defined := variadicToMap(keyValue...)

    reflectedVal := reflect.ValueOf(obj)
    if reflectedVal.Kind() == reflect.Ptr {
        reflectedVal = reflectedVal.Elem()
    }
    err = applyDefinedMap(reflectedVal, defined, checkValue)

    return
}

func checkValue(field, value reflect.Value, fieldName string) (err error) {
    field_int := field.Interface()
    value_int := value.Interface()

    if reflect.DeepEqual(field_int, value_int) {
        return
    }

    msg := fmt.Sprintf("value of field '%v' equals %v, expected %v", fieldName, field_int, value_int)
    err = errors.New(msg)

    return
}
