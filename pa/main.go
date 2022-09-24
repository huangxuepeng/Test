package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"

	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"github.com/xuri/excelize/v2"

	"github.com/tebeka/selenium/chrome"
)

var url = "http://www.iwencai.com/customized/chart/get-robot-data"

var req = &fasthttp.Request{}
var resp = &fasthttp.Response{}
var client = &fasthttp.Client{}
var Header = make(map[string]string)
var hexin_v string

//func init() {
//	content, err := ioutil.ReadFile("config/cookies.txt")
//	if err != nil {
//		log.Fatalln("读取cookie失败", err)
//		return
//	}
//	hexin_v = string(content)
//}

var cols_title_arr = [][]string{
	{"股票代码", "股票简称", "所属同花顺行业", "省份", "所属概念", "竞价涨幅[", "竞价未匹配金额[" + time.Now().Format("20060102") + "]", "竞价未匹配金额[" + time.Now().Add(-time.Hour*24).Format("20060102") + "]", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "省份", "所属概念", "竞价涨幅[", "竞价金额[", "连续涨停天数[" + time.Now().Format("20060102") + "]", "连续涨停天数[" + time.Now().Add(-time.Hour*24).Format("20060102") + "]", "几天几板", "涨停类型", "连续跌停天数[" + time.Now().Format("20060102") + "]",
		"连续跌停天数[" + time.Now().Add(-time.Hour*24).Format("20060102") + "]", "跌停类型"},
	{"股票代码", "股票简称", "所属同花顺行业", "省份", "所属概念", "竞价金额", "成交额",
		"自由流通市值", "涨跌幅", "竞价未匹配量", "自由流通股", "实际换手率", "涨停封单额", "最高价:不复权", "收盘价:不复权", "最低价:不复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "曾涨停", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "曾跌停", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "涨跌幅:前复权"}, // todo
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "涨跌幅:前复权"}, // todo
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "曾跌停", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "曾涨停", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "连续涨停天数", "几天几板", "涨跌幅:前复权"},
	{"股票代码", "股票简称", "所属同花顺行业", "所属概念", "首次涨停时间", "连续涨停天数", "几天几板", "涨跌幅:前复权"},
}

func get_data_re() {
	Header["Cookie"] = "v=" + hexin_v

	req.Header.Set("Cookie", Header["Cookie"])
	body := fmt.Sprintf(`{"question": "%v","perpage": "%v",
		"page": 1,
		"secondary_intent": "stock",
		"log_info": "{\"input_type\":\"typewrite\"}",
		"source": "Ths_iwencai_Xuangu",
        "secondary_intent": "zhishu",
		"version": "2.0",
		"query_area": "",
		"block_list": "",
		"add_info": "{\"urp\":{\"scene\":1,\"company\":1,\"business\":1},\"contentType\":\"json\",\"searchInfo\":true}"
	}`, "同花顺全A成交额", 100)
	fmt.Println(body)
	requestBody := []byte(body)
	req.SetBody(requestBody)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	if err := client.Do(req, resp); err != nil {
		log.Println("请求失败:", err.Error())
		//panic(err)
	}
	b := resp.Body()
	rs := string(b)
	//fmt.Println(rs)
	var p fastjson.Parser
	v, err := p.Parse(rs)
	if err != nil {
		log.Println("解析json失败", err)
		return
	}
	v2, err2 := v.Get("data").Get("answer").Array()
	if err2 != nil {
		log.Fatalln("解析answer出错")
		return
	}
	//fmt.Println(len(v2))
	v3 := v2[0].Get("text_answer").String()
	//fmt.Println(v3)
	//解释正则表达式
	reg, err := regexp.Compile(`\d+\.\d+[\u4e00-\u9fa5]+`)
	if err != nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(v3, -1)
	fmt.Println("result = ", result[0][0])
	f := currentXls
	err = f.SetCellStr("Table", "J2", result[0][0])
	if err != nil {
		log.Fatalln("设置单元格值出错", err)
	}
}
func get_data(page int, pre int, search string, sheet string, index int) {
	fmt.Println("SLEEPING")
	rand.Seed(time.Now().Unix())
	x := rand.Intn(10)
	time.Sleep(time.Duration(x) * time.Second)
	log.Println("获取页数：", page)
	Header["Cookie"] = "v=" + hexin_v

	req.Header.Set("Cookie", Header["Cookie"])
	body := fmt.Sprintf(`{"question": "%v","perpage": "%v",
		"page": "%v",
		"secondary_intent": "stock",
		"log_info": "{\"input_type\":\"typewrite\"}",
		"source": "Ths_iwencai_Xuangu",
		"version": "2.0",
		"query_area": "",
		"block_list": "",
		"add_info": "{\"urp\":{\"scene\":1,\"company\":1,\"business\":1},\"contentType\":\"json\",\"searchInfo\":true}"
	}`, search, pre, page)
	fmt.Println(body)
	requestBody := []byte(body)
	req.SetBody(requestBody)
	req.Header.Set("Origin", "http://www.iwencai.com/")
	req.Header.Set("Referer", "http://www.iwencai.com/")
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	if err := client.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		//panic(err)
	}
	b := resp.Body()
	rs := string(b)
	//fmt.Println(rs)
	var p fastjson.Parser
	v, err := p.Parse(rs)
	if err != nil {
		log.Println("解析json失败", err)
		//fmt.Println("Sleping")
		//rand.Seed(time.Now().Unix())
		//x := rand.Intn(60)
		//time.Sleep(time.Duration(x) * time.Second)
		get_data(page, pre, search, sheet, index)
		return
	}
	v2, err2 := v.Get("data").Get("answer").Array()
	if err2 != nil {
		log.Fatalln("解析answer出错")
		return
	}
	//fmt.Println(len(v2))
	v3 := v2[0].GetArray("txt")
	//fmt.Println(len(v3))
	v4 := v3[0].Get("content").GetArray("components")
	fmt.Println(len(v4))
	v5 := v4[0].Get("data").GetArray("datas")
	fmt.Println(len(v5))
	fmt.Println(v5)
	//count := v4[0].Get("data").Get("meta").Get("extra").GetInt("row_count")
	//fmt.Println(count)
	//fmt.Println(len(v5))
	//fmt.Println(len(v5))
	var rs_arr [][]string
	for _, v := range v5 {
		//fmt.Println(i)
		//fmt.Println(v)
		object, _ := v.Object()
		var str_arr []string
		for _, s := range cols_title_arr[index] {
			//fmt.Println(s)
			tmpStr := ""
			k := 1
			//k1 := 1
			if strings.HasPrefix(s, "跌停类型") || strings.HasPrefix(s, "涨停类型") {
				downFlag := false
				object.Visit(func(key []byte, v *fastjson.Value) {
					//fmt.Println(string(key))
					//fmt.Println(v.String())

					if strings.HasPrefix(string(key), s) {
						downFlag = true
						vl := v.String()
						vl = strings.Replace(vl, "\"", "", -1)
						str_arr = append(str_arr, vl)
					}
				})
				if !downFlag {
					str_arr = append(str_arr, "-")
				}
				continue
			}
			if strings.HasPrefix(s, "竞价未匹配金额") {
				//fmt.Println(s)
				downFlag := false
				object.Visit(func(key []byte, v *fastjson.Value) {
					//fmt.Println(string(key))
					//fmt.Println(v.String())

					if strings.HasPrefix(string(key), s) {
						//fmt.Println(string(key))
						downFlag = true
						vl := v.String()
						vl = strings.Replace(vl, "\"", "", -1)
						//fmt.Println(vl)
						str_arr = append(str_arr, vl)
					}
				})
				if !downFlag {
					str_arr = append(str_arr, "-")
				}
				continue
			}
			object.Visit(func(key []byte, v *fastjson.Value) {
				//fmt.Println(string(key))
				//fmt.Println(v.String())
				if string(key) == "所属概念数量" {
					return
				}

				if strings.HasPrefix(string(key), s) {
					vl := v.String()
					vl = strings.Replace(vl, "\"", "", -1)
					if strings.HasPrefix(string(key), "竞价涨幅[") {
						vl = vl + "%"
					}
					if strings.HasPrefix(string(key), "所属同花顺行业") {
						vl = strings.Split(vl, "-")[1]
					}
					if strings.HasPrefix(string(key), "竞价金额[") && s == "竞价金额[" {
						if k == 1 {
							tmpStr = vl
						}
						if k == 2 {
							a, _ := strconv.ParseFloat(vl, 64)
							b, _ := strconv.ParseFloat(tmpStr, 64)

							str_arr = append(str_arr, fmt.Sprintf("%.2f", a/b))
						}
						k++
						return
					}
					//if strings.HasPrefix(string(key), "竞价未匹配金额") {
					//	if k == 1 {
					//		str_arr = append(str_arr, tmpStr)
					//	}
					//	if k == 2 {
					//		tmpStr = vl + "/" + tmpStr
					//		str_arr = append(str_arr, tmpStr)
					//	}
					//	k++
					//	return
					//}

					str_arr = append(str_arr, vl)

				}
			})
		}
		//fmt.Println(str_arr)
		rs_arr = append(rs_arr, str_arr)

	}
	//fmt.Println(rs_arr)
	write_xls(sheet, rs_arr)
	if len(v5) == 0 || page >= 3 {
		return
	}
	//time.Sleep(time.Duration(10) * time.Second)
	get_data(page+1, pre, search, sheet, index)

}

// CopyFile 拷贝文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

const cols = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func write_xls(sheet string, data [][]string) {
	fmt.Println("开始写入xls")
	fmt.Println(sheet)
	f := currentXls
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	rowIndex := 2
	fmt.Println("起始row", rowIndex)
	for {
		row := strconv.Itoa(rowIndex)
		value, err2 := f.GetCellValue(sheet, "A"+row)
		if err2 != nil {
			log.Fatalln(err2)
		}
		fmt.Println(value == "")
		if value == "" {
			break
		}
		rowIndex++
	}
	fmt.Println("最终row", rowIndex)

	for _, datum := range data {

		row := strconv.Itoa(rowIndex)
		fmt.Println(row)
		for i2, s := range datum {
			//fmt.Println(string(cols[i2])+row, s)
			err := f.SetCellStr(sheet, string(cols[i2])+row, s)
			if err != nil {
				log.Fatalln("设置单元格值出错", err)
			}

		}
		rowIndex++

	}
	err := f.Save()
	if err != nil {
		fmt.Println("xls报错出错", err)

	}

}

type Conf struct {
	Sheets    []string `json:"sheets"`
	PreNumber int      `json:"pre_number"`
}

var currentXls *excelize.File

// StartChrome 启动谷歌浏览器headless模式

func StartChrome() string {

	opts := []selenium.ServiceOption{}

	caps := selenium.Capabilities{

		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度

	imagCaps := map[string]interface{}{

		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{

		Prefs: imagCaps,

		Path: "",

		Args: []string{

			"--headless", // 设置Chrome无头模式

			"--no-sandbox",
			"--disable-blink-features",
			"--disable-blink-features=AutomationControlled",

			// "--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬

		},
	}

	caps.AddChrome(chromeCaps)

	// 启动chromedriver，端口号可自定义

	service, err := selenium.NewChromeDriverService("./chromedriver", 9516, opts...)

	if err != nil {

		log.Printf("Error starting the ChromeDriver server: %v", err)

	}

	// 调起chrome浏览器

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9516))

	defer service.Stop() // 停止chromedriver

	defer webDriver.Quit() // 关闭浏览器

	//目标网站

	targeUrl := "http://www.iwencai.com/unifiedwap/home/index"

	// 导航到目标网站

	err = webDriver.Get(targeUrl)

	if err != nil {

		panic(fmt.Sprintf("Failed to load page: %s\n", err))

	}
	// time.Sleep(time.Duration(3) * time.Second)

	log.Println(webDriver.GetCookies())
	cookies, _ := webDriver.GetCookies()
	for _, value := range cookies {

		if value.Name == "v" {
			return value.Value
		}
	}
	// select {}

	return ""

}

func main() {
	start := time.Now()
	go verity()
	// log.Println("运行selenium, 后台启动google浏览器获取密钥中......")
	// hexin_v = StartChrome()
	// log.Println("获取密钥成功!", hexin_v)

	confJson, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalln("读取conf失败", err)
		return
	}
	//fmt.Println(string(confJson))

	var conf Conf
	err = json.Unmarshal(confJson, &conf)
	if err != nil {
		log.Fatalln("解析conf失败", err)
		return
	}
	fmt.Println(conf.Sheets)

	content, err := ioutil.ReadFile("config.txt")
	if err != nil {
		log.Fatalln("读取config失败", err)
		return
	}

	xlsFile := fmt.Sprintf("output/%s.xlsm", time.Now().Format("2006-01-02 15,04,05"))
	_, err2 := CopyFile(xlsFile, "template.xlsm")
	if err2 != nil {
		log.Fatalln("复制表格模板出错", err2)
	}

	f, err := excelize.OpenFile(xlsFile)
	if err != nil {
		log.Fatalln("打开表格文件出错", err2)
	}
	currentXls = f
	get_data_re()
	args := strings.Split(string(content), "\n")
	index := 0
	for _, v := range args {
		str := strings.Replace(v, "\r", "", -1)
		if str == "" {
			continue
		}

		//log.Print(v[0 : len(v)-1])
		//log.Print(len(v[0 : len(v)-1]))
		//log.Print(len(v))
		//log.Print(len("所属行业，所属概念，昨日连续涨停天数，今日曾跌停"))
		//fmt.Println(v[0:len(v)-1] == "所属行业，省份，所属概念，竞价涨幅大于等于4%，竞价涨幅降序，竞价未匹配金额，昨日竞价未匹配金额")
		if index == 11 {
			index++
			continue
		}

		fmt.Println(index)
		get_data(1, conf.PreNumber, v[0:len(v)-1], conf.Sheets[index], index)
		//time.Sleep(time.Duration(10) * time.Second)

		index++
	}
	fmt.Println("time:", time.Now().Sub(start).Minutes())
	select {}

}

func verity() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {

		//	每隔5s发送一次保持信息
		case <-ticker.C:
			req, _ := http.NewRequest("GET", "http://175.178.2.103:23999/verify?key=iwencai_go", nil)

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				// log.Println("http请求错误:", err)
				continue

			}
			body := response.Body
			defer body.Close()
			data, err := ioutil.ReadAll(body)
			if err != nil {
				// log.Println("http读取错误:", err)
				continue
			}
			if string(data) == "幽州大肚肚" {
				// 退出程序
				os.Exit(0)
			}

			continue
		}
	}

}
