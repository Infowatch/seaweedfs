package pb

type EventErrorType int

const (
	TrivialOnError EventErrorType = iota
	FatalOnError
	RetryForeverOnError
)

// MetadataFollowOption is used to control the behavior of the metadata following
// process. Part of it is used as a cursor to resume the following process.
type MetadataFollowOption struct {
	ClientName             string
	ClientId               int32
	ClientEpoch            int32
	SelfSignature          int32
	PathPrefix             string
	AdditionalPathPrefixes []string
	DirectoriesToWatch     []string
	StartTsNs              int64
	StopTsNs               int64
	EventErrorType         EventErrorType
}
