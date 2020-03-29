package main


import (
	"bytes"
	"fmt"
	"os"

	"protobuf/model"

	"github.com/golang/protobuf/jsonpb"
)

var user1 = &model.User {
	Id:			"u001",
	Name:		"Dharma Sastra",
	Password:	"password",
	Gender:		model.UserGender_MALE,
}

var userList = &model.UserList {
	List: []*model.User {
		user1,
	},
}

var garage1 = &model.Garage {
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.2212847,
		Longtitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main()  {
	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())

	// =========== as json string
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)
	
	// =========== as json string to object proto
	protoObject := new(model.GarageList)

	err2 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
   
	fmt.Printf("# ==== As String\n       %v \n", protoObject.String())
}