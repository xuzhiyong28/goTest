package protobuf

import "fmt"
import "github.com/golang/protobuf/proto"

func Demo1() {
	person := &Person{
		Name:  "Jack",
		Age:   18,
		Hobby: []string{"sing", "dance", "basketball", "rap"},
	}
	binaryData, err := proto.Marshal(person) //序列化
	if err != nil {
		fmt.Println("proto.Marshal err:", err)
	}

	//反序列化
	newPerson := &Person{}
	err = proto.Unmarshal(binaryData, newPerson)
	if err != nil {
		fmt.Println("proto.Unmarshal err:", err)
	}

	fmt.Println("序列化前的原始数据:", person)
	fmt.Println("反序列化得到数据:", newPerson)
}
