package api

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"k8s-api/config"
	"k8s-api/model"
	"net/http"
	"strings"
)
//获取pod的信息
func GetConDescription() []model.ConDescription {
	url:=config.Conf("url")+"api/v1/namespaces/default/pods"
	request,_:=http.NewRequest("GET",url,nil)
	request.Header.Add("Authorization","Bearer "+config.Conf("token"))
	response,_:=Client().Do(request)
	defer response.Body.Close()
	result,_:=ioutil.ReadAll(response.Body)
	j,_:=simplejson.NewJson(result)
	answer:=j.Get("items")
	rows,_:=answer.Array()
	var con []model.ConDescription
	for index,_:=range rows{
		var c1 model.ConDescription
		node:=answer.GetIndex(index)
		c1.Id=index+1
		c1.Name=node.Get("metadata").Get("name").MustString()
		c1.NameSpace=node.Get("metadata").Get("namespace").MustString()
		c1.NodeName=node.Get("spec").Get("nodeName").MustString()
		c1.Status=node.Get("status").Get("phase").MustString()
		img:=node.Get("status").Get("containerStatuses")
		c1.RestartCount=img.GetIndex(0).Get("restartCount").MustInt()
		StartedAt:=node.Get("status").Get("startTime").MustString()
		replacer:=strings.NewReplacer("T"," ","Z","").Replace(StartedAt)
		c1.StartedAt=replacer
		con=append(con,c1)
	}
	return con
}
//创建service
func CreateSvc(name string,port string) error {
	yaml,err:=yaml.JSONToYAML(SvcYaml(name,port))
	if err!=nil{
		return err
	}
	fmt.Println(string(yaml))
	body:=bytes.NewBuffer(yaml)
	url:=config.Conf("url")+"api/v1/namespaces/default/services"
	request,_:=http.NewRequest("POST",url,body)
	request.Header.Add("Content-Type","application/yaml")
	request.Header.Add("Authorization","Bearer "+config.Conf("token"))
	response,_:=Client().Do(request)
	defer response.Body.Close()

	return nil
}
//创建deployment
func CreateDep(name string,replicas string,image string,env string,volum string)error  {
	yaml,err:=yaml.JSONToYAML(DYaml(name,replicas,image,env,volum))
	if err!=nil{
		return err
	}
	fmt.Println(string(yaml))
	body:=bytes.NewBuffer(yaml)
	url:=config.Conf("url")+"apis/apps/v1/namespaces/default/deployments"
	request,_:=http.NewRequest("POST",url,body)

	request.Header.Add("Content-Type","application/yaml")
	request.Header.Add("Authorization","Bearer "+config.Conf("token"))
	response,_:=Client().Do(request)

	defer response.Body.Close()
	return nil
}
