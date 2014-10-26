package struct_mommy_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/fuxi/struct-mommy"
)

type testSingleBool struct{ Field bool }
type testSingleByte struct{ Field byte }
type testSingleComplex64 struct{ Field complex64 }
type testSingleComplex128 struct{ Field complex128 }
type testSingleFloat32 struct{ Field float32 }
type testSingleFloat64 struct{ Field float64 }
type testSingleInt struct{ Field int }
type testSingleInt8 struct{ Field int8 }
type testSingleInt16 struct{ Field int16 }
type testSingleInt32 struct{ Field int32 }
type testSingleInt64 struct{ Field int64 }
type testSingleRune struct{ Field rune }
type testSingleString struct{ Field string }
type testSingleUint struct{ Field uint }
type testSingleUint8 struct{ Field uint8 }
type testSingleUint16 struct{ Field uint16 }
type testSingleUint32 struct{ Field uint32 }
type testSingleUint64 struct{ Field uint64 }
type testSingleUintptr struct{ Field uintptr }

type testMultipleSimpleFields struct {
    FieldInt   int8
    FieldStr   string
    FieldFloat float64
}

type testAnonymousEmbeddedStructs struct {
    testSingleInt
    testSingleString
}

type testNamedEmbeddedStructs struct {
    FInt testSingleInt
    FStr testSingleString
}

func TestStructWithSingleFieldBool(t *testing.T) {
    obj := testSingleBool{}

    ok := false
    for i := 0; i <= 100; i++ {
        struct_mommy.Make(&obj)
        ok = obj.Field

        if ok {
            break
        }
    }

    assert.True(t, ok)

    ok = true
    for i := 0; i <= 100; i++ {
        struct_mommy.Make(&obj)
        ok = obj.Field

        if !ok {
            break
        }
    }

    assert.False(t, ok)
}

func TestStructWithSingleFieldByte(t *testing.T) {
    obj := testSingleByte{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldComplex64(t *testing.T) {
    obj := testSingleComplex64{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldComplex128(t *testing.T) {
    obj := testSingleComplex128{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldFloat32(t *testing.T) {
    obj := testSingleFloat32{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldFloat64(t *testing.T) {
    obj := testSingleFloat64{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldInt(t *testing.T) {
    obj := testSingleInt{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldInt8(t *testing.T) {
    obj := testSingleInt8{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldInt16(t *testing.T) {
    obj := testSingleInt16{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldInt32(t *testing.T) {
    obj := testSingleInt32{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldInt64(t *testing.T) {
    obj := testSingleInt64{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldRune(t *testing.T) {
    obj := testSingleRune{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldString(t *testing.T) {
    obj := testSingleString{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUint(t *testing.T) {
    obj := testSingleUint{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUint8(t *testing.T) {
    obj := testSingleUint8{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUint16(t *testing.T) {
    obj := testSingleUint16{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUint32(t *testing.T) {
    obj := testSingleUint32{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUint64(t *testing.T) {
    obj := testSingleUint64{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestStructWithSingleFieldUintptr(t *testing.T) {
    obj := testSingleUintptr{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.Field)
}

func TestSimpleInt(t *testing.T) {
    var obj uint8 = 0
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj)
}

func TestSimpleStr(t *testing.T) {
    var obj string = ""
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj)
}

func TestMultipleSimpleFields(t *testing.T) {
    obj := testMultipleSimpleFields{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.FieldInt)
    assert.NotEmpty(t, obj.FieldStr)
    assert.NotEmpty(t, obj.FieldFloat)
}

func TestAnonymousEmbeddedStructs(t *testing.T) {
    obj := testAnonymousEmbeddedStructs{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.testSingleInt.Field)
    assert.NotEmpty(t, obj.testSingleString.Field)
}

func TestNamedEmbeddedStructs(t *testing.T) {
    obj := testNamedEmbeddedStructs{}
    struct_mommy.Make(&obj)
    assert.NotEmpty(t, obj.FInt.Field)
    assert.NotEmpty(t, obj.FStr.Field)
}
