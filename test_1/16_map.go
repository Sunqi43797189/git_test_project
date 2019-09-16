package main

import "fmt"

func makeMap() {
	agesV1 := make(map[string]int)
	agesV1["tom"] = 22
	agesV1["jerry"] = 24

	agesV2 := map[string]int{
		"tom":   22,
		"jerry": 24,
	}

	agesV3 := map[string]int{}
	fmt.Println(agesV1, agesV2, agesV3)

	delete(agesV1, "alice")
	delete(agesV1, "tom")
	agesV1["liming"] = 25
	age := agesV1["xiaoming"]
	agesV1["xiaoming"] = agesV1["xiaoming"] + 1
	fmt.Println(age) //可以 不存在 返回value类型的零值
	fmt.Println(agesV1)

	age, ok := agesV1["mary"]
	if !ok {
		fmt.Println("mary is not in map")
	}
}

func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, x_v := range x {
		y_v, ok := y[k]
		if !ok || x_v != y_v {
			return false
		}
	}
	return true
}
