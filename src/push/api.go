package push

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	AppID  string `json:"appid"`
	User   string `json:"user"`
}

func PushText(msg, configFileName, tokenFileName string) (r *Result, err error) {
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

	fp, err := os.ReadFile(configFileName)
	check(err)
	var c Config
	var t AccessToken
	err = json.Unmarshal(fp, &c)
	check(err)
	_, err = t.getTokenFromFile(tokenFileName)
	if err != nil {
		_, errs := t.getNewToken(c.ID, c.Secret)
		check(errs)
	}
	err = t.writeTokenToFile(tokenFileName)
	check(err)

	r, err = t.pushText(msg, c.User, c.AppID)
	check(err)
	return r, nil
}
