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

    fmt.Println("FieldA:", obj.FieldA)
    fmt.Println("FieldB:", obj.FieldB)
    // Output:
    // FieldA: -1.5426473e+38
    // FieldB: 42
}

func ExampleDefine() {
    obj := struct {
        FieldA float32
        FieldB uint8
    }{}

    structmommy.SetSeed(26)
    structmommy.Make(&obj, structmommy.Define("FieldA", 2.0))

    fmt.Println("FieldA:", obj.FieldA)
    fmt.Println("FieldB:", obj.FieldB)
    // Output:
    // FieldA: 2
    // FieldB: 42
}

func ExamplePartialEqual_ok() {
    obj := struct {
        FieldA float32
        FieldB uint8
    }{
        FieldA: 1.2,
        FieldB: 42,
    }

    err := structmommy.PartialEqual(obj, "FieldA", 1.2, "FieldB", 42)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("All fields as expected!")
    }

    fmt.Println("FieldA:", obj.FieldA)
    fmt.Println("FieldB:", obj.FieldB)
    // Output:
    // All fields as expected!
    // FieldA: 1.2
    // FieldB: 42
}

func ExamplePartialEqual_notEqual() {
    obj := struct {
        FieldA float32
        FieldB uint8
    }{
        FieldA: 1.2,
        FieldB: 42,
    }

    err := structmommy.PartialEqual(obj, "FieldA", 123)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("All fields as expected!")
    }

    fmt.Println("FieldA:", obj.FieldA)
    fmt.Println("FieldB:", obj.FieldB)
    // Output:
    // Error: value of field 'FieldA' equals 1.2, expected 123
    // FieldA: 1.2
    // FieldB: 42
}
