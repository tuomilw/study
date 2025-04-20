package task2

import (
	"fmt"
	"math"
)

/*
*定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
 */
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/*
*使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
 */
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (employee *Employee) PrintInfo() {
	fmt.Println("Name:", employee.Name, "Age:", employee.Age, "EmployeeID:", employee.EmployeeID)
}
