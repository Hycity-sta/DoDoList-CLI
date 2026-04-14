package i18n

import (
	"embed"
	"encoding/json"
	"fmt"

	"dodolist/config"
)

//go:embed locales/*.json
var localeFS embed.FS

var (
	currentLang = config.LanguageEnglish
	locales     = map[string]map[string]string{}
)

func init() {
	loadLocale(config.LanguageEnglish)
	loadLocale(config.LanguageChinese)
}

func loadLocale(lang string) {
	data, err := localeFS.ReadFile("locales/" + lang + ".json")
	if err != nil {
		return
	}

	entries := make(map[string]string)
	if err := json.Unmarshal(data, &entries); err != nil {
		return
	}
	locales[lang] = entries
}

func SetLang(lang string) {
	// 这里负责切换当前全局语言状态。
	if config.IsSupportedLanguage(lang) {
		currentLang = lang
	}
}

func Lang() string {
	// 对外返回当前正在使用的语言代码。
	return currentLang
}

func T(key string, args ...any) string {
	// 优先查当前语言，查不到就回退到英文，最后才回传 key。
	template := lookup(currentLang, key)
	if template == "" {
		template = lookup(config.LanguageEnglish, key)
	}
	if template == "" {
		template = key
	}
	if len(args) == 0 {
		return template
	}
	return fmt.Sprintf(template, args...)
}

func lookup(lang, key string) string {
	if entries, ok := locales[lang]; ok {
		return entries[key]
	}
	return ""
}
