package main

import (
    "fmt"
)

func merge(a, b []int) []int {
  size, i, j := len(a)+len(b), 0, 0
  result := make([]int, size)
  for k := 0; k < size; k++ {
    switch true {
      case i == len(a):
          result[k] = b[j]
          j += 1
      case j == len(b):
         result[k] = a[i]
         i += 1
      case a[i] > b[j]:
         result[k] = b[j]
         j += 1
      case a[i] < b[j]:
          result[k] = a[i]
          i += 1
      case a[i] == b[j]:
          result[k] = b[j]
          j += 1
                  }
      }
   return result
  }
  
func MergeSort(numbers []int) []int {
    if len(numbers) < 2 {
       return numbers
     }
    middle := int(len(numbers) / 2)
  return merge(MergeSort(numbers[middle:]), MergeSort(numbers[:middle]))
}

func main() {
 a := []int{2, 1, 3, 4, 50, 78, 32, 33, 45, 45, 45, 91, 99, 37, 101, 101, 104, 106, 110, 116}
 fmt.Println(MergeSort(a))
}