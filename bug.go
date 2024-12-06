func myFunc(a []int) { 
  for i := range a {
    if a[i] == 0 {
      a = append(a[:i], a[i+1:]...) // this is wrong, it modifies a in a loop that iterates on a
    }
  }
}