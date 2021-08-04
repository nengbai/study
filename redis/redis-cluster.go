import "github.com/chasex/redis-go-cluster"

func main() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	cluster.Do("SET", "foo", "bar")
	cluster.Do("INCR", "mycount", 1)
	cluster.Do("LPUSH", "mylist", "foo", "bar")
	cluster.Do("HMSET", "myhash", "f1", "foo", "f2", "bar")

	reply, err := Int(cluster.Do("INCR", "mycount", 1))
	reply, err := String(cluster.Do("GET", "foo"))
	reply, err := Strings(cluster.Do("LRANGE", "mylist", 0, -1))
	reply, err := StringMap(cluster.Do("HGETALL", "myhash"))

	batch := cluster.NewBatch()
	err = batch.Put("LPUSH", "country_list", "France")
	err = batch.Put("LPUSH", "country_list", "Italy")
	err = batch.Put("LPUSH", "country_list", "Germany")
	err = batch.Put("INCRBY", "countries", 3)
	err = batch.Put("LRANGE", "country_list", 0, -1)
	reply, err = cluster.RunBatch(batch)

	var resp int
	for i := 0; i < 4; i++ {
		reply, err = redis.Scan(reply, &resp)
	}

	countries, err := Strings(reply[0], nil)

}
