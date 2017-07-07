package Server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"strconv"

	"github.com/golang/protobuf/proto"
)

/**
 * Simple data representation use to exange data between participants.
 */
type MessageData struct {
	Name  string      /* The name of the parameter... **/
	Value interface{} /** Can be string, double, integer, bytes, or JSON structure **/
}

/**
 * Dont confuse this class with the myceliUs.Message class... this class
 * is use by the hub to keep the connection information associated whit
 * the myceliUs.Message.
 */
type message struct {
	from   connection
	to     []connection // One ore more destinnation for a message...
	msg    *Message     // The underlying message...
	caller interface{}  // a generic pointer holder.

	// asynchronous execution...
	successCallback  func(rspMsg *message, caller interface{})
	errorCallback    func(errMsg *message, caller interface{})
	progressCallback func(tranfertMsg *message, index int, total int, caller interface{})
}

func (self *message) GetBytes() []byte {
	data, err := proto.Marshal(self.msg)
	if err != nil {
		// t.Fatal("marshaling error: ", err)
		log.Printf("Error: ", err)
	}
	return data
}

func (self *message) GetId() string {
	var id string

	if *self.msg.Type == Message_REQUEST {
		id = self.msg.GetRqst().GetId()
	} else if *self.msg.Type == Message_RESPONSE {
		id = self.msg.GetRsp().GetId()
	} else if *self.msg.Type == Message_ERROR {
		id = self.msg.GetErr().GetId()
	} else if *self.msg.Type == Message_EVENT {
		// The id is the name of the event.
		id = self.msg.GetEvt().GetName()
	} else if *self.msg.Type == Message_TRANSFER {
		// The id is the name of the event.
		id = self.msg.GetId()
	}
	return id
}

func (self *message) GetMethod() string {
	method := ""

	if *self.msg.Type == Message_REQUEST {
		method = self.msg.GetRqst().GetMethod()
	}

	return method
}

/**
  Create a new Message from message data. The Message structure is generate
  by google protobuffer.
*/
func NewMessageFromData(data []byte, from connection) (*message, error) {
	m := new(message)
	m.from = from
	m.msg = &Message{}
	err := proto.Unmarshal(data, m.msg)
	return m, err
}

/**
 * Create a new event message to be sent over the network in a response to a request.
 */
func NewEvent(code int32, name string, eventData []*MessageData) (*Event, error) {
	evt := new(Event)
	evt.Name = &name
	evt.Code = &code
	for _, r := range eventData {
		msgEvtData := new(Data)
		msgEvtData.Name = proto.String(r.Name)
		msgEvtData.Type = new(Data_DataType)

		switch val := r.Value.(type) {
		// string
		case string:
			*msgEvtData.Type = Data_STRING
			msgEvtData.DataBytes = []byte(val)
		// Numeric value here
		case bool:
			*msgEvtData.Type = Data_BOOLEAN
			msgEvtData.DataBytes = strconv.AppendBool(msgEvtData.DataBytes, val)
		case int, int8, int16, int32, int64:
			*msgEvtData.Type = Data_INTEGER
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgEvtData.DataBytes = buf.Bytes()
		case float32, float64:
			*msgEvtData.Type = Data_DOUBLE
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgEvtData.DataBytes = buf.Bytes()
		// Bytes
		case []byte:
			*msgEvtData.Type = Data_BYTES
			msgEvtData.DataBytes = []byte(val)
		// Default will be Jsonification... what a word!!!
		default:
			// In that case I will considere the type as struct...
			*msgEvtData.Type = Data_JSON_STR
			b, err := json.Marshal(val)
			GetServer().GetEventManager().appendEventData(evt, string(b))
			if err != nil {
				return nil, err
			}
			msgEvtData.DataBytes = b

		}

		// The resulting message...
		evt.EvtData = append(evt.EvtData, msgEvtData)
	}

	return evt, nil
}

/**
 * Create a new error message to be sent over the network in a response to a request.
 */
func NewErrorMessage(id string, code int32, errMsg string, errData []byte, to []connection) *message {
	// Create the protobuffer message...
	m := new(message)
	m.to = to

	// Set the type to error
	m.msg = new(Message)

	m.msg.Type = new(Message_MessageType)
	*m.msg.Type = Message_ERROR

	index := int32(-1)
	total := int32(1)
	m.msg.Index = &index
	m.msg.Total = &total
	m.msg.Err = new(Error)
	m.msg.Err.Id = &id
	m.msg.Err.Code = &code
	m.msg.Err.Message = &errMsg
	m.msg.Err.Data = errData

	return m
}

/**
 * Create a request to send.
 */
func NewRequestMessage(id string, method string, params []*MessageData, to []connection, successCallback func(rspMsg *message, caller interface{}), progressCallback func(rspMsg *message, index int, total int, caller interface{}), errorCallback func(errMsg *message, caller interface{})) (*message, error) {

	// Create the protobuffer message...
	m := new(message)

	m.to = to

	// Set the type to request
	m.msg = new(Message)

	index_ := int32(-1)
	total := int32(1)
	m.msg.Index = &index_
	m.msg.Total = &total

	m.msg.Type = new(Message_MessageType)
	*m.msg.Type = Message_REQUEST

	// Initialyse the response message...
	m.msg.Rqst = new(Request)
	//id := uuid.NewRandom().String()
	m.msg.Rqst.Id = &id
	m.msg.Rqst.Method = &method
	m.msg.Rqst.Params = make([]*Data, len(params))

	// The callback to excute if the request is
	// no execute by the application server itself...
	m.successCallback = successCallback
	m.errorCallback = errorCallback
	m.progressCallback = progressCallback

	// Initialisation of response results...
	index := 0
	// The result will be
	var err error
	for _, r := range params {

		msgParam := new(Data)
		msgParam.Name = proto.String(r.Name)
		msgParam.Type = new(Data_DataType)

		switch val := r.Value.(type) {
		// string
		case string:
			*msgParam.Type = Data_STRING
			msgParam.DataBytes = []byte(val)
		case bool:
			*msgParam.Type = Data_BOOLEAN
			msgParam.DataBytes = strconv.AppendBool(msgParam.DataBytes, val)
		// Numeric value here
		case int, int8, int16, int32, int64:
			*msgParam.Type = Data_INTEGER
			buf := new(bytes.Buffer)
			err = binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgParam.DataBytes = buf.Bytes()
		// Float
		case float32, float64:
			*msgParam.Type = Data_DOUBLE
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgParam.DataBytes = buf.Bytes()
		// Bytes
		case []byte:
			*msgParam.Type = Data_BYTES
			msgParam.DataBytes = []byte(val)
		// Default will be Jsonification... what a word!!!
		default:
			// In that case I will considere the type as struct...
			*msgParam.Type = Data_JSON_STR
			var b []byte
			b, err = json.Marshal(val)
			if err != nil {
				return nil, err
			}
			msgParam.DataBytes = b
		}

		// The resulting message...
		m.msg.Rqst.Params[index] = msgParam
		index++
	}

	return m, err
}

/**
 * Create a resonpse to a request. The data contain the action result.
 */
func NewResponseMessage(id string, results []*MessageData, to []connection) (*message, error) {

	// Create the protobuffer message...
	m := new(message)
	m.to = to

	// Set the type to response
	m.msg = new(Message)

	index_ := int32(-1)
	total := int32(1)
	m.msg.Index = &index_
	m.msg.Total = &total

	m.msg.Type = new(Message_MessageType)
	*m.msg.Type = Message_RESPONSE

	// Initialyse the response message...
	m.msg.Rsp = new(Response)
	m.msg.Rsp.Id = &id
	m.msg.Rsp.Results = make([]*Data, len(results))

	// Initialisation of response results...
	index := 0
	// The result will be
	var err error
	for _, r := range results {

		msgResult := new(Data)
		msgResult.Name = proto.String(r.Name)
		msgResult.Type = new(Data_DataType)

		switch val := r.Value.(type) {
		// string
		case string:
			*msgResult.Type = Data_STRING
			msgResult.DataBytes = []byte(val)
		case bool:
			*msgResult.Type = Data_BOOLEAN
			msgResult.DataBytes = strconv.AppendBool(msgResult.DataBytes, val)
		// Numeric value here
		case int, int8, int16, int32, int64:
			*msgResult.Type = Data_INTEGER
			buf := new(bytes.Buffer)
			err = binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgResult.DataBytes = buf.Bytes()
		// Float
		case float32, float64:
			*msgResult.Type = Data_DOUBLE
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, val)
			if err != nil {
				log.Println("binary.Write failed:", err)
				return nil, err
			}
			msgResult.DataBytes = buf.Bytes()
		// Bytes
		case []byte:
			*msgResult.Type = Data_BYTES
			msgResult.DataBytes = []byte(val)
		// Default will be Jsonification... what a word!!!
		default:
			// In that case I will considere the type as struct...
			*msgResult.Type = Data_JSON_STR
			var b []byte
			b, err = json.Marshal(val)
			if err != nil {
				return nil, err
			}
			msgResult.DataBytes = b
		}

		// The resulting message...
		m.msg.Rsp.Results[index] = msgResult
		index++
	}
	return m, err
}
