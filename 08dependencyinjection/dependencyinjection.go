package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s" , name)
}

//func main() {
//	Greet(os.Stdin, "Elliot")
//}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "World")
}

func main()  {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler)))
}