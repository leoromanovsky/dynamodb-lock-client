package dynamolock

// DynamoLock is the interface for this package.
// It can be used by upsteam projects to provide stub implementations
// in unit tests.
type DynamoLock interface {
	AcquireLock(key string, opts ...AcquireLockOption) (*Lock, error)
	SendHeartbeat(lockItem *Lock) error
	CreateTable(tableName string, provisionedThroughput *dynamodb.ProvisionedThroughput, opts ...CreateTableOption) (*dyn\|
amodb.CreateTableOutput, error)
	ReleaseLock(lockItem *Lock, opts ...ReleaseLockOption) (bool, error)
	Get(key string, opts ...GetOptions) (*Lock, error)
	Close()
}
