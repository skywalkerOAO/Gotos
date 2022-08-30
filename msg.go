package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*Message
 * @Author Hex575A
 * @Description function Description
 * @Date 8:49 2022/8/30
 * @Param Msg string 消息体
 * @Param Code int 代号
 * @return
 */
func Message(Msg string, Code int) string {
	type Warn struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	warn := &Warn{
		Msg:  Msg,
		Code: Code,
	}
	res, _ := json.Marshal(&warn)
	return string(res)
}

/*DingText
 * @Author Hex575A
 * @Description Text模式的钉钉WebHook方法
 * @Date 9:03 2022/8/30
 * @Param userIds []string 用户id（gwid）
 * @Param phoneNum []string 手机号
 * @Param contentBody string 消息体
 * @Param atAll bool 是否@所有人
 * @Param keyWord string 安全关键字
 * @Param Hook string WebHook地址
 * @return
 */
func DingText(userIds []string, phoneNum []string, contentBody string, atALL bool, keyWord string, Hook string) {
	type At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}
	type Text struct {
		Content string `json:"content"`
	}
	type Context struct {
		At      *At    `json:"at"`
		Text    *Text  `json:"text"`
		MsgType string `json:"msgtype"`
	}
	keywords := keyWord
	s1 := userIds
	s2 := phoneNum
	s3 := keywords + contentBody
	s4 := "text"
	dingContext := Context{
		MsgType: s4,
	}

	DingAt := new(At)
	DingAt.AtUserIds = s1
	DingAt.AtMobiles = s2
	DingAt.IsAtAll = atALL
	dingContext.At = DingAt

	DingTexts := new(Text)
	DingTexts.Content = s3
	dingContext.Text = DingTexts

	jsonDing, err := json.Marshal(dingContext)
	if err != nil {
		fmt.Println(err.Error())
	}
	webHook := Hook

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(webHook,
		"application/json",
		bytes.NewReader(jsonDing))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(string(body))
	}

}

/*DingMarkDown
 * @Author Hex575A
 * @Description function Description
 * @Date 9:19 2022/8/30
 * @Param userIds []string 用户id（gwid）
 * @Param phoneNum []string 手机号
 * @Param title string 标题
 * @Param contentBody string 消息体
 * @Param atAll bool 是否@所有人
 * @Param keyWord string 安全关键字
 * @Param Hook string WebHook地址
 * @return
 */
func DingMarkDown(userIds []string, phoneNum []string, title string, contentBody string, atALL bool, keyWord string, Hook string) {
	type At struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}
	type Markdown struct {
		Text  string `json:"text"`
		Title string `json:"title"`
	}
	type Context struct {
		At       *At       `json:"at"`
		Markdown *Markdown `json:"markdown"`
		MsgType  string    `json:"msgtype"`
	}
	keywords := keyWord
	s1 := userIds
	s2 := phoneNum
	s3 := keywords + contentBody
	s4 := "markdown"
	s5 := title
	dingContext := Context{
		MsgType: s4,
	}

	DingAt := new(At)
	DingAt.AtUserIds = s1
	DingAt.AtMobiles = s2
	DingAt.IsAtAll = atALL
	dingContext.At = DingAt

	DingTexts := new(Markdown)
	DingTexts.Text = s3
	DingTexts.Title = s5
	dingContext.Markdown = DingTexts

	jsonDing, err := json.Marshal(dingContext)
	if err != nil {
		fmt.Println(err.Error())
	}
	webHook := Hook

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(webHook,
		"application/json",
		bytes.NewReader(jsonDing))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(string(body))
	}

}

func DingLink(text string, title string, picUrl string, messageUrl string, keyWord string, Hook string) {

	type Link struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	}
	type Context struct {
		Link    *Link  `json:"link"`
		MsgType string `json:"msgtype"`
	}
	keywords := keyWord
	s1 := picUrl
	s2 := title
	s3 := keywords + text
	s4 := "link"
	s5 := messageUrl
	dingContext := Context{
		MsgType: s4,
	}

	LinkText := new(Link)
	LinkText.Text = s3
	LinkText.Title = s2
	LinkText.PicUrl = s1
	LinkText.MessageUrl = s5
	dingContext.Link = LinkText

	jsonDing, err := json.Marshal(dingContext)
	if err != nil {
		fmt.Println(err.Error())
	}
	webHook := Hook

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(webHook,
		"application/json",
		bytes.NewReader(jsonDing))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(string(body))
	}

}
