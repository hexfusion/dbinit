package dbinit

import (
	"fmt"

	"go.etcd.io/etcd/mvcc/backend"
	"go.etcd.io/etcd/pkg/fileutil"
)

var (
	membersBucketName = []byte("members")
	clusterBucketName = []byte("cluster")
)

// DBConfig holds the config for dbinit
type DBConfig struct {
	// Name of the db file
	Name string
	// Path of db file
	Path string
}

// Create generates a blank etcd db file
func Create(c DBConfig) error {
	if err := fileutil.TouchDirAll(c.Path); err != nil {
		return fmt.Errorf("create db directory error: %v", err)
	}
	dbFilePath := fmt.Sprintf("%s/%s", c.Path, c.Name)
	be := backend.NewDefaultBackend(dbFilePath)
	defer be.Close()
	mustCreateBackendBuckets(be)

	return nil
}

func mustCreateBackendBuckets(be backend.Backend) {
	tx := be.BatchTx()
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket(membersBucketName)
	tx.UnsafeCreateBucket(clusterBucketName)
}
