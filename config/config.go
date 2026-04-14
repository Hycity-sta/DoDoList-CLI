package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// 这里维护项目内置支持的语言常量。
	LanguageEnglish = "en"
	LanguageChinese = "zh"
)

type Config struct {
	Language string `json:"language"`
}

var current = Config{
	Language: LanguageEnglish,
}

func IsSupportedLanguage(lang string) bool {
	// 只允许切换到当前项目已知的语言集合。
	return lang == LanguageEnglish || lang == LanguageChinese
}

func Load() error {
	// 程序启动时先尝试读取配置文件。
	cfg, err := read()
	if err != nil {
		return err
	}
	current = cfg
	return nil
}

func Save() error {
	// 把当前配置写回磁盘，确保语言选择能跨进程保留。
	return write(current)
}

func SetLanguage(lang string) error {
	// 先校验语言是否合法，再修改并持久化。
	if !IsSupportedLanguage(lang) {
		return fmt.Errorf("invalid language: %s", lang)
	}
	current.Language = lang
	return Save()
}

func Language() string {
	// 对外返回当前配置里的语言。
	return current.Language
}

func ConfigPath() (string, error) {
	// 配置文件固定放在可执行文件同级目录。
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(exe), "config.json"), nil
}

func read() (Config, error) {
	path, err := ConfigPath()
	if err != nil {
		return current, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := write(current); err != nil {
				return current, err
			}
			return current, nil
		}
		return current, err
	}

	if len(data) == 0 {
		if err := write(current); err != nil {
			return current, err
		}
		return current, nil
	}

	cfg := current
	if err := json.Unmarshal(data, &cfg); err != nil {
		return current, err
	}
	if !IsSupportedLanguage(cfg.Language) {
		cfg.Language = LanguageEnglish
		if err := write(cfg); err != nil {
			return current, err
		}
	}
	return cfg, nil
}

func write(cfg Config) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}
