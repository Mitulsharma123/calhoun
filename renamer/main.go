package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func findmatch(filename string, total int) (string,error) {
  pieces := strings.Split(filename, ".")
  //fmt.Println(pieces)
  ext := pieces[len(pieces)-1]
  //fmt.Println(ext)
  tmp := strings.Join(pieces[0:len(pieces)-1], ".")
  //fmt.Println(tmp)
  pieces = strings.Split(tmp, "_")
  //fmt.Println(pieces)
  name := strings.Join(pieces[0:len(pieces)-1], "_")
  //fmt.Println(name)
  number , err := strconv.Atoi(pieces[len(pieces)-1])
  //fmt.Println(number)
  if err != nil{
    return "", fmt.Errorf("%d didn't match any pattern")
  }
  return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}

func main(){
  //filename := "bday_001.txt"

  /*newName, err := findmatch(filename, 4)
  if err != nil {
    fmt.Println("no match found")
    os.Exit(1)
  }else {
    fmt.Println(newName)
  }*/


  files, err := os.ReadDir("./sample")
  if err != nil {
    panic(err)
  } 
  for _, file := range files{
    if file.IsDir(){
      fmt.Println("Dir: ", file.Name())
  } else {
      tmp, err := findmatch(file.Name(), 4)
      fmt.Println("match:", tmp, err)
    }
  }
}
