package database

func InitiateAllDatabase() {
	CreateRedisClient()
	ConnectMongoDatabase()
}
