package util

import (
	"encoding/json"
	"hash/fnv"
	"html/template"
	"os"
	"strings"
	"time"
)

// RequireJSON ...
func RequireJSON(filepath string, refer interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&refer)
	if err != nil {
		return err
	}
	return nil
}

// TemplateFuncCSSAndJavascript ...
func TemplateFuncCSSAndJavascript(tags []string, tagType string) template.HTML {
	t := time.Now()
	ts := t.Format("20060102150405")
	var resultStr string
	for _, value := range tags {
		if tagType == "css" {
			resultStr += `<link rel="stylesheet" type="text/css" href="/static/css/` + value + `?t=` + ts + `" />`
		} else if tagType == "js" {
			resultStr += `<script type="text/javascript" src="/static/js/` + value + `?t=` + ts + `"></script>`
		}
	}
	return template.HTML(resultStr)
}

// GetAppIDByHost ...
func GetAppIDByHost(host string) string {
	host = strings.Split(host, ".")[0]
	appID := strings.Split(host, "-")[2]
	return appID
}

// ChkUserAuth ...
// 세션에 들어있는 로그인한 사용자의 권한들을 조회해서 authcd 가 있는지 체크한다.
// (사용자가 해당 권한을 가지고 있는지 체크한다.)
func ChkUserAuth(authcd string, userAuth []string) bool {
	var chkResult bool
	for _, v := range userAuth {
		if v == authcd {
			chkResult = true
			break
		}
	}
	return chkResult
}

// MakeHash32 ...
func MakeHash32(key string) uint32 {
	hash32 := fnv.New32a()
	hash32.Write([]byte(key))
	return hash32.Sum32()
}
