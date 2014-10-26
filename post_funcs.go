package struct_mommy

import (
    "errors"
    "fmt"
    "log"
    "reflect"
)

func variadicToMap(key_value ...interface{}) map[string]interface{} {
    if len(key_value)%2 != 0 {
        panic(errors.New("Number of keys does not match number of values"))
    }

    defined := map[string]interface{}{}

    var key string
    for i, kv := range key_value {
        if i%2 == 0 {
            key = kv.(string)
        } else {
            defined[key] = kv
        }
    }

    return defined
}

func Define(key_value ...interface{}) PostFunc {
    defined := variadicToMap(key_value...)

    post_func := func(obj interface{}) (err error) {
        reflected_val := reflect.ValueOf(obj).Elem()
        if reflected_val.Kind() != reflect.Struct {
            msg := fmt.Sprintf("Define on %s does not make sence", reflected_val.Kind())
            return errors.New(msg)
        }

        for key, value := range defined {
            f := reflected_val.FieldByName(key)
            new_val := reflect.ValueOf(value)

            if new_val.Type() != f.Type() {
                if new_val.Type().ConvertibleTo(f.Type()) {
                    new_val = new_val.Convert(f.Type())
                } else {
                    log.Printf("Type problem: %v vs %v. Skip", new_val.Type(), f.Type())
                    continue
                }
            }
            f.Set(new_val)
        }

        return
    }

    return post_func
}
