package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Rectangle 的额外方法：判断是否为正方形
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// Rectangle 的额外方法：获取对角线长度
func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Circle 的额外方法：获取直径
func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Circle 的额外方法：判断点是否在圆内
func (c Circle) ContainsPoint(x, y float64) bool {
	distance := math.Sqrt(x*x + y*y)
	return distance <= c.Radius
}

// 通用函数：打印形状信息
func PrintShapeInfo(shape Shape, name string) {
	fmt.Printf("\n=== %s 信息 ===\n", name)
	fmt.Printf("面积: %.2f\n", shape.Area())
	fmt.Printf("周长: %.2f\n", shape.Perimeter())

	// 类型断言，调用特定类型的方法
	switch s := shape.(type) {
	case Rectangle:
		fmt.Printf("是否为正方形: %v\n", s.IsSquare())
		fmt.Printf("对角线长度: %.2f\n", s.Diagonal())
		fmt.Printf("宽度: %.2f, 高度: %.2f\n", s.Width, s.Height)
	case Circle:
		fmt.Printf("直径: %.2f\n", s.Diameter())
		fmt.Printf("半径: %.2f\n", s.Radius)
		fmt.Printf("点(1,1)是否在圆内: %v\n", s.ContainsPoint(1, 1))
	}
}

// 通用函数：比较两个形状的面积
func CompareShapes(shape1, shape2 Shape, name1, name2 string) {
	area1 := shape1.Area()
	area2 := shape2.Area()

	fmt.Printf("\n=== 形状比较 ===\n")
	fmt.Printf("%s 面积: %.2f\n", name1, area1)
	fmt.Printf("%s 面积: %.2f\n", name2, area2)

	if area1 > area2 {
		fmt.Printf("%s 的面积更大\n", name1)
	} else if area1 < area2 {
		fmt.Printf("%s 的面积更大\n", name2)
	} else {
		fmt.Printf("两个形状的面积相等\n")
	}
}

// 通用函数：计算形状数组的总面积
func CalculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for i, shape := range shapes {
		area := shape.Area()
		total += area
		fmt.Printf("形状 %d 面积: %.2f\n", i+1, area)
	}
	return total
}

func main() {
	fmt.Println("=== Go语言接口与面向对象编程示例 ===")

	// 创建 Rectangle 实例
	rectangle := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}

	// 创建 Circle 实例
	circle := Circle{
		Radius: 4.0,
	}

	// 调用 Shape 接口方法
	fmt.Println("1. 调用接口方法:")
	fmt.Printf("矩形面积: %.2f\n", rectangle.Area())
	fmt.Printf("矩形周长: %.2f\n", rectangle.Perimeter())
	fmt.Printf("圆形面积: %.2f\n", circle.Area())
	fmt.Printf("圆形周长: %.2f\n", circle.Perimeter())

	// 使用通用函数打印形状信息
	PrintShapeInfo(rectangle, "矩形")
	PrintShapeInfo(circle, "圆形")

	// 比较两个形状
	CompareShapes(rectangle, circle, "矩形", "圆形")

	// 创建形状数组并计算总面积
	fmt.Println("\n=== 形状数组操作 ===")
	shapes := []Shape{rectangle, circle}
	fmt.Printf("形状数组长度: %d\n", len(shapes))

	totalArea := CalculateTotalArea(shapes)
	fmt.Printf("总面积: %.2f\n", totalArea)

	// 演示接口的多态性
	fmt.Println("\n=== 接口多态性演示 ===")
	var shape Shape

	// 将 Rectangle 赋值给 Shape 接口
	shape = rectangle
	fmt.Printf("Shape接口调用Area(): %.2f\n", shape.Area())

	// 将 Circle 赋值给 Shape 接口
	shape = circle
	fmt.Printf("Shape接口调用Area(): %.2f\n", shape.Area())

	// 类型断言示例
	fmt.Println("\n=== 类型断言示例 ===")
	if rect, ok := shape.(Rectangle); ok {
		fmt.Printf("当前形状是矩形，宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	} else if circ, ok := shape.(Circle); ok {
		fmt.Printf("当前形状是圆形，半径: %.2f\n", circ.Radius)
	}

	// 使用类型开关
	fmt.Println("\n=== 类型开关示例 ===")
	switch s := shape.(type) {
	case Rectangle:
		fmt.Printf("类型: 矩形, 宽度: %.2f, 高度: %.2f\n", s.Width, s.Height)
	case Circle:
		fmt.Printf("类型: 圆形, 半径: %.2f\n", s.Radius)
	default:
		fmt.Println("未知类型")
	}

	fmt.Println("\n=== 程序执行完成 ===")
}
