1.  Get cassandra image.
    Installing Cassandra on Windows is a bit of hastle. In order to avoid this, we can use cassandra docker image.

        $ docker pull cassandra

            Using default tag: latest
            latest: Pulling from library/cassandra

Image successfully pulled.

        > docker ps

            CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS                                         NAMES
            286e121d4c98   cassandra              "docker-entrypoint.s…"   56 seconds ago   Up 52 seconds   7000-7001/tcp, 7199/tcp, 9042/tcp, 9160/tcp 👈	hungry_goldwasser
            013ea201dd0d   postgres               "docker-entrypoint.s…"   2 days ago       Up 2 minutes    0.0.0.0:5432->5432/tcp                        db

2.  Enter the cassandra container and run the DB.

- `docker exec it <container name> cqlsh`

        > docker exec -it hungry_goldwasser cqlsh 👈

            Connected to Test Cluster at 127.0.0.1:9042
            [cqlsh 6.0.0 | Cassandra 4.0.3 | CQL spec 3.4.5 | Native protocol v5]
            Use HELP for help.

        cqlsh> describe keyspaces

            system       system_distributed  system_traces  system_virtual_schema
            system_auth  system_schema       system_views

3.  Create a keyspace for the project.

        cqlsh> describe keyspaces

            system       system_distributed  system_traces  system_virtual_schema
            system_auth  system_schema       system_views

        cqlsh> CREATE KEYSPACE rest_api 👈

            ... WITH replication = {
            ... 'class': 'NetworkTopologyStrategy',
            ... 'datacenter1': 1
            ... } ;

        cqlsh> describe keyspaces

            rest_api 👈   system_auth         system_schema  system_views
            system        system_distributed  system_traces  system_virtual_schema

4.  Create a table.

        cqlsh> CREATE TABLE users (
                    ... FirstName text,
                    ... LastName text,
                    ... Email text,
                    ... PRIMARY KEY (Email));

        cqlsh:rest_api> describe tables

            users 👈

        cqlsh:rest_api> select * from users;

            email | firstname | lastname
            -------+-----------+----------

You can inspect your table by running the following:

        > describe users

            CREATE TABLE rest_api.users (
                email text PRIMARY KEY,
                firstname text,
                lastname text
            ) WITH additional_write_policy = '99p'
                AND bloom_filter_fp_chance = 0.01
                AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
                AND cdc = false
                AND comment = ''
                AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
                AND compression = {'chunk_length_in_kb': '16', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
                AND crc_check_chance = 1.0
                AND default_time_to_live = 0
                AND extensions = {}
                AND gc_grace_seconds = 864000
                AND max_index_interval = 2048
                AND memtable_flush_period_in_ms = 0
                AND min_index_interval = 128
                AND read_repair = 'BLOCKING'
                AND speculative_retry = '99p';

5.  Set up `go sql driver` to connect the `go application` with `cassandra db`.

        > go get github.com/gocql/gocql

6.  Bonus: don't forget to `Install/Update Tools` on VS code.

- `ctrl` + `shift` + `P`

  - `Go: Install/Update Tools`

        Tools environment: GOPATH=C:\...\go
        Installing 8 tools at C:\...\go\bin in module mode

            go-outline
            gotests
            gomodifytags
            impl
            goplay
            dlv
            staticcheck
            gopls

        Installing github.com/ramya-rao-a/go-outline@latest (C:\...\go\bin\go-outline.exe) SUCCEEDED
        Installing github.com/cweill/gotests/gotests@latest (C:\...\go\bin\gotests.exe) SUCCEEDED
        Installing github.com/fatih/gomodifytags@latest (C:\...\go\bin\gomodifytags.exe) SUCCEEDED
        Installing github.com/josharian/impl@latest (C:\...\go\bin\impl.exe) SUCCEEDED
        Installing github.com/haya14busa/goplay/cmd/goplay@latest (C:\...\go\bin\goplay.exe) SUCCEEDED
        Installing github.com/go-delve/delve/cmd/dlv@latest (C:\...\go\bin\dlv.exe) SUCCEEDED
        Installing honnef.co/go/tools/cmd/staticcheck@latest (C:\...\go\bin\staticcheck.exe) SUCCEEDED

## Troubleshoot

### Context 1

Error in handling POST request.

        > go run cmd/main2.go

            Starting server...
            Server is listening on port 8080 👈
            2022/03/26 14:39:00 http: panic serving 127.0.0.1:56782: runtime error: invalid
            goroutine 21 [running]:
            net/http.(*conn).serve.func1()

### Reason 1

- In the tutorial, both the Cassandra db and server are running on the computer.
- In my project, server is running on the laptop, but the cassandra db is running in a docker container.

  - Port needs configured properly.

        > docker inspect hungry_goldwasser

            ...
            "Ports": {
                "7000/tcp": null,
                "7001/tcp": null,
                "7199/tcp": null,
                "9042/tcp": null,
                "9160/tcp": null
            },
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "NetworkID": "94e5...f1151",
                    "EndpointID": "ee2...1f190",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2", 👈
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "DriverOpts": null
                }
            }
            ...

### Solution 1

1.  Set the right IP and port in the server.

        > docker

            CON ID  IMAGE       COMMAND                  PORTS                                           NAMES
            8eae    cassandra   "docker-entrypoint.s…"   7000-7001/tcp, 7199/tcp, 9042/tcp, 9160/tcp 👈 hungry_goldwasser

        📂/RESTAPI/pkg/Database/DatabaseConnection.go
        connection.cluster = gocql.NewCluster("172.17.0.2") // 👈 Docker container port IP.

        📂/RESTAPI/pkg/Router/Router.go
        log.Fatal(http.ListenAndServe(":7000", r)) // ":

        http://172.17.0.2:7199/api/users/updateUser

<br>

### Context 2

The port (`7000`) used in docker cassandra-db is in use on host.

        2022/03/26 20:09:27 listen tcp :7000: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.



        > netstat -a

            PROTOCAL  LOCAL ADDRESS    EXTERNAL ADDRESS   STATUS
            TCP       0.0.0.0:7680     Willow:0           LISTENING

### Reason 2

Docker port and host port are not properly mapped.

### Solution 2

Bind docker container port with host port.

        > docker run -d --name=cassandra-db ⚓ -p 7199:7199 👈 cassandra

            a7ff...e9b06

        > docker images

            REPOSITORY             TAG       IMAGE ID       CREATED      SIZE
            cassandra              latest    9d9ad7e447bd   5 days ago   341MB

        > docker ps

            CONTAINER ID   IMAGE       PORTS                                             NAMES
            8cfad1630f82   cassandra   7000-7001/tcp, 7199/tcp,  0.0.0.0:7199->7199/tcp  cassandra-db

<br>
<br>

### References

- CPoP (2020) Connect Go To Cassandra (GoCQL Driver) https://www.youtube.com/watch?v=2oVa7JLaO5E&list=PLle8fNcWU5av-TANvnRXqiM5BECBY-nvj&index=4&ab_channel=CPoP
- https://stackoverflow.com/questions/64088072/why-does-vs-code-not-auto-import-packages-using-go
- bind: Only one usage of each socket address (protocol/network address/port) is normally permitted. https://stackoverflow.com/questions/41836209/only-one-usage-of-each-socket-address-protocol-network-address-port-is-normall
