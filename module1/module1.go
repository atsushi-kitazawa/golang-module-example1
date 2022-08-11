package module1

type Foo1 struct{}

func (f *Foo1) foo() string {
	return "foo function"
}

func New() *Foo1 {
	return &Foo1{}
}
