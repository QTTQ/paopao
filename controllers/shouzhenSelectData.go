// /*
//  * @Author: QTTQ
//  * @Date: 2018-10-23 11:19:50
//  * @LastEditors: QTTQ
//  * @LastEditTime: 2018-12-30 15:41:59
//  * @Email: 1321510155@qq.com
//  */

// package controllers

// import (
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"paopao/models"
// )

// // type ApiRes struct {
// // 	Code uint        `json:"code"`
// // 	Msg  string      `json:"msg"`
// // 	Data interface{} `json:"data"`
// // }
// var slc =[]int{111,222,3333}
// func RemoveRepByMap(slc []int) []int {
// 	result := []int{}
// 	tempMap := map[int]byte{}  // 存放不重复主键
// 	for _, e := range slc{
// 		l := len(tempMap)
// 			fmt.Println(tempMap,l,"............")
// 		tempMap[e] = 0
// 		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
// 			result = append(result, e)
// 		}
// 	}
// 	return result
// }
// func GetShouzhenSelectData(c *gin.Context) {
// 	data, _ := models.GetShouzhenSelectData()
// 	mapBox:=map[string][]string{}
// 	// mapList:=[]interface{}
// 	mapList:=[]map[string][]string{}

// 	for _, value := range data {
// 		mapBox[value.Type]=append(mapBox[value.Type],value.Category)
// 	}
// 	// RemoveRepByMap(slc)
// 	for k,v:=range mapBox{
// 		mapList=append(mapList,map[string][]string{k:v})
// 	}
// 	c.JSON(http.StatusOK, ApiRes{
// 		Code: 0,
// 		Msg:  "获取文章成功",
// 		Data: gin.H{
// 			"data": mapList,
// 		},
// 	})
// 	return
// }

/*
 * @Author: QTTQ
 * @Date: 2018-10-23 11:19:50
 * @LastEditors: QTTQ
 * @LastEditTime: 2018-12-30 15:41:59
 * @Email: 1321510155@qq.com
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
)

func GetShouzhenSelectData(c *gin.Context) {
	data, err := models.GetShouzhenSelectData()
	res := map[string][]string{}
	for _, value := range data {
		if len(res[value.Type]) == 0 {
			res[value.Type] = []string{}
		}
		res[value.Type] = append(res[value.Type], value.Category)
		// res=append(res,map[string][]string{value.Type:value.Category})
	}
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "获取数据错误",
			})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "0", "msg": "成功", "results": res})
	return
}
