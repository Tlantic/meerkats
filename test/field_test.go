package meerkats__test

import (
	"testing"
	. "github.com/Tlantic/meerkats"
)

var key string = "test"
var valueTests = []struct {
	input       interface{}
	input_kind	FieldType
	output		interface{}
	output_kind	FieldType
	string		string
}{
	{nil, TypeInterface, nil, TypeInterface, "<nil>"},
	{"qwerty", String, "qwerty", String, "qwerty"},
	{int8(8), TypeInt8, int64(8), Int64, "8"},
	{int16(8), TypeInt16, int64(8), Int64, "8"},
	{int32(8), TypeInt32, int64(8), Int64, "8"},
	{int(8), Int, int64(8), Int64, "8"},
	{int64(8), Int64, int64(8), Int64, "8"},

	{uint8(8), TypeUint8, uint64(8), Uint64, "8"},
	{uint16(8), TypeUint16, uint64(8), Uint64, "8"},
	{uint32(8), TypeUint32, uint64(8), Uint64, "8"},
	{uint(8), Uint, uint64(8), Uint64, "8"},
	{uint64(8), Uint64, uint64(8), Uint64, "8"},

	{float32(8.33333333), Float32, float32(8.33333333), Float32, "8.333333"},
	{float64(8.33333333), Float64, float64(8.33333333), Float64, "8.33333333"},
}


func TestNew(t *testing.T) {
	for _, vt := range valueTests {
		f := NewField(key, vt.input)

		if f.Type != vt.output_kind {
			t.Fail()
		}
	}
}

func TestField_GetType(t *testing.T) {
	for _, vt := range valueTests {
		f := NewField(key, vt.input)
		if f.Type != vt.output_kind {
			t.Fail()
		}
		if f.GetType() != vt.output_kind {
			t.Fail()
		}
	}
}

func TestField_Get(t *testing.T) {
	for _, vt := range valueTests {
		f := NewField(key, vt.input)
		if f.Get() != vt.output {
			t.Fail()
		}
	}
}
func TestField_Set(t *testing.T) {
	for _, vt := range valueTests {
		f := NewField(key, nil)
		f.Set(vt.input)
		if ( f.Get() != vt.output )  {
			t.Fail()
		}
	}
}

func TestField_SetString(t *testing.T) {
	sample := "qwerty"
	v := NewField(key, nil)
	v.SetString(sample)
	if v.ValueString != sample {
		t.Fail()
	}
}
func TestField_GetString(t *testing.T) {
	sample := "qwerty"
	v := NewField(key, nil)
	v.SetString(sample)
	if v.GetString() != sample {
		t.Fail()
	}
}

func TestField_SetBool(t *testing.T) {
	sample := true
	v := NewField(key, nil)
	v.SetBool(sample)
	if v.ValueBool != sample {
		t.Fail()
	}
}
func TestField_GetBool(t *testing.T) {
	sample := true
	v := NewField(key, nil)
	v.SetBool(sample)
	if v.GetBool() != sample {
		t.Fail()
	}
}

func TestField_SetInt(t *testing.T) {
	sample := int(16)
	v := NewField(key, nil)
	v.SetInt(sample)
	if v.ValueInt64 != int64(sample) {
		t.Fail()
	}
}
func TestField_GetInt(t *testing.T) {
	sample := int(16)
	v := NewField(key, nil)
	v.SetInt(sample)
	if ( v.GetInt() != sample) {
		t.Fail()
	}
}

func TestField_SetInt64(t *testing.T) {
	sample := int64(16)
	v := NewField(key, nil)
	v.SetInt64(sample)

	if ( v.ValueInt64 != int64(sample)) {
		t.Fail()
	}
}
func TestField_GetInt64(t *testing.T) {
	sample := int64(16)
	v := NewField(key, nil)
	v.SetInt64(sample)

	if ( v.GetInt64() != sample) {
		t.Fail()
	}
}

func TestField_SetUint(t *testing.T) {
	sample := uint(16)
	v := NewField(key, nil)
	v.SetUint(sample)

	if ( v.ValueUint64 != uint64(sample)) {
		t.Fail()
	}
}
func TestField_GetUint(t *testing.T) {
	sample := uint(16)
	v := NewField(key, nil)
	v.SetUint(sample)

	if ( v.GetUint() != sample) {
		t.Fail()
	}
}

func TestField_SetUint64(t *testing.T) {
	sample := uint64(16)
	v := NewField(key, nil)
	v.SetUint64(sample)

	if ( v.ValueUint64 != uint64(sample)) {
		t.Fail()
	}
}
func TestField_GetUint64(t *testing.T) {
	sample := uint64(16)
	v := NewField(key, nil)
	v.SetUint64(sample)

	if ( v.GetUint64() != uint64(sample)) {
		t.Fail()
	}
}

func TestField_SetFloat32(t *testing.T) {
	sample := float32(16.666666)
	v := NewField(key, nil)
	v.SetFloat32(sample)

	if ( v.ValueFloat32 != sample) {
		t.Fail()
	}
}
func TestField_GetFloat32(t *testing.T) {
	sample := float32(16.666666)
	v := NewField(key, nil)
	v.SetFloat32(sample)

	if ( v.GetFloat32() != sample) {
		t.Fail()
	}
}

func TestField_SetFloat64(t *testing.T) {
	sample := float64(16.666666)
	v := NewField(key, nil)
	v.SetFloat64(sample)

	if ( v.ValueFloat64 != sample) {
		t.Fail()
	}
}
func TestField_GetFloat64(t *testing.T) {
	sample := float64(16.666666)
	v := NewField(key, nil)
	v.SetFloat64(sample)

	if ( v.GetFloat64() != sample) {
		t.Fail()
	}
}

func TestField_SetInterface(t *testing.T) {
	sample := struct{}{}
	v := NewField(key, nil)
	v.SetInterface(sample)

	if ( v.ValueInterface != sample) {
		t.Fail()
	}
}
func TestField_GetInterface(t *testing.T) {
	sample := struct{}{}
	v := NewField(key, nil)
	v.SetInterface(sample)

	if ( v.GetInterface() != sample) {
		t.Fail()
	}
}

func TestField_String(t *testing.T) {
	for _, vt := range valueTests {
		f := NewField(key, vt.input)
		if f.String() != vt.string {
			t.Logf("Expected: %s but got %s\n", vt.string, f.String())
			t.Fail()
		}
	}
}
