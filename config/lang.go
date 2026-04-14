package config

const (
	// 这里维护项目内置支持的语言常量。
	LanguageEnglish = "en"
	LanguageChinese = "zh"
)

func IsSupportedLanguage(lang string) bool {
	// 只允许切换到当前项目已知的语言集合。
	return lang == LanguageEnglish || lang == LanguageChinese
}
