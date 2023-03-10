package template

var envSecret string = `apiVersion: v1
kind: Secret
metadata:
  name: {{ .Release.Name }}-env-secret
  namespace: {{ .Release.namespace }}
  labels:
    app.kubernetes.io/instance: {{ .Release.Name }}
stringData:
  {{- range name, value := .Values.env  }}
  {{ $name }}: {{ $value | quote }}
  {{- end }}
`

func FormatEnvSecret() string {
	return envSecret
}
