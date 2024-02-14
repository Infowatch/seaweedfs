package pb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Infowatch/seaweedfs/weed/glog"
	"github.com/Infowatch/seaweedfs/weed/util"
)

const (
	Max_Message_Size = 1 << 30 // 1 GB
)

func ParseServerAddress(server string, deltaPort int) (newServerAddress string, err error) {

	host, port, parseErr := hostAndPort(server)
	if parseErr != nil {
		return "", fmt.Errorf("server port parse error: %v", parseErr)
	}

	newPort := int(port) + deltaPort

	return util.JoinHostPort(host, newPort), nil
}

func hostAndPort(address string) (host string, port uint64, err error) {
	colonIndex := strings.LastIndex(address, ":")
	if colonIndex < 0 {
		return "", 0, fmt.Errorf("server should have hostname:port format: %v", address)
	}
	port, err = strconv.ParseUint(address[colonIndex+1:], 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("server port parse error: %v", err)
	}

	return address[:colonIndex], port, err
}

func ServerToGrpcAddress(server string) (serverGrpcAddress string) {

	host, port, parseErr := hostAndPort(server)
	if parseErr != nil {
		glog.Fatalf("server address %s parse error: %v", server, parseErr)
	}

	grpcPort := int(port) + 10000

	return util.JoinHostPort(host, grpcPort)
}

func GrpcAddressToServerAddress(grpcAddress string) (serverAddress string) {
	host, grpcPort, parseErr := hostAndPort(grpcAddress)
	if parseErr != nil {
		glog.Fatalf("server grpc address %s parse error: %v", grpcAddress, parseErr)
	}

	port := int(grpcPort) - 10000

	return util.JoinHostPort(host, port)
}

func WithMasterClient(streamingMode bool, master ServerAddress, grpcDialOption any, waitForReady bool, fn func(client any) error) error {
	return fmt.Errorf("WithMasterClient not implemented")
}

func WithVolumeServerClient(streamingMode bool, volumeServer ServerAddress, grpcDialOption any, fn func(client any) error) error {
	return fmt.Errorf("WithVolumeServerClient not implemented")

}

func WithBrokerClient(streamingMode bool, broker ServerAddress, grpcDialOption any, fn func(client any) error) error {
	return fmt.Errorf("WithBrokerClient not implemented")
}

func WithOneOfGrpcMasterClients(streamingMode bool, masterGrpcAddresses map[string]ServerAddress, grpcDialOption any, fn func(client any) error) (err error) {
	return fmt.Errorf("WithOneOfGrpcMasterClients not implemented")
}

func WithBrokerGrpcClient(streamingMode bool, brokerGrpcAddress string, grpcDialOption any, fn func(client any) error) error {
	return fmt.Errorf("WithBrokerGrpcClient not implemented")
}

func WithFilerClient(streamingMode bool, signature int32, filer ServerAddress, grpcDialOption any, fn func(client any) error) error {
	return fmt.Errorf("WithFilerClient not implemented")
}

func WithGrpcFilerClient(streamingMode bool, signature int32, filerGrpcAddress ServerAddress, any, fn func(client any) error) error {
	return fmt.Errorf("WithGrpcFilerClient not implemented")
}

func WithOneOfGrpcFilerClients(streamingMode bool, filerAddresses []ServerAddress, grpcDialOption any, fn func(client any) error) (err error) {
	return fmt.Errorf("WithOneOfGrpcFilerClients not implemented")
}
