package thing

type Thing interface {
	ThingToDo() string
}

type CoolThing struct {
	Name string
}

func (ct CoolThing) ThingToDo() string {
	return ct.Name
}
