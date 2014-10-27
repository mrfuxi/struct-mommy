package struct_mommy_test

import (
    "fmt"

    "github.com/mrfuxi/struct-mommy"
)

type MyStruct struct {
    FieldA float32
    FieldB uint8
}

func ExampleMake() {
    obj := MyStruct{}

    struct_mommy.SetSeed(26)
    struct_mommy.Make(&obj)

    fmt.Printf("FieldA %v\nFieldB %v", obj.FieldA, obj.FieldB)
    // Output:
    // FieldA -1.5426473e+38
    // FieldB 42
}
