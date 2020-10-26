package templater

import (
	"os"
)

const basePrmission = 511

func mkdirIfNoExist(pth string) error {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return os.Mkdir(pth, basePrmission)
	}
	return nil
}

func (c *Config) inIgnore(pth string) bool {
	for _, e := range c.IgnorePatterns {
		if endWith(pth, e) {
			return true
		}
	}
	return false
}

func endWith(str, end string) bool {
	i := len(str) - len(end)
	if i < 0 || str[i:] != end {
		return false
	}
	return true
}
