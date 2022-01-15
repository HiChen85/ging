package cmd

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const mainFile = "package yourPackageName\n\nimport (\n\t\"demo/routers\"\n\t\"github.com/gin-gonic/gin\"\n)\n\nfunc main() {\n\n    gin.SetMode(gin.DebugMode)\n    // Gin has already add two middleware to its default engine, which are Logger() and Recovery()\n    // engine.Use(Logger(), Recovery())\n    r := gin.Default()\n\n    // load templates dir. All pages are stored in this dir\n    r.LoadHTMLGlob(\"templates/*\")\n\n    // load static files\n    // set router.Static(\"/static\", \"/var/www\") if you want to deploy your app\n    r.Static(\"/static\", \"./static\")\n\n    routers.SetDefaultRouter(r)\n    routers.SetGroupedRouter(r)\n\n    // execute this web service on port 8000\n    r.Run(\"localhost:8000\")\n}"

//go:embed dog.jpg
var imageFile []byte

var templateContent = map[string]string{
	"models":    "package models\n\ntype User struct {\n\tUsername string `form:\"username\" json:\"user\" bson:\"user\"`\n\tPassword string `form:\"password\" json:\"password\" bson:\"password\"`\n}\n",
	"handlers":  "package handlers\n\nimport (\n\t\"demo/models\"\n\t\"github.com/gin-gonic/gin\"\n\t\"net/http\"\n\t\"strconv\"\n)\n\nfunc Index(c *gin.Context) {\n\tif c.Request.Method == \"POST\" {\n\t\tc.JSON(http.StatusOK, gin.H{\n\t\t\t\"Status\": 200,\n\t\t\t\"Method\": \"POST\",\n\t\t})\n\t}\n\tif c.Request.Method == \"GET\" {\n\t\tc.JSON(http.StatusOK, gin.H{\n\t\t\t\"Status\": 200,\n\t\t\t\"Method\": \"GET\",\n\t\t})\n\t}\n}\n\nfunc Login(c *gin.Context) {\n\tif c.Request.Method == \"GET\" {\n\t\tc.HTML(http.StatusOK, \"login.html\", nil)\n\t} else if c.Request.Method == \"POST\" {\n\t\tvar user models.User\n\t\tif err := c.Bind(&user); err != nil {\n\t\t\tc.JSON(http.StatusBadRequest, gin.H{\n\t\t\t\t\"message\": err.Error(),\n\t\t\t})\n\t\t\treturn\n\t\t}\n\t\tif user.Username == \"admin\" && user.Password == \"123456\" {\n\t\t\tc.JSON(http.StatusOK, gin.H{\n\t\t\t\t\"name\":   user.Username,\n\t\t\t\t\"status\": \"Successful\",\n\t\t\t})\n\t\t} else {\n\t\t\tc.JSON(http.StatusBadRequest, gin.H{\n\t\t\t\t\"error\": \"Username or password error\",\n\t\t\t})\n\t\t\treturn\n\t\t}\n\t}\n}\n\nfunc ReturnInfo(c *gin.Context) {\n\tif c.Request.Method == \"GET\" {\n\t\tname := c.Param(\"name\")\n\t\tage, _ := strconv.Atoi(c.Param(\"age\"))\n\t\tc.String(http.StatusOK, \"Hi %v you are %v years old, next year, you will be %v years old\", name, age, age+1)\n\t} else {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\n\t\t\t\"message\": \"method error\",\n\t\t})\n\t}\n}\n\nfunc URLQueryHandler(c *gin.Context) {\n\t// the URL Query can use c.Query to get the value\n\tif c.Request.Method == \"GET\" {\n\t\tname := c.Query(\"name\")\n\t\tage, _ := strconv.Atoi(c.Query(\"age\"))\n\t\tif name == \"\" || age == 0 {\n\t\t\tc.JSON(http.StatusBadRequest, gin.H{\n\t\t\t\t\"message\": \"Query parameter error\",\n\t\t\t})\n\t\t}\n\t\tc.String(http.StatusOK, \"Hi %v you are %v years old, last year you're %v\", name, age, age-1)\n\t} else {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\n\t\t\t\"message\": \"method error\",\n\t\t})\n\t}\n}\n",
	"routers":   "package routers\n\nimport (\n\t\"demo/handlers\"\n\t\"github.com/gin-gonic/gin\"\n)\n\nfunc SetDefaultRouter(r *gin.Engine) {\n\t// default Router\n\t// this is a basic router that when you access this project with localhost:xxxx, this router will execute\n\tr.GET(\"/\", handlers.Index)\n\tr.GET(\"/login\", handlers.Login)\n\tr.POST(\"/login\", handlers.Login)\n}\n\nfunc SetGroupedRouter(r *gin.Engine) {\n\t// router group\n\tv1 := r.Group(\"/v1\")\n\t{\n\t\t// imagine group v1 is for all GET requests\n\t\t// this api will handle the request with parameters in path\n\t\t\n\t\t// the URL should be like localhost:8000/v1/hello/Tom/19\n\t\tv1.GET(\"/hello/:name/:age\", handlers.ReturnInfo)\n\t\t// localhost:8000/v1/hello/user?name=Jerry&age=19\n\t\tv1.GET(\"/hello/user\", handlers.URLQueryHandler)\n\t}\n\t\n}\n",
	"templates": "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>LoginToGin</title>\n</head>\n<body>\n    <img class=\"dog\" src=\"../static/dog.jpg\" style=\"height: 300px; width: 500px\">\n    <form action=\"http://localhost:8000/login\" method=\"POST\">\n        Username:<input type=\"text\" name=\"username\"> <br/>\n        Password:<input type=\"password\" name=\"password\"> <br/>\n        <input type=\"submit\" value=\"Login\">\n    </form>\n</body>\n</html>",
	"static":    string(imageFile),
	"config":    "package config\n\nimport \"fmt\"\n\nconst (\n\tMYSQL_USERNAME = \"admin\"\n\tMYSQL_PASSWORD = \"xxxxx\"\n\tMYSQL_HOST     = \"localhost\"\n\tMYSQL_PORT     = \"3306\"\n\tMYSQL_DATABASE = \"employees\"\n\tMYSQL_TIMEOUT  = \"5s\"\n)\n\nvar MYSQL_DSN = fmt.Sprintf(\"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=%v\",\n\tMYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE, MYSQL_TIMEOUT)\n\nconst (\n\tMONGO_USERNAME   = \"admin\"\n\tMONGO_PASSWORD   = \"xxxxxxx\"\n\tMONGO_HOST       = \"localhost\"\n\tMONGO_PORT       = \"27017\"\n\tMONGO_DATABASE   = \"users\"\n\tMONGO_COLLECTION = \"user\"\n)\n\nvar MONGODB_URI = fmt.Sprintf(\"mongodb://%v:%v@%v:%v/test?authSource=%v\",\n\tMONGO_USERNAME, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT, MONGO_USERNAME)\n",
}

func InitProject(modName string) {
	goModCMD := exec.Command("go", "mod", "init", modName)
	output, err := goModCMD.CombinedOutput()
	//err := goModCMD.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	log.Println("go mod init successfully......")
}

func InstallDependencies() {
	// install gin dependencies
	getGinCMD := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin")
	output, err := getGinCMD.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	
	// install gorm dependencies, gorm is the default database operation library
	getGormCMD := exec.Command("go", "get", "-u", "gorm.io/gorm")
	output, err = getGormCMD.CombinedOutput()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	getGormCMD = exec.Command("go", "get", "-u", "gorm.io/driver/mysql")
	output, err = getGormCMD.CombinedOutput()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	
	// install mongodb driver for this template
	getMongoDBCMD := exec.Command("go", "get", "-u", "go.mongodb.org/mongo-driver/mongo")
	output, err = getMongoDBCMD.CombinedOutput()
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	log.Println("Dependencies installed successfully......")
}

func CreateDirs() {
	basicSettings := map[string]string{
		"modelsDir":      "models",
		"staticFilesDir": "static",
		"templatesDir":   "templates",
		"handlersDir":    "handlers",
		"routersDir":     "routers",
		"configDir":      "config",
	}
	
	for k := range basicSettings {
		_, err := os.ReadDir(basicSettings[k])
		if err != nil {
			log.Printf("No %v Dir, creating it now!\n", basicSettings[k])
			e := os.MkdirAll(basicSettings[k], os.ModePerm)
			if e != nil {
				log.Fatal(e)
			}
		} else {
			log.Println("Dir exists... do not need to create...")
		}
	}
}

func createMainFile(dirName string) {
	err := ioutil.WriteFile(dirName+"/main.go", []byte(mainFile), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("main.go file created......")
}

func createTemplate(dirName string) {
	var templatesFile = map[string]string{
		"models":    "user.go",
		"templates": "login.html",
		"handlers":  "handlers.go",
		"routers":   "routers.go",
		"static":    "dog.jpg",
		"config":    "db_config.go",
	}
	if name, ok := templatesFile[dirName]; ok {
		err := ioutil.WriteFile(dirName+"/"+name, []byte(templateContent[dirName]), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, "created ......")
	}
}
