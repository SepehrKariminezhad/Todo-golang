package dbmanegment

var mysqlrepo = new(MysqlManegment)
var redisrepo = new(RedisManegmant)

type Mydb struct{

}

func (db *Mydb)Connect()error{
	err := mysqlrepo.Connect()
	if err != nil{
		return err
	}
	err = redisrepo.Connect()
	if err != nil{
		return err
	}
	return nil
}

func (db *Mydb)Insert(id uint64 , l string ){
	mysqlrepo.Insert(id , l)
	redisrepo.Set(id , l)
}

func (db *Mydb)Delete(id uint64){
	mysqlrepo.Delete(id)
	redisrepo.Delete(id)
}

func (db *Mydb)Read_all(){
	mysqlrepo.Read_all()
}

func (db *Mydb)Read(id uint64){
	if redisrepo.Rexists(id) {
		redisrepo.Get(id)
	}else{
		val := mysqlrepo.Read(id)
		redisrepo.Set(id , val)
	}
}

func (db *Mydb)Update(id uint64 , l string){
	mysqlrepo.Update(id , l)
	redisrepo.Set(id , l)
}