package pkg

func foo(ghost xs seq[int]) {
  // fails: left-hand side has a type that doesn't match the right-hand side
  //:: ExpectedOutput(type_error)
  assert true in xs
}
