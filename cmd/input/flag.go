package input

import (
	"fmt"
	"github.com/spf13/chartgen/cmd/utils"
	"github.com/spf13/pflag"
	"os"
	"strconv"
	"strings"
)

func AddToFlagSet(flagSet *pflag.FlagSet) {
	for _, item := range Checks {
		if item.Type == STRING {
			flagSet.String(item.Name, item.Value.(string), item.Usage)
		} else if item.Type == INT {
			flagSet.Int(item.Name, item.Value.(int), item.Usage)
		} else if item.Type == BOOL {
			flagSet.Bool(item.Name, item.Value.(bool), item.Usage)
		} else if item.Type == FLOAT {
			flagSet.Float64(item.Name, item.Value.(float64), item.Usage)
		}
	}
}

func Parse(flagSet *pflag.FlagSet) {
	for _, f := range Checks {
		flag := flagSet.Lookup(f.Name)
		if flag == nil {
			fmt.Println("系统存在未知异常，参数%s没有找到", f.Name)
			continue
		}

		if f.Type == STRING {
			str := GetString(flag.Usage)
			if utils.NotEmpty(str) {
				flagSet.Set(flag.Name, strings.TrimSpace(str))
			} else if utils.IsEmpty(flag.Value.String()) {
				fmt.Println("参数" + flag.Name + "没有输入任何值")
				os.Exit(1)
			}
		} else if f.Type == INT {
			flagSet.Set(flag.Name, strconv.Itoa(GetInt(flag.Usage)))
		} else if f.Type == BOOL {
			flagSet.Set(flag.Name, strconv.FormatBool(GetBool(flag.Usage)))
		} else if f.Type == FLOAT {
			flagSet.Set(flag.Name, strconv.FormatFloat(GetFloat(flag.Usage), 'f', 32, 64))
		}
	}
}
