package chart

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
	"cuelang.org/go/pkg/strings"
	"fmt"
	"github.com/spf13/cobra"
)

type ChartInfo struct {
	ApiVersion  string `json:"apiVersion"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	AppVersion  string `json:"appVersion"`
}

func FormatChartYaml(cmd *cobra.Command) string {
	var (
		c *cue.Context
		v cue.Value
	)

	val := ChartInfo{
		ApiVersion:  "v2",
		Name:        GetStringOrDefault(cmd, "appName"),
		Description: "A Helm chart for Kubernetes",
		Type:        "application",
		Version:     GetStringOrDefault(cmd, "chartVersion"),
		AppVersion:  GetStringOrDefault(cmd, "appVersion"),
	}

	c = cuecontext.New()
	v = c.Encode(val)
	bytes, _ := yaml.Encode(v)
	fmt.Println(string(bytes))
	return string(bytes)
}

func GetStringOrDefault(cmd *cobra.Command, name string) string {
	value, _ := cmd.Flags().GetString(name)
	value = strings.TrimSpace(value)
	if value == "" {
		return cmd.Flag(name).DefValue
	}
	return value
}
