// package main

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"net/mail"
// 	"net/smtp"
// )

// func main() {
// 	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// 	host := "smtp.mail.com"
// 	email := "2695009886@qq.com"
// 	password := "pndftuyzutbcdhff"
// 	toEmail := "2695009886@qq.com"

// 	from := mail.Address{"发送人", email}
// 	to := mail.Address{"接收人", toEmail}

// 	header := make(map[string]string)
// 	header["From"] = from.String()
// 	header["To"] = to.String()
// 	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("邮件标题2")))
// 	header["MIME-Version"] = "1.0"
// 	header["Content-Type"] = "text/html; charset=UTF-8"
// 	header["Content-Transfer-Encoding"] = "base64"

// 	body := "我是黄雪朋."

// 	message := ""
// 	for k, v := range header {
// 		message += fmt.Sprintf("%s: %s\r\n", k, v)
// 	}
// 	message += "\r\n" + b64.EncodeToString([]byte(body))

// 	auth := smtp.PlainAuth(
// 		"",
// 		email,
// 		password,
// 		host,
// 	)

// 	err := smtp.SendMail(
// 		host+":25",
// 		auth,
// 		email,
// 		[]string{to.Address},
// 		[]byte(message),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// package main

// import (
// 	"log"
// 	"net/smtp"

// 	"github.com/jordan-wright/email"
// )

// func main() {
// 	// 简单设置 log 参数
// 	log.SetFlags(log.Lshortfile | log.LstdFlags)

// 	em := email.NewEmail()
// 	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
// 	em.From = "xx <2695009886@qq.com>"

// 	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
// 	em.To = []string{"2695009886@qq.com", "3493665801@qq.com", "2081212291@qq.com", "320028161@qq.com", "1598934549@qq.com", "1649336090@qq.com", "447909290@qq.com"}

// 	// 设置主题
// 	em.Subject = "黄雪朋给你发邮件了"

// 	// 简单设置文件发送的内容，暂时设置成纯文本
// 	em.Text = []byte("hello world， 我是黄雪朋, 给你发邮件了！仅仅是测试,, 不要紧张")
// 	// em.HTML = ``
// 	//设置服务器相关的配置
// 	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "2695009886@qq.com", "pndftuyzutbcdhff", "smtp.qq.com"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("send successfully ... ")
// }

package main

import (
	"log"
	"net/smtp"
	"sync"
	"time"

	"github.com/jordan-wright/email"
)

func main() {

	// 简单设置l og 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 创建有5 个缓冲的通道，数据类型是  *email.Email
	ch := make(chan *email.Email, 5)
	// 连接池
	p, err := email.NewPool(
		"smtp.qq.com:25",
		3, // 数量设置成 3 个
		smtp.PlainAuth("", "2695009886@qq.com", "pndftuyzutbcdhff", "smtp.qq.com"),
	)

	if err != nil {
		log.Fatal("email.NewPool error : ", err)
	}

	// sync 包，控制同步
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			// 若 ch 无数据，则阻塞， 若 ch 关闭，则退出循环
			for e := range ch {
				// 超时时间 10 秒
				err := p.Send(e, 10*time.Second)
				if err != nil {
					log.Printf("p.Send error : %v , e = %v , i = %d\n", err, e, i)
				}
			}
		}()
	}

	for i := 0; i < 5; i++ {
		e := email.NewEmail()
		// 设置发送邮件的基本信息
		e.From = "xx <2695009886@qq.com>"
		// e.To = []string{"2695009886@qq.com", "3493665801@qq.com", "2081212291@qq.com", "320028161@qq.com", "1598934549@qq.com", "1649336090@qq.com", "447909290@qq.com"}
		e.To = []string{"1611941024@qq.com"}
		e.Subject = "嘿嘿"
		e.Text = []byte("你好, 我是黄雪朋!")
		ch <- e
	}

	// 关闭通道
	close(ch)
	// 等待子协程退出
	wg.Wait()
	log.Println("send successfully ... ")
}

//实现并发发邮件
