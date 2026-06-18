package main
import (
	"fmt"
	"log"
	"haishen.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	mes, err := greetings.Hello("haishen")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(mes)
}

