package codec

import (
    "time"
    "github.com/huoarter/gohangout/protoLogEvent"
    proto "github.com/gogo/protobuf/proto"
)

type ProtobufDecoder struct {
}

func (pd *ProtobufDecoder) Decode(value []byte) map[string]interface{} {
    rst := make(map[string]interface{})
    rst["@timestamp"] = time.Now()
    ple := &protoLogEvent.ProtoLogEvent{}
    err := proto.Unmarshal(value, ple)
    if err != nil {
    	return map[string]interface{}{
    		"@timestamp": time.Now(),
    		"message":    string(value),
    	}
    }
    rst["loggerFqcn"] = ple.LoggerFqcn
    rst["marker"] = ple.Marker
    rst["level"] = ple.Level
    rst["loggerName"] = ple.LoggerName
    rst["message"] = ple.Message
    rst["timeMillis"] = ple.TimeMillis
    rst["thrown"] = ple.Thrown
    rst["thrownProxy"] = ple.ThrownProxy
    rst["contextMap"] = ple.ContextMap
    rst["contextStack"] = ple.ContextStack
    rst["threadName"] = ple.ThreadName
    rst["source"] = ple.Source
    rst["includeLocation"] = ple.IncludeLocation
    rst["endOfBatch"] = ple.EndOfBatch
    rst["containerMeta"] = ple.ContainerMeta
    rst["nanoTime"] = ple.NanoTime
    return rst
}
