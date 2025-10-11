import ("fmt"
        "strconv"
       )

func main() {
   var n int
    fmt.Scan(&n)
    
    
    qq := strconv.Itoa(n)
    a, err := strconv.Atoi(qq)
    
    if err != nil {
		fmt.Println("Ошибка преобразования:", err)
    } 
    fmt.Println( a % 10)
}