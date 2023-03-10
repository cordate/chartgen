package chart

import (
	"fmt"
	"github.com/spf13/chartgen/cmd/chart/template"
	"github.com/spf13/cobra"
	"os"
)

type ChartFile struct {
	FileName string
	Content  string
}

type Chart struct {
	ChartInfo  ChartFile
	Values     ChartFile
	HelmIgnore ChartFile
	HelpersTpl ChartFile
	EnvSecret  ChartFile
	Template   []ChartFile
}

func BuildChart(cmd *cobra.Command) {
	// format.FormatExample()
	InitValueConfig()
	chart := Chart{
		ChartInfo:  ChartFile{"chart/Chart.yaml", FormatChartYaml(cmd)},
		Values:     ChartFile{"chart/values.yaml", FormatValueYaml(cmd)},
		HelmIgnore: ChartFile{"chart/.helmignore", FormatHelmIgnore()},
		HelpersTpl: ChartFile{"chart/templates/_helpers.tpl", template.FormatHelpsTpl()},
		EnvSecret:  ChartFile{"chart/templates/env-secret.yaml", template.FormatEnvSecret()},
	}
	chart.Template = make([]ChartFile, 0)
	for _, item := range template.FormatResources() {
		chart.Template = append(chart.Template, ChartFile{
			FileName: item.Path,
			Content:  item.Content,
		})
	}
	formatChart(chart, cmd)
}

func formatChart(chart Chart, cmd *cobra.Command) {
	// 准备chart包目录
	os.Mkdir("chart", os.ModePerm)
	os.Mkdir("chart/templates", os.ModePerm)
	// 创建基础文件
	createFile(chart.ChartInfo)
	createFile(chart.Values)
	createFile(chart.HelmIgnore)
	createFile(chart.HelpersTpl)
	createFile(chart.EnvSecret)
	// 生成模板文件
	for _, item := range chart.Template {
		createFile(item)
	}
}

func createFile(content ChartFile) {
	file, err := os.OpenFile(content.FileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(content.Content)
	file.WriteString(content.Content)
}
