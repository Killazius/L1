package main

import "fmt"

type Human struct {
	name string
	Age  int
}

func (h Human) Info() string {
	return fmt.Sprintf("Human %s, %d!", h.name, h.Age)
}

// Action имеет доступ к полям Human и методам. При вызове Action.Info(), будет вызван Action.Human.Info()
type Action struct {
	Human
}

func (a Action) DoAction() string {
	return fmt.Sprintf("Action %s, Do Action for %d!", a.name, a.Age)
}
func main() {
	h := Human{
		name: "killazius",
		Age:  20,
	}
	fmt.Println(h.Info())
	action := Action{
		Human: h,
	}
	fmt.Println(action.DoAction())
	fmt.Println(action.Info())
}
