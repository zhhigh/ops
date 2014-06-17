package business

import(
	"ops/mongo"
)

type Business struct{
	Dbmongo *mongo.MongoDBConn
}

func New()(*Business){
	return &Business{}
}

func (b *Business)DBInit(server string){
	b.Dbmongo = mongo.NewMongoDBConn()
	b.Dbmongo.Connect(server)
}



func  (b *Business)DBDestroy(){
   b.Dbmongo.Stop()
}
