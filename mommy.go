// Package struct_mommy implement object generator for test pourposes
//
// Call Make to fill in your struct with random data.
// Use PostMake functions to define some of fields (Define), save struct to DB or whatever you please.
//
// ToDo: Let gererate multiple objects at once!
package struct_mommy

import (
    "bytes"
    "math/rand"
    "reflect"
    "testing/quick"
    "time"
)

var (
    random *rand.Rand
    chars  = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

const (
    radomStringSize = 20
)

func init() {
    seed := rand.NewSource(time.Now().UnixNano())
    random = rand.New(seed)
}

// Make fills in given object with radom data.
// Afterwords, if specified, PostMake functions are applied in order on the object.
// If any function will return no nil error, no other function is execute and error is returned
func Make(obj interface{}, post_funcs ...func(interface{}) error) (err error) {
    obj_type := reflect.TypeOf(obj)
    kind := obj_type.Elem().Kind()

    switch kind {
    case reflect.Struct:
        make_struct(obj)
    default:
        val := reflect.ValueOf(obj).Elem()
        make_simple(val)
    }

    for _, post_func := range post_funcs {
        err = post_func(obj)

        if err != nil {
            return
        }
    }

    return
}

func random_string() string {
    b := bytes.NewBuffer(make([]byte, radomStringSize))
    for i := 0; i < radomStringSize; i++ {
        x := chars[random.Intn(len(chars))]
        b.WriteByte(x)
    }
    return b.String()
}

func make_simple(val reflect.Value) {
    // Radom string from quick.Value is not readable
    if val.Kind() == reflect.String {
        val.SetString(random_string())
        return
    }

    new_val, _ := quick.Value(val.Type(), random)
    val.Set(new_val)
}

func make_struct(obj interface{}) {
    s := reflect.ValueOf(obj).Elem()

    for i := 0; i < s.NumField(); i++ {
        field := s.Field(i)
        make_simple(field)
    }
}
