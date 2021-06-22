package urfave

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/urfave/cli/v2"
	"os"
)

//教程:https://blog.csdn.net/sd653159/article/details/83381786

func Demo1(){
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Println("BOOM!")
		return nil
	}
	err := app.Run(os.Args);
	if err != nil {
		fmt.Println(err)
	}
}

/***
go run main.go --port 7777
go run main.go lang 许志勇
go run main.go xang 许志勇
go run main.go pre-seal
 */
func Demo2(){
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "GoTool"
	//程序的用途描述
	oApp.Usage = "To save the world"
	//程序的版本号
	oApp.Version = "1.0.0"
	var host string
	var debug bool
	//设置参数
	oApp.Flags = []cli.Flag{
		//参数类型string,int,bool
		&cli.StringFlag{
			Name:        "host",         //参数名字
			Value:       "127.0.0.1",      //参数默认值
			Usage:       "Server Address", //参数功能描述
			Destination: &host,            //接收值的变量
		},
		&cli.IntFlag{
			Name:        "port,p",
			Value:       8888,
			Usage:       "Server port",
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &debug,
		},
	}
	//该程序执行的代码
	oApp.Action = func(c *cli.Context) error {
		fmt.Printf("host=%v \n",host)
		fmt.Printf("port=%v \n",c.Int("port")) //不使用变量接收，直接解析
		fmt.Printf("host=%v \n",debug)
		return nil
	}
	//可以设置多命令
	oApp.Commands = []*cli.Command{
		{
			//命令全称
			Name: "lang",
			//命令简写
			Aliases:[]string{"l"},
			//命令详细描述
			Usage:"Setting language",
			Action: func(c *cli.Context) error {
				fmt.Printf("encoding=%v \n",c.Args().First())
				return nil
			},
		},
		{
			Name: "xang",
			Aliases:[]string{"x"},
			Usage:"Setting x",
			Action: func(c *cli.Context) error {
				fmt.Printf("encoding=%v \n",c.Args().First())
				return nil
			},
		},
		{
			Name: "pre-seal",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name:  "miner-addr",
					Value: "t01000",
					Usage: "specify the future address of your miner",
				},
				&cli.StringFlag{
					Name:  "sector-size",
					Value: "2KiB",
					Usage: "specify size of sectors to pre-seal",
				},
			},
			Action: func(c *cli.Context) error {
				minerAddr := c.String("miner-addr")
				maddr, err := address.NewFromString(minerAddr)
				if err == nil {
					fmt.Println(maddr)
				}
				sectorSizeInt, err := units.RAMInBytes(c.String("sector-size"))
				if err == nil {
					fmt.Println(sectorSizeInt)
					sectorSize := abi.SectorSize(sectorSizeInt)
					fmt.Println(sectorSize)
				}
				return nil
			},
		},
	}

	//启动
	if err := oApp.Run(os.Args); err != nil {
		fmt.Println(err)
	}
 }