package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"log"
	"net/http"
	"strconv"
	"tdengineDemo/comom"
	"tdengineDemo/model"
	"tdengineDemo/tdengine"
	"time"
)

var redisCilent *redis.Client

func InitRedis() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisCilent = client

}

func main() {

	var cofig = tdengine.TaosConfig{
		Host:       "192.168.157.6",
		Port:       6030,
		User:       "root",
		Pass:       "taosdata",
		DB:         "iot",
		DriverName: "taosSql",
	}

	tdengine.Connect(cofig) //链接taos
	InitRedis()
	router := gin.Default()
	userRouter := router.Group("/api")
	{
		userRouter.GET("/list", getlist)
		userRouter.GET("/create", create)
		userRouter.GET("/createBacth", createBacth)
	}
	router.Run(":1130")

}

func getlist(c *gin.Context) {

	val := comom.Get("my_api_key")
	if val != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "my_api_key",
			"code": 200,
		})
		return
	}

	//var pool = comom.NewCachePool(time.Minute * 1) //缓存过期时间为 2 秒
	//
	//if !pool.Check("my_api_key") { // 检查请求是否超出限制
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": "my_api_key",
	//		"code": 200,
	//	})
	//	return
	//}
	sql := `SELECT * FROM tb_node_device_record LIMIT 10000; `
	rest := []model.TbNodeDeviceRecordDb{}
	errs := tdengine.SqlxBD.DB.Select(&rest, sql)
	if errs != nil {
		log.Fatalln(" gormDb failed to select from table, err:", errs)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": rest,
		"code": 200,
	})

	comom.Set("my_api_key", "111111", 60)
	return
}

func create(c *gin.Context) {
	var pool = comom.NewCachePool(time.Second * 1) //缓存过期时间为 2 秒

	if !pool.Check("my_api_keys") { // 检查请求是否超出限制
		c.JSON(http.StatusOK, gin.H{
			"data": "my_api_key",
			"code": 200,
		})
		return
	}
	item := model.TbNodeDeviceRecord{
		CreateTimeStr: time.Now().UnixMilli(),
		ConId:         "1111111111",
		Online:        0,
		AlarmState:    2,
		PowerState:    12,
		Brightness:    64,
		Temperature:   45,
		PowerF:        44,
		Kwh:           0.0998,
		InputVoltage:  34,
		InputCurrent:  2,
		InputPower:    9,
		LeakCurrent:   0,
		RatedPower:    221,
		PoleLeakV:     382,
		PoleLeakC:     34.3333,
		Tags: model.TbNodeDeviceRecordTags{
			NodeId:   "0000040014A1",
			LedNo:    1,
			NodeType: "15",
		},
	}
	var stableName, tableName string
	stableName = "tb_node_device_record"
	tableName = "Node_" + item.Tags.NodeId + "_" + strconv.Itoa(item.Tags.LedNo)
	err := tdengine.SqlxBD.Insert(stableName, tableName, item, item.Tags)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "",
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "",
		"code": 200,
		"msg":  "创建成功",
	})
	return

}

// 批量插入某数据库
func createBacth(c *gin.Context) {
	data := []interface{}{}
	for i := 0; i < 100; i++ {
		item := model.TbNodeDeviceRecord{
			CreateTimeStr: time.Now().UnixMilli(),
			ConId:         "0000000010i6",
			Online:        0,
			AlarmState:    2,
			PowerState:    12,
			Brightness:    64,
			Temperature:   45,
			InputVoltage:  34,
			InputCurrent:  2,
			InputPower:    9,
			PowerF:        44,
			LeakCurrent:   0,
			RatedPower:    221,
			PoleLeakV:     382,
			PoleLeakC:     34.3333,
			Kwh:           0.0998,
		}
		//sleep 1毫秒
		time.Sleep(time.Duration(1) * time.Millisecond) //不间隔的话会时间搓一样
		data = append(data, item)

	}

	var tableName string
	tableName = "Node_" + "0000040014A1" + "_" + strconv.Itoa(1)
	err := tdengine.SqlxBD.InsertBatch(tableName, data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "",
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "",
		"code": 200,
		"msg":  "创建成功",
	})
	return

}

func RedisClient(conId string) {

	// 设置限流器的参数
	limit := 2                        // 最大请求数量
	interval := 10                    // 时间间隔（秒）
	key := conId + "_" + "my_api_key" // 限流器的键名

	// 获取当前时间戳
	now := int64(time.Now().Unix())

	// 使用 Redis 的事务功能
	tx := redisCilent.TxPipeline()
	defer tx.Close()

	// 将当前时间戳添加到有序集合中
	tx.ZAdd(key, &redis.Z{Score: float64(now), Member: now})

	// 移除过期的时间戳
	tx.ZRemRangeByScore(key, "-inf", fmt.Sprintf("%d", (now-int64(interval))))

	// 获取当前有序集合的成员数量
	countCmd := tx.ZCard(key)

	// 提交事务
	_, err := tx.Exec()
	if err != nil {
		log.Fatalf("unable to update rate limiter: %v", err)
	}

	// 检查请求是否超过了限制
	if countCmd.Val() > int64(limit) {
		log.Printf("rate limit exceeded")
		return
	}

	// 处理请求
	log.Printf("processing request")
}
