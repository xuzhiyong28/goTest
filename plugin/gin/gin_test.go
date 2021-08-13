package gin

import (
	"example/plugin/gin/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestDemoBase(t *testing.T) {
	r := gin.Default()

	// http://localhost:8080
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	// http://localhost:8080/someGet/xuzhiyong
	r.GET("/someGet/:name", func(c *gin.Context) {
		name := c.Param("name")
		// 返回字符串方式
		//c.String(http.StatusOK, name)
		// 返回json方式
		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{"status_code": http.StatusOK, "status": "ok"},
			"message": "http is ok , return = " + name,
		})
	})

	// http://localhost:8080/user?name=xuzhiyong
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "许志勇")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// curl -X POST http://127.0.0.1:8080/form -H "Content-Type:application/x-www-form-urlencoded" -d "username=xuzhiyong&password=123456"
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	// 文件上传 - 单个
	r.MaxMultipartMemory = 8 << 20 //限制上传大小
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		c.SaveUploadedFile(file, createPathFile(file.Filename))
		c.String(http.StatusOK, file.Filename)
	})

	// 文件上传 - 多个
	r.POST("/muitupload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, createPathFile(file.Filename)); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestDemoGroup(t *testing.T) {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		// http://localhost:8080/v1/login?name=xuzhiyong
		v1.GET("/login", func(c *gin.Context) {
			name := c.DefaultQuery("name", "jack")
			c.String(200, fmt.Sprintf("hello %s\n", name))
		})
		v1.GET("/submit", func(c *gin.Context) {
			name := c.DefaultQuery("name", "lily")
			c.String(200, fmt.Sprintf("hello %s\n", name))
		})
	}
	v2 := r.Group("v2")
	{
		// http://localhost:8080/v2/login?name=xuzhiyong
		v2.GET("/login", func(c *gin.Context) {
			name := c.DefaultQuery("name", "jack")
			c.String(200, fmt.Sprintf("hello %s\n", name))
		})
		v2.GET("/submit", func(c *gin.Context) {
			name := c.DefaultQuery("name", "lily")
			c.String(200, fmt.Sprintf("hello %s\n", name))
		})
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// json类型的数据绑定
func TestStructJsonDemo(t *testing.T) {

	// curl http://127.0.0.1:8080/loginJSON -H 'Content-Type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST
	r := gin.Default()
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// 判断密码是否正确
		if json.User != "root" || json.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestStructFormDemo(t *testing.T) {
	r := gin.Default()
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 各种数据格式的响应
func TestStructResponseDemo(t *testing.T) {
	r := gin.Default()
	// 返回json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})

	// 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	// XML响应
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "abc"})
	})

	// YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "zhangsan"})
	})

	// protobuf格式,谷歌开发的高效存储读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		person := &proto.Person{
			Name:  "Jack",
			Age:   18,
			Hobby: []string{"sing", "dance", "basketball", "rap"},
		}
		c.ProtoBuf(http.StatusOK, person)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestHtmlResponseDemo(t *testing.T) {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	// http://localhost:8080/index
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", gin.H{"title": "我是测试", "ce": "123456"})
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestRedirectDemo(t *testing.T) {
	// 重定向
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 注册中间件 - 所有请求都经过此中间件
func TestUseAllDemo(t *testing.T) {
	r := gin.Default()
	// 注册中间件
	r.Use(func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	})

	// 额外用{}括起来没什么意思，只是为了规范
	{
		r.GET("/ce", func(c *gin.Context) {
			req, exists := c.Get("request")
			if exists {
				fmt.Println("request:", req)
				c.JSON(http.StatusOK, gin.H{"request": req})
			}
		})
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestUsePartDemo(t *testing.T) {
	r := gin.Default()
	midFun := func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
	r.GET("/ce", midFun, func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func TestCookieBaseDemo(t *testing.T) {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			// secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 模拟实现权限验证中间件
// 有2个路由，login和home
// login用于设置cookie
// home是访问查看信息的请求
// 在请求home之前，先跑中间件代码，检验是否存在cookie
func TestCookieProjectDemo(t *testing.T) {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login success !")
	})

	AuthMidUse := func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
	r.GET("/home", AuthMidUse, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

type People struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 2、在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

/*
   curl -X GET "http://127.0.0.1:8080/testing?name=&age=12&address=beijing"
   curl -X GET "http://127.0.0.1:8080/testing?name=lmh&age=12&address=beijing"
   curl -X GET "http://127.0.0.1:8080/testing?name=adz&age=12&address=beijing"
*/
func TestValidatorDemo(t *testing.T) {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NotNullAndAdmin", func(fl validator.FieldLevel) bool {
			fmt.Println(fl)
			return true
		})
	}
	r.GET("/5lmh", func(c *gin.Context) {
		var person proto.Person
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}


func TestLogDemo(t *testing.T){
	gin.DisableConsoleColor()
	f , _ := os.Create(createPathFile("gin.log"))
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func createPathFile(fileName string) string {
	path, _ := os.Getwd()
	var ostype = runtime.GOOS
	filePath := ""
	if ostype == "windows" {
		filePath = path + "\\" + fileName
	} else if ostype == "linux" {
		filePath = path + "/" + fileName
	}
	return filePath
}
