package struct_mommy_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/mrfuxi/struct-mommy"
)

func TestPartiallyRandom(t *testing.T) {
    obj := testMultipleSimpleFields{}

    struct_mommy.Make(&obj, struct_mommy.Define("FieldFloat", 2.0))

    assert.NotEmpty(t, obj.FieldInt)
    assert.NotEmpty(t, obj.FieldStr)
    assert.Equal(t, obj.FieldFloat, 2.0)
}

func TestPartiallyRandomTypeConversion(t *testing.T) {
    obj := testMultipleSimpleFields{}

    struct_mommy.Make(&obj, struct_mommy.Define("FieldInt", 2.0, "FieldStr", "1"))

    assert.Equal(t, obj.FieldInt, 2)
    assert.Equal(t, obj.FieldStr, "1")
    assert.NotEmpty(t, obj.FieldFloat)
}

func TestPartiallyRandomBasicString(t *testing.T) {
    var obj string = ""

    err := struct_mommy.Make(&obj, struct_mommy.Define("Blah", 123))

    assert.Equal(t, "Define on string does not make sense. Only structs are accepted", err.Error())
    assert.NotEmpty(t, obj)
}
