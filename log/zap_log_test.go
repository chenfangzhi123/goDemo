package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stderr", "./demo.log"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sinoaisdfjlkasjdfajs;dlfkja;lskdj")
	defer logger.Sync()
	for i := 0; i < 10; i++ {
		logger.Info("logger construction succeeded", zap.String("url", "asdklfj"))
		logger.Error("logger construction error", zap.String("url", "asdklfj"))

	}
	fmt.Printf("sinoaisdfjlkasjdfajs;dlfkja;lskdj")
}
