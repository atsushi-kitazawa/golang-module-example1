package module1

type Foo1 struct{}

func (f *Foo1) Foo() string {
	return "foo function"
}

func New() *Foo1 {
	return &Foo1{}
}
