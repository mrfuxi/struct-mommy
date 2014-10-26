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
    RadomStringSize = 20
)

func init() {
    seed := rand.NewSource(time.Now().UnixNano())
    random = rand.New(seed)
}

func Make(obj interface{}) {
    obj_type := reflect.TypeOf(obj)
    kind := obj_type.Elem().Kind()

    switch kind {
    case reflect.Struct:
        make_struct(obj)
    default:
        val := reflect.ValueOf(obj).Elem()
        make_simple(val)
    }
}

func random_string() string {
    b := bytes.NewBuffer(make([]byte, RadomStringSize))
    for i := 0; i < RadomStringSize; i++ {
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
