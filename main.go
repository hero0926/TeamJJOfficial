package main

import (
	"TeamJJOfficial/util"
	"html/template"

	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

const secretKey = "secret"

func main() {
	app := iris.Default()
	appSetting(app)

	// make jwtAuth
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// 메인
	// 세로 캐루셀 프로필
	// 회원가입 - 로그인
	// 메일문의
	app.Get("/", func(ctx iris.Context) {

		//ctx.ViewData("message", "Welcome")
		ctx.View("index.html")

	})

	needsAuth := app.AllowMethods(iris.MethodOptions)
	{

		needsAuth.Use(jwtHandler.Serve)

	}

	// 팀소개
	// 단순 팀소개(DB연동)
	// 팀이름 / 소개 / 연혁

	// 개인소개
	// 개인 프로필, 소개, 프로젝트(DB연동)
	// 회원가입 - 로그인시 CRUD

	// 팀프로젝트소개
	// 프로젝트 이름, 소개, 사진(DB연동)
	// 회원가입 - 로그인시 CRUD. 여기 등록한걸 개인소개에 선택할수 있게

	// 메일
	// 메일전송

	// + Additional?
	// vue등 프론트엔드 신기술 써보기
	// 메일서버 직접 만들기( xxx@teamjj.co? )
	// aws 에 ec2 s3으로 이용하기
	// 회원가입/로그인 jwt 이용

	app.Run(iris.Addr(":8080"))
}

// appSetting 함수는 프로젝트 관련된 전체적인 구조 셋팅을 한다.
func appSetting(app *iris.Application) {

	tmpl := iris.HTML("./web/views", ".html").Delims("[[", "]]")
	tmpl.AddFunc("CSS", func(tags []string) template.HTML {
		return util.TemplateFuncCSSAndJavascript(tags, "css")
	})
	tmpl.AddFunc("SCRIPT", func(tags []string) template.HTML {
		return util.TemplateFuncCSSAndJavascript(tags, "js")
	})

	app.RegisterView(tmpl)
	//requestpath, systempath
	app.StaticWeb("/static", "static")

}
