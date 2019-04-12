package Server

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"code.myceliUs.com/Utility"
)

/**
 * The message processor processes the incomming message received.
 */
type MessageProcessor struct {

	// Run until abortedByEnvironment is false
	abortedByEnvironment chan bool

	// The received message.
	m_incomingChannel chan *message

	// The message to be sent..
	m_outgoingChannel chan *message

	// serialyse request sending...
	m_sendRequest            chan *message
	m_receiveRequestResponse chan *message
	m_receiveRequestError    chan *message
}

/**
 * Create a new message processor.
 */
func newMessageProcessor() *MessageProcessor {
	p := new(MessageProcessor)

	// Incoming message channel
	p.m_incomingChannel = make(chan *message)

	// Outgoing message channel
	p.m_outgoingChannel = make(chan *message)

	// Request channel's
	p.m_sendRequest = make(chan *message)
	p.m_receiveRequestResponse = make(chan *message)
	p.m_receiveRequestError = make(chan *message)

	// Channel to stop the message proecessing.
	p.abortedByEnvironment = make(chan bool)

	return p
}

func (this *MessageProcessor) run() {

	// Request processing when the server act as client to other server (service)...
	go func(outgoingChannel chan (*message), sendRequest chan (*message), receiveRequestResponse chan (*message), receiveRequestError chan (*message)) {

		// Keep track of pending request here...
		var pendingRequest = make(map[string]*message)

		for {
			select {

			case m := <-sendRequest:
				// Set the request...
				if m.tryNb > 0 {
					pendingRequest[*m.msg.Rqst.Id] = m
					this.m_outgoingChannel <- m
					// Decrease the number of try.
					m.tryNb--
					// If not answer is receive from the server the request
					// will be resend after one second.
					go func(m *message, outgoingChannel chan (*message)) {
						timer := time.NewTimer(5 * time.Second)
						<-timer.C
						outgoingChannel <- m
					}(m, this.m_outgoingChannel)
				}

			case m := <-receiveRequestResponse:
				// Here I will execute the successCallback if some is define.
				rqst := pendingRequest[m.GetId()]
				if rqst != nil {
					// Remove the request from the list.
					delete(pendingRequest, m.GetId())
					m.tryNb = 0 // No more necessary...

					if rqst.successCallback != nil {
						// Call the success callback.
						rqst.successCallback(m, rqst.caller)
					}
				}

			case m := <-receiveRequestError:
				rqst := pendingRequest[m.GetId()]
				if rqst != nil {
					delete(pendingRequest, m.GetId())
					if rqst.errorCallback != nil {
						// Call the error callback.
						rqst.errorCallback(m, rqst.caller)
					}
				}

			}
		}
	}(this.m_outgoingChannel, this.m_sendRequest, this.m_receiveRequestResponse, this.m_receiveRequestError)

	// keep the list of pending message.

	// Those collection are access by help of channel.
	pendingMsg := make(map[string][]*message)
	pendingMsgChunk := make(map[string][][]byte)

	// Main message processing loop...
	for {
		select {
		case m := <-this.m_incomingChannel:
			// Process the incomming message.
			go this.processIncomming(m, pendingMsg, pendingMsgChunk)
		case m := <-this.m_outgoingChannel:
			// Process the outgoing message.
			go this.processOutgoing(m, pendingMsg)
		case done := <-this.abortedByEnvironment:
			if done {
				return
			}
		}
	}
}

/**
 * That function determine the max message size
 */
func getMaxMessageSize() int {
	return 17740
}

//////////////////////////////////////////////////////////////////////////////////
// Synchronize access function.
//////////////////////////////////////////////////////////////////////////////////

////////////////////////////////  Response function //////////////////////////////
/**
 * Always use that function to process a response, don't send the resonponse
 * directly with the connection.
 */
func (this *MessageProcessor) appendResponse(m *message) {
	// I will keep the reference to the request to be able
	// to made the action later.
	this.m_outgoingChannel <- m
}

////////////////////////////// Pending message ///////////////////////////////////z
func (this *MessageProcessor) createPendingMessages(m *message, pendingMsg map[string][]*message) {

	//Get the max size.
	maxSize := getMaxMessageSize()

	// So here I will chunk the file into smaler section and
	// send the messages to the client.
	count := len(m.GetBytes()) / maxSize

	// Round up here.
	if len(m.GetBytes())%maxSize > 0 {
		count++
	}

	id := m.GetId()

	// Create the message array
	pendingMsg[id] = make([]*message, count)

	for i := 0; i < count; i++ {
		// So here I w ill create the slice.
		var bytesSlice []byte
		var startIndex = i * maxSize

		if startIndex+maxSize < len(m.GetBytes()) {
			bytesSlice = m.GetBytes()[startIndex : startIndex+maxSize]
		} else {
			bytesSlice = m.GetBytes()[startIndex:]
		}

		transferMsg := new(message)
		transferMsg.to = m.to
		transferMsg.msg = new(Message)
		transferMsg.msg.Id = &id
		index_ := int32(i)
		total := int32(count)

		transferMsg.msg.Index = &index_
		transferMsg.msg.Total = &total

		transferMsg.msg.Type = new(Message_MessageType)
		*transferMsg.msg.Type = Message_TRANSFER

		transferMsg.msg.Data = make([]byte, len(bytesSlice))
		copy(transferMsg.msg.Data, bytesSlice)

		pendingMsg[id][i] = transferMsg
	}
	// So here I will start the pending message processing.
	this.processPendingMessage(id, pendingMsg)
}

//////////////////////////////////////////////////////////////////////////////////
// Processing functions.
//////////////////////////////////////////////////////////////////////////////////

/**
 * Process Is use to execute the action associated whit the request.
 */
func (this *MessageProcessor) processIncomming(m *message, pendingMsg map[string][]*message, pendingMsgChunk map[string][][]byte) {
	msg := m.msg

	if *msg.Type == Message_REQUEST {

		// I will create the new action
		a := newAction(msg.GetRqst().GetMethod(), m)

		// Now the parameters.
		for _, param := range msg.GetRqst().GetParams() {

			// Set the parameter type and name
			a.ParamTypeNames = append(a.ParamTypeNames, param.GetTypeName())
			a.ParamNames = append(a.ParamNames, param.GetName())
			if param.GetType() == Data_DOUBLE {
				val, err := strconv.ParseFloat(string(param.GetDataBytes()), 64)
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_INTEGER {

				val, err := strconv.ParseInt(string(param.GetDataBytes()), 10, 64)
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_BOOLEAN {
				val, err := strconv.ParseBool(string(param.GetDataBytes()))
				if err != nil {
					panic(err)
				}
				a.Params = append(a.Params, val)
			} else if param.GetType() == Data_STRING {
				val := string(param.GetDataBytes())
				a.Params = append(a.Params, val)

			} else if param.GetType() == Data_BYTES {
				a.Params = append(a.Params, param.GetDataBytes())
			} else if param.GetType() == Data_JSON_STR {
				val := string(param.GetDataBytes())
				val_, err := b64.StdEncoding.DecodeString(val)
				if err == nil {
					val = string(val_)
				}

				// Only registered type will be process sucessfully here.
				// how the server will be able to know what to do otherwise.
				if strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]") {

					// It contain an array of values to be init
					var values interface{}
					if param.GetTypeName() == "[]string" {
						values = make([]string, 0)
					} else {
						values = make([]interface{}, 0)
					}

					err = json.Unmarshal([]byte(val), &values)
					if err == nil {
						p, err := Utility.InitializeStructures(values.([]interface{}), param.GetTypeName(), setEntityFct)
						if err == nil {
							a.Params = append(a.Params, p.Interface())
						} else {
							//log.Println("Error:", err)
							// Here I will try to create a the array of object.
							if err.Error() == "NotDynamicObject" {
								p, err := Utility.InitializeArray(values.([]interface{}))
								if err == nil {
									if p.IsValid() {
										a.Params = append(a.Params, p.Interface())
									} else {
										// here i will set an empty generic array.
										a.Params = append(a.Params, make([]interface{}, 0))
									}
								} else {
									log.Panicln("---> 388 fail to initialize array")
								}
							}
						}
					} else {
						log.Panicln("---> 391 fail to unmarshal json!")
					}

				} else {
					// It contain an object.
					var valMap map[string]interface{}
					err = json.Unmarshal([]byte(val), &valMap)
					if err == nil {
						p, err := Utility.InitializeStructure(valMap, setEntityFct)
						if err != nil {
							log.Println("Error:", err)
							a.Params = append(a.Params, valMap)
						} else {
							a.Params = append(a.Params, p.Interface())
						}
					} else {
						// I will set a nil value to the parameter in that case.
						a.Params = append(a.Params, nil)
					}
				}
			}
		}

		go a.execute()

	} else if *msg.Type == Message_RESPONSE {
		// If the response is in the pending message I will process the next message.
		if _, ok := pendingMsg[msg.Rsp.GetId()]; ok { // if this.isPending(msg.Rsp.GetId()) {
			//do something here
			this.processPendingMessage(msg.Rsp.GetId(), pendingMsg)
		} else {
			// Here I received a response from the client so I will process it.
			this.m_receiveRequestResponse <- m
		}

	} else if *msg.Type == Message_ERROR {
		// An error was encounter by the client.
		// here error was received.
		this.m_receiveRequestError <- m

	} else if *msg.Type == Message_EVENT {
		// When the client throw an event this is the place where
		// I handle it.
		evt := msg.GetEvt()
		// I will process the event.
		GetServer().GetEventManager().BroadcastEvent(evt)

	} else if *msg.Type == Message_TRANSFER {
		total := int(msg.GetTotal())
		messageId := msg.GetId()
		index := int(msg.GetIndex())
		chunk := pendingMsgChunk[messageId]

		//LogInfo("---> Transfert message: ", messageId, ":", index, "/", total)
		if chunk != nil {
			// So here it's not the first message receive for the file.
			chunk[index] = msg.GetData()
			if index == total-1 {
				// In that case it's the last message.
				data := make([]byte, 0) // create the buffer that will contain the data.
				for i := 0; i < total; i++ {
					data = append(data, chunk[i]...)
				}

				// Release the memory for that message.
				delete(pendingMsgChunk, messageId)

				// and process the action...
				originMsg, err := NewMessageFromData(data, m.from)
				if err == nil {
					this.m_incomingChannel <- originMsg
				} else {
					log.Println("Error: ", err)
				}
			}

		} else {
			// The chunk is not there so I will insert it.
			// -- firt i will create a new array whit the necessary space.
			container := make([][]byte, total, total)
			container[0] = msg.GetData()
			pendingMsgChunk[messageId] = container

		}

		// Here I will send back an empty response to tell the other end
		// that the message is process in order to continue the transfer.
		to := make([]*WebSocketConnection, 1)
		to[0] = m.from
		result := make([]*MessageData, 0)
		resultMsg, _ := NewResponseMessage(messageId, result, to)
		this.m_outgoingChannel <- resultMsg
	}
}

/**
 * Process outgoing message to be sent to the client.
 */
func (this *MessageProcessor) processOutgoing(m *message, pendingMsg map[string][]*message) {
	// Get the max message size.
	maxSize := getMaxMessageSize()

	// Here I will send the message to the client.
	if *m.msg.Type == Message_REQUEST || *m.msg.Type == Message_RESPONSE {

		for i := 0; i < len(m.to); i++ {
			if m.to[i] == nil {
				// Local message here no need to send over socket.
				if *m.msg.Type == Message_RESPONSE {
					this.m_receiveRequestResponse <- m
				} else if *m.msg.Type == Message_REQUEST {
					this.m_incomingChannel <- m
				}
			} else {
				if len(m.GetBytes()) < maxSize {
					m.to[i].Send(m.GetBytes())
				} else {
					// so here I will split the message in multiple part
					// and send it.
					this.createPendingMessages(m, pendingMsg)
				}
			}
		}

	} else if *m.msg.Type == Message_EVENT || *m.msg.Type == Message_ERROR {
		// Event
		for i := 0; i < len(m.to); i++ {
			if m.to[i] == nil {
				// Local message here no need to send over socket.
				this.m_incomingChannel <- m
			} else {
				if len(m.GetBytes()) < maxSize {
					m.to[i].Send(m.GetBytes())
				} else {
					// so here I will split the message in multiple part
					// and send it.
					this.createPendingMessages(m, pendingMsg)
				}
			}
		}
	} else if *m.msg.Type == Message_TRANSFER {
		// Transfer
		for i := 0; i < len(m.to); i++ {
			if m.to[i] != nil {
				if m.to[i].IsOpen() {
					m.to[i].Send(m.GetBytes())
				}
			}
		}
	}
}

/**
 * Process pending message one by one.
 */
func (this *MessageProcessor) processPendingMessage(id string, pendingMsg map[string][]*message) {
	messages := pendingMsg[id]
	if len(messages) > 0 {
		msg := messages[0]
		if *msg.msg.Type == Message_TRANSFER {
			// I will now remove the first item from the array.
			pendingMsg[id] = make([]*message, 0)
			for i := 1; i < len(messages); i++ {
				pendingMsg[id] = append(pendingMsg[id], messages[i])
			}
			this.m_outgoingChannel <- msg
		}
	} else {
		// Remove the pendeing message from the list.
		delete(pendingMsg, id)
	}
}
