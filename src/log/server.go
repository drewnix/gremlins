package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

func RegisterHandlers(e *echo.Echo) {
	e.GET("/log", func(c echo.Context) error {
		msg, err := ioutil.ReadAll(c.Request().Body)
		if err != nil || len(msg) == 0 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid request"}
		}
		write(string(msg))
		return nil
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
