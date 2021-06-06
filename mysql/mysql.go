package mysql

import (
	"GeekTime_Go/conf"
	"GeekTime_Go/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
	"sync"
)

var once sync.Once

// GetDbConnect 获取数据库连接
// NOTE：这里采用原生sql，未使用orm库
// NOTE：error库采用pkg/errors
func GetDbConnect() (*dao, error) {
	var db *sql.DB
	var err error

	// 初始化
	once.Do(func() {
		db, err = sql.Open(conf.DBDriver, fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=true", conf.DBUser, conf.DBPass, conf.DBName))
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

// Ping 验证数据库连通性
// NOTE: 当前写法存在耦合
func (d *dao) Ping() error {
	return d.db.Ping()
}

// GetServiceRetryByAll 获取重试表全部数据
func (d *dao) GetServiceRetryByAll() ([]model.ServiceRetry, error) {
	res := make([]model.ServiceRetry, 0)

	rows, err := d.db.Query("SELECT * FROM service_retry")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("GetServiceRetryIDByAll 发现sql.ErrNoRows")
		} else {
			return res, errors.Wrap(err, "GetServiceRetryIDByAll 查询失败")
		}
	}

	for rows.Next() {
		var row model.ServiceRetry
		// 按顺序解析
		err := rows.Scan(&row.ID, &row.BizID, &row.Type, &row.Data, &row.Status, &row.RetryNum, &row.TraceID, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt)
		if err != nil {
			return res, errors.Wrap(err, "GetServiceRetryIDByAll Scan 出错")
		}

		res = append(res, row)
	}

	return res, nil
}

// UpdateServiceRetryStatusByID 更新重试表指定ID记录状态
func (d *dao) UpdateServiceRetryStatusByID(status, ID int64) error {
	_, err := d.db.Exec("UPDATE service_retry SET status=? WHERE id=?", status, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("UpdateServiceRetryStatusByID 发现sql.ErrNoRows")
		} else {
			return errors.Wrap(err, "UpdateServiceRetryStatusByID 更新失败")
		}
	}

	return nil
}
