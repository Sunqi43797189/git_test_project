package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MemberOfficialMsgs struct {
	MemberId int `bson:"member_id"`
	MsgId int `bson:"msg_id"`
}

func main(){
	host := "127.0.0.1"
	port := "27017"
	dbName := "yidui_test_msg"
	session, err := mgo.Dial(host + ":" + port)
	defer session.Close()
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	//连接数据库
	session.SetMode(mgo.Monotonic, true)
	yidui_test_db := session.DB(dbName)
	collection := yidui_test_db.C("member_official_msgs")


	err = collection.Insert(&MemberOfficialMsgs{MemberId:1, MsgId:2})
	msg := MemberOfficialMsgs{}
	collection.Find(bson.M{"member_id": 1}).One(&msg)
	fmt.Println(msg)


	collection.Update(bson.M{"member_id": 1}, bson.M{"$set": bson.M{"msg_id": 5}})
	collection.Find(bson.M{"member_id": 1}).One(&msg)
	fmt.Println(msg)

	collection.Remove(bson.M{"member_id": 1})
	collection.RemoveAll(bson.M{"msg_id": 1})
	collection.Find(bson.M{"member_id": 1}).One(&msg)
	fmt.Println(msg)
}