# Kafka Replication Hammer

This is a simple application design to test a small database creating data on its 3 tables.

The software support both Microsoft SQL Server and MySQL/MariaDB.

## Configuration

In the `root` directory you will find a mockup  `.env` file called `example.env`. Copy it as just `.env` and fill the values as stated.

## Command line arguments

```
    -client    Runs the software as if it's running on client replica.
                In this mode it just insert records int the access table.
                (Requires -c).
    -c         Client id. If the software is running as a client the id
                of the client is required to differentiate which software
                instance is generating which records.
    -u         Maximum number of users to create in a run (defaults 10).
    -apu       Maximum number of articles per user to create in a run
                (defaults 10).
    -acc       maximum number of access records to create in a run 
                (defaults 100).
    -help      Prints the help.
```

## Compilation

To compile this software on any O.S. you must have [GoLang](https://go.dev/) installed on your machine. 

The go the the `root` directory of the project and type:

```bash
go build . -o hammer
```
It will compile for the O.S. you have installed on your machine.

## Tables structure

Below you can find the structure for the 3 tables.

To access the SQL scripts that create the empty database and its tables check the project's `scripts` directory.

### Table Accesses

| Field Name | Type | Len | Nulls | Pk  | Generated |
|------------|------| --- | ----- | --- | ----------|
| id         | int  |     | No    | Yes | Yes |
| ip         | varchar | 15 | No | No | No |
| origin_id  | int | | No | No | No |
| changed_date | datetime | | No | No | Yes |

### Table Users

| Field Name | Type | Len | Nulls | Pk  | Generated |
|------------|------| --- | ----- | --- | ----------|
| id         | int  |     | No    | Yes | Yes |
| first_name | varchar | 255 | No | No | No |
| last_name  | varchar | 255 | No | No | No |
| email | varchar | 255 | No | No | No |
| password | varchar | 255 | No | No | No |
| changed_date | datetime | | No | No | Yes |

### Table Articles

| Field Name | Type | Len | Nulls | Pk  | Generated |
|------------|------| --- | ----- | --- | ----------|
| id         | int  |     | No    | Yes | Yes |
| user_id | int | | No | No | No |
| title  | varchar | 255 | No | No | No |
| context_text | text | | No | No | No |
| changed_date | datetime | | No | No | Yes |

## Tables Diagram

Bellow you can check the relationships between the tables.

```
    _________________                   _________________
    | Table         |                   | Table         |
    | Users         |    id < user_id   | Articles      |
    |               | <---------------- |               |
    |               |                   |               |
    -----------------                   -----------------

    _________________
    | Table         |
    | Accesses      |
    |               |
    |               |
    -----------------

```

## Empty tables

To empty the database tables after running a test just type the following SQL commands in your favorite database IDE.

For SQL Server type:

```sql
TRUNCATE TABLE [dbo].accesses;
TRUNCATE TABLE [dbo].articles;
DELETE FROM [dbo].users;
GO
```

For MySQL/MariaDB type:
```sql
TRUNCATE TABLE KafkaReplication.accesses;
TRUNCATE TABLE KafkaReplication.articles;
DELETE FROM KafkaReplication.users;
```



