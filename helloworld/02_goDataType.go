// 1. Go Data Type
// - 1. bool
// - 2. 문자열 : 한번 생성되면 수정될 수 없는 Immutable 타입
// - 3. 정수형 : int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
// - 4. Float & 복소수 : float32 float64 complex64 complex128
// - 5. ETC : byte (unit8과 동일하며 바이트코드에 사용), rune (int32와 동일하며 유니코드 코드포인트에 사용)

// 2. 문자열
// - 1. Back Quote(``) : Raw String Literal이라 부르며, 이 안의 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖음.
//                       예를 들어, 문자열 안에 \n이 있을 경우 NewLine으로 해석되지 않는다. Back Quote는 복수 라인의 문자열 표현에 자주 사용됨
// - 2. Double Quatation("") : Interpreted String Literal이라 부르며, Escape 문자열은 특별한 의미로 해석됨.
//                             \n 의 경우 NewLine으로 해석 됨. 여러 라인을 쓰기 위해선 + 연산자를 이용함.
package main

import "fmt"

func main() {
	// Raw Wtring Literal 복수라인
	rawLiteral := `아리랑\n
아리랑\n
                     아라리요`

	// Interpreted String rawLiteral
	interLiteral := "아리랑 아리랑\n 아라리요"
	//또는 아래처럼 + 연산자를 이용
	// interLiteral := "아리랑 아리랑\n" + "아라리요"
	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}

// 3. 데이터타입 변환 (Type Conversion)
// 데이터 타입 변환을 위해 T(v)와 같이 표현하고 Type Conversion이라고 부름.
// 예를 들어, 정수 100을 float으로 변경하기 위해 float32(100), 문자열을 바이트배열로 변경하기 위해 []byte("ABC")로 표현 할 수 있음.

/**
** 주의할 점은! Go에서 타입 변환은 명시적으로 지정해 주어야 함. 즉, int에서 unit으로 implicit 변환이 일어나지 않고 runtime error 발생.
**/

func main2() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
