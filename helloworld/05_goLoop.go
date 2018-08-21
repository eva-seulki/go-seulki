// 1. for 문
// Go에서 반복문은 for 루프 뿐
// - 특징 : "초기값;조건식;증감"을 둘러싸는 괄호()를 생략하는데, 괄호쓰면 에러남

package main

// 실행하고자하는 함수를 mainX -> main으로 바꿔서 실행하자.
func main0() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

// 2. for문 - 조건식만 쓰는 for Loop
// 초기값과 증감식을 생략하고 조건식만 사용할 수 있는데, 마치 다른 언어의 while Loop 처럼 쓰임

func main1() {
	n := 1
	for n < 100 {
		n *= 2
		//if n > 90 {
		//  break
		//}
	}
	println(n)
}

// 3. for문 - 무한루프
// 무한루프를 만들려면 "초기값; 조건식; 증감"을 모두 생략하면 됨. 빠져나오려면 Ctrl+c
//func main2() {
//for {
//  println("Infinite loop")
//}
//}

// 4. for range 문
// Collection에서 한 element씩 가져와 차례로 for 블럭을 실행하며, 다른언어의 foreach와 비슷함
// "for 인덱스, 요소값 := range 컬렉션"

func main3() {
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}

// 5. break, continue, togo
// break : for Loop에서 즉시 빠져나옴 (switch, select문에도 사용)
// continue : for Loop 중간에서 나머지 문장을 실행하지 않고 for Loop 시작부분으로 바로 감
// goto : 기타 임의의 문장으로 이동 (for Loop와 상관없이 사용)

func main4() {
	var a = 11
	for a < 15 {
		if a == 5 {
			a += a
			continue // for Loop 시작으로
		}
		a++
		if a > 10 {
			break // Loop를 빠져나옴
		}
	}
	if a == 11 {
		goto END // goto 사용 예
	}
	println(a)

END:
	println("End")
}

/**
** break문은 보통 단독으로 사용되지만, 경우에 따라 "break 레이블"과 같이 사용하여 지정된 레이블로 이동할 수도 있다.
** break의 "레이블"은 보통 현재의 for 루프를 바로 위에 적게 되는데, 이러한 "break 레이블"은 현재의 루프를 빠져나와 지정된 레이블로 이동하고,
** break문의 직속 for 루프 전체의 다음 문장을 실행하게 한다. 아래 예제는 언뜻 보기에 무한루프를 돌 것 같지만,
** 실제로는 OK를 출력하고 프로그램을 정상 종료한다. 이는 "break L1" 문이 for 루프를 빠져나와 L1 레이블로 이동한 후,
** break가 있는 현재 for 루프를 건너뛰고 다음 문장인 println() 으로 이동하기 때문이다.
**/

func main5() {
	i := 0

L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("OK")
}
