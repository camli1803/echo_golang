package config


type config struct {
    MongoDBHost string
    MongoDBUser string
    MongoDBPwd  string
    Database    string
}

var (
    AppConfig config
)

func Init() {
    AppConfig = config{
        MongoDBHost: localhost:27017,
        MongoDBUser: ""
        MongoDBPwd: ""
        Database: "demo"
    }
}