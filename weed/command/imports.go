package command

import (
	_ "net/http/pprof"

	_ "github.com/Infowatch/seaweedfs/weed/remote_storage/azure"
	_ "github.com/Infowatch/seaweedfs/weed/remote_storage/gcs"
	_ "github.com/Infowatch/seaweedfs/weed/remote_storage/s3"

	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/azuresink"
	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/b2sink"
	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/filersink"
	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/gcssink"
	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/localsink"
	_ "github.com/Infowatch/seaweedfs/weed/replication/sink/s3sink"

	_ "github.com/Infowatch/seaweedfs/weed/filer/arangodb"
	_ "github.com/Infowatch/seaweedfs/weed/filer/cassandra"
	_ "github.com/Infowatch/seaweedfs/weed/filer/elastic/v7"
	_ "github.com/Infowatch/seaweedfs/weed/filer/etcd"
	_ "github.com/Infowatch/seaweedfs/weed/filer/hbase"
	_ "github.com/Infowatch/seaweedfs/weed/filer/leveldb"
	_ "github.com/Infowatch/seaweedfs/weed/filer/leveldb2"
	_ "github.com/Infowatch/seaweedfs/weed/filer/leveldb3"
	_ "github.com/Infowatch/seaweedfs/weed/filer/mongodb"
	_ "github.com/Infowatch/seaweedfs/weed/filer/mysql"
	_ "github.com/Infowatch/seaweedfs/weed/filer/mysql2"
	_ "github.com/Infowatch/seaweedfs/weed/filer/postgres"
	_ "github.com/Infowatch/seaweedfs/weed/filer/postgres2"
	_ "github.com/Infowatch/seaweedfs/weed/filer/redis"
	_ "github.com/Infowatch/seaweedfs/weed/filer/redis2"
	_ "github.com/Infowatch/seaweedfs/weed/filer/redis3"
	_ "github.com/Infowatch/seaweedfs/weed/filer/sqlite"
	_ "github.com/Infowatch/seaweedfs/weed/filer/tarantool"
	_ "github.com/Infowatch/seaweedfs/weed/filer/tikv"
	_ "github.com/Infowatch/seaweedfs/weed/filer/ydb"

	_ "github.com/Infowatch/seaweedfs/weed/credential/filer_etc"
)
