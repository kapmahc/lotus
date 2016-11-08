package web

import (
	"log/syslog"
	"math/rand"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/viper"
)

// OpenLogger open syslog writer
func OpenLogger(tag string) (*syslog.Writer, error) {
	priority := syslog.LOG_DEBUG
	if IsProduction() {
		priority = syslog.LOG_INFO
	}
	return syslog.New(priority, tag)
}

//IsProduction is production mode?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}

//Shell exec shell command
func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}

//RandomStr randome string
func RandomStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}
