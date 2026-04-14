package i18n

var currentLang = "en"

func SetLang(lang string) {
	currentLang = lang
}

func Lang() string {
	return currentLang
}

func T(key string) string {
	return key
}
