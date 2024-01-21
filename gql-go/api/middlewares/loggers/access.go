package loggers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4x2Hach1/learning/gql-go/api/utils"
)

type User struct {
	Email string
	Name  string
	ID    string
}

type AccessLog struct {
	StartTime  time.Time
	R          *http.Request
	User       *User
	ReqBodyBuf *[]byte
	ResBody    *InterceptWriter
}

func (log *AccessLog) InfoAccessLog() {
	utils.LoggingInfo()("access-log",
		"user", *log.User,
		"http_method", log.R.Method,
		"path", log.R.URL,
		"remote_addr", log.R.RemoteAddr,
		"user_agent", log.R.UserAgent(),
		"request", string(*log.ReqBodyBuf),
		"status", log.ResBody.status,
		"response", log.ResBody.body.String(),
		"response_duration_second", fmt.Sprintf("%.3f", time.Since(log.StartTime).Seconds()),
	)
}
