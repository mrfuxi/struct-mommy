package structmommy

import (
    "errors"
    "fmt"
    "reflect"
    "strings"
)

// Definition of function that can be applied specific field
type fieldActionFunc (func(field, value reflect.Value, field_name string) error)

// Define is post Make function. Lets define values of struct fields
func Define(keyValue ...interface{}) func(interface{}) error {
    defined := variadicToMap(keyValue...)

    postFunc := func(obj interface{}) (err error) {
        reflectedVal := reflect.ValueOf(obj).Elem()
        return applyDefinedMap(reflectedVal, defined, setNewValue)
    }

    return postFunc
}

func applyDefinedMap(reflectedVal reflect.Value, defined map[string]interface{}, actionFunc fieldActionFunc) (err error) {
    for key, value := range defined {
        err = applyDefinedField(reflectedVal, key, value, actionFunc)
        if err != nil {
            break
        }
    }

    return
}

func applyDefinedField(reflectedVal reflect.Value, fieldName string, value interface{}, actionFunc fieldActionFunc) (err error) {
    var field reflect.Value

    //If field has dots, try to go though struct
    if strings.Contains(fieldName, ".") {
        fields := strings.SplitN(fieldName, ".", 2)
        directField := fields[0]
        furtherFields := fields[1]

        if field, err = fieldByName(reflectedVal, directField); err != nil {
            return
        }

        return applyDefinedField(field, furtherFields, value, actionFunc)
    }

    if err = hasFields(reflectedVal); err != nil {
        return
    }

    if field, err = fieldByName(reflectedVal, fieldName); err != nil {
        return
    }

    newVal := reflect.ValueOf(value)

    if newVal.Type() != field.Type() {
        if newVal.Type().ConvertibleTo(field.Type()) {
            newVal = newVal.Convert(field.Type())
        } else {
            err = fmt.Errorf(
                "could not convert between %v (src) and %v (dst)",
                newVal.Type(),
                field.Type(),
            )
            return
        }
    }

    err = actionFunc(field, newVal, fieldName)

    return
}

func setNewValue(field, value reflect.Value, fieldName string) error {
    field.Set(value)

    return nil
}

func fieldByName(reflectedVal reflect.Value, fieldName string) (field reflect.Value, err error) {
    field = reflectedVal.FieldByName(fieldName)
    if !field.IsValid() {
        err = fmt.Errorf(
            "could not find field '%s' in %s struct",
            fieldName,
            reflectedVal.Type(),
        )
        return
    }
    return
}

func hasFields(reflectedVal reflect.Value) (err error) {
    if reflectedVal.Kind() != reflect.Struct {
        err = fmt.Errorf(
            "define on %s does not make sense; only structs are accepted",
            reflectedVal.Kind(),
        )
        return
    }
    return
}

func variadicToMap(keyValue ...interface{}) map[string]interface{} {
    if len(keyValue)%2 != 0 {
        panic(errors.New("number of keys does not match number of values"))
    }

    defined := map[string]interface{}{}

    var key string
    for i, kv := range keyValue {
        if i%2 == 0 {
            key = kv.(string)
        } else {
            defined[key] = kv
        }
    }

    return defined
}
