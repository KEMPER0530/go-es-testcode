package domain

import (
	"os"
	"time"
)

type ResultStatus struct {
	Code  int
	DBSts string
	Time  time.Time
	Host  string
}

func NewResultStatus(code int) ResultStatus {
	var resultStatus ResultStatus
	switch code {
	case 200:
		resultStatus = ResultStatus{Code: 200, DBSts: "CONNECTED"}
	case 404:
		resultStatus = ResultStatus{Code: 404, DBSts: "UNCONNECTED"}
	default:
		resultStatus = ResultStatus{Code: 500, DBSts: "UNCONNECTED"}
	}
	resultStatus.Time = time.Now().In(time.FixedZone("JST", 9*60*60))
	host, _ := os.Hostname()
	resultStatus.Host = host
	return resultStatus
}
