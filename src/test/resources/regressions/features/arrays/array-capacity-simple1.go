package pkg

func test1() {
  var a [42]int
  assert cap(a) == 42
  assert cap(a) == len(a)
}

func test2(a [12]int) {
  assert cap(a) == 12
  assert cap(a) == len(a)
}

requires cap(a) == 14
func test3(a [12]int) {
  assert false
}

func test4(a [12]int) {
  assert cap(a) != 21
}
