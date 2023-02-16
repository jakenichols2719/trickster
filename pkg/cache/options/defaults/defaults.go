/*
 * Copyright 2018 The Trickster Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package defaults

import "github.com/trickstercache/trickster/v2/pkg/cache/providers"

const (
	// DefaultCacheProvider is the default cache providers for any defined cache
	DefaultCacheProvider = "memory"
	// DefaultCacheProviderID is the default cache providers ID for any defined cache
	// and should align with DefaultCacheProvider
	DefaultCacheProviderID = providers.Memory
	// DefaultUseCacheChunking determines if caches use chunking when no value is set in config.
	DefaultUseCacheChunking = false
	// DefaultTimeseriesChunkFactor determines the extent of timeseries chunks, based on timeseries query step duration.
	DefaultTimeseriesChunkFactor = 420
	// DefaultByterangeChunkSize determines the size of byterange chunks.
	DefaultByterangeChunkSize = 4096
)
