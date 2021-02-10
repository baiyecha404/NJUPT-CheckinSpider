const puppeteer = require('puppeteer')
const dotenv  = require('dotenv')
dotenv.config()


const timeout = (delay) => {
    return new Promise(resolve => setTimeout(resolve, delay));
}

const answer = [
  3,//不在校 off-campus
  3,//其他 other reasons
  0,//体温低于37.3度
  0,//正常 Normal
  0//绿色 Green
]

async function main() {
    //const browser = await puppeteer.launch({headless: false});
    const browser = await puppeteer.launch({});
    const page = await browser.newPage();
    page.on('dialog', async dialog => {
      await dialog.accept();
    });
    await page.goto('http://rzfw.njupt.edu.cn/cas/login');
    await page.type('#username', process.env.NJUPT_USERNAME)
    await page.type('#ppassword', process.env.NJUPT_PASSWORD)
    await page.click('#dl')

    await timeout(0.5 * 1000)
    await page.goto('http://datav.njupt.edu.cn/feiyan_api/h5/html/index/casLogin.html');
    
    await timeout(0.5 * 1000)
    await page.goto('http://datav.njupt.edu.cn/feiyan_api/h5/html/daka/daka-multi.html#question-form')
    

    //await page.click('.button')  for editing
    await timeout(1 * 1000)
  

    await page.$$eval('input[name="jrszxq"]', inputs => { inputs[3].click()})
    await page.$$eval('input[name="jrszxqs"]', inputs => { inputs[3].click()})
    await page.$$eval('input[name="dqtw"]', inputs => { inputs[0].click()})
    await page.$$eval('input[name="jtswtw"]', inputs => { inputs[0].click()})
    await page.$$eval('input[name="skmys"]', inputs => { inputs[0].click()})
    await timeout(1 * 1000)
    await page.click('.button')
    await timeout(3 * 1000)
    console.log('Done')
    await browser.close()
}

main()