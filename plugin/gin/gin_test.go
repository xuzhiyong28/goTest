package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	r.GET("/someProtoBuf" , func(c *gin.Context) {
		person := &Person{
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

func TestRedirectDemo(t *testing.T){
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
