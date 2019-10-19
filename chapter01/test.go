package main

import "fmt"
import "math"
import "unsafe"

func main() {

	// 1. 기본 초기화 변수
	//  정수 타입 : 0, 실수(소수점) :0.0, 문자열 :"", Boolean : True, False
	// 변수명 : 숫자 첫글자 X, 대소문자 구분 o, 문자,숫자, 밑줄, 특수기호 사용 가능
	// 변수 및 상수 : 함수 내외 사용 가능

	var strHelloWorld string = "Hello, World!"
	var nTest int = 1234
	var strTest string = "hello"
	var dTest float64 = 11.4

	fmt.Println(strHelloWorld, nTest, strTest)

	fmt.Println("a : ", dTest)

	// 여러개 변수 한번에 선언
	var (
		name   string = "ddolsoon"
		age    int32  = 27
		height int32  = 100
		width  int32  = 50
	)

	// 병렬 할당 ( parallel assignment )
	var x, y, z int = 10, 20, 30
	fmt.Println("x,y,z : ", x, y, z)

	fmt.Println("name : ", name)
	fmt.Println("age : ", age)
	fmt.Println("height : ", height)
	fmt.Println("width : ", width)

	// 짧은 선언 (:=) ==> var 를 사용하지 않고, 변수 초기화 가능
	shortX := 10
	fmt.Println(shortX)

	// 정수 타입
	var num1 int = 32
	var num2 int = -15
	var num3 int = 010  // 8진수로 저장
	var num4 int = 0xFF // 16진수로 저장
	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)

	// 실수 타입
	var fNum1 float32 = 0.1
	var fNum2 float32 = .35
	fmt.Println("num1 : ", fNum1)
	fmt.Println("num2 : ", fNum2)

	// 복소수 타입
	var cNum1 complex64 = 1 + 2i
	var cNum2 complex64 = 4.23 + 2.35i
	fmt.Println("cNum1 : ", cNum1)
	fmt.Println("cNum2 : ", cNum2)

	// 바이트 타입  (16진수, 문자값 저장) : 바이너리 파일 Read/Write 용도 또는 데이터 암호화/복호화 용도
	// 문자열은 byte에 저장 불가능
	var bNum1 byte = 10
	var bNum2 byte = 0x10
	var bNum3 byte = 'A'
	//var bNum4 byte = "compile error"
	fmt.Println("bNum1 : ", bNum1)
	fmt.Println("bNum2 : ", bNum2)
	fmt.Println("bNum3 : ", bNum3)

	// 룬 타입 (유니코드(UTF-8) 문자코드 저장 )
	// 룬 타입 또한, 문자열은 저장 불가능
	var rStr1 rune = '한'
	var rStr3 rune = '\U0000d55c' //한

	fmt.Println("rStr1 : ", string(rStr1))
	fmt.Println("rStr3 : ", string(rStr3))

	// 오버플로우와 언더플로우
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	// var overflow uint8 = 0 - 1 // 컴파일 에러

	// 변수 크기 구하기
	var vSize1 int8 = 1
	var vSize2 int16 = 1
	var vSize3 int32 = 1
	var vSize4 int64 = 1
	fmt.Println(unsafe.Sizeof(vSize1))
	fmt.Println(unsafe.Sizeof(vSize2))
	fmt.Println(unsafe.Sizeof(vSize3))
	fmt.Println(unsafe.Sizeof(vSize4))

	// 문자열 사용하기
	var strText1 = "heelo, world!"
	fmt.Println(strText1)
	var strText2 = `
					SELECT *
					FROM ddolsoon_home d 
					WHERE d.weight >= 90;`
	fmt.Println(strText2)

	// 문자열 길이 구하기
	fmt.Println(len(strText1))

	// 문자열 연산하기
	var strOperation1 string = "한글"
	var strOperation2 string = "한글"
	var strOperation3 string = "Go"
	fmt.Println(strOperation1 == strOperation1)
	fmt.Println(strOperation1 + strOperation2)
	fmt.Println(strOperation1 + "창제")
	fmt.Printf("%c \n", strOperation3[0])

	// Go는 문자열 수정 불가능
	// var strError = "hello, world!"
	// strError[0] = 'z' // 에러

	// 부울 사용하기
	var bBool bool = true
	var bBool2 bool = false
	fmt.Println(bBool)
	fmt.Println(bBool2)

	// 상수 사용하기
	const constAge int = 10
	const constName string = "test const1"
	const (
		constAge2  int    = 1
		constName2 string = "test const2"
	)

	fmt.Printf("const example 1 : %d %s \n", constAge, constName)
	fmt.Printf("csont example 2 : %d %s \n", constAge2, constName2)

	// 열거형 사용하기 (iota 키워드 사용)
	// 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음 타입
	const (
		Sunday = iota // 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 7
	)
	fmt.Println(numberOfDays)

	// 연산자 (Go 고유 연산자만 정리)
	// reference 연산자
	refA := 1
	refB := &refA
	fmt.Println(refB)

	// dereference 연산자
	fmt.Println(*refB)

	// Go 채널 수신 연산자
	channel := make(chan int)

	go func() {
		channel <- 10000 // channel 에 10000 을 보냄
	}()

	var chanA int
	chanA = <-channel // channel 에서 값을 가져와서 chanA에 대입
	fmt.Println(chanA)
}
