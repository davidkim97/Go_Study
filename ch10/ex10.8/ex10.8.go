package main

import "fmt"

type ColorType int // 별칭 ColorType을 선언하고 const 열거값 정의

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// 각 ColorType 열거값에 따른 문자열을 반환하는 함수
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoritColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorit color is", colorToString(getMyFavoritColor()))
}
