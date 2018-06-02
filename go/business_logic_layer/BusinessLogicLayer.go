package business_logic_layer

import (
	"github.com/lampo100/botnet_p2p/models"
	"github.com/lampo100/botnet_p2p/p2p_layer"
	"log"
	"math/rand"
	"time"
	"sync"
)

var myNode models.Node
var pingedNodes []models.Node

var messagesChannel chan models.Message
var terminateChannel chan struct{}
var hasTerminated chan struct{}
var nextLayerTerminated chan struct{}

var mainTerminateChannel chan struct{}

var mutex = &sync.Mutex{}

func InitLayer(port uint32, terminate chan struct{}, thisTerminated chan struct{}) (bool, error) {
	rand.Seed(time.Now().UnixNano())
	terminateChannel = make(chan struct{})
	mainTerminateChannel = terminate
	hasTerminated = thisTerminated
	nextLayerTerminated = make(chan struct{})
	messagesChannel = make(chan models.Message, 16)
	node, err := generateSelfNode(port)
	if err != nil {
		return false, err
	}
	myNode = node
	p2p_layer.InitLayer(myNode, messagesChannel, terminateChannel, nextLayerTerminated)
	go messageListener()
	log.Println("[BL] Initialized")
	return true, nil
}

func messageListener() {
	messageHandlerLoop()
	log.Println("[BL] Leaving network")
	LeaveNetwork()
	close(terminateChannel)
	<-nextLayerTerminated
	log.Println("[BL] Terminated")
	hasTerminated <- struct{}{}
}

func messageHandlerLoop() {
	for {
		select {
		case msg := <-messagesChannel:
			switch msg.Type {
			case models.Message_FOUND_NODES:
				go handleFoundNodes(msg)
				break
			case models.Message_FIND_NODE:
				go handleFindNode(msg)
				break
			case models.Message_PING:
				go handlePingMessage(msg)
				break
			case models.Message_PING_RESPONSE:
				go handlePingResponse(msg)
				break
			}
		case <-mainTerminateChannel:
			return
		}
	}
}

func JoinNetwork(bootstrapNode models.Node) error {
	err := p2p_layer.FindNode(myNode, bootstrapNode, myNode.Guid)
	if err != nil {
		return err
	}
	log.Printf("[BL] Joined network with bootstrap at %v\n", bootstrapNode)
	return nil
}

func LeaveNetwork() {
	p2p_layer.LeaveNetwork()
}

func handleFoundNodes(msg models.Message) {
	foundNodesMsg := msg.GetFoundNodes().Nodes
	foundNodes := make([]models.Node, 0, len(foundNodesMsg))
	log.Printf("[BL] Got FoundNodes : %v\n", foundNodes)
	p2p_layer.AddNodeToRoutingTable(msg.Sender.ToNode())
	for _, f := range foundNodesMsg {
		foundNodes = append(foundNodes, f.ToNode())
	}

	for _, node := range foundNodes {
		log.Printf("[BL] Pinging node: %v", node.Guid)
		p2p_layer.Ping(myNode, node)
		mutex.Lock()
		pingedNodes = append(pingedNodes, node)
		mutex.Unlock()
	}
	<-time.After(10 * time.Second)
	mutex.Lock()
	for _, node := range pingedNodes {
		log.Printf("[BL] Timeout pinging node: %v", node)
		p2p_layer.RemoveFromRoutingTable(node)
	}
	pingedNodes = pingedNodes[:0]
	mutex.Unlock()
}

func handleFindNode(msg models.Message) {
	p2p_layer.AddNodeToRoutingTable(msg.Sender.ToNode())
	nodeGUID := models.GuidFromString(msg.GetFindNode().Guid)
	p2p_layer.FoundNodes(myNode, msg.Sender.ToNode(), nodeGUID)
}

func handlePingMessage(msg models.Message) {
	p2p_layer.AddNodeToRoutingTable(msg.Sender.ToNode())
	p2p_layer.PingResponse(myNode, msg.Sender.ToNode())
}

func handlePingResponse(msg models.Message) {
	mutex.Lock()
	for i, node := range pingedNodes {
		if node.Equals(msg.Sender.ToNode()) {
			p2p_layer.AddNodeToRoutingTable(msg.Sender.ToNode())
			pingedNodes = append(pingedNodes[:i], pingedNodes[i+1:]...)
			break
		}
	}
	mutex.Unlock()
}

func generateSelfNode(port uint32) (models.Node, error) {
	ip, err := getRemoteIP()
	if err != nil {
		return models.Node{}, err
	}
	isNAT, err := checkNAT()
	if err != nil {
		return models.Node{}, err
	}

	node := models.Node{
		Host:  ip,
		IsNAT: isNAT,
		Guid:  models.GenerateGUID(),
		Port:  port,
	}
	return node, nil
}
