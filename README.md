# Service Inventory Go

This project is example service inventory management system written in Golang.

## Table of Contents

List of all readme contents:

- [Key Features](#key-features)
- [Built With](#built-with)
- [Deployment](#deployment)
- [API Reference](#api-reference)
- [Changelog](#changelog)
- [Authors](#authors)

## Key Features

List of all features/functionalities handled by this app/service:

 - Report Barang Masuk 
   - Retrieve All Barang Masuk
   - Retrieve Detail Barang Masuk
 - Report Barang Keluar
   - Retrieve All Barang Keluar
   - Retrieve Detail Barang Keluar
 - Report Penjualan
   - Retrieve All Penjualan
   - Retrieve Detail Penjualan
 - Report Stok Barang
   - Retrieve All Stok Barang
 - Report Value Barang
   - Retrieve All Value Barang 
 - Generate Report to CSV Format (For Generating CSV, sample availaible in Frontend) 

 - Frontend Repository `https://github.com/RezaNurRochmat13/Ijah-Inventory-System-Frontend`
 - DB Design  `https://drive.google.com/open?id=1wzZW0Pj7ceKIFnkauKLQC9i2frl4u-oS`

## Built With
List of all core technologies used to build this app/service along with their functions, versions (if any), and link to their online pages, grouped based on development stack:

### Backend
- [Golang](https://golang.org/) - Programming Language [v1.12.4]
- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP Routing Framework for Golang [v1.3.0]
- [Dep Package Management](https://github.com/golang/dep) - Dep Package Management For Golang [v0.5.1]
- [Viper Configuration Tools](https://github.com/spf13/viper) - Viper Configuration Tools [v1.3.2]
- [Go SQL Driver](github.com/mattn/go-sqlite3) - SQL Driver For SQLite DB [v1.4.1]
- [Docker](https://www.docker.com/) - Container Environment For Running Services [v18.09.6]

### Frontend
- [Angular](https://angular.io/) - JS Framework [v.7.0.0]
- [Bootstrap](http://getbootstrap.com/) - CSS Framework [v4.0.0]

### Automation Test
- For Running All Test, run command go test

### Database
- [SQLite 3](https://www.sqlite.org/index.html) - Core Database [v3]

## Deployment

List all required steps to deploy this app/service in server, like environment variables, server requirements, amount of compute resources (CPU, RAM), and dependency services that will communicate with this app/service.

### Running Application Using Docker
Before building and ship your application in Docker, follow this steps below
  - Compile your application to binary code using shell script in manifest. Running command sh manifest/compile.sh
  - After compiling, running command sudo docker build -t $IMAGES_NAME:$TAGS .
  - After successfully building your apps in Docker, we can run command sudo docker --rm -it -p $PORTS:$PORTS $IMAGES_NAME
  - Try to access your endpoints
  
### Running Application in Debugging Mode
For running in debugging mode, running shell script in manifest folder using command sh manifest/run_locally.sh in your terminal.


## API Reference

Depending on the size of this app/service APIs, if it is small and simple enough the reference docs can be added to this README. For medium to larger app/service please use Gitlab wiki or provide a link to where the API reference docs live.

For API Docs, using Google Docs https://docs.google.com/document/d/1XdxDOrk2CgvJaTId4g_4Jy7t4rcUf0Ig2rItvai8XoI/edit?usp=sharing

## Changelog

List of all released versions of this app/service along with their version logs, sorted from newest to oldest:

- 0.2.1
  - CHANGE: Update docs (module code remains unchanged)
- 0.2.0
  - CHANGE: Remove `setDefaultXYZ()`
  - ADD: Add `init()`
- 0.1.1
  - FIX: Crash when calling `baz()`
- 0.1.0
  - The first proper release
  - CHANGE: rename `foo()` to `bar()`
- 0.0.1
  - Work in progress

## Authors
List of all people working on this app/service along with their respective roles & contact emails:

- [Reja Nur Rochmat](mailto:rezanur@uii.ac.id) - Software Engineer