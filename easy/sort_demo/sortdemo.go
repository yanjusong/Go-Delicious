package main

import (
	"fmt"
	"sort"
)

func sortBuildType() {
	nums := []int{3, 4, 1, 2, -1}
	s := sort.IntsAreSorted(nums)
	fmt.Println("Sorted: ", s)

	sort.Ints(nums)
	fmt.Println(nums)
	s = sort.IntsAreSorted(nums)
	fmt.Println("Sorted: ", s)

	// ----------------------------
	names := []string{"tom", "bob", "andy"}
	sort.Strings(names)
	fmt.Println(names)
}

type Student struct {
	age   int
	score int
	namec string
}

type StudentVec []Student

func (s StudentVec) Len() int {
	return len(s)
}

func (s StudentVec) Less(i, j int) bool {
	if s[i].score == s[j].score {
		return s[i].age < s[j].age
	} else {
		return s[i].score > s[j].score
	}
}

func (s StudentVec) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortCostomStruct() {
	students := []Student{{12, 90, "bob"}, {11, 90, "tom"}, {13, 95, "jerry"}}
	sort.Sort(StudentVec(students))
	fmt.Println(students)
}

func sortCostomComparator() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)
}

func sortMap() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func main() {
	sortBuildType()

	fmt.Println("\n---------------------------------")
	sortCostomStruct()

	fmt.Println("\n---------------------------------")
	sortCostomComparator()

	fmt.Println("\n---------------------------------")
	sortMap()
}
