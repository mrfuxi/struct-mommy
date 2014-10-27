package structmommy_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/mrfuxi/struct-mommy"
)

func TestPartiallyRandom(t *testing.T) {
    obj := testMultipleSimpleFields{}

    structmommy.Make(&obj, structmommy.Define("FieldFloat", 2.0))

    assert.NotEmpty(t, obj.FieldInt)
    assert.NotEmpty(t, obj.FieldStr)
    assert.Equal(t, obj.FieldFloat, 2.0)
}

func TestPartiallyRandomTypeConversion(t *testing.T) {
    obj := testMultipleSimpleFields{}

    structmommy.Make(&obj, structmommy.Define("FieldInt", 2.0, "FieldStr", "1"))

    assert.Equal(t, obj.FieldInt, 2)
    assert.Equal(t, obj.FieldStr, "1")
    assert.NotEmpty(t, obj.FieldFloat)
}

func TestPartiallyRandomInvalidField(t *testing.T) {
    obj := testMultipleSimpleFields{}

    err := structmommy.Make(&obj, structmommy.Define("INVALID", 1))

    assert.NotEmpty(t, obj.FieldInt, 2)
    assert.NotEmpty(t, obj.FieldStr, "1")
    assert.NotEmpty(t, obj.FieldFloat)
    assert.EqualError(t, err, "could not find field 'INVALID' in structmommy_test.testMultipleSimpleFields struct")
}

func TestPartiallyRandomBasicString(t *testing.T) {
    var obj string

    err := structmommy.Make(&obj, structmommy.Define("Blah", 123))

    assert.EqualError(t, err, "define on string does not make sense; only structs are accepted")
    assert.NotEmpty(t, obj)
}

func TestPartiallyRandomEmbeddedNamedStruct(t *testing.T) {
    obj := testNamedEmbeddedStructs{}

    err := structmommy.Make(
        &obj,
        structmommy.Define("FInt.Field", 123),
    )

    assert.NoError(t, err)
    assert.Equal(t, 123, obj.FInt.Field)
}

func TestPartiallyRandomEmbeddedNamedStructInvalidName(t *testing.T) {
    obj := testNamedEmbeddedStructs{}

    err := structmommy.Make(
        &obj,
        structmommy.Define("FInt.INVALID", 123),
    )

    assert.EqualError(t, err, "could not find field 'INVALID' in structmommy_test.testSingleInt struct")
}

func TestPartiallyRandomEmbeddedNamedStructInvalidType(t *testing.T) {
    obj := testNamedEmbeddedStructs{}

    err := structmommy.Make(
        &obj,
        structmommy.Define("FInt.Field.X", 123),
    )

    assert.EqualError(t, err, "define on int does not make sense; only structs are accepted")
}

func TestPartiallyRandomDeepEmbeddedNamedStruct(t *testing.T) {
    obj := struct {
        A struct {
            B struct {
                C struct {
                    Value string
                }
            }
        }
    }{}

    err := structmommy.Make(
        &obj,
        structmommy.Define("A.B.C.Value", "Bazinga!"),
    )

    assert.NoError(t, err)
    assert.Equal(t, obj.A.B.C.Value, "Bazinga!")
}

func TestPartiallyRandomSetStruct(t *testing.T) {
    obj := struct {
        A struct {
            B struct {
                C testSingleString
            }
        }
    }{}

    newVal := testSingleString{Field: "Test"}
    err := structmommy.Make(
        &obj,
        structmommy.Define("A.B.C", newVal),
    )

    assert.NoError(t, err)
    assert.Equal(t, obj.A.B.C, newVal)
}
