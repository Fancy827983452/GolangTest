package model

import "time"

type visitRecord struct {
	id            uint32    //ID
	applicantKey  string    //申请人
	respondantKey string    //被申请人
	Purpose       string    //申请目的
	State         string    //申请状态
	Time          time.Time //申请时间
}
