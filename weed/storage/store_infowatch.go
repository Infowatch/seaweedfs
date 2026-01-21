package storage

import (
	"fmt"

	"github.com/Infowatch/seaweedfs/weed/storage/needle"
)

func (s *Store) IsExists(i needle.VolumeId, n *needle.Needle) (bool, error) {
	if v := s.findVolume(i); v != nil {
		v.dataFileAccessLock.RLock()
		defer v.dataFileAccessLock.RUnlock()

		if v.nm == nil || v.DataBackend == nil {
			return false, ErrNotReady
		}
		nv, ok := v.nm.Get(n.Id)
		if !ok || nv.Offset.IsZero() {
			return false, nil
		}
		readSize := nv.Size
		if readSize.IsDeleted() {
			return false, nil
		}
		if readSize == 0 {
			return false, nil
		}
		return true, nil
	}
	return false, fmt.Errorf("volume %d not found", i)
}
