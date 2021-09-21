package apis

import (
	"os"
)



var LIBRARY_URL = os.Getenv("LIBRARY_URL") // 图书馆网址
var CAPTCHA_NEW_BREAKER_URL = os.Getenv("CAPTCHA_BREAKER_NEW_URL")
var ZF_URL = os.Getenv("ZF_URL")
var CARD_URL = os.Getenv("CARD_URL")
var ZF_Main_URL = os.Getenv("ZF_URL")
var ZF_BK_URL = os.Getenv("ZF_URL_BK")
var CANTEEN_URL = os.Getenv("CANTEEN_URL")
