//go:build !linux
// +build !linux

package stats

import "github.com/Infowatch/seaweedfs/weed/pb/volume_server_pb"

func fillInMemStatus(status *volume_server_pb.MemStatus) {
	return
}
