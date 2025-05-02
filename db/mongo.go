package db
import (
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
var DBClient *mongo.Client

func ConnectToDb()(client *mongo.Client, err error){
	fmt.Println("connecting to mongo db database")
	var mongodb_url string
	if mongodb_url = os.Getenv("MONGO_URI"); mongodb_url == ""{
		panic("you must connect to a valid database")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodb_url).SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	DBClient = client

	return
}
