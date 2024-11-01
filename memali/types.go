package memali

// Структура для int8
type Int8Field struct {
	Name  string
	Value int8
}

// Структура для int16
type Int16Field struct {
	Name  string
	Value int16
}

// Структура для int32
type Int32Field struct {
	Name  string
	Value int32
}

// Структура для int64
type Int64Field struct {
	Name  string
	Value int64
}

type IntField struct {
	Name  string
	Value int
}

// Структура для uint8
type Uint8Field struct {
	Name  string
	Value uint8
}

// Структура для uint16
type Uint16Field struct {
	Name  string
	Value uint16
}

// Структура для uint32
type Uint32Field struct {
	Name  string
	Value uint32
}

// Структура для uint64
type Uint64Field struct {
	Name  string
	Value uint64
}

type UintField struct {
	Name  string
	Value uint
}

// Структура для float32
type Float32Field struct {
	Name  string
	Value float32
}

// Структура для float64
type Float64Field struct {
	Name  string
	Value float64
}

type Complex64Field struct {
	Name  string
	Value complex64
}

type Complex128Field struct {
	Name  string
	Value complex128
}

type BoolField struct {
	Name  string
	Value bool
}

// Структура для string
type StringField struct {
	Name  string
	Value string
}

type ArrayField struct {
	Name  string
	Value []interface{}
}

type SliceField struct {
	Name  string
	Value []interface{}
}

type MapField struct {
	Name  string
	Value map[interface{}]interface{}
}

type StructField struct {
	Name   string
	Fields []interface{}
}

type PointerField struct {
	Name  string
	Value interface{}
}

type FuncType struct {
	Name  string
	Value interface{}
}

type ChanType struct {
	Name  string
	Value interface{}
}
