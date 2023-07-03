package constants

import "time"

const (
	TimePingInterval = 30 * time.Second

	ErrOpenConnection = "Can't connect to database"
	ErrPingConnection = "Database not responding"
	SucPingConnection = "Database connected"

	MaxIdleConn = 5
	MaxOpenConn = 100
)
