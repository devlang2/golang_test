package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	//	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/golibs/orm"
	_ "github.com/go-sql-driver/mysql"
	"wins21.co.kr/soc/aptx/events"
)

const (
	YYYYMMDDHH24MISS = "2006-01-02 15:04:05"
	YYYYMMDD         = "2006-01-02"
)

var logFiletrans []events.FiletransLog

func init() {
	initDatabase("root:sniper123!@#@tcp(127.0.0.1:3306)/aptxm?charset=utf8&parseTime=true&loc=Asia%2FSeoul")
}

type Engine struct {
	interval int
	c        chan []events.FiletransLog
}

type Form struct {
	Evtgrp int
}

type Filter struct {
	ScoreMin       int
	ScoreMax       int
	TransType      map[int]bool
	SensorCode     int
	IppoolSrcGcode int
	Ext1           int
	Ext2           int
	Ext3           int // Yara
	Ext4           int
}

func NewEngine(interval int) *Engine {
	return &Engine{
		interval: interval,
		c:        make(chan []events.FiletransLog),
	}
}

func (this *Engine) Start() error {
	log.Println("Starting engine..")

	o := orm.NewOrm()
	query := `
        select
            t.rdate,
            sensor_code,
            ippool_src_gcode,
            ippool_src_ocode,
            ippool_dst_gcode,
            ippool_dst_ocode,
            concat('', session_id) session_id,
            category1,
            category2,
            protocol,
            inet_ntoa(src_ip) src_ip,
            src_port,
            lower(src_country) src_country,
            inet_ntoa(dst_ip) dst_ip,
            dst_port,
            lower(dst_country) dst_country,
            domain,
            url,
            trans_type,
            filename,
            t.filesize,
            t.md5,
            mail_sender,
            mail_recipient,
            mail_contents_type,
            mail_contents,
            t1.score,
            t1.filetype,
            t1.category,
            t1.private_type,
            t1.private_string,
            t1.local_vaccine,
            t1.malware_name,
            t1.ext1,
            t1.ext2
            ,t1.ext3,t1.ext4,t1.ext5,t1.ext6,t1.ext7
        from log_event_filetrans t left outer join pol_file_md5 t1 ON t1.md5 = t.md5
        where  (t.rdate >= ? and t.rdate <= ?)
    `

	go func() {
		for {
			t := time.Now()
			args := make([]string, 0)
			args = append(args, t.Format(YYYYMMDD)+" 00:00:00", t.Format(YYYYMMDD)+" 23:59:59")
			logs := events.FiletransLogs{}
			_, err := o.Raw(query, args).QueryRows(&logs)
			if err != nil {
				log.Println(err.Error())
			}
			sort.Sort(sort.Reverse(logs))
			logFiletrans = logs
			time.Sleep(time.Duration(this.interval) * time.Second)
		}
	}()

	return nil
}

func initDatabase(connStr string) {
	orm.RegisterDataBase("default", "mysql", connStr, 3, 3)
}

func GetFiletransLog(res http.ResponseWriter, req *http.Request) {

	evtgrp := strings.TrimSpace(req.URL.Query().Get("evtgrp"))
	scoreMin, _ := strconv.Atoi(strings.TrimSpace(req.URL.Query().Get("score_min")))
	scoreMax, _ := strconv.Atoi(strings.TrimSpace(req.URL.Query().Get("score_max")))
	selectedSensors, _ := strconv.Atoi(strings.TrimSpace(req.URL.Query().Get("selectedSensors")))
	ippool_src_gcode, _ := strconv.Atoi(strings.TrimSpace(req.URL.Query().Get("ippool_src_gcode")))
	ext1, _ := strconv.Atoi(req.URL.Query().Get("ext1"))
	ext2, _ := strconv.Atoi(req.URL.Query().Get("ext2"))
	ext3, _ := strconv.Atoi(req.URL.Query().Get("ext3"))
	ext4, _ := strconv.Atoi(req.URL.Query().Get("ext4"))
	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 5
	}

	filter := Filter{
		TransType: map[int]bool{
			1: true,
			2: true,
			3: true,
			4: true,
		},
		ScoreMin:       scoreMin,
		ScoreMax:       scoreMax,
		SensorCode:     selectedSensors,
		IppoolSrcGcode: ippool_src_gcode,
		Ext1:           ext1,
		Ext2:           ext2,
		Ext3:           ext3,
		Ext4:           ext4,
	}

	if filter.ScoreMax == 0 {
		filter.ScoreMax = 100
	}

	if evtgrp == "1" {
		delete(filter.TransType, 1)
		delete(filter.TransType, 2)
	} else if evtgrp == "2" {
		delete(filter.TransType, 3)
		delete(filter.TransType, 4)
	}
	//	spew.Dump(filter)

	logs := events.FiletransLogs{}
	//	for _, r := range logFiletrans {
	//		if _, ok := filter.TransType[r.TransType]; ok {
	//			if r.Score >= filter.ScoreMin && r.Score <= filter.ScoreMax {
	//				r.RdateStr = r.Rdate.Format(YYYYMMDDHH24MISS)
	//				if filter.SensorCode == 0 && filter.IppoolSrcGcode == 0 {
	//					logs = append(logs, r)
	//					if len(logs) >= limit {
	//						break
	//					}
	//				} else {
	//					if filter.SensorCode == r.SensorCode {
	//						logs = append(logs, r)
	//						if len(logs) >= limit {
	//							break
	//						}
	//						continue
	//					}
	//					if filter.IppoolSrcGcode == r.IppoolSrcGcode {
	//						logs = append(logs, r)
	//						if len(logs) >= limit {
	//							break
	//						}
	//						continue
	//					}
	//				}
	//			}
	//		}
	//	}

	for _, r := range logFiletrans {
		if _, ok := filter.TransType[r.TransType]; !ok {
			continue
		}
		if r.Score < filter.ScoreMin || r.Score > filter.ScoreMax {
			continue
		}
		if filter.SensorCode > 0 && filter.SensorCode != r.SensorCode {
			continue
		}
		if filter.IppoolSrcGcode > 0 && filter.IppoolSrcGcode != r.IppoolSrcGcode {
			continue
		}
		if filter.Ext1 > 0 && filter.Ext1 != r.Ext1 {
			continue
		}
		if filter.Ext2 > 0 && filter.Ext2 != r.Ext2 {
			continue
		}
		if filter.Ext3 > 0 && r.Ext3 < 1 {
			continue
		}
		if filter.Ext4 > 0 && filter.Ext4 != r.Ext4 {
			continue
		}

		r.RdateStr = r.Rdate.Format(YYYYMMDDHH24MISS)
		logs = append(logs, r)
		if len(logs) >= limit {
			break
		}
		continue
	}

	j, _ := json.Marshal(logs)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(j))
}
