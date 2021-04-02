package push

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"text/template"
)

type Result struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

func (t *AccessToken) pushText(msg string, usr string, agentid string) (resultJson *Result, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok {
			resultJson = nil
			err = panicErr
		} else {
			panic(errors.New("转换错误"))
		}
	}()
	URL := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + t.Token
	msgJson := `{
		"touser" : "{{.usr}}",
		"toparty":"",
		"totag":"",
		"msgtype":"text",
		"agentid":"{{.agentid}}",
		"text":{
			"content": "{{.msg}}"
		},
		"safe":"0"
		}`
	msgMap := map[string]string{"msg": msg, "usr": usr, "agentid": agentid}

	tp := template.Must(template.New("Json").Parse(msgJson))

	by := bytes.Buffer{}
	err = tp.Execute(&by, msgMap)
	check(err)

	resp, err := http.Post(URL, "application/json", &by)
	check(err)

	result, err := io.ReadAll(resp.Body)
	check(err)

	resultJson = new(Result)

	err = json.Unmarshal(result, resultJson)
	check(err)

	return resultJson, nil
}
