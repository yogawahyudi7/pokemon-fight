<div align="center">
  <a href="https://raw.githubusercontent.com/yogawahyudi7/POKEMON-FIGHT/develop/docs/logo.png">
    <img src="https://raw.githubusercontent.com/yogawahyudi7/POKEMON-FIGHT/develop/docs/logo.png">
  </a>
</div>
<div>
#  POKEMON-FIGHT:CLUB

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

## Table of Content
- [POKEMON-FIGHT:CLUB](#pokemon-fightclub)
  - [Table of Content](#table-of-content)
  - [Features](#features)
    - [Endpoints](#endpoints)
    - [API Documentation](#api-documentation)
  - [System Design](#system-design)
    - [ERD](#erd)
    - [Layered Architecture](#layered-architecture)
  - [Getting Started](#getting-started)
    - [Installing](#installing)
  - [Authors](#authors)

## Features
- JWT Authentication
- Layered Architecture
- Dependency Injection
- Parameters Validation
- Blacklist Feature
- API Consume - [POKE-API](https://pokeapi.co/)
- Multi Role Middleware (Bos, Operasional, Pengedar)

### Endpoints
- [x] Register Bos
- [x] Register Operasional
- [x] Register Pengedar
- [x] Login
- [x] Logout
- [x] Get all pokemon - [POKE-API](https://pokeapi.co/)
- [x] Get pokemon by Name & ID - [POKE-API](https://pokeapi.co/)
- [x] Add Season
- [x] Get All Season
- [x] Add Competition Pokemon
- [x] Get Competitions Pokemon
- [x] Get Scores Pokemon
- [x] Add Blacklist pokemon
- [x] Get Blacklist pokemon


### API Documentation
![API Documentation](https://raw.githubusercontent.com/yogawahyudi7/POKEMON-FIGHT/develop/docs/postman.png)
Application Programming Interface is available at [POSTMAN.](https://documenter.getpostman.com/view/16411992/2s93Jrx5NY)

## System Design

### ERD
![Pokemon Fight Club - ERD](https://raw.githubusercontent.com/yogawahyudi7/POKEMON-FIGHT/develop/docs/erd3.png)

### Layered Architecture
![Pokemon Fight Club - Layered Architecture](https://raw.githubusercontent.com/yogawahyudi7/POKEMON-FIGHT/develop/docs/layeredStructure.png)

## Getting Started

Below we describe how to start this project
### Installing

You must download and install `Go`, follow [this instruction](https://golang.org/doc/install) to install.

After Golang installed, Follow this instructions
```bash
$ git clone https://github.com/yogawahyudi7/POKEMON-FIGHT.git
$ go run main.go
```

Go to `http://localhost:7000/` to [start this application.](http://localhost:7000/)

## Authors

- [@yogawahyudi7](https://github.com/yogawahyudi7) - Developer

