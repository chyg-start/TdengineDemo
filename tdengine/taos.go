package tdengine

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/taosdata/driver-go/v3/taosSql"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var SqlxBD *TDengineClient

type TDengineClient struct {
	DB *sqlx.DB
}
type TaosConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Pass       string `yaml:"password"`
	DB         string `yaml:"db"`
	DriverName string `yaml:"driver_name"`
}

func Connect(config TaosConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Pass, config.Host, strconv.Itoa(config.Port), config.DB)
	db, err := sqlx.Connect(config.DriverName, dsn)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	SqlxBD = &TDengineClient{DB: db}
	return nil
}

// 查询
func (tc *TDengineClient) Query(sql string) (*sql.Rows, error) {
	rows, err := tc.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// 执行语句
func (tc *TDengineClient) Exec(sql string) error {
	_, err := tc.DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

// 插入数据
// 有 tags实体判定是通过超级表建表的
func (tc *TDengineClient) Insert(stableName, tableName string, dest, tags interface{}) error {
	var sql string
	field, val := tc.getStructFiedelsInsertSql(dest)
	if stableName != "" && tags != nil {
		_, tagVal := tc.getStructFiedelsInsertSql(tags)
		sql = `INSERT INTO  %s (%s) USING %s TAGS (%s) VALUES (%s);`
		sql = fmt.Sprintf(sql, tableName, field, stableName, tagVal, val)
	} else {
		sql = `INSERT INTO %s (%s) VALUES (%s);`
		sql = fmt.Sprintf(sql, tableName, field, val)
	}
	_, err := tc.DB.Exec(sql)
	if err != nil {
		fmt.Println(err, sql)
		return err
	}
	return nil
}

// 批量插入
func (tc *TDengineClient) InsertBatch(tableName string, dest []interface{}) (err error) {
	var sql, fields, values string
	for _, item := range dest {
		field, vals := tc.getStructFiedelsInsertSql(item)
		values += fmt.Sprintf(` (%s) `, vals)
		fields = field
	}
	sql = fmt.Sprintf(`INSERT INTO %s (%s) VALUES  %s ;`, tableName, fields, values)
	_, err = tc.DB.Exec(sql)
	if err != nil {
		fmt.Println(err, sql)
		return err
	}
	return nil
}

// 通过结构体获取 字段名称 和值
func (tc *TDengineClient) getStructFiedelsInsertSql(item interface{}) (field, value string) {
	fields := make([]string, 0)
	values := make([]string, 0)
	s := reflect.ValueOf(item)
	t := reflect.TypeOf(item)
	for i := 0; i < s.NumField(); i++ {
		tag := s.Type().Field(i).Tag.Get("db") //关键标签
		if tag != "" && s.Field(i).CanInterface() {
			fields = append(fields, tag)
			var val = s.Field(i).Interface()
			if len(fmt.Sprintf("%v", val)) > 0 {
				var typeName = t.Field(i).Type.Kind()
				if typeName == reflect.String { //在Tdengine中
					values = append(values, fmt.Sprintf(`"%v"`, val))
				} else {
					values = append(values, fmt.Sprintf("%v", val))
				}
			} else {
				values = append(values, "")
			}
		}
	}
	field = strings.Join(fields, ",")
	value = strings.Join(values, ",")
	return
}
