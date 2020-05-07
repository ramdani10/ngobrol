# Ngobrol

Ngobrol yuk!

# Getting Started

## Requirements

- [Golang](https://golang.org/)
- [Postgresql 11.6](https://www.postgresql.org/)

## Installation

Clone the repository

    git clone git@github.com:imansuparman/ngobrol.git

Switch to the repo folder

    cd ngobrol
    
Create file configuration

    cd params
    cp ngobrol.toml.sample ngobrol.toml
    
Change database configuration (`params/ngobrol.toml`)

    [postgre]
    is_enabled = true
    is_migration_enable = true
    host = "your_db_host"
    port = 5432
    name = "ngobrol"
    username = "your_username"
    password = "your_password"
    pool_size = 5
    
Install all the dependencies using composer

    go build

Run Migration
    
    ./ngobrol migrate
    
Start the local development server

    ./ngobrol

You can now access the server at http://localhost:4500

# How to Use

1. Get list messages
    
    ```
    curl --location --request GET 'http://127.0.0.1:4500/api/v1/messages' --header 'Content-Type: application/json'
   ```
    
2. Create a new message

    ```
     curl --location --request POST 'http://127.0.0.1:4500/api/v1/messages' \
     --header 'Content-Type: application/json' \
     --header 'Content-Type: text/plain' \
     --data-raw '{
     	"message": "Hello World"
     }'
   ```

# Documentation
  
    vim api/ngobrol.yaml
    
# Contributing

1. Clone it!
2. Create your feature branch: `git flow feature start my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git flow feature publish my-new-feature`
5. Submit a pull request 🙂