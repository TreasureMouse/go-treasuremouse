// Code generated by "stringer -type=logEventType"; DO NOT EDIT.

package main

import "strconv"

const _logEventType_name = "logEventChainInsertlogEventChainInsertSidelogEventHeaderChainInsertlogEventMinedBlocklogEventDownloaderStartlogEventDownloaderDonelogEventDownloaderFailedlogEventIntervallogEventBeforelogEventAfter"

var _logEventType_index = [...]uint8{0, 19, 42, 67, 85, 108, 130, 154, 170, 184, 197}

func (i logEventType) String() string {
	if i < 0 || i >= logEventType(len(_logEventType_index)-1) {
		return "logEventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _logEventType_name[_logEventType_index[i]:_logEventType_index[i+1]]
}
