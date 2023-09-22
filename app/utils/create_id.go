package utils

import "github.com/lucsky/cuid"

func CreateId() string {
	return cuid.New()[0:10]
}
