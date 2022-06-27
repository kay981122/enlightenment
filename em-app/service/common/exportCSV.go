package common

import (
	"database/sql"
	"em-app/global"
	"em-app/model/bussiness"
	"em-app/model/system"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type ExportCSVService struct {
}

var (
	pageSize int64 = 5000
)

func (exportCSVService *ExportCSVService) GenerateCSV(table string) (err error) {
	fmt.Println("开始处理：", table)
	// 启动延迟导出
	go delayExport(table)
	return
}
func delayExport(table string) {
	db := global.Db.Model(&bussiness.Domain{})
	db2 := global.Db
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
	var filePath = "d:/tmp"
	var fileName = time.Now().Format("20060102150405") + "_" + table + ".csv"
	fileName = filepath.Join(filePath, fileName)
	err = os.MkdirAll(filePath, os.ModeDir)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	// 记录初始化文件
	mySnow, _ := global.NewSnowFlake(0, 0)
	id, _ := mySnow.NextId()
	date := global.GetCurrentTime(global.TimeTemplates[0])
	exportCSVProgress := system.ExportCSVProgress{
		Id:         id,
		UserId:     "00000",
		Module:     table,
		Status:     "0",
		CreateTime: date,
		UpdateTime: date,
	}
	db2.Create(exportCSVProgress)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	if err != nil {
		log.Panic(err.Error())
		return
	}
	// 记录正在生成文件
	exportCSVProgress.Status = "1"
	exportCSVProgress.UpdateTime = global.GetCurrentTime(global.TimeTemplates[0])
	exportCSVProgress.FilePath = fileName
	db2.Save(&exportCSVProgress)
	// 先写好第一行数据--列名称
	w := csv.NewWriter(f)
	w.Write(columns)
	w.Flush()
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
				exportCSVProgress.Status = "3"
				exportCSVProgress.UpdateTime = global.GetCurrentTime(global.TimeTemplates[0])
				db2.Save(&exportCSVProgress)
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
			exportCSVProgress.Status = "3"
			exportCSVProgress.UpdateTime = global.GetCurrentTime(global.TimeTemplates[0])
			db2.Save(&exportCSVProgress)
			return
		}
		writeToCSV(f, totalValues, count)
	}
	fmt.Println("处理完毕：" + fileName)
	// 记录生成完成
	exportCSVProgress.Status = "2"
	exportCSVProgress.UpdateTime = global.GetCurrentTime(global.TimeTemplates[0])
	db2.Save(&exportCSVProgress)
}
func writeToCSV(f *os.File, totalValues [][]string, count int64) {
	w := csv.NewWriter(f)
	for _, i := range totalValues {
		w.Write(i)
	}
	w.Flush()
	fmt.Printf("处理进度:第%v行数据\n", count)
}
