package i18n

var currentLang = "en"

func SetLang(lang string) {
	// 这里负责切换当前全局语言状态。
	currentLang = lang
}

func Lang() string {
	// 对外返回当前正在使用的语言代码。
	return currentLang
}

func T(key string) string {
	// 现在先直接回传 key，后面再接真正的翻译表。
	return key
}
