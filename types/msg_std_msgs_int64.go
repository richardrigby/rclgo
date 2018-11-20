package types
/////////////////////////////////////////////////////
//// THE CONTENT OF THIS FILE HAS BEEN AUTOGENERATED
/////////////////////////////////////////////////////
// #cgo CFLAGS: -I/opt/ros/bouncy/include
// #cgo LDFLAGS: -L/opt/ros/bouncy/lib -lrcl -lrosidl_generator_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c
// #include "msg_types.h"
import "C"
import "unsafe"


type StdMsgsInt64 struct {
	data    *C.std_msgs__msg__Int64
	MsgType MessageTypeSupport
}

func (msg *StdMsgsInt64) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsInt64) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsInt64) InitMessage() {
	msg.data = C.init_std_msgs_msg_Int64()
	msg.MsgType = GetMessageTypeFromStdMsgsInt64()
}

func (msg *StdMsgsInt64) SetInt64(data int64) {
	//TODO: to implement the setter
	msg.data.data = C.long(data)
}

func (msg *StdMsgsInt64) GetInt64() int64 {
	return int64(msg.data.data)
}

func (msg *StdMsgsInt64) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue:=""
	return myRetValue
}

func (msg *StdMsgsInt64) DestroyMessage() {
	C.destroy_std_msgs_msg_Int64(msg.data)
}

func GetMessageTypeFromStdMsgsInt64() MessageTypeSupport {
	return MessageTypeSupport{C.get_message_type_from_std_msgs_msg_Int64()}
}
