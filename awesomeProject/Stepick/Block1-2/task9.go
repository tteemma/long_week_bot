package main

import "fmt"

type testStruct struct {
	On    bool
	Ammo  int
	Power int
}

func newTestStruct(on bool, ammo int, power int) *testStruct {
	return &testStruct{On: on, Ammo: ammo, Power: power}
}

func (t *testStruct) Shoot() bool {
	if t.On {
		if t.Ammo > 0 {
			t.Ammo--
			return true
		}
		return false
	}
	return false
}
func (t *testStruct) RideBike() bool {
	if t.On {
		if t.Power > 0 {
			t.Power--
			return true
		}
		return false
	}
	return false
}

func main() {
	testStruct := newTestStruct(true, 1, 1)

	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}
