package utils

import "sync"

var MessagesMutex sync.Mutex
var SubscribersMutex sync.Mutex
