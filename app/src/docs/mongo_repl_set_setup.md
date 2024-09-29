# SETUP STANDALONE MONGOD INTO REPLICA SET

## Procedure[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#procedure "Permalink to this heading")

1

### Shut down the standalone instance.[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#shut-down-the-standalone-instance. "Permalink to this heading")

Use [`mongosh`](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh) to [connect](https://www.mongodb.com/docs/mongodb-shell/connect/#std-label-mdb-shell-connect) to your `mongod` instance.

```
mongosh
```

Switch to the `admin` database and run [`shutdown`.](https://www.mongodb.com/docs/manual/reference/command/shutdown/#mongodb-dbcommand-dbcmd.shutdown)

```
use admindb.adminCommand(
    {
        shutdown: 1,     
        comment: "Convert to cluster"
    }
)
```

2

### Name the replica set.[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#name-the-replica-set. "Permalink to this heading")

If you configure your `mongod` instance from the command line, use the [`--replSet`](https://www.mongodb.com/docs/manual/reference/program/mongod/#std-option-mongod.--replSet) option to set a name for your replica set.

A typical command line invocation might include:

|Purpose|Option|
|---|---|
|Cluster name|[`--replSet`](https://www.mongodb.com/docs/manual/reference/program/mongod/#std-option-mongod.--replSet)|
|Network details|[`--port`](https://www.mongodb.com/docs/manual/reference/program/mongod/#std-option-mongod.--port)|
|Data path|[`--dbpath`](https://www.mongodb.com/docs/manual/reference/program/mongod/#std-option-mongod.--dbpath)|

Update the example code with the settings for your deployment.

```
mongod --replSet rs0 \       
        --port 27017 \       
        --dbpath /path/to/your/mongodb/dataDirectory \       --authenticationDatabase "admin" \       
        --username "adminUserName" \       
        --password
```

If you use a configuration file to start `mongodb`, add a `replication` section to your configuration file. Edit the `replSetName` value to set the name of your replica set.

```
replication:   
  replSetName: rs0
```

## Note

### Optional

You can specify the data directory, replica set name, and the IP binding in the `mongod.conf` [configuration file](https://www.mongodb.com/docs/manual/reference/configuration-options/), and start the [`mongod`](https://www.mongodb.com/docs/manual/reference/program/mongod/#mongodb-binary-bin.mongod) with the following command:

```
mongod --config /etc/mongod.conf
```

3

### Initialize the replica set.[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#initialize-the-replica-set. "Permalink to this heading")

To initialize the replica set, use `mongosh` to reconnect to your server instance. Then, run [`rs.initiate()`.](https://www.mongodb.com/docs/manual/reference/method/rs.initiate/#mongodb-method-rs.initiate)

```
rs.initiate()
```

You only have to initiate the replica set once.

To view the replica set configuration, use [`rs.conf()`.](https://www.mongodb.com/docs/manual/reference/method/rs.conf/#mongodb-method-rs.conf)

To check the status of the replica set, use [`rs.status()`.](https://www.mongodb.com/docs/manual/reference/method/rs.status/#mongodb-method-rs.status)

4

### Add nodes to the replica set.[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#add-nodes-to-the-replica-set. "Permalink to this heading")

The new replica set has a single, primary node. The next step is to add new nodes to the replica set. Review the documentation on clusters before you add additional nodes:

- [Replica set architectures](https://www.mongodb.com/docs/manual/core/replica-set-architectures/#std-label-replica-set-deployment-overview)
    
- [Adding members to a replica set](https://www.mongodb.com/docs/manual/tutorial/expand-replica-set/#std-label-server-replica-set-deploy-expand)
    

When you are ready to add nodes, use [`rs.add()`.](https://www.mongodb.com/docs/manual/reference/method/rs.add/#mongodb-method-rs.add)

5

### Update Your Application Connection String.[![](https://www.mongodb.com/docs/manual/assets/link.svg)](https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/#update-your-application-connection-string. "Permalink to this heading")

After you convert the mongod to a replica set, update the [connection string](https://www.mongodb.com/docs/manual/reference/connection-string/#std-label-mongodb-uri) used by your applications to the connection string for your replica set. Then, restart your applications.


## Problem issuing

<!-- https://www.mongodb.com/community/forums/t/cant-start-the-mongod-service-after-modify-the-mongo-conf-file/230168/6 -->