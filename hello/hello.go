package main
import (
	"fmt"
	"log"
	"haishen.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string {"haishen", "yudonglai", "guanyu", "zhangfei"}
	mes, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(mes)
}

