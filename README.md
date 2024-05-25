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

### Installation
```bash
git clone https://github.com/praveenmahasena/chat_auth.git
```

### Migration
This app uses [postgresql](https://www.postgresql.org/) as the database driver and before doing the migration we ask you to create a database named **chat_app** by loging in to your postgresql client and running the following commend `CREATE DATABASE chat_app;`.

Then with the help of [sqltool](https://github.com/praveenmahasena/sqltool/) you could setup all the other tables.


### Start
```bash
make server && ./bin/server
```
