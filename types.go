package meerkats

type FieldType uint

const (
	TypeInterface FieldType = iota
	TypeBool
	TypeInt
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt64
	TypeUint
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint64
	TypeUintptr
	TypeFloat32
	TypeFloat64
	TypeComplex64
	TypeComplex128
	TypeString
	TypeUnsafePointer
)