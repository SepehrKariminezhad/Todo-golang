package dbmanegment

import (
	"context"
	"fmt"
    "Todo_app/tools"
	"github.com/redis/go-redis/v9"
)


type RedisManegmant struct {
	Rdb *redis.Client
}



func (r *RedisManegmant) Connect ()error{
	r.Rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", 
        DB:       0,  
	})
	_ , err := r.Rdb.Ping(context.Background()).Result()
	if err != nil{
		return err
	}
    return nil
}


func (r *RedisManegmant) Set (id uint64 , l string){ 
    key := tools.Inttostr(id)
	err := r.Rdb.Set(context.Background(), key, l, 0).Err()
    if err != nil {
        fmt.Printf("Failed to save value in redis %s\n" , err.Error())
    }
}

func (r *RedisManegmant) Get (id uint64) string {
    key := tools.Inttostr(id)
	val , err := r.Rdb.Get(context.Background(), key).Result()
    if err != nil {
        fmt.Printf("Failed to get value in redis %s\n" , err.Error())
    }
	fmt.Printf("ID: %s        Log: %s\n", key, val)
    return val
}

func (r *RedisManegmant) Delete (id uint64) {
    key := tools.Inttostr(id) 
	_ , err := r.Rdb.Del(context.Background(), key).Result()
    if err != nil {
        fmt.Printf("Failed to delete key in redis %s\n" , err.Error())
    }
}


func (r *RedisManegmant) Rexists (tmpkey uint64) bool{
    tmpkey2 := tools.Inttostr(tmpkey)
	var flag bool = false
	exists, err := r.Rdb.Exists(context.Background(), tmpkey2).Result()
    if err != nil {
		fmt.Printf("error while finding key in redis %s\n" , err.Error())    
	}

    if exists > 0 {
        flag = true
    } else {
        fmt.Printf("The key %v does not exist in Redis.\n" , tmpkey2)
    }
	return flag
}
