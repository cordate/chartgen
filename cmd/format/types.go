package format

type CommonParam struct {
	ReleaseName string `json:"releaseName"`
	Namespace   string `json:"namespace"`
	Labels      string `json:"labels"`
}

var GlobalParam = CommonParam{
	ReleaseName: "{{ .Release.Name }}",
	Namespace:   "{{ .Release.namespace }}",
	Labels:      "{{- include \"example.labels\" . | nindent 4 }}",
}
