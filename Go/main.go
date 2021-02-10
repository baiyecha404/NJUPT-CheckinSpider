package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

var (
	Username string
	Password string
)

func main() {
	flag.PrintDefaults()
	flag.StringVar(&Username, "u", "","Your NJUPT Username")
	flag.StringVar(&Password, "p", "","Your NJUPT Password")
	flag.Parse()

	if Username == "" || Password == "" {
		panic("You need to provide Username and Password")
	}

	// disable headless for debugging
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		//chromedp.Flag("headless", false),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancel()


	task1, cancel := context.WithTimeout(ctx, 20 * time.Second)
	defer cancel()
	err := chromedp.Run(task1,
		chromedp.Navigate(`http://rzfw.njupt.edu.cn/cas/login`),
		chromedp.WaitVisible(`#username`, chromedp.ByID),
		chromedp.WaitVisible(`#ppassword`, chromedp.ByID),
		chromedp.SetValue(`#username`, Username, chromedp.ByID),
		chromedp.SetValue(`#ppassword`, Password, chromedp.ByID),
		chromedp.Click(`#dl`, chromedp.NodeVisible),
		chromedp.Sleep(1 * time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login Done")


	task2, cancel := context.WithTimeout(ctx, 20 * time.Second)
	defer cancel()
	err = chromedp.Run(task2,
		chromedp.Navigate(`http://datav.njupt.edu.cn/feiyan_api/h5/html/index/casLogin.html`),
		chromedp.Sleep(2 * time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Visiting Done")


	task3, cancel := context.WithTimeout(ctx, 20 * time.Second)
	defer cancel()

	var  res  string
	err = chromedp.Run(task3,
		chromedp.Navigate(`http://datav.njupt.edu.cn/feiyan_api/h5/html/daka/daka-multi.html#question-form`),
		chromedp.WaitVisible(`a.button`),
		//chromedp.Click(`a.button`,chromedp.BySearch), // for editing
		chromedp.Sleep(2 * time.Second),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(2) > div.card-content > div > div > li:nth-child(5) > label > input[type=radio]").click();a="1"`,&res),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(3) > div.card-content > div > div > li:nth-child(5) > label > input[type=radio]").click();a="1"`,&res),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(4) > div.card-content > div > div > li:nth-child(2) > label > input[type=radio]").click();a="1"`,&res),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(5) > div.card-content > div > div > input[type=text]").value="36.5";a="1"`,&res),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(6) > div.card-content > div > div > li:nth-child(2) > label > input[type=radio]").click();a="1"`,&res),
		chromedp.Evaluate(`document.querySelector("#question-form > ul > li:nth-child(7) > div.card-content > div > div > li:nth-child(2) > label > input[type=radio]").click();a="1"`,&res),
		chromedp.Evaluate(`confirm = function () { return true};a="1"`,&res),
		chromedp.Click(`a.button`,chromedp.BySearch),
		chromedp.Sleep(4 * time.Second),
		)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Submitting form Done")
}
