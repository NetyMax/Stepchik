import "fmt"

func main() {

	var a int = 2
	var b int = 2
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a

	b = b * b

	c := a + b
	fmt.Println(c)
}