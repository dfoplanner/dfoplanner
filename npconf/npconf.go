package npconf


// Neople uses a ini accent for their configuration
// This package implements a simple parser for those format
// TODO: Add error handling

type NpConf struct {
	Properties []*Property `@@*`
}

type Property struct {
	Key string `"[" @Ident "]"`
	Value *Value `@@*`
}

type Value struct {
	String *string `@StringPointer`
}
