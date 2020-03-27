package codec

import (
	"time"
    "reflect"
    "github.com/huoarter/gohangout/protoLogEvent"
    "github.com/golang/protobuf/proto"
)

type ProtobufDecoder struct {
}
func StructToMap(obj interface{}) map[string]interface{}{
        obj1 := reflect.TypeOf(obj)
        obj2 := reflect.ValueOf(obj)

        var data = make(map[string]interface{})
        for i := 0; i < obj1.NumField(); i++ {
                data[obj1.Field(i).Name] = obj2.Field(i).Interface()
        }
        return data
}

func (pd *ProtobufDecoder) Decode(value []byte) map[string]interface{} {
    rst := make(map[string]interface{})
    rst["@timestamp"] = time.Now()
    ple := &protoLogEvent.ProtoLogEvent{}
    err := proto.Unmarshal(value, ple)
    data := StructToMap(*ple)
    if err != nil {
    	return map[string]interface{}{
    		"@timestamp": time.Now(),
    		"message":    string(value),
    	}
    }
    for k,v := range data {
        if _,ok := data[k]; ok {
            rst[k] = v 
        }
    }
    return rst
}
