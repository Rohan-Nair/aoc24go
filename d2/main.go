package main
import (
 "fmt"
 "os"
 "strings"
 "strconv"
 "log"
)

func main() {
 p1()
}

func p1() {
 data, err := os.ReadFile("ip.txt")
 if err != nil {
  log.Fatal(err)
 }
 lines := strings.Split(string(data), "\n")
 var ans int
 for _, line := range lines{
  line = strings.TrimSpace(line)
  if line == ""{
   continue
  }

  parts:= strings.Fields(line)
  newparts := make([]int, len(parts))
  for i, str := range parts {
   num, err := strconv.Atoi(str)
   if err != nil {
    fmt.Printf("Error converting %s to an integer: %v\n", str, err)
    continue
   }
   newparts[i] = num
  }
  // check if the array is either monotonic increasing or monotonic decreasing 
  if monoInc(newparts) == true || monoDec(newparts) == true {
   ans += 1
  }
 }   
 fmt.Println(ans)
}

func monoInc (arr []int) bool{
 for i := 0; i < len(arr)-1; i++{
  if arr[i] > arr[i+1]{
   return false
  }
  val := abs(arr[i] - arr[i+1])
  if val > 3 || val == 0{
   return false
  }
 }
 return true
}

func monoDec (arr []int) bool{
 for i := 0; i < len(arr)-1; i++{
  if arr[i] < arr[i+1]{
   return false
  }
  val := abs(arr[i] - arr[i+1])
  if val > 3 || val == 0{
   return false
  }
 }
 return true
}

func abs (val int) int{
 if val < 0 {
  return val * -1; 
 }
 return val 
}
