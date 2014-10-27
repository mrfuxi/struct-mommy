package structmommy_test

import (
    "fmt"

    "github.com/mrfuxi/struct-mommy"
)

func ExampleMake() {
    obj := struct {
        FieldA float32
        FieldB uint8
    }{}

    structmommy.SetSeed(26)
    structmommy.Make(&obj)

    fmt.Printf("FieldA %v\nFieldB %v", obj.FieldA, obj.FieldB)
    // Output:
    // FieldA -1.5426473e+38
    // FieldB 42
}

func ExampleDefine() {
    obj := struct {
        FieldA float32
        FieldB uint8
    }{}

    structmommy.SetSeed(26)
    structmommy.Make(&obj, structmommy.Define("FieldA", 2.0))

    fmt.Printf("FieldA %v\nFieldB %v", obj.FieldA, obj.FieldB)
    // Output:
    // FieldA 2
    // FieldB 42
}
