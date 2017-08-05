package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

type WeiXinController struct {
	beego.Controller
}

func (c *WeiXinController) WeixinHandler() {
	config := &wechat.Config{
		AppID:          "wx365306ef17d56404",
		AppSecret:      "06e82680e896b2a0beaa50cf1a6068c3",
		Token:          "thisisxudeningweixintesttoken83722",
		EncodingAESKey: "",
	}
	wc := wechat.NewWechat(config)

	server := wc.GetServer(c.Ctx.Request, c.Ctx.ResponseWriter)

	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		msg.Content = "test"
		text := message.NewText(msg.Content)
		return &message.Reply{message.MsgTypeText, text}
	})

	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Send()

}
