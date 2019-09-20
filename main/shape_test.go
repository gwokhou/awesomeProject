package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 2 * math.Pi * 10

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	// 创建匿名结构体，声明匿名结构体切片
	areaTestCases := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	// 每个testCase作为一个子测试，明确testCase的名称和值。
	// “当测试用例不是一系列操作，而是事实的断言时，测试才清晰明了。”
	for _, testCase := range areaTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.hasArea {
				t.Errorf("%#v got %.2f want %.2f", testCase.shape, got, testCase.hasArea)
			}
		})
	}
}
