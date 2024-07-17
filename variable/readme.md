## 변수 

-  값을 저장하는 메모리 공간을 가리키는 이름 
- 프로그램이란 결국 데이터를 연산, 조작하는 일 
    - CPU가 메모리에서 가져온 값을 변수를 통해서 code에서 data 값을 조작한다. 
    - 변수를 통해 메모리에 접근하는 것이다. 
<br>

#### 변수 선언 

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
<br>

#### 변수 값 변경 

```
a = 20 
```

- 위와 같은 방식으로 변수 값 변경 
- a가 가리키는 메모리 공간에 20 대입 
