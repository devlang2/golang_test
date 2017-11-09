package events

import (
	"time"
)

type FiletransLog struct {
	Rdate            time.Time `json:"-"`
	RdateStr         string    `json:"rdate"`
	SensorCode       int       `json:"sensor_code"`
	IppoolSrcGcode   int       `json:"ippool_src_gcode"`
	IppoolSrcOcode   int       `json:"ippool_src_ocode"`
	IppoolDstGcode   int       `json:"ippool_dst_gcode"`
	IppoolDstOcode   int       `json:"ippool_dst_ocode"`
	SessionId        string    `json:"session_id"`
	Category1        int       `json:"category1"`
	Category2        int       `json:"category2"`
	Protocol         int       `json:"protocol"`
	SrcIp            string    `json:"src_ip"`
	SrcPort          int       `json:"src_port"`
	SrcCountry       string    `json:"src_country"`
	DstIp            string    `json:"dst_ip"`
	DstPort          int       `json:"dst_port"`
	DstCountry       string    `json:"dst_country"`
	Domain           string    `json:"domain"`
	Url              string    `json:"url"`
	TransType        int       `json:"trans_type"`
	Filename         string    `json:"filename"`
	Filesize         int       `json:"filesize"`
	Md5              string    `json:"md5"`
	MailSender       string    `json:"mail_sender"`
	MailRecipient    string    `json:"mail_recipient"`
	MailContentsType string    `json:"mail_contents_type"`
	MailContents     string    `json:"mail_contents"`
	Score            int       `json:"score"`
	Filetype         int       `json:"filetype"`
	Category         int       `json:"category"`
	PrivateType      int       `json:"private_type"`
	PrivateString    string    `json:"private_string"`
	LocalVaccine     string    `json:"local_vaccine"`
	MalwareName      string    `json:"malware_name"`
	Ext1             int       `json:"ext1"`
	Ext2             int       `json:"ext2"`
	Ext3             int       `json:"ext3"`
	Ext4             int       `json:"ext4"`
	Ext5             int       `json:"ext5"`
	Ext6             int       `json:"ext6"`
	Ext7             int       `json:"ext7"`
}

type FiletransLogs []FiletransLog

func (this FiletransLogs) Len() int {
	return len(this)
}
func (this FiletransLogs) Less(i, j int) bool {
	return this[i].Rdate.Before(this[j].Rdate)
}
func (this FiletransLogs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
