package main

import (
	"3gnx/dao"
	"3gnx/routers"
	"github.com/gin-contrib/cors"
)

func main() {
	r := routers.SetUpRouter()
	// 注册路由
	//r := routers.SetUpRouter()
	//store := cookie.NewStore([]byte("secret-key"))
	//store.Options(sessions.Options{
	//	MaxAge:   3600, // 设置过期时间为1小时，单位为秒
	//	HttpOnly: true, // 设置仅HTTP访问
	//})

	// 注册会话中间件
	//r.Use(sessions.Sessions("session", store))
	//router := gin.Default()
	// 使用cors中间件，允许所有来源访问
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},              // 允许所有来源
		AllowMethods:     []string{"*"},              // 允许的请求方法
		AllowHeaders:     []string{"*"},              // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"}, // 允许暴露的响应头
		AllowCredentials: true,                       // 允许携带凭证（如Cookie）
	}))

	//创建连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	//r.POST("/user/register-eee", func(c *gin.Context) {
	//	// 获取当前会话
	//	session := sessions.Default(c)
	//
	//	// 如果会话已过期，则重新设置会话信息
	//	if session.Get("username") == nil {
	//		// 设置会话值
	//		session.Set("username", "exampleuser")
	//
	//		// 设置会话过期时间为1小时
	//		session.Options(sessions.Options{
	//			MaxAge:   3600, // 1小时，单位为秒
	//			HttpOnly: true,
	//		})
	//
	//		// 保存会话
	//		err := session.Save()
	//		if err != nil {
	//			// 处理保存会话时的错误
	//			c.String(http.StatusInternalServerError, "Failed to save session")
	//			return
	//		}
	//
	//		c.String(http.StatusOK, "Session has been set")
	//	} else {
	//		c.String(http.StatusOK, "Session already exists")
	//	}
	//	sessionID, _ := c.Cookie("session")
	//	//sessionID := middles.GetSessionId(c)
	//	var requestData struct {
	//		Email string `form:"email" json:"email"`
	//	}
	//	if err := c.ShouldBind(&requestData); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
	//		return
	//	}
	//	// 调用UserRegister函数进行注册
	//	result := server.SendEmail(requestData.Email, sessionID, false)
	//
	//	// 根据注册结果返回相应的数据给前端
	//	// 封装注册结果为RestBean对象
	//	fmt.Println("kkk")
	//	fmt.Println(session)
	//	fmt.Println(c.Cookie("session"))
	//	fmt.Println(sessions.Default(c).Get("session"))
	//	var restBeanRegister *models.RestBean
	//	if result == "" {
	//		restBeanRegister = models.SuccessRestBeanWithData("邮件已发送，请注意查收")
	//
	//	} else {
	//		restBeanRegister = models.FailureRestBeanWithData(http.StatusBadRequest, result)
	//	}
	//	//返回注册结果给前端
	//	c.JSON(restBeanRegister.Status, restBeanRegister)
	//})

	r.Run(":8080")
}
