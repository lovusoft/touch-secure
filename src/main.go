package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
	"strings"
	"touch-secure/src/entity"
)

func main() {
	//user := entity.User{}
	//user.Name = "长相思"
	//user.Password = "各自安好"
	//user.Email = "letitgo@qq.com"
	//user.Role = "COR"
	//fmt.Println(action.Join(user))

	//entity.Question{}.RadomOneQuestion()

	//var qs entity.Question
	//qs.Question = "去年今日此门中"
	//qs.Answer = "人面桃花相映红"
	//entity.Question{}.AddQuestions(qs)
	//var op entity.Opition
	//op.QuestionID = 1
	//op.Option = "人面不知何处去"
	//entity.Opition{}.AddOption(op)
	//var op1 entity.Opition
	//op1.QuestionID = 1
	//op1.Option = "桃花依旧笑春风"
	//entity.Opition{}.AddOption(op1)
	//var op2 entity.Opition
	//op2.QuestionID = 1
	//op2.Option = "人面桃花相映红"
	//entity.Opition{}.AddOption(op1)

	router := gin.Default()
	router.Use(Cors())
	router.Use(TLSHandler())
	router.POST("/getQuestion", func(context *gin.Context) {
		var question entity.Question
		question = entity.Question{}.RadomOneQuestion()
		fmt.Println(question)
		context.PureJSON(http.StatusOK, gin.H{"question": question})
	})

	router.POST("/addQuestion", func(context *gin.Context) {
		var question entity.Question
		questionName := context.PostForm("question")
		if questionName != "" {
			question.Name = questionName
			entity.Question{}.AddQuestions(question)
			context.PureJSON(http.StatusOK, gin.H{"res": true})
		} else {
			context.PureJSON(http.StatusOK, gin.H{"res": false})
		}
	})

	//_ = router.RunTLS(":1314", "src/ssl.pem", "src/ssl.key")
	_ = router.Run(":1314")
}

//Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}

// TLSHandler tls中间件
func TLSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
