// Package structmommy implement object generator for test purposes
//
// Call Make to fill in your struct with random data.
// Use PostMake functions to define some of fields (Define), save struct to DB or whatever you please.
//
// ToDo: Let generate multiple objects at once!
package structmommy

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
func Make(obj interface{}, postFuncs ...func(interface{}) error) (err error) {
    objType := reflect.TypeOf(obj)
    kind := objType.Elem().Kind()

    switch kind {
    case reflect.Struct:
        makeStruct(obj)
    default:
        val := reflect.ValueOf(obj).Elem()
        makeSimple(val)
    }

    for _, postFunc := range postFuncs {
        err = postFunc(obj)

        if err != nil {
            return
        }
    }

    return
}

// SetSeed let you set/reset seed on rand object used to generate data.
//
// Using it is completely optional as on init seed is set to current timesstamp (UnixNano).
func SetSeed(seed int64) {
    seedSource := rand.NewSource(seed)
    random = rand.New(seedSource)
}

func randomString() string {
    b := bytes.NewBuffer(make([]byte, radomStringSize))
    for i := 0; i < radomStringSize; i++ {
        x := chars[random.Intn(len(chars))]
        b.WriteByte(x)
    }
    return b.String()
}

func makeSimple(val reflect.Value) {
    // Radom string from quick.Value is not readable
    if val.Kind() == reflect.String {
        val.SetString(randomString())
        return
    }

    newVal, _ := quick.Value(val.Type(), random)
    val.Set(newVal)
}

func makeStruct(obj interface{}) {
    s := reflect.ValueOf(obj).Elem()

    for i := 0; i < s.NumField(); i++ {
        field := s.Field(i)
        makeSimple(field)
    }
}
