package kc

import (
	"context"
	"fmt"
	"github.com/123456/c_code"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"net/url"
	"regexp"
	"strings"

	"io/ioutil"
	"log"
	"os"
	"time"
)

type KcBody struct {
	Ctx context.Context
	//CtxTime10Second func() context.Context
	Headless bool
	Close    context.CancelFunc
	UA       string
	TempDir  string
	NoDel    bool
}

//关闭浏览器
func (c *KcBody) Off() {
	c.Close()
}

func (c *KcBody) SetCookieNavigate(_url string, cookie string) chromedp.Tasks {
	cookies := []string{}
	split := strings.Split(cookie, ";")
	for _, v := range split {
		_one_cookie := strings.Split(strings.TrimSpace(v), "=")
		if len(_one_cookie) == 2 {
			cookies = append(cookies, _one_cookie...)
		}
	}
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	parse, _ := url.Parse(_url)
	name := "." + c_code.DomainRootName(parse.Host)
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				success, err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain(name).
					//WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", cookies[i], cookies[i+1])
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(_url),
		// read the returned values
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			//cookies, err := network.GetAllCookies().Do(ctx)
			//if err != nil {
			//	return err
			//}
			//
			////for i, cookie := range cookies {
			////	log.Printf("chrome cookie %d: %+v", i, cookie)
			////}

			return nil
		}),
	}
}

//得到Cookie 字符串
func (c *KcBody) GetCookieStr() (cookie string) {

	chromedp.Run(c.Ctx, chromedp.Evaluate(`document.cookie`, &cookie))
	return
}

func (c *KcBody) CtxTimeOutCtx(duration time.Duration) context.Context {
	timeout, _ := context.WithTimeout(c.Ctx, duration)
	return timeout
}
func (c *KcBody) CtxTimeOut(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.Ctx, duration)
}

func (c *KcBody) GetIMG(sel, save_file string) (isok bool) {
	js_sel := sel
	if !strings.HasPrefix(sel, "document.") {
		js_sel = `document.querySelector('` + sel + `')`
	}

	runjs := `
	var img = ` + js_sel + `;
	img.crossOrigin = "Anonymous";
	var c = document.createElement('canvas');
	
	c.height = img.naturalHeight;
	c.width = img.naturalWidth;
	var ctx = c.getContext('2d');
	ctx.drawImage(img, 0, 0, c.width, c.height);
	c.toDataURL();`

	var buf string
	chromedp.Run(c.Ctx,
		chromedp.EvaluateAsDevTools(runjs,
			&buf,
		))
	if !strings.Contains(buf, "base64,") {
		time.Sleep(time.Millisecond * 150)
		chromedp.Run(c.Ctx,
			chromedp.EvaluateAsDevTools(runjs,
				&buf,
			))
	}
	if !strings.Contains(buf, "base64,") {
		time.Sleep(time.Millisecond * 150)
		chromedp.Run(c.Ctx,
			chromedp.EvaluateAsDevTools(runjs,
				&buf,
			))
	}

	//保存验证码
	ibase := buf
	re := `data:i.*?/.{3,7};base64,`
	if !regexp.MustCompile(re).MatchString(ibase) {
		log.Println("不是base格式")
		return
	}
	ibase = regexp.MustCompile(re).ReplaceAllString(ibase, "")
	//decodeString, err := base64.RawStdEncoding.DecodeString(ibase)
	decodeString, err := c_code.Bs64Decode(ibase)
	if err != nil {
		log.Println("格式转换失败")
		return
	}
	err = ioutil.WriteFile(save_file, []byte(decodeString), os.ModePerm)
	if err != nil {
		log.Println("写入文件失败")
		return
	}
	img, err := c_code.IsImg(save_file)
	if err != nil {
		return
	}
	if !img {
		return
	}
	//检测图片是否存在
	return true

}

//获得新标签得数据
func (c *KcBody) GetNewTarget(_ctx2 context.Context) (target_ctx context.Context, target_ctx_close context.CancelFunc) {

	//err := chromedp.Run(_ctx2, DelWebdriver())
	//fmt.Println(err)
	ch := chromedp.WaitNewTarget(_ctx2, func(info *target.Info) bool {
		return info.URL != ""
	})
	//return
	return chromedp.NewContext(_ctx2, chromedp.WithTargetID(<-ch))

}

//获取HTML 数据
func (c *KcBody) Html(_ctx2 context.Context) string {
	var _ht string
	chromedp.Run(_ctx2, chromedp.OuterHTML("html", &_ht))
	//	chromedp.Run(_ctx2, chromedp.Evaluate(`document.querySelector("html").outerHTML`, &_ht))
	return _ht
}

// 启动浏览器
func (c *KcBody) Up() {
	if c.UA == "" {
		c.UA = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36"
	}
	if c.TempDir == "" {
		c.TempDir = "c:\\click-temp\\"
	}
	if c.NoDel {
		dir, _ := c_code.IsDir(c.TempDir)
		if !dir {
			os.MkdirAll(c.TempDir, os.ModePerm)
		}
	} else {
		os.RemoveAll(c.TempDir)
		os.MkdirAll(c.TempDir, os.ModePerm)
	}
	dir := c.TempDir
	if c.NoDel {

	} else {
		var err error
		dir, err = ioutil.TempDir(c.TempDir, "chromedp-example")
		if err != nil {
			panic(err)
		}
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent(c.UA),
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", c.Headless),
		//chromedp.Flag("start-maximized", true),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("incognito", true),
		chromedp.Flag("window-size", "1380,900"),
		chromedp.UserDataDir(dir),
		chromedp.Flag("disable-infobars", true),
		//chromedp.Flag("disable-infobars", true),
		chromedp.Flag("excludeSwitches", `['enable-automation', 'load-extension']`),
	)
	chromelist := []string{
		`C:\Users\123\Desktop\Chrome-bin\chrome.exe`,
		`C:\Users\Administrator\Desktop\chrome\Chrome-bin\chrome.exe`,
		`C:\Users\Administrator\AppData\Local\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
	}
	for _, chrome_v := range chromelist {
		if c_code.IsFile(chrome_v) {
			opts = append(opts, chromedp.ExecPath(chrome_v))
			break
		}
	}

	var allocCtx context.Context
	var ctxxx context.Context
	allocCtx, c.Close = chromedp.NewExecAllocator(context.Background(), opts...)

	// also set up a custom logger
	ctxxx, c.Close = chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	//程序允许得总耗时
	c.Ctx, c.Close = context.WithTimeout(ctxxx, time.Hour*90000)
	//page.AddScriptToEvaluateOnNewDocument()

	chromedp.Run(c.Ctx, DelWebdriver())

}

//func login_captcha(ctx context.Context) {
//	chromedp.ListenTarget(ctx, func(ev interface{}) {
//		switch ev := ev.(type) {
//		case *network.EventResponseReceived:
//			resp := ev.Response
//			if len(resp.Headers) != 0 {
//				r_url, e := url.QueryUnescape(resp.URL)
//				if e != nil {
//					return
//				}
//				if _, ok := resp.Headers["Content-Type"]; ok {
//					fmt.Println(resp.Headers["Content-Type"])
//					//下载图片
//					p_url, e := url.Parse(r_url)
//					content_type := strings.ToLower(resp.Headers["Content-Type"].(string))
//					if strings.Contains(content_type, "image") {
//
//						if !strings.HasSuffix(p_url.Path, "vcode") {
//							return
//						}
//
//						fmt.Println(p_url, e, resp.Headers["Content-Type"])
//						go func() {
//							if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
//								body, err := network.GetResponseBody(ev.RequestID).Do(ctx)
//								if err != nil {
//
//									return err
//								}
//								err = ioutil.WriteFile(gui.CaptchaName, body, os.ModePerm)
//
//								//respBody <- string(body)
//								//return err
//								return err
//							})); err != nil {
//								return
//							}
//						}()
//					}
//					////监听AJAX请求登录请求
//					//if strings.Contains(content_type, "text/html") && strings.HasSuffix(p_url.Path, "login") {
//					//	go func() {
//					//		if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
//					//			body, err := network.GetResponseBody(ev.RequestID).Do(ctx)
//					//			if err != nil {
//					//
//					//				return err
//					//			}
//					//
//					//			logninok <- string(body)
//					//
//					//			return err
//					//		})); err != nil {
//					//			return
//					//		}
//					//	}()
//					//}
//				}
//			}
//
//		}
//	})
//}
