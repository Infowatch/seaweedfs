package needle

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/Infowatch/seaweedfs/weed/storage/types"
)

const (
	NeedleChecksumSize = 4
	PairNamePrefix     = "Seaweed-"
)

/*
* A Needle means a uploaded and stored file.
* Needle file size is limited to 4GB for now.
 */
type Needle struct {
	Cookie Cookie   `comment:"random number to mitigate brute force lookups"`
	Id     NeedleId `comment:"needle id"`
	Size   Size     `comment:"sum of DataSize,Data,NameSize,Name,MimeSize,Mime"`

	DataSize     uint32 `comment:"Data size"` //version2
	Data         []byte `comment:"The actual file data"`
	Flags        byte   `comment:"boolean flags"` //version2
	NameSize     uint8  //version2
	Name         []byte `comment:"maximum 255 characters"` //version2
	MimeSize     uint8  //version2
	Mime         []byte `comment:"maximum 255 characters"` //version2
	PairsSize    uint16 //version2
	Pairs        []byte `comment:"additional name value pairs, json format, maximum 64kB"`
	LastModified uint64 //only store LastModifiedBytesLength bytes, which is 5 bytes to disk
	Ttl          *TTL

	Checksum   CRC    `comment:"CRC32 to check integrity"`
	AppendAtNs uint64 `comment:"append timestamp in nano seconds"` //version3
	Padding    []byte `comment:"Aligned to 8 bytes"`
}

func (n *Needle) String() (str string) {
	str = fmt.Sprintf("%s Size:%d, DataSize:%d, Name:%s, Mime:%s Compressed:%v", formatNeedleIdCookie(n.Id, n.Cookie), n.Size, n.DataSize, n.Name, n.Mime, n.IsCompressed())
	return
}

func CreateNeedleFromRequest(r *http.Request, fixJpgOrientation bool, sizeLimit int64, bytesBuffer *bytes.Buffer) (n *Needle, originalSize int, contentMd5 string, e error) {
	e = fmt.Errorf("CreateNeedleFromRequest is not implemented")
	return
}

func (n *Needle) ParsePath(fid string) (err error) {
	length := len(fid)
	if length <= CookieSize*2 {
		return fmt.Errorf("Invalid fid: %s", fid)
	}
	delta := ""
	deltaIndex := strings.LastIndex(fid, "_")
	if deltaIndex > 0 {
		fid, delta = fid[0:deltaIndex], fid[deltaIndex+1:]
	}
	n.Id, n.Cookie, err = ParseNeedleIdCookie(fid)
	if err != nil {
		return err
	}
	if delta != "" {
		if d, e := strconv.ParseUint(delta, 10, 64); e == nil {
			n.Id += Uint64ToNeedleId(d)
		} else {
			return e
		}
	}
	return err
}

func GetAppendAtNs(volumeLastAppendAtNs uint64) uint64 {
	return max(uint64(time.Now().UnixNano()), volumeLastAppendAtNs+1)
}

func (n *Needle) UpdateAppendAtNs(volumeLastAppendAtNs uint64) {
	n.AppendAtNs = max(uint64(time.Now().UnixNano()), volumeLastAppendAtNs+1)
}

func ParseNeedleIdCookie(key_hash_string string) (NeedleId, Cookie, error) {
	if len(key_hash_string) <= CookieSize*2 {
		return NeedleIdEmpty, 0, fmt.Errorf("KeyHash is too short.")
	}
	if len(key_hash_string) > (NeedleIdSize+CookieSize)*2 {
		return NeedleIdEmpty, 0, fmt.Errorf("KeyHash is too long.")
	}
	split := len(key_hash_string) - CookieSize*2
	needleId, err := ParseNeedleId(key_hash_string[:split])
	if err != nil {
		return NeedleIdEmpty, 0, fmt.Errorf("Parse needleId error: %v", err)
	}
	cookie, err := ParseCookie(key_hash_string[split:])
	if err != nil {
		return NeedleIdEmpty, 0, fmt.Errorf("Parse cookie error: %v", err)
	}
	return needleId, cookie, nil
}

func (n *Needle) LastModifiedString() string {
	return time.Unix(int64(n.LastModified), 0).Format("2006-01-02T15:04:05")
}

func max(x, y uint64) uint64 {
	if x <= y {
		return y
	}
	return x
}
