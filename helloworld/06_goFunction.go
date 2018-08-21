// 1. Go Function
// func 함수명 (파라미터) 리턴타입 {}

// 2. Pass By Reference
// 파라미터 전달 방식 : Pass By Value, Pass By Reference

// - 1. Pass By Value
//      아래 예제에서 msg 값 "Hello" 문자열이 복사되어 say() 함수에 전달됨
//      만약 파라미터 msg 값이 say() 함수 내에서 변경되더라도 호출함수 main()에서의 msg 변수는 변함이 없음

package main

func main1() {
	msg := "Hello"
	say1(msg)
}

func say1(msg string) { // say함수, 문자열 msg 파라미터 한개, 리턴 없음, 리턴타입 없음
	println(msg)
}

// - 2. Pass By Reference
//      아래 예제에서 msg 변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시하게 됨
//      포인터라 불리는 이 용법을 사용하면 함수에 msg 변수의 값을 복사하지 않고 msg 변수의 주소를 전달함
//      피호출 함수 say()에서는 *string과 같이 파라미터가 포인터임을 표시하고, 이때 say()함수의 msg 는 문자열을 갖는 메모리영역 주소를 갖게 됨
//      msg 주소에 데이터를 쓰기 위해서는 *msg = ""와 같이 앞에 *를 붙이는데 이를 흔히 Deferencing 이라고 함

func main2() {
	msg := "Hello"
	say2(&msg)
	println(msg) //변경된 메시지 출력
}

func say2(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

// - 3. Variadic Function (가변인자함수)
//      함수에 고정된 수의 파라미터가 아닌 다양한 숫자의 파라미터 전달 시 가변 파라미터를 나타내는 ... 을 사용함
func main3() {
	say3("This", "is", "a", "book")
	say3("Hi")
}

// _(underscore) 변수는 이것이 필요하지 않다고 알려주는 데 사용함
// for문 내부에서 사용하지 않는 변수를 넣을 경우 컴파일 에러가 발생하기 때문에 _ 변수를 쓴 것임
func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

// - 4. 함수 리턴 값
/** 특징
**     1. 복수 개의 리턴도 가능함 (물론 리턴값은 없을 수도, 한개일 수도 있음)
**     2. Named Return Parameter 기능 제공 (리턴되는 값들을 함수에 정의된 리턴 파라미터들에 할당할 수 있는 기능)
**/
func main4() {
	total := sum4(1, 7, 3, 5, 9)
	println(total)
}

func sum4(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

// 복수개의 리턴을 위해 리턴 타입들을 괄호() 안에 적어 줌
func main5() {
	count, total := sum5(1, 7, 3, 5, 9)
	println(count, total)

	cnt, tot := sum6(1, 3, 2, 4)
	println(cnt, tot)
}

func sum5(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 개수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

// Named Return Parameter들에 리턴값을 할당하여 리턴하면, 리턴 값이 여러개일 때 코드 가독성을 높일 수 있음
// 빈 return 문 생략 시 에러 발생
func sum6(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
