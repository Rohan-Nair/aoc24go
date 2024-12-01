package main

import (
  "fmt"
  "os"
  "strings"
  "sort"
  "strconv"
  "log"
)

func main(){
  data, err:= os.ReadFile("ip1.txt")
  if err != nil {
    log.Fatal(err)
  }
  lines := strings.Split(string(data), "\n")
  var list1 []int
  var list2 []int

  for _, line := range lines {
    // trimming the leading or trailing whitespace
    line = strings.TrimSpace(line)
    if line == ""{
      continue
    }

    // splitting the line into parts
    parts := strings.Fields(line)

    // converting to integers
    first, err1 := strconv.Atoi(parts[0])
    second, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil{
      fmt.Println("Error in converting to int")
      return
    }

    list1 = append(list1, first)
    list2 = append(list2, second)

  }

  // sorting list1 and list2 
  sort.Ints(list1)
  sort.Ints(list2)

  var sum int 
  for i := 0; i < len(list1); i++ {
    sum += max(list1[i], list2[i]) - min(list1[i], list2[i])
  }

  fmt.Println(sum)

}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func min (a, b int) int {
  if a < b {
    return a
  }
  return b
}
