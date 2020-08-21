package main
import (
    "fmt"
)
type Animal interface{
    Eat()
    Speak()
    Move()
}
type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow Cow) Eat(){
    fmt.Printf("food : grass\n")
}

func (cow Cow) Speak(){
    fmt.Printf("noise : moo\n")
}

func (cow Cow) Move(){
    fmt.Printf("locomotion : walk\n")
}


func (bird Bird) Eat(){
    fmt.Printf("food : worms\n")
}

func (bird Bird) Speak(){
    fmt.Printf("noise : peep\n")
}

func (bird Bird) Move(){
    fmt.Printf("locomotion : fly\n")
}

func (snake Snake) Eat(){
    fmt.Printf("food : mice\n")
}

func (snake Snake) Speak(){
    fmt.Printf("noise : hsss\n")
}

func (snake Snake) Move(){
    fmt.Printf("locomotion : slither\n")
}

func main(){
    var (
        command, names, types string
    )
    cow:=Cow{}
    bird:=Bird{}
    snake:=Snake{}
    listNames :=make(map[string]Animal)

    for {
        fmt.Println("New command: newanimal or query")
        _, _ = fmt.Scan(&command)
        if(command == "newanimal"){
            fmt.Println("Enter names animal and types: ")
            _, _ = fmt.Scan(&names)
            _, _  = fmt.Scan(&types)
            if types == "cow" {
                listNames[names] = cow
            } else if types == "bird" {
                listNames[names] = bird
            } else if types == "snake" {
                listNames[names] = snake
            }else{
                fmt.Println("ERROR! enter correct types!")
                continue
            }
            fmt.Println("Created It!")
        }else if(command == "query"){
            fmt.Println("Enter names animal and types")
            _, _ = fmt.Scan(&names)
            _, _ = fmt.Scan(&types)
            any, ok:= listNames[names]
            if !ok {
                fmt.Println("Incorrect")
            }else{
                fmt.Print("Name :", names, " ")
                if types == "eat" {
                    any.Eat()
                } else if types == "move" {
                    any.Move()
                } else if types == "speak" {
                    any.Speak()
                } else {
                    fmt.Println("Incorrect input. Try again")
                }
            }
        }else{
            break
        }
    }
}