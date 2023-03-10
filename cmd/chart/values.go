package chart

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
	"fmt"
	"github.com/spf13/chartgen/cmd/input"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/**
 * 镜像配置处理
 */
type ImageItem struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
}

type ServiceItem struct {
	Type  string           `json:"type"`
	Ports []v1.ServicePort `json:"ports"`
}

type Volumes struct {
	v1.Volume
	Enabled bool
}

type ValueEntity struct {
	ReplicaCount       int                       `json:"replicaCount"`
	Image              ImageItem                 `json:"image"`
	ImagePullSecrets   []v1.LocalObjectReference `json:"imagePullSecrets"`
	NameOverride       string                    `json:"nameOverride"`
	FullnameOverride   string                    `json:"fullnameOverride"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodSecurityContext v1.SecurityContext        `json:"podSecurityContext"`
	SecurityContext    v1.SecurityContext        `json:"securityContext"`
	Ports              []v1.ContainerPort        `json:"ports"`
	Service            ServiceItem               `json:"service"`
	Resources          v1.ResourceRequirements   `json:"resources"`
	NodeSelector       v1.NodeSelector           `json:"nodeSelector"`
	Affinity           v1.Affinity               `json:"affinity"`
	Env                map[string]string         `json:"env"`
	Tolerations        []v1.Toleration           `json:"tolerations"`
}

func getPorts() []v1.ContainerPort {
	ports := make([]v1.ContainerPort, 0)
	for added := input.GetBool("请输入是否添加容器暴露端口："); added; added = input.GetBool("请输入是否继续添加容器暴露端口") {
		ports = append(ports, v1.ContainerPort{
			Name:          input.GetString("请输入容器暴露端口的别名(http)"),
			ContainerPort: int32(input.GetInt("请输入容器暴露的端口")),
			Protocol:      "TCP",
		})
	}

	return ports
}

func getServicePorts() []v1.ServicePort {
	ports := make([]v1.ServicePort, 0)
	ports = append(ports, v1.ServicePort{
		Name:       input.GetString("请输入Service端口的别名(http)"),
		TargetPort: intstr.FromString(input.GetString("请输入容器暴露端口的别名(http)")),
		Protocol:   "TCP",
		Port:       int32(input.GetInt("请输入服务暴露的端口(80)")),
	})

	for added := input.GetBool("请输入是否继续添加服务暴露端口"); added; added = input.GetBool("请输入是否继续添加服务暴露端口") {
		ports = append(ports, v1.ServicePort{
			Name:       input.GetString("请输入容器暴露端口的别名(http)"),
			TargetPort: intstr.FromString(input.GetString("请输入容器暴露端口的别名(http)")),
			Protocol:   "TCP",
			Port:       int32(input.GetInt("请输入服务暴露的端口(80)")),
		})
	}

	return ports
}

func addEnvParams() map[string]string {
	member := make(map[string]string)
	for added := input.GetBool("请输入是否继续添加环境变量"); added; added = input.GetBool("请输入是否继续添加环境变量") {
		member[input.GetString("请输入环境变量Key")] = input.GetString("请输入环境变量的默认值")
	}

	return member
}

var ValueConfig ValueEntity

func InitValueConfig() ValueEntity {
	ValueConfig = ValueEntity{
		ReplicaCount: input.GetInt("pod副本数"),
		Image: ImageItem{
			Repository: input.GetString("输入镜像地址"),
			PullPolicy: "IfNotPresen",
		},
		ImagePullSecrets:   nil,
		NameOverride:       "",
		FullnameOverride:   "",
		PodAnnotations:     nil,
		PodSecurityContext: v1.SecurityContext{},
		SecurityContext:    v1.SecurityContext{},
		Ports:              getPorts(),
		Service: ServiceItem{
			Type:  "ClusterIp",
			Ports: getServicePorts(),
		},
		Resources:    v1.ResourceRequirements{},
		NodeSelector: v1.NodeSelector{},
		Affinity:     v1.Affinity{},
		Env:          addEnvParams(),
		Tolerations:  make([]v1.Toleration, 0),
	}

	return ValueConfig
}

func FormatValueYaml(cmd *cobra.Command) string {
	var (
		c *cue.Context
		v cue.Value
	)

	c = cuecontext.New()
	v = c.Encode(ValueConfig)
	bytes, _ := yaml.Encode(v)
	fmt.Println(string(bytes))
	return string(bytes)
}
