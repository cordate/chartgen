package template

import (
	"fmt"
	"github.com/spf13/chartgen/cmd/input"
	"io/ioutil"
	"os"
	"strings"
)

type TemplateFile struct {
	Path    string
	Content string
}

func FormatResources() []TemplateFile {
	tmp := make([]TemplateItem, 0)
	for ok := input.GetBool("请输入是否增加资源"); ok; ok = input.GetBool("请输入是否继续添加资源") {
		name := input.GetString("请输入需要添加的模板(" + strings.Join(ListAvailableResource(tmp), ",") + ")")
		if _, exsits := Template[name]; !exsits {
			fmt.Println("输入的资源名称非法，请重新输入")
			continue
		}
		// 存在性检查
		check := false
		for _, item := range ListAvailableResource(tmp) {
			if item == name {
				check = true
				break
			}
		}
		if !check {
			fmt.Println("输入的资源已经存在，请勿重复输入")
			continue
		}

		tmp = append(tmp, Template[name])
	}

	result := make([]TemplateFile, 0)
	for _, item := range tmp {
		data, err := ioutil.ReadFile(item.Path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, TemplateFile{
			Path:    item.Dest + item.Name + item.Suffix,
			Content: string(data),
		})
	}

	return result
}

func ListAvailableResource(selected []TemplateItem) []string {
	selectedMap := make(map[string]TemplateItem)
	for _, item := range selected {
		selectedMap[item.Name] = item
	}
	result := make([]string, 0)
	for name, item := range Template {
		if item.Multiable {
			result = append(result, name)
			continue
		}
		if _, ok := selectedMap[name]; !ok {
			result = append(result, name)
		}
	}

	return result
}
