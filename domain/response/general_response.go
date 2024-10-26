package response

import "time"

type Response struct {
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}
