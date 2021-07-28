package main

import (
	"os"

	bst "github.com/adityadevvanshi/learning-go/BinarySearchTree"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	bst.Logger = logger
}

func main() {
	InitLogger()
	binarytree := &bst.Node{}
	binarytree.BuildBST()
	binarytree.InOrderTraversal()
}
