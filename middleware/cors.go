package middleware

import (
	"github.com/gin-gonic/gin"
	// "strings"
	"net/http"
)

func Cors() gin.HandlerFunc {
	// return func(c *gin.Context) {
	//     method := c.Request.Method
	//     origin := c.Request.Header.Get("Origin")
	//     var headerKeys []string
	//     for k, _ := range c.Request.Header {
	//         headerKeys = append(headerKeys, k)
	//     }
	//     headerStr := strings.Join(headerKeys, ", ")
	//     if headerStr != "" {
	//         headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
	//     } else {
	//         headerStr = "access-control-allow-origin, access-control-allow-headers"
	// 	}
	// 	fmt.Println("............")
	//     if origin != "" {
	// 		fmt.Println("............")
	//         //下面的都是乱添加的-_-~
	//         // c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//         c.Header("Access-Control-Allow-Origin", "*")
	//         c.Header("Access-Control-Allow-Headers", headerStr)
	//         c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//         // c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	//         c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//         // c.Header("Access-Control-Max-Age", "172800")
	//         c.Header("Access-Control-Allow-Credentials", "true")
	//         c.Set("content-type", "application/json")
	//     }
	//     //放行所有OPTIONS方法
	//     if method == "OPTIONS" {
	//         c.JSON(http.StatusOK, "Options Request!")
	//     }
	//     c.Next()
	// }
	return func(c *gin.Context) {
		w := c.Writer
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,X-Auth-Token")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
			//放行所有OPTIONS方法
	    method := c.Request.Method
	    if method == "OPTIONS" {
	        c.JSON(http.StatusOK, "Options Request!")
	    }
		c.Next()
	}
}
