package mysql

import (
	"GeekTime_Go/conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"sync"
)

var once sync.Once

// GetMysqlDbConnect 获取数据库连接
// NOTE：这里采用原生sql，未使用orm库
// NOTE：error库采用pkg/errors
func GetMysqlDbConnect() (*dao, error) {
	var db *sql.DB
	var err error

	// 初始化
	once.Do(func() {
		db, err = sql.Open(conf.DBDriver, fmt.Sprintf("%v:%v@/%v", conf.DBUser, conf.DBPass, conf.DBName))
		if err != nil {
			err = errors.Wrap(err, "GetMysqlDbConnect 连接数据库失败")
		} else {
			Dao.db = db
		}
	})
	if err != nil {
		return &dao{}, err
	}

	// TODO 配置文件变化重新设置连接

	// 测试连通性
	err = Dao.db.Ping()
	if err != nil {
		return &dao{}, errors.Wrap(err, "GetMysqlDbConnect Ping数据库失败")
	}

	return &Dao, nil
}

func (d *dao) Ping() error {
	return d.db.Ping()
}
