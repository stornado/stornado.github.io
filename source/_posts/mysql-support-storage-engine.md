---
title: MySQL 8.0 Supported Storage Engines
date: 2019-12-16 15:07:42
tags:
---

# MySQL 8.0 Supported Storage Engines

- **`InnoDB`**: The default storage engine in MySQL 8.0. `InnoDB` is a transaction-safe (ACID compliant) storage engine for MySQL that has commit, rollback, and crash-recovery capabilities to protect user data. `InnoDB` row-level locking (without escalation to coarser granularity locks) and Oracle-style consistent nonlocking reads increase multi-user concurrency and performance. `InnoDB` stores user data in clustered indexes to reduce I/O for common queries based on primary keys. To maintain data integrity, `InnoDB` also supports `FOREIGN KEY` referential-integrity constraints.
- **`MyISAM`**: These tables have a small footprint. **Table-level locking** limits the performance in read/write workloads, so it is often used in read-only or read-mostly workloads in Web and data warehousing configurations.
- **`Memory`**: Stores all data in RAM, for fast access in environments that require quick lookups of non-critical data. This engine was formerly known as the `HEAP` engine. Its use cases are decreasing; `InnoDB` with its buffer pool memory area provides a general-purpose and durable way to keep most or all data in memory, and `NDBCLUSTER` provides fast key-value lookups for huge distributed data sets.
- **`CSV`**: Its tables are really text files with comma-separated values. CSV tables let you import or dump data in CSV format, to exchange data with scripts and applications that read and write that same format. Because CSV tables are not indexed, you typically keep the data in `InnoDB` tables during normal operation, and only use CSV tables during the import or export stage.
- **`Archive`**: These compact, unindexed tables are intended for storing and retrieving large amounts of seldom-referenced historical, archived, or security audit information.
- **`Blackhole`**: The Blackhole storage engine accepts but does not store data, similar to the Unix `/dev/null` device. Queries always return an empty set. These tables can be used in replication configurations where DML statements are sent to slave servers, but the master server does not keep its own copy of the data.
- **`NDB`** (also known as **`NDBCLUSTER`**): This clustered database engine is particularly suited for applications that require the highest possible degree of uptime and availability.
- **`Merge`**: Enables a MySQL DBA or developer to logically group a series of identical `MyISAM` tables and reference them as one object. Good for VLDB environments such as data warehousing.
- **`Federated`**: Offers the ability to link separate MySQL servers to create one logical database from many physical servers. Very good for distributed or data mart environments.
- **`Example`**: This engine serves as an example in the MySQL source code that illustrates how to begin writing new storage engines. It is primarily of interest to developers. The storage engine is a “stub” that does nothing. You can create tables with this engine, but no data can be stored in them or retrieved from them.

<!-- more -->

## Storage Engines Feature Summary

| Feature                                    | MyISAM       | Memory           | InnoDB       | Archive      | NDB          |
| ------------------------------------------ | ------------ | ---------------- | ------------ | ------------ | ------------ |
| **B-tree indexes**                         | Yes          | Yes              | Yes          | No           | No           |
| **Backup/point-in-time recovery** (note 1) | Yes          | Yes              | Yes          | Yes          | Yes          |
| **Cluster database support**               | No           | No               | No           | No           | Yes          |
| **Clustered indexes**                      | No           | No               | Yes          | No           | No           |
| **Compressed data**                        | Yes (note 2) | No               | Yes          | Yes          | No           |
| **Data caches**                            | No           | N/A              | Yes          | No           | Yes          |
| **Encrypted data**                         | Yes (note 3) | Yes (note 3)     | Yes (note 4) | Yes (note 3) | Yes (note 3) |
| **Foreign key support**                    | No           | No               | Yes          | No           | Yes (note 5) |
| **Full-text search indexes**               | Yes          | No               | Yes (note 6) | No           | No           |
| **Geospatial data type support**           | Yes          | No               | Yes          | Yes          | Yes          |
| **Geospatial indexing support**            | Yes          | No               | Yes (note 7) | No           | No           |
| **Hash indexes**                           | No           | Yes              | No (note 8)  | No           | Yes          |
| **Index caches**                           | Yes          | N/A              | Yes          | No           | Yes          |
| **Locking granularity**                    | Table        | Table            | Row          | Row          | Row          |
| **MVCC**                                   | No           | No               | Yes          | No           | No           |
| **Replication support** (note 1)           | Yes          | Limited (note 9) | Yes          | Yes          | Yes          |
| **Storage limits**                         | 256TB        | RAM              | 64TB         | None         | 384EB        |
| **T-tree indexes**                         | No           | No               | No           | No           | Yes          |
| **Transactions**                           | No           | No               | Yes          | No           | Yes          |
| **Update statistics for data dictionary**  | Yes          | Yes              | Yes          | Yes          | Yes          |



**Notes:**

1. Implemented in the server, rather than in the storage engine.

2. Compressed MyISAM tables are supported only when using the compressed row format. Tables using the compressed row format with MyISAM are read only.

3. Implemented in the server via encryption functions.

4. Implemented in the server via encryption functions; In MySQL 5.7 and later, data-at-rest tablespace encryption is supported.

5. Support for foreign keys is available in MySQL Cluster NDB 7.3 and later.

6. InnoDB support for FULLTEXT indexes is available in MySQL 5.6 and later.

7. InnoDB support for geospatial indexing is available in MySQL 5.7 and later.

8. InnoDB utilizes hash indexes internally for its Adaptive Hash Index feature.

9. See the discussion later in this section.



## MyISAM Storage Engine Features

| Feature                                                                                           | Support                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **B-tree indexes**                                                                                | Yes                                                                                                                                                       |
| **Backup/point-in-time recovery** (Implemented in the server, rather than in the storage engine.) | Yes                                                                                                                                                       |
| **Cluster database support**                                                                      | No                                                                                                                                                        |
| **Clustered indexes**                                                                             | No                                                                                                                                                        |
| **Compressed data**                                                                               | Yes (Compressed MyISAM tables are supported only when using the compressed row format. Tables using the compressed row format with MyISAM are read only.) |
| **Data caches**                                                                                   | No                                                                                                                                                        |
| **Encrypted data**                                                                                | Yes (Implemented in the server via encryption functions.)                                                                                                 |
| **Foreign key support**                                                                           | No                                                                                                                                                        |
| **Full-text search indexes**                                                                      | Yes                                                                                                                                                       |
| **Geospatial data type support**                                                                  | Yes                                                                                                                                                       |
| **Geospatial indexing support**                                                                   | Yes                                                                                                                                                       |
| **Hash indexes**                                                                                  | No                                                                                                                                                        |
| **Index caches**                                                                                  | Yes                                                                                                                                                       |
| **Locking granularity**                                                                           | Table                                                                                                                                                     |
| **MVCC**                                                                                          | No                                                                                                                                                        |
| **Replication support** (Implemented in the server, rather than in the storage engine.)           | Yes                                                                                                                                                       |
| **Storage limits**                                                                                | 256TB                                                                                                                                                     |
| **T-tree indexes**                                                                                | No                                                                                                                                                        |
| **Transactions**                                                                                  | No                                                                                                                                                        |
| **Update statistics for data dictionary**                                                         | Yes                                                                                                                                                       |



Each `MyISAM` table is stored on disk in two files. The files have names that begin with the table name and have an extension to indicate the file type. The data file has an `.MYD` (`MYData`) extension. The index file has an `.MYI`(`MYIndex`) extension. The table definition is stored in the MySQL data dictionary.





`InnoDB` is a general-purpose storage engine that balances high reliability and high performance. In MySQL 8.0, `InnoDB` is the default MySQL storage engine. Unless you have configured a different default storage engine, issuing a **`CREATE TABLE`** statement without an `ENGINE=` clause creates an `InnoDB` table.

## Key Advantages of InnoDB

- Its **DML** operations follow the **ACID** model, with **transactions** featuring **commit**, **rollback**, and **crash-recovery**capabilities to protect user data.
- Row-level **locking** and Oracle-style **consistent reads** increase multi-user concurrency and performance.
- `InnoDB` tables arrange your data on disk to optimize queries based on **primary keys**. Each `InnoDB` table has a primary key index called the **clustered index** that organizes the data to minimize I/O for primary key lookups.
- To maintain data **integrity**, `InnoDB` supports **`FOREIGN KEY`** constraints. With foreign keys, inserts, updates, and deletes are checked to ensure they do not result in inconsistencies across different tables.



## InnoDB Storage Engine Features

| Feature                                                                                           | Support                                                                                                                            |
| ------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| **B-tree indexes**                                                                                | Yes                                                                                                                                |
| **Backup/point-in-time recovery** (Implemented in the server, rather than in the storage engine.) | Yes                                                                                                                                |
| **Cluster database support**                                                                      | No                                                                                                                                 |
| **Clustered indexes**                                                                             | Yes                                                                                                                                |
| **Compressed data**                                                                               | Yes                                                                                                                                |
| **Data caches**                                                                                   | Yes                                                                                                                                |
| **Encrypted data**                                                                                | Yes (Implemented in the server via encryption functions; In MySQL 5.7 and later, data-at-rest tablespace encryption is supported.) |
| **Foreign key support**                                                                           | Yes                                                                                                                                |
| **Full-text search indexes**                                                                      | Yes (InnoDB support for FULLTEXT indexes is available in MySQL 5.6 and later.)                                                     |
| **Geospatial data type support**                                                                  | Yes                                                                                                                                |
| **Geospatial indexing support**                                                                   | Yes (InnoDB support for geospatial indexing is available in MySQL 5.7 and later.)                                                  |
| **Hash indexes**                                                                                  | No (InnoDB utilizes hash indexes internally for its Adaptive Hash Index feature.)                                                  |
| **Index caches**                                                                                  | Yes                                                                                                                                |
| **Locking granularity**                                                                           | Row                                                                                                                                |
| **MVCC**                                                                                          | Yes                                                                                                                                |
| **Replication support** (Implemented in the server, rather than in the storage engine.)           | Yes                                                                                                                                |
| **Storage limits**                                                                                | 64TB                                                                                                                               |
| **T-tree indexes**                                                                                | No                                                                                                                                 |
| **Transactions**                                                                                  | Yes                                                                                                                                |
| **Update statistics for data dictionary**                                                         | Yes                                                                                                                                |



## Benefits of Using InnoDB Tables

You may find `InnoDB` tables beneficial for the following reasons:

- If your server crashes because of a hardware or software issue, regardless of what was happening in the database at the time, you don't need to do anything special after restarting the database. `InnoDB` **crash recovery** automatically finalizes any changes that were committed before the time of the crash, and undoes any changes that were in process but not committed. Just restart and continue where you left off.
- The `InnoDB` storage engine maintains its own **buffer pool** that caches table and index data in main memory as data is accessed. Frequently used data is processed directly from memory. This cache applies to many types of information and speeds up processing. On dedicated database servers, up to 80% of physical memory is often assigned to the buffer pool.
- If you split up related data into different tables, you can set up **foreign keys** that enforce **referential integrity**. Update or delete data, and the related data in other tables is updated or deleted automatically. Try to insert data into a secondary table without corresponding data in the primary table, and the bad data gets kicked out automatically.
- If data becomes corrupted on disk or in memory, a **checksum** mechanism alerts you to the bogus data before you use it.
- When you design your database with appropriate **primary key** columns for each table, operations involving those columns are automatically optimized. It is very fast to reference the primary key columns in **`WHERE`** clauses, **`ORDER BY`** clauses, **`GROUP BY`** clauses, and **join** operations.
- Inserts, updates, and deletes are optimized by an automatic mechanism called **change buffering**. `InnoDB` not only allows concurrent read and write access to the same table, it caches changed data to streamline disk I/O.
- Performance benefits are not limited to giant tables with long-running queries. When the same rows are accessed over and over from a table, a feature called the **Adaptive Hash Index** takes over to make these lookups even faster, as if they came out of a hash table.
- You can compress tables and associated indexes.
- You can create and drop indexes with much less impact on performance and availability.
- Truncating a **file-per-table** tablespace is very fast, and can free up disk space for the operating system to reuse, rather than freeing up space within the **system tablespace** that only `InnoDB` can reuse.
- The storage layout for table data is more efficient for **`BLOB`** and long text fields, with the **DYNAMIC** row format.
- You can monitor the internal workings of the storage engine by querying **INFORMATION_SCHEMA** tables.
- You can monitor the performance details of the storage engine by querying **Performance Schema** tables.
- You can freely mix `InnoDB` tables with tables from other MySQL storage engines, even within the same statement. For example, you can use a **join** operation to combine data from `InnoDB` and **`MEMORY`** tables in a single query.
- `InnoDB` has been designed for CPU efficiency and maximum performance when processing large data volumes.
- `InnoDB` tables can handle large quantities of data, even on operating systems where file size is limited to 2GB.