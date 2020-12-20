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
	_, err = client.PostForm(urls["check"],
		url.Values{"username": {os.Args[1]}, "password": {os.Args[2]}})
	if err != nil {
		fmt.Println("check error")
	}
	resp, err := client.Get(urls["main"])
	if err != nil {
		fmt.Println("get main error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil error")
	}

	// 获取表单参数
	re := regexp.MustCompile(`"created":[0-9]+`)
	findCreated := re.FindAll(body, -1)
	created := strings.Split(string(findCreated[0]), ":")[1]
	re = regexp.MustCompile(`"id":[0-9]+`)
	findId := re.FindAll(body, -1)
	id := strings.Split(string(findId[0]), ":")[1]
	date := getDate()
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
			"uid":          {"12462"},
			"date":         {date},
			"tw":           {"2"},
			"sfcxtz":       {"0"},
			"sfyyjc":       {"0"},
			"jcjgqr":       {"0"},
			"jcjg":         {},
			"sfjcbh":       {"0"},
			"sfcxzysx":     {"0"},
			"qksm":         {},
			"remark":       {},
			"address":      {"北京市海淀区北太平庄街道北京邮电大学北京邮电大学海淀校区"},
			"area":         {"北京市  海淀区"},
			"province":     {"北京市"},
			"city":         {"北京市"},
			"geo_api_info": {`"type":"complete","info":"SUCCESS","status":1,"$Da":"jsonp_812014_","position":{"Q":39.96306,"R":116.35577,"lng":116.35577,"lat":39.96306},"message":"Get ipLocation success.Get address success.","location_type":"ip","accuracy":null,"isConverted":true,"addressComponent":{"citycode":"010","adcode":"110108","businessAreas":[{"name":"北下关","id":"110108","location":{"Q":39.955976,"R":116.33873,"lng":116.33873,"lat":39.955976}},{"name":"小西天","id":"110108","location":{"Q":39.957147,"R":116.364058,"lng":116.364058,"lat":39.957147}},{"name":"西直门","id":"110102","location":{"Q":39.942856,"R":116.34666099999998,"lng":116.346661,"lat":39.942856}}],"neighborhoodType":"科教文化服务;学校;高等院校","neighborhood":"北京邮电大学","building":"","buildingType":"","street":"西土城路","streetNumber":"10号院","country":"中国","province":"北京市","city":"","district":"海淀区","township":"北太平庄街道"},"formattedAddress":"北京市海淀区北太平庄街道北京邮电大学北京邮电大学海淀校区","roads":[],"crosses":[],"pois":[]`},
			"created":      {created},
			"sfzx":         {"0"},
			"sfjcwhry":     {"0"},
			"sfcyglq":      {"0"},
			"gllx":         {},
			"glksrq":       {},
			"jcbhlx":       {},
			"jcbhrq":       {},
			"sftjwh":       {"0"},
			"sftjhb":       {"0"},
			"fxyy":         {},
			"bztcyy":       {},
			"fjsj":         {"20200820"},
			"sfjchbry":     {"0"},
			"sfjcqz":       {},
			"jcqzrq":       {},
			"jcwhryfs":     {},
			"jchbryfs":     {},
			"xjzd":         {},
			"sfsfbh":       {"0"},
			"jhfjsftjwh":   {"0"},
			"jhfjsftjhb":   {"0"},
			"szsqsfybl":    {"0"},
			"sfygtjzzfj":   {"0"},
			"gtjzzfjsj":    {},
			"sfsqhzjkk":    {"0"},
			"sqhzjkkys":    {},
			"id":           {id},
			"gwszdd":       {},
			"sfyqjzgc":     {},
			"jrsfqzys":     {},
			"jrsfqzfy":     {},
		})
	if err != nil {
		fmt.Println("post form init error")
	}
}
