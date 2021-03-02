package setting

import (
	"os"
	"time"

	"github.com/unknwon/com"
)

var (
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = os.Getenv("RUN_MODE")
}

func LoadServer() {
	RunMode = os.Getenv("RUN_MODE")

	HTTPPort = com.StrTo(os.Getenv("PORT")).MustInt()
	ReadTimeout = time.Duration(com.StrTo(os.Getenv("READ_TIMEOUT")).MustInt()) * time.Second
	WriteTimeout = time.Duration(com.StrTo(os.Getenv("WRITE_TIMEOUT")).MustInt()) * time.Second
}

func LoadApp() {
	JwtSecret = os.Getenv("JWT_SECRET")
	PageSize = com.StrTo(os.Getenv("PAGE_SIZE")).MustInt()
}
