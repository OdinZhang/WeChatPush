package push

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

//AccessToken access_token返回数据
type AccessToken struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	Time      int64  `json:"time"`
}

//判断出错
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//GetNewToken 获取新的access_token
func (t *AccessToken) getNewToken(id, secret string) (str string, err error) {
	//从panic中恢复以返回错误
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok {
			str = ""
			err = panicErr
		} else {
			panic(errors.New("转换错误"))
		}
	}()

	var URL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + id + "&corpsecret=" + secret

	resp, err := http.Get(URL)
	check(err)

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	check(err)

	err = json.Unmarshal(body, &t)
	check(err)

	if t.Errcode != 0 {
		panic(errors.New("获取access_token出错"))
	}

	t.Time = time.Now().Unix()

	return t.Token, nil
}

func (t *AccessToken) getTokenFromFile(filePath string) (token string, err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok {
			token = ""
			err = panicErr
		} else {
			panic(errors.New("转换错误"))
		}
	}()

	fp, err := os.ReadFile(filePath)
	check(err)

	err = json.Unmarshal(fp, t)
	check(err)

	tm := time.Now().Unix()

	if tm-t.Time > t.ExpiresIn {
		return "", errors.New("超时")
	}

	return t.Token, nil
}

func (t *AccessToken) writeTokenToFile(filepath string) (err error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if panicErr, ok := e.(error); ok {
			err = panicErr
		} else {
			panic(errors.New("转换出错"))
		}
	}()

	js, err := json.Marshal(t)
	check(err)

	err = os.WriteFile(filepath, js, 0666)
	check(err)
	return nil
}
