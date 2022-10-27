package main

import (
	"github.com/playwright-community/playwright-go"
	"log"
)

func her(err error, text string) { // стандартный обработчик ошибок. 1 строка кода вместо трех
	if err != nil {
		log.Fatalf(text+" ", err)
	}
}
func main() {
	// var timeoutscr float64 = 10000
	headless := false // поставь true для отключения видимости браузера
  browser_options := playwright.BrowserTypeLaunchOptions{Headless: &headless} // браузер будет запускаться в видимом режиме
	var anim playwright.ScreenshotAnimations = "disabled" // отключим анимацию в элементе
	var waituntstat playwright.WaitUntilState = "commit"  // не будем ждать загрузки страницы
	path := "/Screen/fear&greed.png"  // путь для сохранения скриншота	
	pw, err := playwright.Run()                                                 // поехали
	her(err, "could not start playwright:")
	browser, err := pw.Chromium.Launch(browser_options) // выбираем Chromium
	her(err, "could not launch browser:")
	page, err := browser.NewPage() // создаем вкладку
	her(err, "could not create page:")
	_, err = page.Goto("https://edition.cnn.com/markets/fear-and-greed", playwright.PageGotoOptions{WaitUntil: &waituntstat}) // открываем страницу
	her(err, "could not goto:")
	// //div[@class='market-fng-gauge market-fng-gauge']
	_, err = page.WaitForSelector(`xpath=id("fear-and-greed-dial")`) // ожидаем подгрузки нужной части страницы
	her(err, "Could not load d page:")
	// locator, _ := page.Locator(`xpath=id("fear-and-greed-dial")`)
	locator, _ := page.Locator(`xpath=//div[@class="market-fng-gauge market-fng-gauge"]`)   // выбираем локатор
	locator.Screenshot(playwright.LocatorScreenshotOptions{Path: &path, Animations: &anim}) // делаем и сохраняем скриншот
	her(err, "could not create screenshot:")
	err = browser.Close()
	her(err, "could not close browser:")
	err = pw.Stop()
	her(err, "could not stop Playwright:")

}
