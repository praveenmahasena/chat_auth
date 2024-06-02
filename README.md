# Chat Authentication System

This repository is a micro service which contains only Authentication and Authorization of a personal chat app.

## Features
- User registration
- Password hashing

## Upcoming Features
- login functionality
- Session management


## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) programming language
- [postgres](https://www.postgresql.org/) Data base
- [Gmail SMTP](https://support.google.com/accounts/search?q=smtp&hl=en&utm_source=google-account&utm_medium=web) SMTP password for email service

## Environment Variables

Make sure you have configure your Environment Variables properly on your `.env` file. The `.env` file contains following properties
```
HOST= # contains host id for postgresql
PORT= # contains port id for postgresql
USER= # contains user name for postgresql DB
PASSWORD= # contains user name for postgresql DB
DBNAME= # contains database name for postgresql DB
SSLMODE= # contains sslmode for postgresql DB
URL= # contains url for mail service link
FROM= # contains emailID which will be used for email service
EPASSWORD= # contains password of the emailID which is used for email service
```

### Installation
```bash
git clone https://github.com/praveenmahasena/chat_auth.git
```

### Migration
This app uses [postgresql](https://www.postgresql.org/) as the database driver and before doing the migration we ask you to create a database named **chat_app** by logging in to your postgresql client and running the following commend `CREATE DATABASE chat_app;`.

Then with the help of [sqltool](https://github.com/praveenmahasena/sqltool/) you could setup all the other tables.


### Start
```bash
make server && ./bin/server
```
