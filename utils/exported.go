package utils

//ExportedType to test promoted methods
type ExportedType struct {
	AnotherType
}

//AnotherType is just a type with an unexported name
type AnotherType struct {
	name string
}

//SetName a setter for name
func (e *AnotherType) SetName(n string) {
	e.name = n
}

//Hi returns a slaute
func (e AnotherType) Hi() string {
	return "Hi " + e.name
}
