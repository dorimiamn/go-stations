package middleware

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/mileusna/useragent"
)

type Log struct{
	Timestamp time.Time
	Latency int64
	Path string
	OS string
}

func Logger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		accessedDate := time.Now()
		h.ServeHTTP(w, r)
		executionTime := time.Since(accessedDate).Milliseconds()
		useragent := useragent.Parse(r.UserAgent())
		log := Log{
			Timestamp: accessedDate,
			Latency: executionTime,
			Path: r.URL.Path,
			OS: useragent.OS,
		}
		jsonLog, _ := json.Marshal(log)
		fmt.Println(string(jsonLog))
	}
	return http.HandlerFunc(fn)
}