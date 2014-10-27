package structmommy

import (
    "errors"
    "fmt"
    "log"
    "reflect"
)

// Define is post Make function. Lets define values of struct fields
func Define(keyValue ...interface{}) func(interface{}) error {
    defined := variadicToMap(keyValue...)

    postFunc := func(obj interface{}) (err error) {
        reflectedVal := reflect.ValueOf(obj).Elem()
        if reflectedVal.Kind() != reflect.Struct {
            msg := fmt.Sprintf(
                "Define on %s does not make sense. Only structs are accepted",
                reflectedVal.Kind(),
            )
            return errors.New(msg)
        }

        for key, value := range defined {
            f := reflectedVal.FieldByName(key)
            newVal := reflect.ValueOf(value)

            if newVal.Type() != f.Type() {
                if newVal.Type().ConvertibleTo(f.Type()) {
                    newVal = newVal.Convert(f.Type())
                } else {
                    log.Printf("Type problem: %v vs %v. Skip", newVal.Type(), f.Type())
                    continue
                }
            }
            f.Set(newVal)
        }

        return
    }

    return postFunc
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
