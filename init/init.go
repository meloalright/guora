package main

import (
	"context"
	"fmt"

	"github.com/meloalright/guora/configuration"
	"github.com/meloalright/guora/database"
	"github.com/meloalright/guora/model"
)

var ctx = context.Background()

func main() {

	database.RDB.FlushDB(ctx)
	fmt.Println("redis flushdb.")

	if (database.SQLITE3DB.HasTable(&model.User{})) {
		fmt.Println("db has the table user, so drop it.")
		database.SQLITE3DB.DropTable(&model.User{})
	}

	if (database.SQLITE3DB.HasTable(&model.Profile{})) {
		fmt.Println("db has the table profile, so drop it.")
		database.SQLITE3DB.DropTable(&model.Profile{})
	}

	if (database.SQLITE3DB.HasTable(&model.Question{})) {
		fmt.Println("db has the table question, so drop it.")
		database.SQLITE3DB.DropTable(&model.Question{})
	}

	if (database.SQLITE3DB.HasTable(&model.Answer{})) {
		fmt.Println("db has the table answer, so drop it.")
		database.SQLITE3DB.DropTable(&model.Answer{})
	}

	if (database.SQLITE3DB.HasTable(&model.Comment{})) {
		fmt.Println("db has the table comment, so drop it.")
		database.SQLITE3DB.DropTable(&model.Comment{})
	}

	if (database.SQLITE3DB.HasTable(&model.Reply{})) {
		fmt.Println("db has the table reply, so drop it.")
		database.SQLITE3DB.DropTable(&model.Reply{})
	}

	if (database.SQLITE3DB.HasTable(&model.Supporter{})) {
		fmt.Println("db has the table supporter, so drop it.")
		database.SQLITE3DB.DropTable(&model.Supporter{})
	}

	database.SQLITE3DB.AutoMigrate(&model.User{})
	database.SQLITE3DB.AutoMigrate(&model.Profile{})
	database.SQLITE3DB.AutoMigrate(&model.Question{})
	database.SQLITE3DB.AutoMigrate(&model.Answer{})
	database.SQLITE3DB.AutoMigrate(&model.Comment{})
	database.SQLITE3DB.AutoMigrate(&model.Reply{})
	database.SQLITE3DB.AutoMigrate(&model.Supporter{})

	p0 := model.Profile{Name: configuration.C.Admin.Name, Desc: "This is " + configuration.C.Admin.Name}
	p0.Create()

	database.SQLITE3DB.Create(&model.User{Mail: configuration.C.Admin.Mail, Password: configuration.C.Admin.Password, Authorized: 1, Type: 100, ProfileID: p0.ID})

	q0 := model.Question{Title: "How to use the Guora?", Desc: `{"blocks":[{"key":"9s4bh","text":"How should we use it?","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}],"entityMap":{}}`, Type: 0, QuestionProfile: p0}
	q1 := model.Question{Title: "这玩意咋用？", Desc: `{"blocks":[{"key":"bvelr","text":"这个 Guora 究竟如何使用？","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}}],"entityMap":{}}`, Type: 0, QuestionProfile: p0}
	q0.Create()
	q1.Create()

	a0 := model.Answer{Content: `{"blocks":[{"key":"798b1","text":"Hello everyone, this is an example of answer.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"72m1a","text":"First of all, Guora is a self hosted web application based on community Q & A.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"9f5rc","text":"When you host the application, you can enjoy the Q & A platform and connect with your working partners. You can also archive your documents here.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"7qk2d","text":"If you're the administrator, you can go to the Admin Page to manage the application.","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":47,"length":10,"style":"ITALIC"}],"entityRanges":[{"offset":47,"length":10,"key":0}],"data":{}},{"key":"64oef","text":"tips: If new users wanna join, only admin is able to create new user in Admin Page right now.","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":72,"length":10,"style":"ITALIC"}],"entityRanges":[{"offset":72,"length":10,"key":0}],"data":{}}],"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"href":"/admin","title":"/admin","url":"/admin"}}}}`, Type: 0, QuestionID: q0.ID, AnswerProfileID: p0.ID}
	a1 := model.Answer{Content: `{"blocks":[{"key":"elb3v","text":"大家好，这是一个回答示例，可以解释我们的问题。","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"5nmjl","text":"首先 Guora 是一个私有部署的社区问答系统。","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"dqjdg","text":"在部署后你可以尽情使用该问答平台，并且和你的工作伙伴社交互动。你也可以在这里归档你的工作文档/学习文档。","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"alaje","text":"如过你是管理员，你可以进入 管理页面 对系统进行管理。","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":14,"length":5,"style":"BOLD"}],"entityRanges":[{"offset":14,"length":4,"key":0}],"data":{}},{"key":"cr81h","text":"ps: 如果想要添加新用户，目前只能通过管理员在  管理页面 手动添加。","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":36,"style":"ITALIC"},{"offset":26,"length":10,"style":"BOLD"}],"entityRanges":[{"offset":26,"length":4,"key":1}],"data":{}}],"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"href":"/admin","title":"/admin","url":"/admin"}},"1":{"type":"LINK","mutability":"MUTABLE","data":{"href":"/admin","title":"/admin","url":"/admin"}}}}`, Type: 0, QuestionID: q1.ID, AnswerProfileID: p0.ID}
	a0.Create()
	a1.Create()

	c0 := model.Comment{Content: "Comment here", Type: 0, AnswerID: a0.ID, CommentProfileID: p0.ID}
	c1 := model.Comment{Content: "在这里留言", Type: 0, AnswerID: a1.ID, CommentProfileID: p0.ID}
	c0.Create()
	c1.Create()

	r0 := model.Reply{Content: "Reply here", Type: 0, CommentID: c0.ID, ReplyFromProfile: p0, ReplyToProfile: p0}
	r1 := model.Reply{Content: "在这里回复", Type: 0, CommentID: c1.ID, ReplyFromProfile: p0, ReplyToProfile: p0}
	r0.Create()
	r1.Create()

	fmt.Println("already restarted.")
}
