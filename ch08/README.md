# Golang

## Chapter 08.

## 상수

### ✅ 상수 소개

- 상수는 변하지 않는 값을 표현 할 때 사용

- 자주 변경되지 않는 숫자값이나 문자열값에 주로 사용된다.

### 1. 상수 선언

- 상수는 변하지 않는 값을 의미.

- 변수는 대입문을 통해서 값을 수시로 바꿀 수 있지만, 상수는 초기화된 값이 변하지 않는다.

- 정수, 실수, 문자열 등 기본 타입값들만 상수로 선언 될 수 있다.

- 구조체, 배열 등 기본타입(Priavate)이 아닌 타입(complex)에는  
  상수를 사용 할 수 없다.

- 상수로 사용 가능한 타입 :

  - 불리언
  - 룬(rune)
  - 정수
  - 실수
  - 복소수
  - 문자열

- 상수 선언 방법

```Go
const ConstValue int = 10
```

### 2. 상수는 언제 사용하는지

- 변하면 안되는 값에 사용

- 코드값을 통해서 숫자에 의미를 부여할 때 사용

  - 코드값으로 사용하기 때문에 값을 그냥 1,2,3... 처럼 1씩증가하도록 정의할 때  
    iota를 사용 하면 편리하다.

    ```Go
    const (
        Red int = iota
        Blue int = iota
        Green int = iota
    )

    /* 상수 목록을 const와 ()로 묶고 iota를 계속 사용하면
     * 0부터 1씩 차례로 증가하면서 값이 초기화 된다.
     * 똑같은 규칙으로 계속 적용된다면 타입과 iota를 생략할 수 있다.
    */
    const(
        C1 uint = iota + 1
        C2
        C3
    )
    // 또는
    const(
        BitFlag1 uint = 1 << iota // 1 = 1 << 0
        BigFlag2                  // 2 = 1 << 1
        BigFlag3                  // 3 = 2 << 2
    )
    ```

### 3. 타입 없는 상수

- 상수 선언 시 타입을 명시하지 않을 수 있다.

- 타입 없는 상숫값은 타입이 정해지지 않은 상태로 사용된다.

```Go
package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	var b int = int(FloatPI) * 100

	fmt.Println(a)
	fmt.Println(b)
}
```

- 타입 없는 상수는 변수에 복사될 때 타입이 정해지기 때문에 여러 타입에 사용되는 상숫값을 사용할 때 편리하다.

### 4. 상수와 리터럴

- 컴퓨터에서 리터럴이란 고정된 값, 값 자체로 쓰인 문구라고 볼 수 있다.

```Go
var str string = "Hello World"
var i int = 0
i = 30
// 위와 같이 고정된 값 자체로 쓰인 문구가 바로 리터럴이다.
```

- Go에서는 상수는 리터럴과 같이 취급한다.  
  컴파일될때 상수는 리터럴로 변환되어 실행 파일에 쓰인다.

- 상수 표현식 역시 컴파일 타임에 실제 결괏값 리터럴로 변환하기 때문에  
  상수 표현식 계산에 CPU 자원을 사용하지 않는다.

```GO
const PI = 3.14
var a int = PI * 100
```

- 상수의 메모리 주소값에 접근 할 수 없는 이유 역시 컴파일 타임에 리터럴로 전환되어서 실행 파일에 값 형태로 쓰이기 때문이다.  
  그래서 동적 할당 메모리 영역을 사용하지 않는다.

### ⭐️ 핵심 요약

1. 상수는 변하지 않는 값이다.

2. 자주 쓰는 고정값에 이름을 부여해서 편리하게 사용할 수 있다.

3. 상수를 코드값으로 사용할 수 있다.

4. iota를 이용하면 증가하는 상수를 편리하게 선언할 수 있다.

5. 타입 없는 상수를 선언하면 타입이 다른 여러 변수에서 사용 할 수 있다.
