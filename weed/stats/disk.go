package stats

import (
	"github.com/Infowatch/seaweedfs/weed/glog"
	"github.com/Infowatch/seaweedfs/weed/pb/volume_server_pb"
)

func NewDiskStatus(path string) (disk *volume_server_pb.DiskStatus) {
	disk = &volume_server_pb.DiskStatus{Dir: path}
	fillInDiskStatus(disk)
	if disk.PercentUsed > 95 {
		glog.V(0).Infof("disk status: %v", disk)
	}
	return
}
