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

func getDate() string {
	year := time.Now().AddDate(0, 0, -1).Format("2006")
	month := time.Now().AddDate(0, 0, -1).Format("01")
	day := time.Now().AddDate(0, 0, -1).Format("02")
	return year + month + day
}

func main() {
	urls := make(map[string]string, 0)
	urls["check"] = "https://app.bupt.edu.cn/uc/wap/login/check"
	urls["main"] = "https://app.bupt.edu.cn/ncov/wap/default/index"
	urls["save"] = "https://app.bupt.edu.cn/ncov/wap/default/save"

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Jar init error")
	}
	client := http.Client{Jar: jar}
	fmt.Println(len(os.Getenv("BUPT_USERNAME")), len(os.Getenv("BUPT_PASSWORD")))
	_, err = client.PostForm(urls["check"],
			url.Values{"username": {os.Getenv("BUPT_USERNAME")}, "password": {os.Getenv("BUPT_PASSWORD")}})
	if err != nil {
		fmt.Println("check error")
	}
	resp, err := client.Get(urls["main"])
	if err != nil {
		fmt.Println("get main error")
	}
	fmt.Println("resp:      "resp)
	fmt.Println("respbody:         "resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil error")
	}

	//获取表单参数
	re := regexp.MustCompile(`"created":[0-9]+`)
	findCreated := re.FindAll(body, -1)
	created := strings.Split(string(findCreated[0]), ":")[1]
	re = regexp.MustCompile(`"id":[0-9]+`)
	findId := re.FindAll(body, -1)
	id := strings.Split(string(findId[0]), ":")[1]
	re = regexp.MustCompile(`"uid":[0-9]+`)
	findUid := re.FindAll(body, -1)
	uid := strings.Split(string(findUid[0]), ":")[1]
	date := getDate()
	fmt.Println(created, id, date, uid)
	//_, err = client.PostForm(urls["save"],
	//	url.Values{
	//		"ismoved":      {"0"},
	//		"jhfjrq":       {},
	//		"jhfjjtgj":     {},
	//		"jhfjhbcc":     {},
	//		"sfxk":         {"0"},
	//		"xkqq":         {},
	//		"szgj":         {},
	//		"szcs":         {},
	//		"zgfxdq":       {"0"},
	//		"mjry":         {"0"},
	//		"csmjry":       {"0"},
	//		"ymjzxgqk": 	{"已接种完成"},
	//		"xwxgymjzqk": 	{"3"},
	//		"date":         {date},
	//		"tw":           {"2"},
	//		"sfcxtz":       {"0"},
	//		"sfjcbh":       {"0"},
	//		"sfcxzysx":     {"0"},
	//		"qksm":         {},
	//		"sfyyjc":       {"0"},
	//		"jcjgqr":       {"0"},
	//		"remark":       {},
	//		"address":      {"海南省海口市秀英区海秀街道长信路泰达比华利山庄"},
	//		"area":         {"海南省 海口市 秀英区"},
	//		"province":     {"海南省"},
	//		"city":         {"海口市"},
	//		"geo_api_info": {`"type":"complete","position":{"Q":20.020156521268,"R":110.26003553602499,"lng":110.260036,"lat":20.020157},"location_type":"html5","message":"Get geolocation success.Convert Success.Get address success.","accuracy":75,"isConverted":true,"status":1,"addressComponent":{"citycode":"0898","adcode":"460105","businessAreas":[],"neighborhoodType":"","neighborhood":"","building":"","buildingType":"","street":"滨海大道","streetNumber":"193号","country":"中国","province":"海南省","city":"海口市","district":"秀英区","towncode":"460105002000","township":"海秀街道"},"formattedAddress":"海南省海口市秀英区海秀街道长信路泰达比华利山庄","roads":[],"crosses":[],"pois":[],"info":"SUCCESS"`},
	//		"created":      {created},
	//		"id":			{id},
	//		"uid":			{uid},
	//		"created_uid":	{"0"},
	//		"sfsqhzjkk":	{"0"},
	//		"sqhzjkkys":	{},
	//		"fxyy":			{},
	//		"sfjcqz":		{},
	//		"jcqzrq":		{},
	//		"sfzx":         {"0"},
	//		"sfjcwhry":     {"0"},
	//		"sfjcwbry": 	{"0"},
	//		"sfcyglq":      {"0"},
	//		"gllx":         {},
	//		"glksrq":       {},
	//		"jcbhlx":       {},
	//		"jcbhrq":       {},
	//		"bztcyy": 		{"3"},
	//		"sftjwh":       {"0"},
	//		"sftjhb":       {"0"},
	//		"sfsfbh":       {"1"},
	//		"xjzd":         {"泰达比华利"},
	//		"jcwhryfs":     {},
	//		"jchbryfs":     {},
	//		"szsqsfybl":    {"0"},
	//		"sfygtjzzfj":   {"0"},
	//		"gtjzzfjsj":    {},
	//		"gtjzzchdfh":	{},
	//		"fjqszgjdq": 	{},
	//		"sfjzxgym":		{"1"},
	//		"sfjzdezxgym": 	{"1"},
	//		"jcjg": 		{},
	//	})
	//if err != nil {
	//	fmt.Println("post form init error")
	//}
}
