package userDao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	sql "github.com/jmoiron/sqlx"
	"web-demo-interface-test/config"
	"web-demo-interface-test/log"
)

var db *sql.DB

func InitMysql(config *config.MySQLConfig) (err error)  {
	log.Info("InitMysql ...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DB)
	db,err = sql.Connect("mysql",dsn)
	if err != nil {
		return
	}
	return 
}

// QueryRowDB 这个查询方法查到的结果，会存在这个dest里面，具体查到的数据根据sql语句确定
func QueryRowDB(dest interface{}, sql string, args ...interface{}) error {
	return db.Get(dest, sql, args...)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}