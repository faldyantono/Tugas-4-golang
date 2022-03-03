package main
import "fmt"

func main(){
  var mahasiswa = map[string]int{
    "aldo": 182,
    "Yosep": 178,
  }
  for key, val := range mahasiswa {
      fmt.Println(key,"\t:", val)
}
}
