package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func findmatch(filename string, total int) (string,error) {
  pieces := strings.Split(filename, ".")
  ext := pieces[len(pieces)-1]
  tmp := strings.Join(pieces[0:len(pieces)-1], ".")
  pieces = strings.Split(tmp, "_")
  name := strings.Join(pieces[0:len(pieces)-1], "_")
  number , err := strconv.Atoi(pieces[len(pieces)-1])
  if err != nil{
    return "", fmt.Errorf("%d didn't match any pattern")
  }
  return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), err
}

func main(){
  files, err := os.ReadDir("./sample")
  if err != nil {
    panic(err)
  } 
  for _, file := range files{
    if file.IsDir(){
      fmt.Println("Dir: ", file.Name())
  } else {
      tmp, err := findmatch(file.Name(), 4)
      fmt.Println(tmp,err)
    }
  }
}
