package routers

import (
	"3gnx/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	//Jwt令牌校验
	//r.Use(func(ctx *gin.Context) {
	//	//登录本身不用校验
	//	if ctx.Request.URL.Path == "/api/user/login" {
	//		return
	//
	//	}
	//	auth := ctx.Request.Header.Get("Authorization")
	//	if auth == "" {
	//		ctx.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}
	//	segs := strings.Split(auth, " ")
	//	if len(segs) != 2 {
	//		ctx.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}
	//	tokenStr := segs[1]
	//	claims := &middles.UserClaims{}
	//	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
	//		return []byte("hhdfsukadgkDGhkDGyughhj"), nil
	//	})
	//	if err != nil {
	//		//有人搞我，这我忍得了，抱紧
	//		ctx.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}
	//	if token == nil || !token.Valid || claims.Username == "" {
	//		ctx.AbortWithStatus(http.StatusBadRequest)
	//		return
	//	}
	//	//now := time.Now()
	//	//if claims.ExpiresAt.Time.Sub(now) < time.Second*50 {
	//	//	//通过校验后，我们刷新过期时间
	//	//	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * 10))
	//	//	tokenStr, err = token.SignedString([]byte("hhdfsukadgkDGhkDGyughhj"))
	//	//	if err != nil {
	//	//		return
	//	//	}
	//	//	ctx.Header("x-jwt-token", tokenStr)
	//	//}
	//
	//})
	//session校验
	store := cookie.NewStore([]byte("secret-key"))
	store.Options(sessions.Options{
		MaxAge:   3600, // 设置过期时间为1小时，单位为秒
		HttpOnly: true, // 设置仅HTTP访问
	})
	// 注册会话中间件
	r.Use(sessions.Sessions("session", store), func(c *gin.Context) {
		// 获取当前会话
		session := sessions.Default(c)

		// 如果会话已过期，则重新设置会话信息
		if session.Get("username") == nil {
			// 设置会话值
			session.Set("username", "exampleuser")

			// 设置会话过期时间为1小时
			session.Options(sessions.Options{
				MaxAge:   3600, // 1小时，单位为秒
				HttpOnly: true,
			})

			// 保存会话
			err := session.Save()
			if err != nil {
				// 处理保存会话时的错误
				c.String(http.StatusInternalServerError, "Failed to save session")
				return
			}

			//c.String(http.StatusOK, "Session has been set")
		} else {
			//c.String(http.StatusOK, "Session already exists")
		}
	})
	V1Group := r.Group("api")
	{
		//V1Group.Use(sessions.Sessions("session", store))
		//middles.GetSessionId(V1Group)
		//待办事项
		//1.用户注册(包含添加用户)
		V1Group.POST("/user/register", controller.UserRegister)
		//2.用户登录
		V1Group.POST("/user/login", controller.UserLogin)
		//3.管理员登录
		V1Group.POST("/manager/login", controller.MangerLogin)
		//4.注册时发送验证码
		V1Group.POST("/user/register-email", controller.SendEmailRegister)
		//5.修改密码验证邮箱时发送验证码
		V1Group.POST("/user/reset-email", controller.SendEmailReSet)
		//6.验证身份
		V1Group.POST("/user/VerifyCode-email", controller.ResetCodeVerify)
		//7.重设密码
		V1Group.POST("/user/reset-password", controller.ResetPassword)
		//8.学生报名
		V1Group.POST("/user/login/apply", controller.StudentApplySuccess)
		//9.查询一面通过学生
		V1Group.GET("/manager/login/getOne", controller.GetApplyStuListOne)
		//10.查询二面通过学生
		V1Group.GET("/manager/login/getTwo", controller.GetApplyStuListTwo)
		//11.学生查询自己一面状况
		V1Group.POST("/user/login/getOne", controller.StuGetApplyStuListOne)
		//12.学生查询自己er面状况
		V1Group.POST("/user/login/getTwo", controller.StuGetApplyStuListTwo)
		//13.设置一面通过
		V1Group.PUT("/manager/login/SetOne", controller.SetOneSuccessfully)
		//14.设置二面通过
		V1Group.PUT("/manager/login/SetTwo", controller.SetTwoSuccessfully)
	}
	return r
}
