package usecase

import (
	"log"
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
	resultStatus := ResultStatus{}

	if code == 200 {
		resultStatus.Code = 200
		resultStatus.DBSts = "CONNECTED"
	} else if code == 404 {
		resultStatus.Code = 404
		resultStatus.DBSts = "UNCONNECTED"
	} else {
		resultStatus.Code = 500
		resultStatus.DBSts = "UNCONNECTED"
	}
	// 日本時間へ変換
	jst, _ := time.LoadLocation("Asia/Tokyo")
	resultStatus.Time = time.Now().In(jst)
	resultStatus.Host, _ = os.Hostname()

	log.Println(resultStatus)
	return resultStatus
}
