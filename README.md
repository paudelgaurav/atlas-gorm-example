# GORM Migration Versioning with Atlas

This repository simplifies the process of managing database migrations using GORM and Atlas, offering tools to generate, apply, and track database migrations effectively.

## Prerequisites

Before proceeding, make sure you have installed:
- [Git](https://git-scm.com/)
- [Go](https://golang.org/dl/)
- [Curl](https://curl.se/) (needed for installing Atlas)

## Installation and Setup

### Clone the Repository

Start by cloning the repository to your local machine:

```bash
git clone  https://github.com/paudelgaurav/atlas-gorm-example.git
cd atlas-gorm-migration-versioning

Install Dependencies
## Install the required Go modules:
- go mod download

Environment Setup
## Create a .env file from the provided sample and update it with your details:
- cp .env.sample .env

Run the Server
## Start the server using the following command:
- go run main.go
```

### Managing Migrations with Atlas

Atlas is utilized here for efficient schema and migration management. 
Below are the steps to install and use Atlas for your database management. 

```bash
Install Atlas 
- curl -sSf https://atlasgo.sh | sh

Migration Commands
## Generate Migration File
- atlas migrate diff --env gorm

Check Migration Status
## Before applying any changes, always check the migration status:
- atlas migrate status --url "mysql://username:password@:port/dbname"

Apply Migrations
## Apply the migrations to update your database schema:
- atlas migrate apply --url "mysql://username:password@:port/dbname"
```





