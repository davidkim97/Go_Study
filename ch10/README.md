# Golang

## Chapter 10.

## switch문

### ✅ switch문 소개

- switch문은 비굣값과 같은 case에 해당하는 문장을 수행한다.

- 변수값에 따라서 다른 명령을 수행해야 하는 경우 유용하다.

### 1. switch문 동작 원리

```Go
switch 비굣값 { // 검사하는 값이 온다.
    case 값 1: // 비교값과 값 1이 같을 때 수행
        문장
    case 값 2: // 비교값과 값 2가 같을 때 수행
        문장
    default: // 만족하는 case가 없을 때 수행
        문장
}
```

### 2. switch문을 언제 쓰는가

- switch문을 이용하면 복잡한 if else 문을 보기 좋게 정리 할 수 있다.

```Go
// if else 문
package main

import "fmt"

func main() {
	day := 3

	if day == 1 {
		fmt.Println("첫째 날 입니다.")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	} else if day == 2 {
		fmt.Println("둘째 날 입니다.")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	} else if day == 3 {
		fmt.Println("셋째 날 입니다.")
		fmt.Println("설계안을 확정하는 날 입니다.")
	} else if day == 4 {
		fmt.Println("넷째 날 입니다.")
		fmt.Println("예산을 확정하는 날 입니다.")
	} else if day == 5 {
		fmt.Println("다섯째 날 입니다.")
		fmt.Println("최종 계약하는 날 입니다.")
	} else {
		fmt.Println("프로젝트를 진행하세요.")
	}
}

// switch 문
package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날 입니다.")
		fmt.Println("오늘은 팀 미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날 입니다.")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날 입니다.")
		fmt.Println("설계안을 확정하는 날 입니다.")
	case 4:
		fmt.Println("넷째 날 입니다.")
		fmt.Println("예산을 확정하는 날 입니다.")
	case 5:
		fmt.Println("다섯째 날 입니다.")
		fmt.Println("최종 계약하는 날 입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요.")
	}
}

```

### 3. 다양한 switch문 형태

- 하나의 case는 하나 이상의 값을 비교할 수 있다.  
  각 값은 쉼표로 구분

- switch문의 동작을 응용하면 단순히 값만 비교하는 게 아닌 if문처럼 true가 되는 조건문을 검사할 수 있다.

- switch 다음에 비굣값을 적지 않을 경우 default값으로 true를 사용한다.

- if문과 마찬가지로 switch문에서도 초기문을 넣을 수 있다.

```Go
switch 초기문; 비교값{
case 값1:
    ...
case 값2:
    ...
default:
    ...
}
```

### 4. const 열거값과 switch

- const 열거값에 따라 수행되는 로직을 변경할 때 switch문을 주로 사용한다.

### 5. break와 fallthrough 키워드

- Go언어에서는 break를 사용하지 않아도 case 하나를 실행 후 자동으로 switch문을 빠져나간다.

- 만약 하나의 case문 실행 후 다음 case문까지 같이 실행하고 싶을 때는 fallthrough 키워드를 사용한다.

- case 마지막에 fallthrough 키워드를 사용하면 다음 case까지 같이 실행된다.

- fallthrough 키워드는 코드를 보는 사람에게 혼동을 일으킬 수 있으니 되도록 사용하지 않기를 권장

### ⭐️ 핵심요약

- switch문은 값에 따라서 다른 문장을 실행 할 때 사용되는 구문이다.

```Go
switch 비교값 {
    case 값1:
        ...
    case 값2:
        ...
    default:
        ...
}
```

```Go
switch { // 비교값을 생략하면 비교값은 true가 된다.
    case 조건문1:
        문장
    case 조건문2:
        문장
    default:
        문장
}
```

```Go
// 초기문이 먼저 실행되고 비교값을 case들과 비교한다.
switch 초기문; 비교값 {
    case 값1:
        문장
    case 값2:
        문장
    default:
        문장
}
```
