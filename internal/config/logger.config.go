package config

var LoggerConfig = []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase",
		"timeKey": "timestamp",
		"functionKey": "origin"
	  }
	}`)
