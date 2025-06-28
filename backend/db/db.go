package db

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
	"backend/structTypes"
    "gopkg.in/yaml.v3"
    _ "github.com/go-sql-driver/mysql"
)



var DB *sql.DB

func InitDB(configPath string) {
    // 读取配置文件
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatalf("读取配置文件失败: %v", err)
    }

    var cfg structTypes.Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        log.Fatalf("解析配置文件失败: %v", err)
    }


    // 构建DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DatabaseConnection.User,
        cfg.DatabaseConnection.Password,
        cfg.DatabaseConnection.Host,
        cfg.DatabaseConnection.Port,
        cfg.DatabaseNames.LocalSql,
    )

    // 连接数据库
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("数据库连接失败: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("数据库不可用: %v", err)
    }

    DB = db
    fmt.Println("MySQL 数据库连接成功！")
}