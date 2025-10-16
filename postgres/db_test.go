package bsdb

import (
	"blockchain_services/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

type AddressTest struct {
	gorm.Model
	Address string
	Key     string
}

func TestDb(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port,
	)

	gormCfg := &gorm.Config{}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	// 迁移 schema
	db.AutoMigrate(&AddressTest{})

	// Create
	db.Create(&AddressTest{Address: "0xE25583099BA105D9ec0A67f5Ae86D90e50036425", Key: "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d"})

	// Read
	var p AddressTest
	db.First(&p, 1)
	// 根据主键查找
	db.First(&p, "address = ?", "0xE25583099BA105D9ec0A67f5Ae86D90e50036425") // 查找
	fmt.Println(" 根据主键查找", p)

	// Update
	db.Model(&p).Update("Key", "0xc22F5A437A43254cc6DB63Bf75b2D3EfC4dC5b49")
	fmt.Println("更新字段", p)
	// Update - 更新多个字段
	db.Model(&p).Updates(AddressTest{Address: "0x670C34bFf1c1Cd71aD4e4B28C4c2338CA33C56eC", Key: "9fd16a2008d879795506f46bb5e7c400ebf3108dc8c235318577b98894e15609"}) // 仅更新非零值字段
	fmt.Println("更新多个字段:", p)
	db.Model(&p).Updates(map[string]interface{}{"Address": "0xE25583099BA105D9ec0A67f5Ae86D90e50036425", "Key": "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d"})
	fmt.Println("更新多个字段:", p)
	// Delete - 删除 p
	db.Delete(&p, 1)

	// 关闭数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("获取底层 sql.DB 失败:", err)
	} else {
		sqlDB.Close() // 关闭连接池
	}
}

func TestFirst(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port,
	)

	gormCfg := &gorm.Config{}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	db.Create(&AddressTest{Address: "0xE25583099BA105D9ec0A67f5Ae86D90e50036425", Key: "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d"})

	// Read
	var p AddressTest
	// 根据主键查找
	db.First(&p, "address = ?", config.EthParams.AdminAddr) // 查找
	fmt.Println("key -> ", p.Key)

}

func TestCheck(t *testing.T) {
	const initSQLFile = "./sql/receipts.sql"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port,
	)
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		t.Fatal("连接失败:", err)
	}

	log.Println("数据库已连接, dsn: ", dsn)

	if !db.Migrator().HasTable(&Receipt{}) {
		fmt.Println("数据库未初始化，正在执行初始化脚本...")

		// 读取 SQL 文件
		sqlBytes, err := os.ReadFile(initSQLFile)
		if err != nil {
			log.Fatalf("无法读取 SQL 文件 %s: %v", initSQLFile, err)
		}

		// 获取底层 *sql.DB 以执行原生 SQL
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("获取底层数据库连接失败:", err)
		}

		// 执行整个 SQL 脚本（注意：不支持 psql 特有语法）
		_, err = sqlDB.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("执行初始化 SQL 失败: %v", err)
		}

		// 创建标记表，表示已初始化
		err = db.AutoMigrate(&Receipt{})
		if err != nil {
			log.Fatal("创建初始化标记表失败:", err)
		}

		fmt.Println("数据库初始化完成！")
	} else {
		fmt.Println("数据库已初始化，跳过 SQL 执行。")
	}

	// 关闭数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("获取底层 sql.DB 失败:", err)
	} else {
		sqlDB.Close() // 关闭连接池
	}
}
