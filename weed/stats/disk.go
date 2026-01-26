package stats

import (
	"github.com/shirou/gopsutil/v3/disk"

	"github.com/Infowatch/seaweedfs/weed/glog"
	"github.com/Infowatch/seaweedfs/weed/pb/volume_server_pb"
)

func NewDiskStatus(path string) *volume_server_pb.DiskStatus {
	res := &volume_server_pb.DiskStatus{Dir: path}

	if stat, err := disk.Usage(path); err == nil {
		res.All = stat.Total

		res.Free = stat.Free
		res.Used = stat.Used

		res.PercentUsed = float32(stat.UsedPercent)
		res.PercentFree = float32((float64(stat.Free) / float64(stat.Total)) * 100.)
	} else {
		glog.V(0).Infof("get disk usage error: %v", err)
	}

	return res
}
