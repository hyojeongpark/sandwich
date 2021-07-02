package main

import "fmt"

type Bread struct {
	val string
}
type StrawberryJam struct {
	opend bool
}
type SpoonOfStrawberry struct {
}
type sandwich struct {
	val string
}

func GetBreads(num int) []*Bread {

	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread(빵)"}
	}
	return breads
}
func openStrawberryJam(jam *StrawberryJam) {
	jam.opend = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}

}
func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += "+ Strawberry Jam(쨈)"

}

func MakeSandwich(bread []*Bread) *sandwich {
	sandwich := &sandwich{}
	for i := 0; i < len(bread); i++ {
		sandwich.val += bread[i].val + "+"
	}
	return sandwich
}

func main() {

	breads := GetBreads(2)

	jam := &StrawberryJam{} //잼을열고

	openStrawberryJam(jam) //딸기잼 뚜껑열기

	spoon := GetOneSpoon(jam)

	PutJamOnBread(breads[0], spoon)
	//PutJamOnBread(breads[0], spoon)
	PutJamOnBread(breads[1], spoon) //잼을 한번만 더 바를때
	// for i := 0; i < len(breads); i++ {
	// 	PutJamOnBread(breads[i], spoon)
	// }

	sandwich := MakeSandwich(breads)
	//jam더 올리기

	fmt.Println(sandwich.val)
}
