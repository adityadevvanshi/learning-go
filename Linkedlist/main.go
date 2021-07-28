package main

import (
	"os"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	ll "github.com/adityadevvanshi/linkedlist/LinkedList"
)
func InitLogger() {
	var encoderCfg = zap.NewProductionEncoderConfig()
	var atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)//for debug change to zap.DebugLevel
	var logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	))
	defer logger.Sync()
	ll.Logger = logger
}

func main(){
	InitLogger()
	llist1 := &ll.NodeList{}
	llist1.InsertAtBeginning(5)
	llist1.InsertAtBeginning(10)
	llist1.InsertAtBeginning(15)

	fmt.Println("First Linked List :")
	llist1.Head.PrintLinkedList()

	llist2:= &ll.NodeList{}
	llist2.InsertAtBeginning(2)
	llist2.InsertAtBeginning(4)

	fmt.Println("Second Linked List :")
	llist2.Head.PrintLinkedList()

	llist1.Mergetwolinkedlists(llist2)

	fmt.Println("Merged Linked List")
	llist1.Head.PrintLinkedList()
}