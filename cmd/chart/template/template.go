package template

type TemplateItem struct {
	Name      string
	Multiable bool
	Path      string
	Dest      string
	Suffix    string
	Config    any
}

var Template map[string]TemplateItem = make(map[string]TemplateItem)

func init() {
	Template["deployment"] = TemplateItem{"deployment", false, "example/templates/deployment.yaml", "chart/templates/", ".yaml", ""}
	Template["configmap"] = TemplateItem{"configmap", false, "example/templates/configmap.yaml", "chart/templates/", ".yaml", ""}
	Template["service"] = TemplateItem{"service", false, "example/templates/service.yaml", "chart/templates/", ".yaml", ""}
	Template["statefulset"] = TemplateItem{"statefulset", false, "example/templates/statefulset.yaml", "chart/templates/", ".yaml", ""}
	Template["_helpers"] = TemplateItem{"_helpers", false, "example/templates/_helpers.tpl", "chart/templates/", ".tpl", ""}
}

func Register(name string, template TemplateItem) {
	Template[name] = template
}
