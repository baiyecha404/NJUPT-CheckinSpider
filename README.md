# NJUPT-CheckinSpider

Auto health-checkin script based on headless chrome using puppeteer/chromedp
Just for self-practicing


## Node

* Nodejs 
* puppeteer

### usage

1. install with `yarn/npm install`

2. modify `.env` config

your creds for `http://rzfw.njupt.edu.cn/cas/login`
```
NJUPT_USERNAME="Q18010123"  
NJUPT_PASSWORD="xxxxx"
```

3. run script `node main.js`


## Go

* Go 
* chromedp

### install/build/release

`go build -o main.exe main.go` for windows or `go build -o main main.go` for linux


### usage

```cmd
go run main.go -u "YOUR_NJUPT_USERNAME" -p "YOUR_NJUPT_USERNAME$"
//main.exe -u YOUR_NJUPT_USERNAME -p YOUR_NJUPT_PASSWORD
//main xxx
```

## lastwords

puppeteer is ez and fun :) 
chromedp is piece of shit XD

