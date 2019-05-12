package configs

import (
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func TestSharedConfig(t *testing.T) {
	config := SharedConfig()
	logs.PrintAsJSON(config, t)
}
