package main

import "fmt"

type Bread struct {
	val string
}
type OrangeJam struct {
	opend bool
}
type SpoonOfOrange struct { //구조체
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
func openOrangeJam(jam *OrangeJam) {
	jam.opend = true
}

func GetOneSpoon(_ *OrangeJam) *SpoonOfOrange {
	return &SpoonOfOrange{}

}
func PutJamOnBread(bread *Bread, jam *SpoonOfOrange) {
	for i := 0; i < 1; i++ { //쨈 바르는 횟수 수정하는 곳
		bread.val += "+ Orange Jam(쨈)"
	}

}

func MakeSandwich(bread []*Bread) *sandwich {
	sandwich := &sandwich{}
	for i := 0; i < len(bread); i++ {
		sandwich.val += bread[i].val + "+"
	}
	return sandwich
}

func main() {

	breads := GetBreads(10) // 빵 두개 꺼내기

	jam := &OrangeJam{}

	openOrangeJam(jam) //딸기잼 뚜껑열기

	spoon := GetOneSpoon(jam) //빵 위에 올린다

	// PutJamOnBread(breads[0], spoon) //잼을 바른다
	// PutJamOnBread(breads[0], spoon) //잼을 한번만 더 바를때
	for i := 0; i < len(breads); i++ {
		PutJamOnBread(breads[i], spoon)
		//빵 갯수에 따라 쨈의 갯수가 달라지기 때문에
		//50번째 라인 getBreads(숫자)의 숫자를 수정해야함
	}

	sandwich := MakeSandwich(breads) //빵을 덮는다
	//jam더 올리기

	fmt.Println(sandwich.val) //완성
}
