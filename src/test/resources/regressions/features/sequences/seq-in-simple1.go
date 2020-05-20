package pkg

func example1(ghost xs seq[int]) {
  ghost b := 2 in xs
}

func example2() {
  assert 42 in seq[1..100]
  assert 2 in seq[int] { 1, 2, 3 }
  assert !(false in seq[bool] { true })
}

requires x in xs && !(y in xs)
func example3(x int, y int, ghost xs seq[int]) {
  assert x != y
}

requires x in xs
requires y in ys
ensures x in xs ++ ys
ensures y in xs ++ ys
func example4(x int, y int, ghost xs seq[int], ghost ys seq[int]) {
}

func example5() {
  ghost xs := seq[1..10]
  ghost ys := seq[bool] { false, false }
  assert 42 in xs in ys
}

func example6(x int, ghost xs seq[int], ghost ys seq[bool]) {
  ghost b1 := x in xs in ys
  ghost b2 := (x in xs) in ys
  assert b1 == b2
}

requires 0 <= n
ensures !(n in seq[0..n])
func example7(n int) {
}

requires 0 < n
ensures n - 1 in seq[0..n]
func example8(n int) {
}
