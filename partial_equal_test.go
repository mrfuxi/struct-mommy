package structmommy_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/mrfuxi/struct-mommy"
)

func TestPartiallyEqualOneField(t *testing.T) {
    obj := testMultipleSimpleFields{
        FieldInt:   1,
        FieldFloat: 2.0,
    }

    err := structmommy.PartialEqual(&obj, "FieldFloat", 2)
    assert.NoError(t, err)
    assert.Equal(t, obj.FieldFloat, 2)
}

func TestPartiallyEqualManyFieldsAllHasToBeOK(t *testing.T) {
    obj := testMultipleSimpleFields{
        FieldInt:   1,
        FieldFloat: 2.0,
        FieldStr:   "ABC",
    }

    err := structmommy.PartialEqual(&obj, "FieldFloat", 2, "FieldStr", "ABC")
    assert.NoError(t, err)
    assert.Equal(t, obj.FieldFloat, 2)
    assert.Equal(t, obj.FieldStr, "ABC")
}

func TestPartiallyEqualErrorMsg(t *testing.T) {
    obj := testMultipleSimpleFields{
        FieldInt:   1,
        FieldFloat: 2.0,
        FieldStr:   "123ABC",
    }

    err := structmommy.PartialEqual(&obj, "FieldStr", "CBAXYZ")
    if assert.Error(t, err) {
        // All values should be present. The rest are details
        assert.Contains(t, err.Error(), "FieldStr")
        assert.Contains(t, err.Error(), "123ABC")
        assert.Contains(t, err.Error(), "CBAXYZ")
    }
}

func TestPartiallyEqualNotPointerToObject(t *testing.T) {
    obj := testMultipleSimpleFields{
        FieldInt:   1,
        FieldFloat: 2.0,
    }

    err := structmommy.PartialEqual(obj, "FieldFloat", 2)
    assert.NoError(t, err)
    assert.Equal(t, obj.FieldFloat, 2)
}

func TestPartiallyEqualEmbeddedStructs(t *testing.T) {
    obj := testNamedEmbeddedStructs{}
    obj.FInt.Field = 123
    obj.FStr.Field = "ABC"

    err := structmommy.PartialEqual(obj, "FInt.Field", 123, "FStr.Field", "ABC")
    assert.NoError(t, err)
    assert.Equal(t, obj.FInt.Field, 123)
    assert.Equal(t, obj.FStr.Field, "ABC")
}

func TestPartiallyEqualEmbeddedStructsAllHasToBeOk(t *testing.T) {
    obj := testNamedEmbeddedStructs{}
    obj.FInt.Field = 123
    obj.FStr.Field = "ABC"

    err := structmommy.PartialEqual(obj, "FInt.Field", 123, "FStr.Field", "XYZ")
    assert.Error(t, err)
    assert.Equal(t, obj.FInt.Field, 123)
    assert.Equal(t, obj.FStr.Field, "ABC")
}
