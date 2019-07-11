// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanstore

import (
	"time"
)

// returns index name with date
func indexWithDate(indexPrefix string, date time.Time) string {
	// spanDate := date.UTC().Format("2006-01-02")
	// return indexPrefix + spanDate
	return indexPrefix
}

// returns archive index name
func archiveIndex(indexPrefix, archiveSuffix string) string {
	return indexPrefix + archiveSuffix
}
