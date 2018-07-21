package pack

type A struct {
	Name string
}

func changeA(a A) {
	a.Name = "0"
}

func changeAByP(a *A) {
	a.Name = "0"
}
