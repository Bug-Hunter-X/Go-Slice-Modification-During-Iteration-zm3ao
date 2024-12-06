func myFunc(a []int) {
  b := make([]int, 0)
  for _, v := range a {
    if v != 0 {
      b = append(b, v)
    }
  }
  a = b
} 