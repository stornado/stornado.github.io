---
title: Redis info 命令简介
date: 2019-12-29 11:55:27
tags:
  - redis
categories:
  - [database, redis]
---

# Redis INFO

**Available since 1.0.0.**

The [INFO](info.html) command returns information and statistics about the server in a format that is simple to parse by computers and easy to read by humans.

The optional parameter can be used to select a specific section of information:

- `server`: General information about the Redis server
- `clients`: Client connections section
- `memory`: Memory consumption related information
- `persistence`: RDB and AOF related information
- `stats`: General statistics
- `replication`: Master/replica replication information
- `cpu`: CPU consumption statistics
- `commandstats`: Redis command statistics
- `cluster`: Redis Cluster section
- `keyspace`: Database related statistics

It can also take the following values:

- `all`: Return all sections
- `default`: Return only the default set of sections

When no parameter is provided, the `default` option is assumed.

<!-- more -->

## Return value

[Bulk string reply](../topics/protocol.html#bulk-string-reply): as a collection of text lines.

Lines can contain a section name (starting with a # character) or a property. All the properties are in the form of `field:value` terminated by `\r\n`.

```shell
/data #  redis-cli info
# Server
redis_version:5.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:c33a24c35f10df99
redis_mode:standalone
os:Linux 4.9.184-linuxkit x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:atomic-builtin
gcc_version:8.3.0
process_id:1
run_id:588256722f81ad322be7c318c45b41a4af82e607
tcp_port:6379
uptime_in_seconds:259678
uptime_in_days:3
hz:10
configured_hz:10
lru_clock:532902
executable:/data/redis-server
config_file:

# Clients
connected_clients:1
client_recent_max_input_buffer:4
client_recent_max_output_buffer:0
blocked_clients:0

# Memory
used_memory:862872
used_memory_human:842.65K
used_memory_rss:2637824
used_memory_rss_human:2.52M
used_memory_peak:4962376
used_memory_peak_human:4.73M
used_memory_peak_perc:17.39%
used_memory_overhead:840070
used_memory_startup:790240
used_memory_dataset:22802
used_memory_dataset_perc:31.39%
allocator_allocated:939408
allocator_active:1310720
allocator_resident:3895296
total_system_memory:4139208704
total_system_memory_human:3.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.40
allocator_frag_bytes:371312
allocator_rss_ratio:2.97
allocator_rss_bytes:2584576
rss_overhead_ratio:0.68
rss_overhead_bytes:-1257472
mem_fragmentation_ratio:3.30
mem_fragmentation_bytes:1837832
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:49694
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1577520526
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:499712
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:0

# Stats
total_connections_received:1378
total_commands_processed:1537
instantaneous_ops_per_sec:0
total_net_input_bytes:25223
total_net_output_bytes:79067
instantaneous_input_kbps:0.00
instantaneous_output_kbps:0.00
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:3
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:54
keyspace_misses:18
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:12879
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:0
master_replid:ba361feadd43179ca84c0232c16922fb44019d7b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:368.440000
used_cpu_user:106.630000
used_cpu_sys_children:0.010000
used_cpu_user_children:0.000000

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=1,expires=0,avg_ttl=0
/data #
/data #
/data #
/data #  redis-cli info all
# Server
redis_version:5.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:c33a24c35f10df99
redis_mode:standalone
os:Linux 4.9.184-linuxkit x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:atomic-builtin
gcc_version:8.3.0
process_id:1
run_id:588256722f81ad322be7c318c45b41a4af82e607
tcp_port:6379
uptime_in_seconds:259761
uptime_in_days:3
hz:10
configured_hz:10
lru_clock:532985
executable:/data/redis-server
config_file:

# Clients
connected_clients:1
client_recent_max_input_buffer:4
client_recent_max_output_buffer:0
blocked_clients:0

# Memory
used_memory:862904
used_memory_human:842.68K
used_memory_rss:2637824
used_memory_rss_human:2.52M
used_memory_peak:4962376
used_memory_peak_human:4.73M
used_memory_peak_perc:17.39%
used_memory_overhead:840070
used_memory_startup:790240
used_memory_dataset:22834
used_memory_dataset_perc:31.42%
allocator_allocated:939408
allocator_active:1310720
allocator_resident:3895296
total_system_memory:4139208704
total_system_memory_human:3.85G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.40
allocator_frag_bytes:371312
allocator_rss_ratio:2.97
allocator_rss_bytes:2584576
rss_overhead_ratio:0.68
rss_overhead_bytes:-1257472
mem_fragmentation_ratio:3.30
mem_fragmentation_bytes:1837832
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:49694
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0

# Persistence
loading:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1577520526
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:499712
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:0

# Stats
total_connections_received:1380
total_commands_processed:1539
instantaneous_ops_per_sec:0
total_net_input_bytes:25260
total_net_output_bytes:82363
instantaneous_input_kbps:0.00
instantaneous_output_kbps:0.00
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:3
expired_stale_perc:0.00
expired_time_cap_reached_count:0
evicted_keys:0
keyspace_hits:54
keyspace_misses:18
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:12879
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0

# Replication
role:master
connected_slaves:0
master_replid:ba361feadd43179ca84c0232c16922fb44019d7b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:368.760000
used_cpu_user:106.730000
used_cpu_sys_children:0.010000
used_cpu_user_children:0.000000

# Commandstats
cmdstat_hmget:calls=1,usec=25,usec_per_call=25.00
cmdstat_llen:calls=4,usec=1030,usec_per_call=257.50
cmdstat_command:calls=4,usec=25443,usec_per_call=6360.75
cmdstat_hgetall:calls=3,usec=1559,usec_per_call=519.67
cmdstat_keys:calls=7,usec=186,usec_per_call=26.57
cmdstat_setnx:calls=4,usec=61,usec_per_call=15.25
cmdstat_pexpire:calls=1,usec=82,usec_per_call=82.00
cmdstat_type:calls=4,usec=150,usec_per_call=37.50
cmdstat_msetnx:calls=2,usec=1403,usec_per_call=701.50
cmdstat_bitfield:calls=1,usec=6,usec_per_call=6.00
cmdstat_mset:calls=2,usec=7565,usec_per_call=3782.50
cmdstat_pfcount:calls=3,usec=117,usec_per_call=39.00
cmdstat_sadd:calls=1,usec=745,usec_per_call=745.00
cmdstat_expire:calls=1,usec=34,usec_per_call=34.00
cmdstat_pttl:calls=5,usec=1327,usec_per_call=265.40
cmdstat_info:calls=22,usec=24510,usec_per_call=1114.09
cmdstat_lpop:calls=9,usec=337,usec_per_call=37.44
cmdstat_ping:calls=1352,usec=18213,usec_per_call=13.47
cmdstat_brpop:calls=17,usec=1049,usec_per_call=61.71
cmdstat_setbit:calls=4,usec=1780,usec_per_call=445.00
cmdstat_decr:calls=1,usec=65,usec_per_call=65.00
cmdstat_bitpos:calls=2,usec=100,usec_per_call=50.00
cmdstat_rpush:calls=3,usec=271,usec_per_call=90.33
cmdstat_ltrim:calls=3,usec=266,usec_per_call=88.67
cmdstat_rpoplpush:calls=7,usec=256,usec_per_call=36.57
cmdstat_get:calls=9,usec=292,usec_per_call=32.44
cmdstat_bitop:calls=1,usec=31,usec_per_call=31.00
cmdstat_hmset:calls=3,usec=2189,usec_per_call=729.67
cmdstat_lpush:calls=15,usec=1000,usec_per_call=66.67
cmdstat_incr:calls=1,usec=1420,usec_per_call=1420.00
cmdstat_mget:calls=5,usec=154,usec_per_call=30.80
cmdstat_pfadd:calls=4,usec=1638,usec_per_call=409.50
cmdstat_rpop:calls=1,usec=13,usec_per_call=13.00
cmdstat_del:calls=12,usec=1866,usec_per_call=155.50
cmdstat_incrby:calls=1,usec=41,usec_per_call=41.00
cmdstat_bitcount:calls=2,usec=759,usec_per_call=379.50
cmdstat_set:calls=2,usec=686,usec_per_call=343.00
cmdstat_lrange:calls=14,usec=766,usec_per_call=54.71
cmdstat_ttl:calls=5,usec=107,usec_per_call=21.40
cmdstat_setex:calls=1,usec=1157,usec_per_call=1157.00

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=1,expires=0,avg_ttl=0
/data #
```



## Notes

Please note depending on the version of Redis some of the fields have been added or removed. A robust client application should therefore parse the result of this command by skipping unknown properties, and gracefully handle missing fields.