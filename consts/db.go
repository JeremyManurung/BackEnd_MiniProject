package consts

const (
	// DB_MAIN = "root:@tcp(host.docker.internal:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB_MAIN = "root:@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB_TEST = "root:@tcp(127.0.0.1:3306)/backend_test?charset=utf8mb4&parseTime=True&loc=Local"
)
