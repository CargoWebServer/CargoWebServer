package Server

import (
	"errors"
	"log"

	"code.myceliUs.com/Utility"
)

/**
 * Interface in go with the SearchEngine.
 * Cargo made use of Xapian as search engine. That library is written in C++ and
 * are interface on the network with help of the service container.
 */
type SearchEngine struct {
	conn connection
}

var (
	searchEngine *SearchEngine
)

/**
 * Create a new search engine and make a connection with it.
 */
func NewSearchEngine(address string) (*SearchEngine, error) {
	engine := new(SearchEngine)

	// Here I will open a new connection with the search engine.
	var err error
	engine.conn, err = GetServer().connect(address)

	return engine, err
}

/**
 * Get the search engine.
 */
func (this *Server) GetSearchEngine() *SearchEngine {
	if searchEngine == nil {
		searchEngine, _ = NewSearchEngine("127.0.0.1:9595")
	}
	return searchEngine
}

/**
 * Initialize security related information.
 */
func (this *SearchEngine) initialize() {

	log.Println("--> Initialize Search Engine")
}

/**
 * Index a entity on the search engine.
 */
func (this *SearchEngine) IndexEntity(dbpath string, entity Entity, language string) error {
	id := Utility.RandomUUID()

	// Call Execute Js function on the service
	method := "ExecuteJsFunction"
	params := make([]*MessageData, 0)

	to := make([]connection, 1)
	to[0] = GetServer().getConnectionById(this.conn.GetUuid())

	// The method to be call on the search engine.
	param0 := new(MessageData)
	param0.Name = "functionSrc"
	param0.Value = "XapianInterface.indexEntity"
	param0.TYPENAME = "Server.MessageData"
	params = append(params, param0)

	// The dbpath
	param1 := new(MessageData)
	param1.Name = "dbpath"
	param1.Value = dbpath
	param1.TYPENAME = "Server.MessageData"
	params = append(params, param1)

	// The entity object.
	param2 := new(MessageData)
	param2.Name = "entity"
	param2.Value = entity.GetObject()
	param2.TYPENAME = "Server.MessageData"
	params = append(params, param2)

	// Other param are related to the entity prototype.
	prototype := entity.GetPrototype()

	// Fields.
	param3 := new(MessageData)
	param3.Name = "fields"
	param3.Value = prototype.Fields
	param3.TYPENAME = "Server.MessageData"
	params = append(params, param3)

	// FieldsType.
	param4 := new(MessageData)
	param4.Name = "fieldstype"
	param4.Value = prototype.FieldsType
	param4.TYPENAME = "Server.MessageData"
	params = append(params, param4)

	param5 := new(MessageData)
	param5.Name = "language"
	param5.Value = language
	param5.TYPENAME = "Server.MessageData"
	params = append(params, param5)

	// Use a channel to synchronize the function.
	wait := make(chan interface{})

	// Create the request.
	rqst, _ := NewRequestMessage(id, method, params, to,
		/** Success callback **/
		func(rspMsg *message, caller interface{}) {
			caller.(chan interface{}) <- nil
		},
		/** Progress callback **/
		nil,
		/** Error callback **/
		func(errMsg *message, caller interface{}) {
			// Here I will return the error message.
			caller.(chan interface{}) <- errors.New(*errMsg.msg.Err.Message)
		}, wait)

	// Send the message to the search engine...
	go func(rqst *message) {
		GetServer().GetProcessor().m_sendRequest <- rqst
	}(rqst)

	result := <-wait
	if result == nil {
		return nil
	}

	return result.(error)
}
