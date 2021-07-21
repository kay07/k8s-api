package middleware

import "k8s-api/model"

func Page(num int,c []model.ConDescription)  model.ConPage{
	//每页默认显示10条
	everyPage:=10
	content:=model.ConPage{}
	content.Amount=len(c)

	//当前不会出现越界情况，直接输出即可
	if num*everyPage<=content.Amount{
		content.CD=c[(num-1)*everyPage:num*everyPage]
		return  content
	}else {
		//当前会出现越界
		content.CD=c[(num-1)*everyPage:]
		return content
	}
}