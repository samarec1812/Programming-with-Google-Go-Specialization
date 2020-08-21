package main

import "fmt"


type Animal struct {
   food string
   locomotion string
   noise string
}

func (beast *Animal) Eat()  {
   fmt.Println(beast.food)
}
func (beast *Animal) Move()  {
     fmt.Println(beast.locomotion)
}
func (beast *Animal) Speak()  {
   fmt.Println(beast.noise)
}


func main() {
   var (
      being, action string
   )
   for {
      fmt.Print(">")
      _, _ = fmt.Scan(&being, &action)
      if being == "cow" {
         cow := Animal{"grass", "walk", "moo"}
         if action == "eat" {
            cow.Eat()
         } else if action == "move" {
            cow.Move()
         } else if action == "speak" {
            cow.Speak()
         } else {
            fmt.Println("ERROR!")
         }
      } else if being == "bird" {
         bird := Animal{"worms", "fly", "peep"}
         if action == "eat" {
            bird.Eat()
         } else if action == "move" {
            bird.Move()
         } else if action == "speak" {
            bird.Speak()
         } else {
            fmt.Println("ERROR!")
         }
      } else if being == "snake" {
         snake := Animal{"mice", "slither", "hsss"}
         if action == "eat" {
            snake.Eat()
         } else if action == "move" {
            snake.Move()
         } else if action == "speak" {
            snake.Speak()
         } else {
            fmt.Println("ERROR!")
         }
      } else {
         fmt.Println("ERROR!")
      }
    }
}



