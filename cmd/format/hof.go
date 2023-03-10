package format

import (
	"fmt"
	"github.com/spf13/chartgen/cmd/input"
	"os"
	"text/template"
)

var scheme = `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .GlobalParam.ReleaseName }}-env-secret
  namespace: {{ .GlobalParam.Namespace }}
  labels:
    {{ .GlobalParam.Labels}}
data:
  {{- if .EnvEnabled }}
  {{ .Data }}
  {{- else }}
  {{- range $name, $value := .KvMap }}
  {{ $name }}: {{ $value }}
  {{- end }}
  {{- end }}
`

type ConfigMapParam struct {
	GlobalParam CommonParam       `json:"globalParam"`
	EnvEnabled  bool              `json:"envEnabled"`
	Data        string            `json:"data"`
	KvMap       map[string]string `json:"kv"`
}

func GetParams() ConfigMapParam {
	result := ConfigMapParam{
		GlobalParam: GlobalParam,
		EnvEnabled:  input.GetBool("请输入是否支持环境变量"),
		Data: `{{- range name, value := .Values.configmap.data }}
  {{ $name }}: {{ $value | quote }}
  {{- end }}`,
	}

	if !result.EnvEnabled {
		kvs := make(map[string]string)
		for ok := input.GetBool("请输入是否添加configmap的data"); ok; ok = input.GetBool("请输入是否添加configmap的data") {
			kvs[input.GetString("请输入Key")] = input.GetString("请输入value")
		}
		result.KvMap = kvs
	}

	return result
}

func FormatExample() {
	t, err := template.New("scheme").Parse(scheme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t.Execute(os.Stdout, GetParams())
}
