package xlsx

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

//https://mp.weixin.qq.com/s/jRN1cUY7Dew4KWeV7ahFnQ

const (
	xlsxfile = "D://demo.xlsx"
)

// 新建表格
func Demo1(){
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("人员收集信息")
	if err != nil {
		fmt.Println(err)
	}
	//添加一行
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "性别"

	//第二行
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "张三"
	cell = row.AddCell()
	cell.Value = "男"

	err = file.Save(xlsxfile)
	if err != nil {
		panic(err.Error())
	}
}

func Demo2(){
	file, err := xlsx.OpenFile(xlsxfile)
	if err == nil {
		file.Sheets[0].Rows[1].Cells[0].Value = "李四"
	}

	//样式
	style := xlsx.NewStyle()
	style.Alignment = xlsx.Alignment{
		Horizontal:   "center",
		Vertical:     "center",
	}
	style.Font.Color = xlsx.RGB_Dark_Red
	style.Fill.BgColor = xlsx.RGB_Dark_Green
	file.Sheets[0].Rows[0].Cells[0].SetStyle(style)

	//保存
	file.Save("D://demo.xlsx")
}

