package main

import "fmt"

func main() {
	fmt.Println("case 1:")
	{
		str := "Go编程"

		// G(1) + 0(1) + 编(3) + 程(3)
		fmt.Println(len(str))

		// 用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。
		// 要想访问中文的话，还是要用rune切片，这样就能按下表访问。
		runeStr := []rune(str)
		fmt.Println(len(runeStr))

		for i, x := range runeStr {
			fmt.Printf("%d: %c\n", i, x)
		}
	}

	fmt.Println("\ncase 2:")
	{
		s := "abcd"
		bs := []byte(s)

		bs[1] = 'B'
		println(string(bs))

		u := "电脑"
		us := []rune(u)

		us[1] = '话'
		println(string(us))
	}

	fmt.Println("\ncase 3:")
	{
		s := "abc汉字"

		// s[0] = '1' // error: string是不可变的

		// byte形式遍历
		for i := 0; i < len(s); i++ { // byte
			fmt.Printf("%c,", s[i])
		}

		fmt.Println()

		// rune形式遍历
		for _, r := range s { // rune
			fmt.Printf("%c,", r)
		}

		fmt.Println()
	}
}
