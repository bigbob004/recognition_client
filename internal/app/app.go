package app

import (
	"FaceRecognitionClient/internal/handler/auth"
	"FaceRecognitionClient/internal/handler/recognition"
	"FaceRecognitionClient/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	indexFileName           = "index.html"
	recognitionFormFileName = "recognition.html"
	trainingFormFileName    = "training.html"
	signInFileName          = "sign_in.html"
	signUpFileName          = "sign_up.html"
	aboutFileName           = "about.html"
)

type App struct {
	recognitionHandler recognition.Handler
	authHandler        auth.Handler
}

func NewApp(recognitionHandler recognition.Handler, authHandler auth.Handler) App {
	return App{recognitionHandler: recognitionHandler, authHandler: authHandler}
}

func (a *App) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true

	router.Use(middleware.HandlerName())

	//Загружаем статические файлы (.ccs)
	router.LoadHTMLGlob("templates/*.html")
	//TODO возможно, можно лишь один раз указать, где искать статические файлы, но у меня как-то не вышло
	router.StaticFS("/static/", http.Dir("templates/static/"))
	router.StaticFS("/home/static/", http.Dir("templates/static/"))

	resourses := router.Group("/home", middleware.TimeoutMiddleware(5*time.Second), userIdentity)
	{

		//Создаём обработчик для корневой страницы
		resourses.GET("/", func(c *gin.Context) {
			c.HTML(
				http.StatusOK,
				indexFileName,
				gin.H{
					"title": "Home Page",
				},
			)
		})

		resourses.GET("/logOut", a.logOut)

		resourses.GET("about", func(c *gin.Context) {
			c.HTML(
				http.StatusOK,
				aboutFileName,
				gin.H{},
			)
		})

		//Создаём переход на страничку с распознаванием
		resourses.GET("/recognition", func(c *gin.Context) {
			c.HTML(
				http.StatusOK,
				recognitionFormFileName,
				gin.H{},
			)
		})

		recognition := resourses.Group("/recognition")
		{
			recognition.POST("/recognize_face", a.recognize)
		}

		//Создаём переход на страничку с распознаванием
		resourses.GET("/training", func(c *gin.Context) {
			c.HTML(
				http.StatusOK,
				trainingFormFileName,
				gin.H{},
			)
		})

		training := resourses.Group("/training")
		{
			training.POST("/train", a.train)
		}

	}

	router.GET("/", middleware.TimeoutMiddleware(5*time.Second), func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			signInFileName,
			gin.H{},
		)
	})

	router.GET("/sign_up", middleware.TimeoutMiddleware(5*time.Second), func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			signUpFileName,
			gin.H{},
		)
	})

	auth := router.Group("/auth", middleware.TimeoutMiddleware(5*time.Second))
	{
		auth.POST("/signIn", a.signIn)
		auth.POST("/signUp", a.signUp)
	}

	return router
}
