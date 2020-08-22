package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/conf"
	"github.com/meloalright/guora/internal/model"
	"github.com/meloalright/guora/internal/service/authorization"
)

var server *httptest.Server
var ss string

func init() {
	var u model.User

	initAll(conf.Config())

	// create http.Handler
	r := gin.Default()
	SetupApiRouter(r)

	// run server using httptest
	server = httptest.NewServer(r)
	// defer server.Close()

	u.ID = 1
	admin, _ := u.Get()
	ss, _ = authorization.Gen(admin)
}

// Web Security
func TestWebSecurity(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.POST("/api/web/security/sign").
		WithJSON(map[string]interface{}{"name": "newuser", "mail": "newuser", "password": "password"}).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.POST("/api/web/security/login").
		WithJSON(map[string]interface{}{"mail": "newuser", "password": "password"}).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.POST("/api/web/security/logout").
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Web Add
func TestWebAdd(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.POST("/api/web/question").
		WithJSON(map[string]interface{}{"title": "A Test Qustion Yeah?", "type": -1, "desc": "Just Test."}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.POST("/api/web/answer").
		WithJSON(map[string]interface{}{"Content": "", "QuestionID": 1, "AnswerProfileID": 2}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.POST("/api/web/comment").
		WithJSON(map[string]interface{}{"Content": "A Web Comment", "type": 1, "answerID": 1}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.POST("/api/web/reply").
		WithJSON(map[string]interface{}{"content": "A Test Qustion Yeah?", "type": 1, "commentID": 1, "replyToProfileID": 1}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

}

// Coment Web Supporters
func TestWebSupporters(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.POST("/api/web/answer/3/supporters").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/web/answer/3/supporters").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// User Rest
func TestUser(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/user/2").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/users").WithQuery("limit", "10").WithQuery("offset", "0").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/users/counts").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/user/2").
		WithJSON(map[string]interface{}{"mail": "test"}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/user/2").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Profile Rest
func TestProfile(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/profile/2").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/profiles").WithQuery("limit", "10").WithQuery("offset", "0").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/profiles/counts").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/profile/2").
		WithJSON(map[string]interface{}{"name": "test (小测试)"}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/profile/2").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Coment Question
func TestQuestion(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/question/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/questions").WithQuery("limit", "10").WithQuery("offset", "0").WithQuery("questionProfileID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/questions/counts").WithQuery("questionProfileID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/question/3").
		WithJSON(map[string]interface{}{"type": -1}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/question/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Coment Answer
func TestAnswer(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/answer/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/answers").WithQuery("limit", "10").WithQuery("offset", "0").WithQuery("questionID", "1").WithQuery("answerProfileID", "1").WithQuery("order", "supporters_counts").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/answers/counts").WithQuery("questionID", "1").WithQuery("answerProfileID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/answer/3").
		WithJSON(map[string]interface{}{"type": -1}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/answer/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Coment Rest
func TestComment(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/comment/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/comments").WithQuery("limit", "10").WithQuery("offset", "0").WithQuery("answerID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/comments/counts").WithQuery("answerID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/comment/3").
		WithJSON(map[string]interface{}{"content": "Test comment."}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/comment/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

// Reply Rest
func TestReply(t *testing.T) {

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	e.GET("/api/rest/reply/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/replies").WithQuery("limit", "10").WithQuery("offset", "0").WithQuery("commentID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.GET("/api/rest/replies/counts").WithQuery("commentID", "1").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.PUT("/api/rest/reply/3").
		WithJSON(map[string]interface{}{"content": "Test reply."}).
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)

	e.DELETE("/api/rest/reply/3").
		WithCookie("ss", ss).
		Expect().
		Status(http.StatusOK).JSON().Object().ContainsKey("status").ValueEqual("status", 200)
}

