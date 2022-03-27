# Connecting to Docker DB from host

        > docker exec -it cassandra-db cat /etc/hosts

            127.0.0.1       localhost
            ::1     localhost ip6-localhost ip6-loopback
            fe00::0 ip6-localnet
            ff00::0 ip6-mcastprefix
            ff02::1 ip6-allnodes
            ff02::2 ip6-allrouters
            172.17.0.2      f28...11a

        > docker exec -it cassandra-db /bin/bash
            root@f28...11a:/# ip -4 -o address
        
                1: lo    inet 127.0.0.1/8 scope host lo\       valid_lft forever preferred_lft forever
                17: eth0    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0\       valid_lft forever preferred_lft forever


# Troubleshoot

## Context

Thunder client test

- `"172.17.0.2:7199"`
- `localhost:7199/api/users/getUser`

`GET` works, but `POST` doesn't. It returns "Connection was forcibly closed by a peer." message.

<br>

### REFERENCES

- How to Get A Docker Container IP Address - Explained with Examples https://www.freecodecamp.org/news/how-to-get-a-docker-container-ip-address-explained-with-examples/