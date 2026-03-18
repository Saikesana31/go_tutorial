package main
import "fmt"

func main(){
	var printtext string = "Hello, World!"
	myfunc(printtext)

}

func myfunc(printtext string){
	fmt.Println(printtext)
}