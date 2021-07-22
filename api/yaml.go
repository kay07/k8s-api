package api

import (
	"encoding/json"
	"strconv"
	"strings"
)

func SvcYaml(name string,port string)[]byte  {
	svc:=[]byte(`{
    "apiVersion":"v1",
    "kind":"Service",
    "metadata":{
        "labels":{
            "app":`+name+`
        },
        "name":`+name+`
    },
    "spec":{
        "ports":[
            {
                "port":`+port+`,
                "targetPort":`+port+`,
                "name":`+name+`
            }
        ],
        "selector":{
            "app":`+name+`
        },
        "type":"NodePort"
    }
}`)
	return  svc
}

func DYaml(name string,replicas string,image string,env string,volum string)  []byte{
	d:=Deployment{}
	d.Kind="Deployment"
	d.ApiVersion="apps/v1"
	d.Name=name
	d.Spec.SelectorMatchLabels.App=name
	d.Spec.Template.SpecTemplateMetadata.SelectorMatchLabelsApp.App=name
	rep,_:=strconv.Atoi(replicas)
	d.Replicas=rep

	var c Containers
	c.Name=name
	c.ImagePullPolicy="Always"
	c.Image=image
	d.Spec.Template.SpecTemplateSpec.ListCon=append(d.Spec.Template.SpecTemplateSpec.ListCon,c)

	//添加环境变量,环境变量kv之间用","分割，环境变量之间用";"分割
	e:=strings.Split(env,";")
		for _,v:=range e{
			val:=strings.Split(v,",")
			if len(val)!=2{
				//不是一个kv键值对
			}else {
				//依次将kv键值对放到env中
				var envv Env
				envv.Name=val[0]
				envv.Value=val[1]
				d.Spec.Template.SpecTemplateSpec.ListCon[0].ListEnv=append(d.Spec.Template.SpecTemplateSpec.ListCon[0].ListEnv,envv)
			}
		}
	//添加数据卷,数据卷kv之间用","分割，数据卷之间用";"分割
	//如设置的值为a,b,c;d,e,f
	//a和d表示名称，b和e表示内部挂载路径；c和f表示外部挂载路径
	v:=strings.Split(volum,";")
	for _,vv:=range v{
		val:=strings.Split(vv,",")
		if len(val)!=3{
			//不是一个kv键值对(a,b,c类型)
		}else {
			//依次将kv键值对放到volumes中
			var vvv VolumeMounts//内部路径
			var vout Volumes//外部路径

			vvv.Name=val[0]//设置名称
			vout.Name=val[0]//设置名称（外部），两个名称需要一致才能匹配上
			vvv.MountPath=val[1]//设置内部路径
			vout.HostPath.Path=val[2]//设置外部路径
			d.Spec.Template.SpecTemplateSpec.ListCon[0].ListVolumeMounts=append(d.Spec.Template.SpecTemplateSpec.ListCon[0].ListVolumeMounts,vvv)
			d.Spec.Template.SpecTemplateSpec.ListVolumes=append(d.Spec.Template.SpecTemplateSpec.ListVolumes,vout)
		}
	}

	//添加外部挂载卷
	byte,_:=json.Marshal(&d)
	return byte
}
//整个json
type Deployment struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata `json:"metadata"`
	Spec `json:"spec"`
}
//解析metadata.name
type Metadata struct {
	Name string `json:"name"`
}
//解析spec
type Spec struct {
	Replicas int `json:"replicas"`
	SelectorMatchLabels `json:"selector"`
	Template  `json:"template"`
}
//解析spec.selector
type SelectorMatchLabels struct {
	 SelectorMatchLabelsApp `json:"matchLabels"`
}
//解析spec.selector.matchlabels
type SelectorMatchLabelsApp struct {
	App string `json:"app"`
}
//sepc.template
type Template struct {
	SpecTemplateMetadata `json:"metadata"`
	SpecTemplateSpec  `json:"spec"`
}
//spec.template.metadata
type SpecTemplateMetadata struct {
	SelectorMatchLabelsApp `json:"labels"`
}
//spec.template.spec
type SpecTemplateSpec struct {
	ListCon []Containers  `json:"containers"`
	ListVolumes []Volumes `json:"volumes"`
}
//spec.template.spec.containers
//容器内部是list
type Containers struct {
	Name string `json:"name"`
	Image string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	//env是list
	ListEnv []Env `json:"env"`
	//volumemounts是list
	ListVolumeMounts []VolumeMounts `json:"volumeMounts"`
}
type Env struct {
	Name string `json:"name"`
	Value string `json:"value"`
}
type VolumeMounts struct {
	MountPath string `json:"mountPath"`
	Name string `json:"name"`
}
//spec.template.spec.volumes
type Volumes struct {
	Name string `json:"name"`
	HostPath  `json:"hostPath"`
}
type HostPath struct {
	Path string `json:"path"`
}