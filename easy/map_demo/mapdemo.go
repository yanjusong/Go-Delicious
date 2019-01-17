package main

import "fmt"

func test1() {
	var map1 map[string]string

	if map1 == nil {
		fmt.Printf("map1 is nil.\n")
	}

	map1 = make(map[string]string)

	if map1 == nil {
		fmt.Printf("map1 is nil.\n")
	}

	map1["tom"] = "Beijing"
	map1["jerry"] = "Shanghai"

	fmt.Println(map1)

	for name := range map1 {
		fmt.Printf("name:%s addr:%s\n", name, map1[name])
	}

	addr1, ok := map1["tom"]
	if ok {
		fmt.Printf("find tom in map1, addr is %s.\n", addr1)
	} else {
		fmt.Printf("can't find tom in map1.\n")
	}

	addr1, ok = map1["Tom"]
	if ok {
		fmt.Printf("find Tom in map1, addr is %s.\n", addr1)
	} else {
		fmt.Printf("can't find Tom in map1.\n")
	}

	map1["tom"] = "Hainan"
	fmt.Println(map1)
}

func test2() {
	map2 := map[string]int{"tom": 18, "jerry": 20, "bob": 21}
	fmt.Println(map2)

	delete(map2, "bob")
	fmt.Println(map2)
}

func test3() {
	m := make(map[string]string)
	m["name"] = "tom"
	m["nickname"] = ""

	nickname, ok := m["nickname"]

	if ok {
		fmt.Printf("nickname is %s. size:%d\n", nickname, len(nickname))
	} else {
		fmt.Printf("can't find nickname in m.\n")
	}
}

func main() {
	// test1()
	// test2()
	test3()
}
