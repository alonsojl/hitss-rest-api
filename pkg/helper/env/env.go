package env

import (
	"hitss/pkg/helper/logger"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Load() {
	setRootDir()
	err := godotenv.Load(os.Getenv("ROOT") + "/.env")
	if err != nil {
		logger.Write(err)
		panic(err)
	}
}

func setRootDir() {
	_, b, _, _ := runtime.Caller(0)
	os.Setenv("ROOT", filepath.Join(filepath.Dir(b), "../../.."))
}
