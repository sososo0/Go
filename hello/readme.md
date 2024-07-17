## 코드 분석해보기 

```
package main 

import "fmt"

func main() {
	//주석
	fmt.Println("hello world")
}
```

- 위의 코드를 분석해보자 
<br>

#### Package 

```
package main 
```

- go는 반드시 package로 시작해야 한다.
- 이는 code가 속한 package가 무엇인지를 나타낸다. 
- package는 코드를 묶는 단위를 말한다. 
- main은 키워드로 특별한 의미가 있는 패키지 명이다. (프로그램 시작점을 포함하는 패키지)

##### 프로그램 시작점? 

- 프로그램 = 실행 파일(하드디스크에 저장) 
- 프로그램 실행 = 프로그램 실행 파일이 Load 되어서 메모리에 올라가서 CPU가 읽어서 명령어를 한 줄씩 실행하는 것 
    - 이때, 어디서 시작하는 가? = starting point (시작지점) 
    - **main은 starting point를 가지고 있음**
    - starting point는 반드시 1개이어야 한다. 
    - 따라서, 모든 프로그램에서 starting point를 가지고 있는 package는 1개이어야 한다. 
    - Go는 main package 1개와 다른 여러개의 package들로 구성되어 있다. 
<br>

#### import 

```
import "fmt"
```

- import는 가져온다는 의미이다. 
- fmt는 패키지 명이다. 
- import "fmt"는 fmt라는 패키지를 사용하겠다는 의미이다. 

##### Package를 왜 가져오는가? 

- package는 코드를 묶는 단위이다. 
- 코드를 묶는 단위이기 때문에 어떠한 기능을 가지고 있다. 
- 따라서, 어떤 패키지의 기능을 사용하고 싶을 때 패키지를 가져온다. 
<br>

#### func main() 

```
func main() {
}
```

- func은 함수를 선언하는 것이다. 
- func 뒤에 나오는 것은 함수명이다. 
    - 여기서도 main은 키워드로 특별한 의미를 가지고 있다. 
    - **main은 프로그램 시작점**이다. 
    - 프로그램이 load되면 main 함수에서부터 시작한다. (main 함수가 끝나면 프로그램이 끝난 것)
- 주석은 //(1줄)과 /*(여러줄)이 있음 

```
func main() {
	//주석
	fmt.Println("hello world")
}
```

- fmt. : fmt 패키지 안에 있는 함수를 사용하겠다. 
- Println : 1줄 출력 (fmt 패키지 안에 있는 함수)
