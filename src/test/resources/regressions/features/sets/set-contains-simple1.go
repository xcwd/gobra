package pkg

func example1(ghost s set[int]) (ghost b bool) {
  b = 42 in s
}

func example2() {
  assert 2 in set[int] { 1, 2, 3 }
  assert !(2 in set[int] { 1, 3 })
  assert 2 in set[int] { 2 } in set[bool] { true }
}

func example3(ghost x int, ghost s set[int], ghost t set[int]) {
  assert x in s ==> x in s union t
  assert x in s ==> x in t ==> x in s intersection t
  assert x in s ==> s subset t ==> x in t
  assert x in s intersection t ==> x in s && x in t
  assert x in s union t ==> x in s || x in t 
}
