package main

import (
	"fmt"
	"log"
	"time"

	"github.com/corpix/uarand"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/firefox"
)

const (
	urlTwitterLogin   = "https://twitter.com/login"
	urlTwitterCompose = "https://twitter.com/compose/tweet"
)

func initWebDriver(browser string, args []string, seleniumAddr string) (selenium.WebDriver, error) {
	caps := selenium.Capabilities{"browserName": browser}

	switch browser {
	case "firefox":
		imagCaps := map[string]interface{}{
			"browser.startup.page": 2,
		}
		firefoxCaps := firefox.Capabilities{
			Prefs: imagCaps,
			Args: []string{
				"--no-sandbox",
				"--user-agent=" + uarand.GetRandom(),
			},
		}
		caps.AddFirefox(firefoxCaps)

	case "chrome":
		fallthrough

	default:
		chromeCaps := chrome.Capabilities{
			Args: []string{
				"--headless",
				"--no-sandbox",
				"--start-maximized",
				"--window-size=1920,1080",
				"--disable-crash-reporter",
				"--hide-scrollbars",
				"--disable-gpu",
				"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			},
		}
		caps.AddChrome(chromeCaps)
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://%s/wd/hub", seleniumAddr))
	if err != nil {
		return nil, err
	}
	return wd, nil
}

func (op *TwitterSeleniumOutputProvider) Login() error {
	// defer wd.Quit()
	op.wd.Get(urlTwitterLogin)
	op.wd.SetImplicitWaitTimeout(2 * time.Second)

	// write email
	emailInput, err := op.wd.FindElement(selenium.ByXPATH, "//input[@name=\"session[username_or_email]\"]")
	if err != nil {
		return err
	}
	if err := emailInput.SendKeys(op.username); err != nil {
		return err
	}

	// write password
	passwordInput, err := op.wd.FindElement(selenium.ByName, `session[password]`)
	if err != nil {
		return err
	}
	// if err := passwordInput.SendKeys("aado33ve79T!"); err != nil {
	if err := passwordInput.SendKeys(op.password); err != nil {
		return err
	}

	// submit the login form
	form, err := op.wd.FindElement(selenium.ByXPATH, "//div[@data-testid='LoginForm_Login_Button']")
	if err != nil {
		return err
	}
	if err := form.Click(); err != nil {
		return err
	}
	return nil
}

func (op *TwitterSeleniumOutputProvider) Tweet(text string) error {
	// go to the twitter compose page
	op.wd.Get(urlTwitterCompose)
	op.wd.SetImplicitWaitTimeout(2 * time.Second)

	// click on the draft editor
	tweetInput, err := op.wd.FindElement(selenium.ByClassName, "DraftEditor-root")
	if err != nil {
		return err
	}
	if err := tweetInput.Click(); err != nil {
		return err
	}

	tweetInputEdit, err := op.wd.FindElement(selenium.ByClassName, "DraftEditor-editorContainer")
	if err != nil {
		return err
	}
	if err := tweetInputEdit.Click(); err != nil {
		return err
	}

	op.wd.SetImplicitWaitTimeout(1 * time.Second)
	// write the text
	textTweet, err := op.wd.FindElement(selenium.ByCSSSelector, "br[data-text=\"true\"]")
	if err != nil {
		return err
	}
	if err := textTweet.SendKeys(text); err != nil {
		return err
	}

	// submit the tweet
	formTweet, err := op.wd.FindElement(selenium.ByCSSSelector, "div[data-testid = 'tweetButton'][role = 'button']")
	if err != nil {
		return err
	}
	if err := formTweet.Click(); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
