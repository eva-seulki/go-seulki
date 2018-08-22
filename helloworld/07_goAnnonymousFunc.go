// 1. 익명함수
//    함수명을 갖지 않는 함수. 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어 사용됨
//    아래 예제를 보면, 익명함수가 선언되어 sum이라는 변수에 할당되고 있음을 보여줌
//    할당된 익명함수를 보면 func 바로 뒤에 함수명이 생략되었음을 알 수 있으며, 이 점 제외하고 함수의 정의 내용은 동일함
//    익명함수가 변수에 할당되면 이후 변수명이 함수명과 같이 취급되어 "변수명(파라미터)" 형식으로 호출할 수 있음
package main

func main1() {
  sum := func(n ...int) int { // 익명함수 정의
    s := 0
    for _, i := range n {
      s += i
    }
    return s
  }

  result := sum(1, 2, 3, 4, 5) // 익명함수 호출
  println(result)
}

// 2. 일급함수
//    Go의 함수는 Java와 다르게 일급함수임. 따라서 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로 사용될 수 있음
//    함수를 다른 함수의 파라미터로 전달하기 위해서는 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법과 직접 다른 함수 호출 파라미터에 함수를 적는 방법 둘 다 됨
func main2() {
  // 변수 add에 익명함수 할당
  add := func (i int, j int) int {
    return i + j
  }

  // add 함수 전달
  r1 := calc(add, 10, 20)
  println(r1)

  // 직접 첫번째 파라미터에 익명함수를 정의함
  r2 := calc (func(x int, y int) int { return x - y }, 10, 20)
  println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
  result := f(a, b)
  return result
}

// 3. type문을 사용한 함수 원형 정의
//    type문은 구조체, 인터페이스 등 Custom Type (혹은 User Defined Type)을 정의하기 위해 사용됨
//    함수 원형 정의하는 데도 사용될 수 있음. 위 예제에서 func(x int, y int) int 함수 원형이 코드상 반복되는 데 type문을 정의하여 간단히 표현
type calculator func(int, int) int // 원형 정의

func calc(f calculator, a int, b int) int { // calculator 원형 사용
  result := f(a, b)
  return result
}
