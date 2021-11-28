package object

type Object struct {
	count int
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Update() {}
