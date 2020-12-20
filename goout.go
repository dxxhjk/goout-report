package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

func getDate(yesterday bool) string {
	year := time.Now()
	month := time.Now()
	day := time.Now()
	if yesterday {
		year = year.AddDate(0, 0, -1)
		month = month.AddDate(0, 0, -1)
		day = day.AddDate(0, 0, -1)
	}
	year1 := year.Format("2006")
	month1 := month.Format("01")
	day1 := day.Format("02")
	return year1 + "-" + month1 + "-" + day1
}

func main() {
	urls := make(map[string]string, 0)
	urls["login"] = "https://auth.bupt.edu.cn/authserver/login?service=https%3A%2F%2Fservice.bupt.edu.cn%2Fsite%2Flogin%2Fcas-login%3Fredirect_url%3Dhttps%253A%252F%252Fservice.bupt.edu.cn%252Fv2%252Fmatter%252Fm_start%253Fid%253D578"
	urls["launch"] = "https://service.bupt.edu.cn/site/apps/launch"

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("jar init error")
	}
	client := http.Client{Jar: jar}

	// 获取lt
	resp, err := client.Get(urls["login"])
	body, _ := ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile(`name="lt" value="[a-zA-Z0-9-]+"`)
	findLT := re.FindAll(body, -1)
	lt := strings.Split(string(findLT[0]), "\"")[3]

	resp, err = client.PostForm(urls["login"], url.Values{
		"username":  {os.Args[1]},
		"password":  {os.Args[2]},
		"execution": {"e1s1"},
		"lt":        {lt},
		"_eventId":  {"submit"},
		"rmShown":   {"1"},
	})
	if err != nil {
		fmt.Println("login error")
	}
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	u, err := url.Parse(urls["launch"])
	cookies := ""
	for _, cookie := range client.Jar.Cookies(u) {
		cookies += cookie.String() + "; "
	}

	formData := `{"app_id":"xxx","form_data":{"1716":{"User_5":"<姓名>","User_7":"<学号>","User_9":"计算机学院（国家示范性软件学院）","User_11":"<手机号>","SelectV2_58":[{"name":"西土城校区","value":"2","default":0,"imgdata":""}],"UserSearch_60":{"uid":197752,"name":"<辅导员姓名>"},"Calendar_62":"` + getDate(false) + `T00:00:00+08:00","Calendar_50":"` + getDate(true) + `T16:00:01.000Z","Calendar_47":"` + getDate(false) + `T15:59:59.000Z","Input_28":"南门","MultiInput_30":"吃饭","Radio_52":{"value":"1","name":"本人已阅读并承诺"},"Validate_63":"","Alert_65":"","Validate_66":"","Alert_67":"","UserSearch_73":{"uid":197752,"name":"<辅导员姓名>","number":"<辅导员工号>"}}}}`
	req, err := http.NewRequest("POST", urls["launch"], strings.NewReader(url.Values{"data": {formData}}.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(req.Cookies())
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("get form init error")
	}
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
