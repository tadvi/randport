package randport

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	minTCPPort         = 0
	maxTCPPort         = 65535
	maxReservedTCPPort = 1024
	maxRandTCPPort     = maxTCPPort - (maxReservedTCPPort + 1)
)

var (
	tcpPortRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// IsAvailable returns a flag indicating whether or not a TCP port is available.
func IsAvailable(port int) bool {
	if port < minTCPPort || port > maxTCPPort {
		return false
	}
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// Get a free, random TCP port between 1025-65535. If no free
// ports are available -1 is returned.
func Get() int {
	for i := maxReservedTCPPort; i < maxTCPPort; i++ {
		p := tcpPortRand.Intn(maxRandTCPPort) + maxReservedTCPPort + 1
		if IsAvailable(p) {
			return p
		}
	}
	return -1
}
