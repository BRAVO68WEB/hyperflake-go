package lib

import "time"

/*
GetCurrentTimestampSinceEpoch returns the current timestamp in milliseconds since a custom epoch.

Parameters:
  - epoch:
    The custom epoch timestamp in milliseconds.

Returns:
- The current timestamp in milliseconds since the custom epoch.
*/
func GetCurrentTimestampSinceEpoch(epoch int64) int64 {
	return time.Now().UnixMilli() - epoch
}

/*
GetCurrentTimestamp returns the current timestamp in milliseconds since the Unix epoch.

Returns:
- The current timestamp in milliseconds since the Unix epoch.
*/
func GetCurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}
