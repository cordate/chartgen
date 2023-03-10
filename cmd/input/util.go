package input

const (
	STRING = iota
	INT
	BOOL
	FLOAT
)

type FlagList struct {
	Name    string
	Usage   string
	Type    int
	Value   any
	Default any
}

var Checks = make([]FlagList, 0)

func init() {
	Checks = append(Checks, FlagList{"appName", "chart包名称", STRING, "", ""})
	Checks = append(Checks, FlagList{"chartVersion", "chart包版本(默认0.1.0)", STRING, "0.1.0", "0.1.0"})
	Checks = append(Checks, FlagList{"appVersion", "应用的版本(默认1.16.0)", STRING, "1.16.0", "1.16.0"})
	// Checks = append(Checks, FlagList{"replicas", "pod副本数", INT, 1})
	// Checks = append(Checks, FlagList{"repository", "镜像地址", STRING, ""})
}
