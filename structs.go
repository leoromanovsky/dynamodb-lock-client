package dynamolock

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type acquireLockOptions struct {
	partitionKey                string
	sortKey                     *string
	data                        []byte
	replaceData                 bool
	deleteLockOnRelease         bool
	refreshPeriod               time.Duration
	additionalTimeToWaitForLock time.Duration
	additionalAttributes        map[string]*dynamodb.AttributeValue
	sessionMonitor              *sessionMonitor
}

type lockClientOptions struct {
	dynamoDBClient   *dynamodb.DynamoDB
	tableName        string
	partitionKeyName string
	sortKeyName      *string
	ownerName        string
	leaseDuration    time.Duration
	heartbeatPeriod  time.Duration
}

type getLockOptions struct {
	partitionKeyName    string
	sortKeyName         *string
	deleteLockOnRelease bool
}

type releaseLockOptions struct {
	lockItem   *Lock
	deleteLock bool
	data       []byte
}

// LockNotGrantedError indicates that an AcquireLock call has failed to
// establish a lock because of its current lifecycle state.
type LockNotGrantedError struct {
	msg string
}

func (e *LockNotGrantedError) Error() string {
	return e.msg
}

type createDynamoDBTableOptions struct {
	provisionedThroughput *dynamodb.ProvisionedThroughput
	tableName             string
	partitionKeyName      string
	sortKeyName           *string
}

type sendHeartbeatOptions struct {
	lockItem              *Lock
	data                  []byte
	deleteData            bool
	leaseDurationToEnsure time.Duration
}

type sessionMonitor struct {
	safeTime time.Duration
	callback func()
}

func (s *sessionMonitor) isLeaseEnteringDangerZone(lastAbsoluteTime time.Time) bool {
	return s.timeUntilLeaseEntersDangerZone(lastAbsoluteTime) <= 0
}

func (s *sessionMonitor) timeUntilLeaseEntersDangerZone(lastAbsoluteTime time.Time) time.Duration {
	return lastAbsoluteTime.Add(s.safeTime).Sub(time.Now())
}

func (s *sessionMonitor) runCallback() {
	if s.callback == nil {
		return
	}

	go s.callback()
}

func (s *sessionMonitor) hasCallback() bool {
	return s.callback != nil
}
