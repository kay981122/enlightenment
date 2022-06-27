package common

import (
	"database/sql"
	"em-app/global"
	"em-app/model/bussiness"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type ExportCSVService struct {
}

var (
	pageSize int64 = 10000
	//ch     = make(chan bool, count)
)

func (exportCSVService *ExportCSVService) GenerateCSV(table string) (err error) {
	fmt.Println("开始处理：", table)
	// 启动延迟导出
	go delayExport(table)
	return
}
func delayExport(table string) {
	db := global.Db.Model(&bussiness.Domain{})
	rows, _ := db.Rows()
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Panic(err.Error())
		return
	}
	// values：一行的所有值，长度==列数
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// 创建文件
	var filePath = "d:/tmp/" + time.Now().Format("20060102150405") + "_" + table + ".csv"
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	if err != nil {
		log.Panic(err.Error())
		return
	}

	// 遍历按pageSize大小写入数据
	var count int64 = 0
	for count < total {
		totalValues := [][]string{}
		var targetCount = count + pageSize
		for rows.Next() {
			var s []string
			err = rows.Scan(scanArgs...) //把每行的内容添加到scanArgs，也添加到了values
			if err != nil {
				log.Panic(err.Error())
				return
			}

			for _, v := range values {
				s = append(s, string(v))
			}
			totalValues = append(totalValues, s)
			count++
			if count >= total || count >= targetCount {
				break
			}
		}
		if err = rows.Err(); err != nil {
			log.Panic(err.Error())
			return
		}
		writeToCSV(f, columns, totalValues, count)
	}
	fmt.Println("处理完毕：" + filePath)
}
func writeToCSV(f *os.File, columns []string, totalValues [][]string, count int64) {
	w := csv.NewWriter(f)
	for a, i := range totalValues {
		if a == 0 {
			w.Write(columns)
			w.Write(i)
		} else {
			w.Write(i)
		}
	}
	w.Flush()
	fmt.Printf("处理进度:第%v行\n", count)
}
