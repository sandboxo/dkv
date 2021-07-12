# Start a full fledged dkv cluster.

dkv_master1:./bin/dkvsrv -role master -db-engine rocksdb -listen-addr 127.0.0.1:8070  -db-folder /tmp/nexus_dkv/1/data  -nexus-node-url "http://127.0.0.1:8021" -nexus-cluster-url "http://127.0.0.1:8021,http://127.0.0.1:8022,http://127.0.0.1:8023" -nexus-log-dir /tmp/nexus_dkv/1/logs -nexus-snap-dir /tmp/nexus_dkv/1/snap -nexus-snapshot-count 100 -nexus-snapshot-catchup-entries 5
dkv_master2:./bin/dkvsrv -role master -db-engine rocksdb -listen-addr 127.0.0.1:8080  -db-folder /tmp/nexus_dkv/2/data  -nexus-node-url "http://127.0.0.1:8022" -nexus-cluster-url "http://127.0.0.1:8021,http://127.0.0.1:8022,http://127.0.0.1:8023" -nexus-log-dir /tmp/nexus_dkv/2/logs -nexus-snap-dir /tmp/nexus_dkv/2/snap -nexus-snapshot-count 100 -nexus-snapshot-catchup-entries 5
dkv_master3:./bin/dkvsrv -role master -db-engine rocksdb -listen-addr 127.0.0.1:8090  -db-folder /tmp/nexus_dkv/3/data  -nexus-node-url "http://127.0.0.1:8023" -nexus-cluster-url "http://127.0.0.1:8021,http://127.0.0.1:8022,http://127.0.0.1:8023" -nexus-log-dir /tmp/nexus_dkv/3/logs -nexus-snap-dir /tmp/nexus_dkv/3/snap -nexus-snapshot-count 100 -nexus-snapshot-catchup-entries 5

dkv_slave1:./bin/dkvsrv -role slave -db-engine badger -diskless -listen-addr 127.0.0.1:8091 -repl-master-addr 127.0.0.1:8080 -repl-poll-interval 100ms
dkv_slave2:./bin/dkvsrv -role slave -db-engine badger -diskless -listen-addr 127.0.0.1:8092 -repl-master-addr 127.0.0.1:8080 -repl-poll-interval 100ms
dkv_slave3:./bin/dkvsrv -role slave -db-engine badger -diskless -db-folder /tmp/nexus_dkv/4/data -listen-addr 127.0.0.1:8093 -repl-master-addr 127.0.0.1:8080 -repl-poll-interval 100ms

