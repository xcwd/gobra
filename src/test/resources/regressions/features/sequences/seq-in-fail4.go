package pkg

// fails since the membership test might fail
func foo(ghost xs seq[int]) {
  //:: ExpectedOutput(assert_error:assertion_error)
  assert 42 in xs
}
