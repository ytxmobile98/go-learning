# Database access

Use Golang to access a MySQL relational database.

See: https://go.dev/doc/tutorial/database-access

## Prereuisite: Install MySQL

This tutorial needs a MySQL database server.

My OS is Ubuntu 22.04, so I followed [this guide](https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/en/) to install MySQL.

1. Download the MySQL APT repository config tool from https://dev.mysql.com/downloads/repo/apt/.
2. Install the MySQL APT repository config tool.
3. Install `mysql-server` using APT:

    ```bash
    sudo apt install mysql-server
    ```

## Tutorial components

The tutorial has the following components:

1. Set up a database (using the file [`create-tables.sql`](./create-tables.sql)).
2. Find and import a database driver. (Driver used: https://github.com/go-sql-driver/mysql)
3. Get a database handle and connect.
4. Query for multiple rows. (_albums by artist_)
5. Query for a single row. (_album by id_)
6. Add data. (_add an album_)