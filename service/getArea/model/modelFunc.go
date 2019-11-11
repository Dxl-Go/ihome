package model

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
)

func GetArea() ([]Area, error) {
	//链接数据库
	var areas []Area

	//从缓存中获取数据  缓存中没有则从redis中获取数据
	//conn,err:=redis.Dial("tcp","192.168.202.136:6379")
	conn := GlobalRedis.Get()
	areaByte,_:=redis.Bytes(conn.Do("get", "areaData"))
	if len(areaByte)==0 {
		//从mysql中获取数据
		if err := GlobalDB.Find(&areas).Error; err != nil {
			return areas, err
		}
		//序列化数据，存入redis中
		areaJson,err:=json.Marshal(areas)
		if err!=nil {
			return nil,err
		}
		_,err=conn.Do("set","areaData",areaJson)
		fmt.Println(err)
		fmt.Println("从mysql中获取 数据")


	}else {
		json.Unmarshal(areaByte,&areas)
		fmt.Println("从redis中获取数据")
	}
	return areas, nil

}
