package gateway

import (
	"fmt"
	"net"
	"runtime"
)

// ErrNoGateway is returned if a valid gateway entry was not
// found in the route table.
type ErrNoGateway struct {
	RouteTable []byte
}

// ErrCantParse is returned if the route table is garbage.
type ErrCantParse struct {
	RouteTable []byte
}

// ErrNotImplemented is returned if your operating system
// is not supported by this package. Please raise an issue
// to request support.
type ErrNotImplemented struct{}

// ErrInvalidRouteFileFormat is returned if the format
// of /proc/net/route is unexpected on Linux systems.
// Please raise an issue.
type ErrInvalidRouteFileFormat struct {
	row string
}

func (e *ErrNoGateway) Error() string {
	return fmt.Sprintf("no gateway found in route table:\n%s", string(e.RouteTable))
}

func (e *ErrCantParse) Error() string {
	return fmt.Sprintf("can't parse route table:\n%s", string(e.RouteTable))
}

func (*ErrNotImplemented) Error() string {
	return "not implemented for OS: " + runtime.GOOS
}

func (e *ErrInvalidRouteFileFormat) Error() string {
	return fmt.Sprintf("invalid row %q in route file: doesn't have 11 fields", e.row)
}

// DiscoverGateway is the OS independent function to get the default gateway
func DiscoverGateway() (ip net.IP, err error) {
	return discoverGatewayOSSpecific()
}

// DiscoverInterface is the OS independent function to call to get the default network interface IP that uses the default gateway
func DiscoverInterface() (ip net.IP, err error) {
	return discoverGatewayInterfaceOSSpecific()
}
