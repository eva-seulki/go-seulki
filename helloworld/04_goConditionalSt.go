

// 1. if 문
// - 특징 : 조건에 괄호 없음. 조건 블럭 시작 브레이스({)를 if 와 같은 라인에 두지 않으면 에러남. 조건은 자바처럼 boolean식

  if k == 1 { // if와 같은 라인!
      println("One")
  } else if k == 2 { // else 문도 시작 브레이스가 같은 라인에 와야 함
      println("Two")
  } else {
      println("Other") // else도 같은 라인!
  }

/**
** if 문에서 조건식 사용 전 간단한 문장(Optional Statement)를 함께 실행할 수 있음. 아래의 val := 1*2가 바로 그것인데,
** 주의할 점은, 이때 정의된 변수는 if문 블럭 안에서만 사용할 수 있음
**/

  if val := i * 2; val < max {
    println(val)
  }
// 아래처럼 사용 시 Scope가 벗어나 error
// val++


/** 예제는 위에 코드 전부 주석쳐야 돌아감 **/

package main
func main() {// 2. switch 문
// 복수개의 case가 있는 경우 콤마로 나열하면 됨
  var name string
  var category = 1

  switch category {
  case 1:
    name = "Paper Book"
  case 2:
    name = "eBook"
  case 3, 4:
    name = "Blog"
  default:
    name = "Other"
  }

  println(name)

  // Expression을 사용한 경우
  switch x := category << 2; x - 1 {
    //...
  }
}

/**
** Go의 특별한 switch문 용법
**
** 1. switch 뒤에 expression이 없을 수 있음 : 이 경우 switch expression을 true로 생각하고 첫 case문으로 이동해 검사함
** 2. case문에 expression을 쓸 수 있음.
** 3. No default fall through : 다른 언어의 경우 break를 쓰지 않는 한 다음 case로 넘어가지만 Go는 넘어가지 않음
** 4. Type switch : 다른 언어의 경우 일반적으로 변수 값을 기준으로 case 분기하나, Go는 변수의 Type에 따라 분기할 수 있음
**/

// 복잡한 if..else 반복문을 단순화하는 데 유용한, switch 뒤에 조건변수 or Expression을 적지 않는 용법
func grade (score int) {
  switch {
  case score >= 90:
    println("A"
  case score >= 80:
    println("B"
  case score >= 70:
    println("C"
  case score >= 60:
    println("D")
  default:
    println("No Hope")
  }
}

// switch 변수의 type 검사
switch v := i.(type) {
case int:
  println("int")
case bool:
  println("bool")
case string:
  println("string")
default:
    println("unknown")
}

// Go 컴파일러가 자동으로 case문 블럭 마지막에 break문을 추가하기 때문에 자동으로 다음 블럭으로 넘어가지 않음.
// 다음 블럭으로 넘어가기 위해선 fallthrough문을 명시해줘야 함.

func check(val int) {
  switch val {
  case 1:
    fmt.println("1 이하")
  case 2:
    fmt.println("2 이하")
  case 3:
    fmt.println("3 이하")
  default:
    fmt.println("default 도달")
  }
}
