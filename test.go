package main

func main() {
	k := 10
	p := &k

	println(*p)

	if k == 10 {
		println("Ten")
	}

	if val := k * 2; val < 21 {
		println(val)
	}

	var name string
	var category = 1

	switch category {
	case 1:
		name = "book"
	case 2:
		name = "other"
	case 3, 4:
		name = "calc"
	default:
		name = "unknown"
	}
	println(name)

	switch {
	case category%2 == 0:
		name = "even"
	case category%2 == 1:
		name = "odd"
	}

	println(name)

	//switch category.(type) {
	//case int:
	//	println("int")
	//case bool:
	//	println("bool")
	//case string:
	//	println("string")
	//}

	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	println(sum)

	names := []string{"a", "b", "c"}

	for index, name := range names {
		println(index, name)
	}

	for e := range names {
		say(string(e + '0'))
	}

	msg := "test"
	say2(&msg)
	println(msg)

	say3("a", "b", "c")
	say3(names...)
}

func say(msg string) {
	println(msg)
}

func say2(msg *string) {
	println(*msg)
	*msg = "new msg"
}

func say3(msg ...string) {
	for k, v := range msg {
		println(k, v)
	}
}
