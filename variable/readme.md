## 변수 

-  값을 저장하는 메모리 공간을 가리키는 이름 
- 프로그램이란 결국 데이터를 연산, 조작하는 일 
    - CPU가 메모리에서 가져온 값을 변수를 통해서 code에서 data 값을 조작한다. 
    - 변수를 통해 메모리에 접근하는 것이다.
<br>

### 변수의 4가지 속성 

- 이름, 값, 주소, 타입 
<br>

#### 숫자 타입 

- uint : 부호가 없는 정수 (모두 +) 
    - uint8 : 1byte(8bit)짜리 
    - uint16 : 2byte 
    - uint32 : 4byte 
    - uint64 : 8byte 
- int : 부호 있는 정수 (+, -) 
    - int8, int16, int32, int64 
- float: 실수
    - float32, float64 
- byte : uint8과 같음, 별칭임 
- rune : 문자 1개, rune이 모이면 string이 됨, int32(4byte)의 alias 타입 
- int : 사이즈가 플랫폼에 따라 달라짐 
    - 32bit 컴퓨터에서는 int32, 64bit 컴퓨터에서는 int64 
uint : 사이즈가 플랫폼에 따라 달라짐 
<br>

#### 그외 타입 

- bool : true, false를 값으로 가지는 변수 
- string : 문자열 
- 배열 : 같은 타입의 요소를 여러개 같는 것 
- 슬라이스 : 사이즈가 정해져 있지 않은 배열 
- 구조체 
- 포인터 
- 함수타입
- 맵
- 인터페이스
- 채널 
<br>

#### 별칭 타입 

- 어떤 타입에 다른 이름을 부여하는 것 

```
type myInt int 
```
- 위와 같이 사용한다. 
- myInt라고 사용해도 되고, int라고 사용해도 된다.
<br> 

#### 변수 선언 

```
var a int = 10 
```
- 가장 기본적인 형태 

```
var a int
```
- 초기값을 안적는 경우 
- 초기값은 default 값 
- 정수의 default 값은 0 

```
var a = 10 
```
- 타입을 생략하는 경우 
- data의 타입과 같아진다. 
- 타입을 생략하면 반드시 값을 적어야 한다. 

```
a := 10 
```
- var와 타입을 생략한 경우 
- := 을 사용해야 한다. 
- 가장 많이 쓰는 방식 
<br>

##### 타입별 기본값 

- 모든 정수 타입 : 0 
- 모든 실수 타입 : 0.0 
- 불리언 : false 
- 문자열 : 빈 문자열 
- 그외 : nil (정의되지 않은 메모리 주소를 나타내는 go 키워드)
<br>

##### 타입 변환 

- 연산의 각 항목의 타입은 **반드시 같아야** 한다. 

ex. 
```
var a int = 10
var b int = 20 
var c float = 1.1
var d int16 = 20 
```
- a+b는 가능 
- a+c는 불가능 
- a+d는 불가능 
<br>

#### 코드 분석 

```
func main() {
	var a int = 10
	var msg string = "hello variable"

	a = 20
	msg = "good morning"

	fmt.Println(msg, a)
}
```

- main 함수의 코드를 뜯어보자. 

```
var a int = 10 
```

- var : 변수 선언 
- a : 변수명 
- int : 타입, 여기서는 정수 타입  
- = : 대입 연산자 (우변의 값을 좌변에 넣음) 

```
var msg string = "hello variable"
```

- var : 변수 선언 
- msg : 변수명 
- string : 타입, 여기서는 문자열 타입  
- = : 대입 연산자 (우변의 값을 좌변에 넣음)

```
a = 20 
```

- 위와 같은 방식으로 변수 값 변경 
- a가 가리키는 메모리 공간에 20 대입 

```
func main() {
	a := 3
	var b float64 = 3.5
	// var c int = b
	// d := a * b

	var c int = int(b)
	d := float64(a) * b

	fmt.Println(c, d)
}
```

- 타입이 다르면 대입 자체도 안된다.(주석 부분) 
- 다른 타입을 대입하기 위해서는 타입 변환 함수를 사용해야 한다. (ex. float64(), int() 같은 함수) 
