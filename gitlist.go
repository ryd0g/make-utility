package main
  
import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    
    fname := "Names.txt"

    textin, err := ioutil.ReadFile(fname)
    if err == nil {
        fmt.Println(string(textin))
    }

    // append text to file
    f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        panic(err)
    }

    fmt.Println("Enter Your Name: ")
    var firstName string
    fmt.Scanln(&firstName)

    fmt.Println("Enter Your GitHub Username: ")
    var gitUsername string
    fmt.Scanln(&gitUsername)

    _, err = f.Write([]byte(firstName + " "))
    _, err = f.Write([]byte(gitUsername))
    if err != nil {
        panic(err)
    }
    f.Close()

    textin, err = ioutil.ReadFile(fname)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(textin))
}