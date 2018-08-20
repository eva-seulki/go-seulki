// 1. 변수
// var parameter-name type
var a int

// 변수 선언과 동시에 초기화
var f float32 = 11.

// 선언된 변수에 값 할당
a = 10
f = 12.0

// 동일 타입 변수가 여러개일 때
var i, j, k int

// 복수 변수 선언 및 초기화
var i, j, k int = 1, 2, 3

/**
** 초기화하지 않으면 Go는 default로 'Zero Value'를 할당함
** 숫자형 0, boolean false, string형 "" ,, Java와 동일함
**/

// 타입추론 기능 : 자바스크립트와 동일
var i = 1 var s = "Hi" // i에는 정수형, s에는 문자열로 할당

/**
** 변수를 선언하는 또다른 방식 'Short Assignment Statement ( := )'
** 즉, var i = 1을 쓰는대신 i:=1로 사용하여 var를 생략할 수 있음
** 단, 이런 표현은 함수(func) 내에서만 사용할 수 있음
**/

//--------------------------------------------------------------------------------------------

// 2. 상수
// 'const'keyword name type Value
const c int = 10
const s string = "Hi"

// 타입 추론
const c = 10
const s = "Hi"

// 상수만 묶어서 지정
const (
  Visa = "Visa"
  Master = "MasterCard"
  Amex = "American Express"
)

// Tip : iota 라는 identifier를 사용하여 순차적으로 증가된 값 할당 가능
const (
  Apple = iota //0
  Grape        //1
  Orange       //2
)

// Go keyword
// 25개의 예약 키워드를 갖으며, 이들은 변수명, 상수명, 함수명 등의 identifier로 사용할 수 없음
// break      default       func      interface   select
// case       defer         go        map         struct
// chan       else          goto      package     switch
// const      fallthrough   if        range       type
// continue   for           import    return      var
