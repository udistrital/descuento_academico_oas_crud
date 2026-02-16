package main

import (
	"net/url"

	"github.com/udistrital/utils_oas/database"
	"github.com/udistrital/utils_oas/security"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/beego/logs"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/descuento_academico_oas_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerrorv2"
	"github.com/udistrital/utils_oas/xray"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		url.QueryEscape(beego.AppConfig.String("PGpass"))+"@"+
		beego.AppConfig.String("PGurls")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")
}

func main() {
	conn, err := database.BuildPostgresConnectionString()
	if err != nil {
		logs.Error("error consultando la cadena de conexión: %v", err)
		return
	}

	err = orm.RegisterDataBase("default", "postgres", conn)
	if err != nil {
		logs.Error("error al conectarse a la base de datos: %v", err)
		return
	}
	////////////////////

	allowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == beego.DEV {
		allowedOrigins = []string{"*"}
		orm.Debug = true // Solo para APIs CRUD
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"}, // ajustar según los métodos usados en el api
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Amzn-Trace-Id"},
		ExposeHeaders:    []string{"Content-Length"}, // agregar otros headers según sea el caso
		AllowCredentials: true,
	}))

	err = xray.InitXRay()
	if err != nil {
		logs.Error("error configurando AWS XRay: %v", err)
	}
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	security.SetSecurityHeaders()
	beego.Run()
}
