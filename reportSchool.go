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
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	return year + month + day
}

func main() {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
			time.Sleep(600 * time.Second)
			main()
		}
	}()
	urls := make(map[string]string, 0)
	urls["check"] = "https://app.bupt.edu.cn/uc/wap/login/check"
	urls["main"] = "https://app.bupt.edu.cn/ncov/wap/default/index"
	urls["save"] = "https://app.bupt.edu.cn/ncov/wap/default/save"

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic("Jar init error")
	}
	client := http.Client{Jar: jar}
	fmt.Println(len(os.Getenv("BUPT_USERNAME")), len(os.Getenv("BUPT_PASSWORD")))
	_, err = client.PostForm(urls["check"],
		url.Values{"username": {os.Getenv("BUPT_USERNAME")}, "password": {os.Getenv("BUPT_PASSWORD")}})
	if err != nil {
		fmt.Println(err)
		panic("check error")
	}
	resp, err := client.Get(urls["main"])
	if err != nil {
		fmt.Println(err)
		panic("get main error")
	}
	if resp == nil {
		panic("get empty resp")
	}
	fmt.Println("resp:      ", resp)
	fmt.Println("respbody:         ", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic("ioutil error")
	}

	//获取表单参数
	re := regexp.MustCompile(`"created":[0-9]+`)
	findCreated := re.FindAll(body, -1)
	created := strings.Split(string(findCreated[0]), ":")[1]
	re = regexp.MustCompile(`"id":[0-9]+`)
	findId := re.FindAll(body, -1)
	id := strings.Split(string(findId[0]), ":")[1]
	re = regexp.MustCompile(`"uid":"[0-9]+"`)
	findUid := re.FindAll(body, -1)
	uid := strings.Split(string(findUid[0]), ":")[1]
	date := getDate()
	fmt.Println(created, id, date, uid)
	_, err = client.PostForm(urls["save"],
		url.Values{
			"ismoved":      {"0"},
			"jhfjrq":       {},
			"jhfjjtgj":     {},
			"jhfjhbcc":     {},
			"sfxk":         {"0"},
			"xkqq":         {},
			"szgj":         {},
			"szcs":         {},
			"zgfxdq":       {"0"},
			"mjry":         {"0"},
			"csmjry":       {"0"},
			"ymjzxgqk": 	{"已接种完成"},
			"xwxgymjzqk": 	{"3"},
			"date":         {date},
			"tw":           {"2"},
			"sfcxtz":       {"0"},
			"sfjcbh":       {"0"},
			"sfcxzysx":     {"0"},
			"qksm":         {},
			"sfyyjc":       {"0"},
			"jcjgqr":       {"0"},
			"remark":       {},
			"address":      {"北京市海淀区北太平庄街道北京邮电大学北京邮电大学海淀校区"},
			"area":         {"北京市  海淀区"},
			"province":     {"北京市"},
			"city":         {"北京市"},
			"geo_api_info": {`"type":"complete","info":"SUCCESS","status":1,"$Da":"jsonp_812014_","position":{"Q":39.96306,"R":116.35577,"lng":116.35577,"lat":39.96306},"message":"Get ipLocation success.Get address success.","location_type":"ip","accuracy":null,"isConverted":true,"addressComponent":{"citycode":"010","adcode":"110108","businessAreas":[{"name":"北下关","id":"110108","location":{"Q":39.955976,"R":116.33873,"lng":116.33873,"lat":39.955976}},{"name":"小西天","id":"110108","location":{"Q":39.957147,"R":116.364058,"lng":116.364058,"lat":39.957147}},{"name":"西直门","id":"110102","location":{"Q":39.942856,"R":116.34666099999998,"lng":116.346661,"lat":39.942856}}],"neighborhoodType":"科教文化服务;学校;高等院校","neighborhood":"北京邮电大学","building":"","buildingType":"","street":"西土城路","streetNumber":"10号院","country":"中国","province":"北京市","city":"","district":"海淀区","township":"北太平庄街道"},"formattedAddress":"北京市海淀区北太平庄街道北京邮电大学北京邮电大学海淀校区","roads":[],"crosses":[],"pois":[]`},
			"created":      {created},
			"id":			{id},
			"sfzx":         {"1"},
			"sfjcwhry":     {"0"},
			"sfjcwbry": 	{"0"},
			"sfcyglq":      {"0"},
			"gllx":         {},
			"glksrq":       {},
			"jcbhlx":       {},
			"jcbhrq":       {},
			"bztcyy": 		{},
			"sftjwh":       {"0"},
			"sftjhb":       {"0"},
			"sfsfbh":       {"0"},
			"xjzd":         {},
			"jcwhryfs":     {},
			"jchbryfs":     {},
			"szsqsfybl":    {"0"},
			"sfygtjzzfj":   {"0"},
			"gtjzzfjsj":    {},
			"gtjzzchdfh":	{},
			"fjqszgjdq": 	{},
			"sfjzxgym":		{"1"},
			"sfjzdezxgym": 	{"1"},
			"jcjg": 		{},
		})
	if err != nil {
		fmt.Println("post form init error")
	}
}