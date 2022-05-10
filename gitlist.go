package main
  
import "fmt"
  
func main() {
  

    fmt.Println("Enter Your Name: ")
  
    var firstName string
  
    fmt.Scanln(&firstName)

    fmt.Println("Enter Your GitHub Username: ")
    var gitUsername string
    fmt.Scanln(&gitUsername)
  

    fmt.Print("Name: " + firstName + " GitHub Username: " + gitUsername)
}