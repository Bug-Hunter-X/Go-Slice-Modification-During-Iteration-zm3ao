# Go Slice Modification During Iteration

This repository demonstrates a common error in Go: modifying a slice while iterating over it using a `for...range` loop. This can lead to unexpected behavior or even panics.

The `bug.go` file contains code that attempts to remove elements from a slice during iteration.  The `bugSolution.go` file shows the correct way to handle this situation, preventing unexpected behavior.