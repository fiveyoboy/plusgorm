# plusgorm

一个增强 gorm 功能的包，功能包含：复制 db 、清理 db、深拷贝 db、db 数据库表结构转 go struct、解析 binglog 日志.......



## 总览

* TableToStruct：表结构转 go struct

* CloneDB：深拷贝 db

* NewDB：创建全新 db

* CleanDB：清楚 db 的 where 条件，避免交叉影响

* ...



## 快速开始

```go
package main

import (
	"github.com/fiveyoboy/plusgorm"
)

func main() {
	pg := plusgorm.NewPlusGorm(&plusgorm.DBConfig{
		Name:     "test",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "123456",
	})
	err := pg.TableToStruct("test_tb", &plusgorm.TableToStructConfig{
		Stdout:         false,
		FilePath:       "./table.go",
		MysqlGoTypeMap: nil,
	})
	if err != nil {
		return
	}
}

```

> 将会在当前目录下 创建 table.go 文件
