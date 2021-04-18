package config

type StoreConfig struct {
	ID   string
	Name string
	Host string
	Port int
	Pass string
}

func GetConfigByStoreID(storeId string) StoreConfig {
	// TODO: Replace this mock struct with actual fetch from config
	return StoreConfig{
		ID:   storeId,
		Name: "oryan-activities-redis-cluster-eu",
		Host: "redis-cluster-host.redis.com",
		Pass: "password1234",
	}
}
