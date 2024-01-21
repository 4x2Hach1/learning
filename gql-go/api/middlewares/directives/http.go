package directives

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/4x2Hach1/learning/gql-go/api/middlewares/loggers"
)

func AuthAndLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log出力の設定
		reqBodyBuf, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(reqBodyBuf))
		// responseの取得用設定
		iw := loggers.NewInterceptWriter(w)
		w = iw

		log := &loggers.AccessLog{
			StartTime:  time.Now(),
			R:          r,
			User:       &loggers.User{},
			ReqBodyBuf: &reqBodyBuf,
			ResBody:    iw,
		}

		authHeader := r.Header.Get("Authorization")
		// 認証headerがない場合
		if authHeader == "" {
			defer func() {
				log.InfoAccessLog()
			}()

			next.ServeHTTP(w, r)
			return
		}

		token, err := DecodeToken(authHeader)
		// 認証headerはあるが無効な値の場合
		if err != nil {
			defer func() {
				log.InfoAccessLog()
			}()

			http.Error(w, "token invalid", http.StatusUnauthorized)
			return
		} else {

			ctx := context.WithValue(r.Context(), JwtClaims{}, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
