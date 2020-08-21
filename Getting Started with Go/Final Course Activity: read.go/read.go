package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
      var fileName string
      _, _ = fmt.Scan(&fileName)

      nameArr := make([]Name, 0)

      confFile, err := ioutil.ReadFile(fileName)
      if err != nil {
      	log.Fatal(err)
	  }
	  nameLines := strings.Split(string(confFile), "\n")
	  for i := 0; i < len(nameLines); i++ {
	  	if nameLines[i] != "" {
	  		nameLine := strings.Split(string(nameLines[i]), " ")
	  		newNameLine := Name{fname: nameLine[0], lname: nameLine[1]}
	  		nameArr = append(nameArr, newNameLine)
		}
	  }
	  for _, names := range nameArr {
	  	fmt.Println(names.fname + " " + names.lname)
	  }
}
