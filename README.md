# NJUPT-CheckinSpider

Auto health-checkin script based on headless chrome using puppeteer/chrome
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

<<<<<<< HEAD
puppeteer is ez and fun :) 
chromedp is piece of shit XD
=======

puppeteer is ez and fun :) 
chromedp is piece of shit XD
>>>>>>> b006cca5f1ba0cdb106514c8598ab8a43629c3dd
