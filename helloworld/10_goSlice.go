// 1. 슬라이스(Slice)
//    배열은 고정된 배열크기 안에 동일한 타입 데이터를 연속적으로 저장하지만, 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등 기능은 없음
//    Slice는 배열에 기초해 만들어졌지만, 배열과 다른점은
//      (1) 고정된 크기를 미리 지정하지 않을 수 있음
//      (2) 차후 크기를 동적으로 변경 가능함
//      (3) 부분 배열을 발췌할 수 있음

// 선언 : 배열처럼 "var v {}T"처럼 하지만 배열과 달리 크기는 지정하지 않음
package main

import "fmt"

func main1() {
	var a []int        // 슬라이스 변수 선언
	a = []int{1, 2, 3} // 슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a)
}

// 1-1. make() : Go에서 Slice를 생성하는 내장함수
// 장점 : make() 함수로 슬라이스 생성 시 슬라이스의 길이(Length)와 용량(Capacity)을 임의로 지정할 수 있음
// 사용 : make(슬라이스 타입, Length, Capacity), Capacity 생략 시 Capacity=Length, len(), cap()함수로 확인
func main2() {
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s)) // len 5, cap 10
}

// Nill Slice : Slice에 길이,용량 미지정하여 default로 길이와 용량이 0으로 생성된 슬라이스. nil과 비교하면 true 리턴
func main3() {
	var s []int

	if s == nil {
		var bool = s == nil
		fmt.Println(bool)
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(s), cap(s)) // 0
}

//--------------------------------------------------------------------------------------------
// 2. 부분슬라이스(Sub-slice)
// 슬라이스에서 일부를 발췌함
// 슬라이스[처음 idx, 마지막 idx] 형식
// 슬라이스 s의 2~4까지의 부분슬라이스 생성 시 s[2:5]로 표현함
// 즉, 첫 인덱스는 Inclusive, 마지막 인덱스는 Exclusive임 (Python과 동일)
// 첫 인덱스 생략(0~4)시 [:5], 마지막 인덱스 생략(2~마지막)시 [2:] 로 표현 가능, 모두 생략 시[:] 슬라이스 전체를 의미함
func main4() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]     // 2, 3, 4
	s = s[1:]      // 3, 4
	fmt.Println(s) // 3, 4
}

//--------------------------------------------------------------------------------------------
// 3. 슬라이스 추가, 병합, 복사
// 배열은 고정된 크기로 인해, 그 크기 이상을 추가할 수 없지만, 슬라이스는 가능함
// 3-1. append() : 슬라이스에 새로운 요소를 추가하는 Go 내장함수, append(슬라이스, 추가 요소 한개부터 여러개까지 콤마로 추가 가능)
func main5() {
	s := []int{0, 1}

	// 한개
	s = append(s, 2) // 0, 1, 2
	// 복수개
	s = append(s, 3, 4, 5) // 0, 1, 2, 3, 4, 5
	fmt.Println(s)
}

// append()의 내부 프로세스 : 슬라이스 용량(capacity)이 아직 남아 있는 경우는 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가하고,
//                        용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array을 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당함
//                        아래 예제에서 길이 0/용량 3의 슬라이스에 1부터 15까지의 숫자를 계속 추가하면서 슬라이스의 길이와 용량이 어떻게 변하는 지를 체크함
//                        코드 실행시 1~3까지는 기존의 용량 3을 사용하고, 4~6까지는 용량 6을, 7~12는 용량 12, 그리고 13~15는 용량 24의 슬라이스가 사용되고 있음을 알 수 있음
func main6() {
	// len = 0, cap = 3 인 슬라이스 생성
	sliceA := make([]int, 0, 3)

	// 요소를 순차적 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(sliceA)
	}
}

// 한 슬라이스를 다른 슬라이스에 병함할 때에도 append()를 사용함. 단 ellipsis(...)를 잊지 말 것
// ellipsis는 해당 슬라이드의 컬렉션을 표현하는 것으로 모든 요소의 집합을 나타냄
func main7() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	// sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6]
}

// 3-2. copy() : 슬라이스를 복사함
//               슬라이스는 실제 배열을 가리키는 포인터 정보만 가지므로, 정확히 말해 copy()는 소스 슬라이스가 갖는 배열 데이터를 타겟 슬라이스가 갖는 배열로 복제하는 것임
func main8() {
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2) // 소스 슬라이스와 Lengh가 같고 Capacity가 2배인 슬라이스 생성
	copy(target, source)
	fmt.Println(target)               // [0, 1, 2]
	println(cap(source), cap(target)) // 3, 6
}

//--------------------------------------------------------------------------------------------
// 4. 슬라이스 내부구조
//    s := []int{1, 2, 3, 4, 5} 실행 시 슬라이드
//    [배열 포인터][Length(6)][Capacity(6)]
//    처음 슬라이스가 생성될 때, 만약 길이와 용량이 지정되었다면, 내부적으로 용량(Capacity)만큼의 배열을 생성하고, 슬라이스 첫번째 필드에 그 배열의 처음 메모리 위치를 지정함
//    그리고, 두번째 길이 필드는 지정된 (첫 배열요소로부터의) 길이를 갖게되고, 세번째 용량 필드는 전체 배열의 크기를 갖음

//    s := s[2:5] 실행시 변경된 슬라이드
//    [배열 포인터][Length(3)][Capacity(4)] *****