package config

const (
	LanguageEnglish = "en"
	LanguageChinese = "zh"
)

func IsSupportedLanguage(lang string) bool {
	return lang == LanguageEnglish || lang == LanguageChinese
}
