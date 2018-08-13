package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {

		ctx.Writef("TeamJJ")

	})

	// 메인
	// 세로 캐루셀 프로필
	// 회원가입 - 로그인
	// 메일문의

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

	app.Run(iris.Addr(":8080"))
}
