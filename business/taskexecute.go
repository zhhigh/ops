/*
Dated:2013-11-16
*/
package business

import(
	"labix.org/v2/mgo/bson"
	"fmt"
	//"encoding/json"
	//"reflect"
)

type Host struct{
	Id_       bson.ObjectId `bson:"_id"`
	Ip   string     `bson:"ip"`
	Ostype  string     `bson:"ostype"`
	user  string     `bson:"user"`
	pwd  string     `bson:"pwd"`
}

type TaskList struct{
	Id_       bson.ObjectId `bson:"_id"`
	ostype  string     `bson:"ostype"`
	command  string     `bson:"command"`
}

/*
func (b *Business)GetHost(ip string)(interface{}){
   host := Host{}
   var  r interface{}
   b.Dbmongo.SetDBName("ops")
   b.Dbmongo.SetTableName("host")
   r =b.Dbmongo.FindOne(bson.M{"ip": ip},host)
   return r
}*/


func (b *Business)GetHost(ip string)(interface{}){
	var  r interface{}
	b.Dbmongo.SetDBName("ops")
	b.Dbmongo.SetTableName("host")
	b.Dbmongo.FindOne(bson.M{"ip": ip},&r)
	//fmt.Println(r)
	return r
}


/*
return : muti rows
*/
func (b *Business)GetTaskList(ostype string)(interface{}){
	//r := []TaskList{}
	//var  r interface{}
	var r []interface{}
	//var r []TaskList
	//r = t
	//fmt.Println(r)
	b.Dbmongo.SetDBName("ops")
	b.Dbmongo.SetTableName("tasklist")
	b.Dbmongo.Find(bson.M{"ostype": ostype},&r)
	//fmt.Println(r)
	return r
}



func (b *Business)GetTaskOne(ostype string)(interface{}){
	//fmt.Println(ostype)
	var r interface{}
	b.Dbmongo.SetDBName("ops")
	b.Dbmongo.SetTableName("tasklist")
	b.Dbmongo.FindOne(bson.M{"ostype": ostype},&r)
	fmt.Println(r)
	return r
}

