# Golang

## Chapter 07.

## 함수

### 1. 함수 정의

```Go
func Add(a int, b int) int {
    return a + b
}
```

- 함수를 호출할 때 입력하는 값을 argument 또는 인수 라고 한다.

- 함수가 외부로부터 입력받는 값을 parameter 또는 매개변수 라고 한다.

- 인수는 매배견수로 복사된다.

- 매개변수와 함수 내에서 선언된 변수는 함수가 종료되면 변수 범위를 벗어나서 접근하지 못한다.

- 함수를 사용하는 이유

  - 함수를 이용해서 중복 코드를 제거하여 코드를 간결하게 한다.

- 멀티 반환 함수

  - 함수는 여러개 반환 할 수 있다.  
    반환 값이 여럿일 때는 반환 타입들을 소괄호로 묶어서 표현한다.

- 변수명을 지정해 반환
  - 함수 선언부에 반환 타입을 적을 때 변수명까지 지정해주면,  
    return문으로 변수를 명시적으로 반환하지 않아도 값을 반환 할 수 있다.

### 2. 재귀 호출

- 재귀 호출(recursive call)이란 함수 안에서 자기 자신 함수를 다시 호출하는 것을 말한다.

### 핵심 요약

1. 함수란 특수한 코드 묶음을 말한다. 함수를 만들면 재 사용 할 수 있다.

2. 함수 정의

```Go
func Add(a int, b int) int {
    return a + b
}
```

3. 멀티 반환 함수는 값을 여러 개 반환 할 수 있다.  
   반환 타입 자리에 소괄호로 여러 반환 타입을 묶어서 표시한다.

4. 재귀 호출을 함수 안에서 같은 함수를 또 호출하는 기법을 말한다.  
   재귀 호출 시에는 항상 탈출 조건을 명확히 해야 한다.
