package input

import (
	"cuelang.org/go/pkg/strconv"
	"cuelang.org/go/pkg/strings"
	"fmt"
	"os"
)

func GetString(hint string) string {
	fmt.Print(hint + ": ")
	var result string
	fmt.Scanln(&result)
	result = strings.TrimSpace(result)
	return result
}

func GetInt(hint string) int {
	fmt.Print(hint + ": ")
	var result string
	fmt.Scanln(&result)
	result = strings.TrimSpace(result)
	if len(result) == 0 {
		return 1
	}
	ret, err := strconv.Atoi(result)
	if err != nil {
		fmt.Print(err)
	}
	return ret
}

func GetFloat(hint string) float64 {
	fmt.Print(hint + ": ")
	var result string
	fmt.Scanln(&result)
	result = strings.TrimSpace(result)
	if len(result) == 0 {
		fmt.Println("输入值非法")
		os.Exit(1)
	}
	ret, err := strconv.ParseFloat(result, 64)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return ret
}

func GetBool(hint string) bool {
	fmt.Print(hint + ": ")
	var result string
	fmt.Scanln(&result)
	result = strings.TrimSpace(result)
	if len(result) == 0 {
		fmt.Println("输入值非法")
		os.Exit(1)
	}
	ret, err := strconv.ParseBool(result)
	if err != nil {
		fmt.Print(err)
	}
	return ret
}
