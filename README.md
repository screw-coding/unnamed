# unnamed
[![Bors enabled](https://bors.tech/images/badge_small.svg)](https://app.bors.tech/repositories/46302)



## Description

## Project Structure

1. It based on [project-layout](https://github.com/golang-standards/project-layout)
2. This project use **go work** mode to manage different modules


## Project Architecture

With Jetbrains products,you can search markdown in settings,then install mermaid extension,then you can see the following graph
```mermaid
flowchart LR
    app-load-balancer-->app-web-server-->app-web
     
    app-web-->app-relation-database
    app-web-->app-cache
    app-web-->app-recommend-engine
    app-web-->app-message-queue
    app-web-->app-full-text-search-engine
   
    app-relation-database-.->comp-tcp-client-server
    app-relation-database-.->comp-tokenizer
    app-relation-database-.->comp-file-store
    app-relation-database-.->comp-filter
    app-relation-database-.->comp-cache
    
    app-recommend-engine-.->comp-tokenizer
    app-recommend-engine-.->comp-file-store
    app-recommend-engine-.->comp-http-client-server
    app-recommend-engine-.->comp-filter
    
    app-cache-.->comp-tcp-client-server
    app-cache-.->comp-cache
    app-cache-.->comp-file-store
    
    app-full-text-search-engine-.->comp-http-client-server
    app-full-text-search-engine-.->comp-file-store
    app-full-text-search-engine-.->comp-filter
    app-full-text-search-engine-.->comp-tokenizer
    
    app-message-queue-.->comp-tcp-client-server
    app-message-queue-.->comp-file-store
    app-message-queue-.->comp-filter
    app-message-queue-.->algo-delay-message
    
    comp-file-store
    comp-http-client-server
    comp-tcp-client-server
    comp-file-cache
    comp-memory-cache
    comp-tokenizer
    comp-filter
    
    
    algo-b-tree-index
    algo-hash-index
    algo-inverted-index
    algo-bitcask
    algo-delay-message
    
    algo-lru
    algo-wal
    
    algo-filter-stop-word
    algo-filter-stemmer
    algo-filter-collaborative
    algo-filter-lowercase
    
    
    comp-http-client-server-.->comp-tcp-client-server
    comp-http-client-server-.->comp-file-cache
    
    comp-memory-cache-.->algo-lru
    comp-memory-cache-.->algo-hash-index
    
    comp-file-store-.->algo-bitcask
    comp-file-store-.->algo-inverted-index
    comp-file-store-.->algo-b-tree-index
    comp-file-store-.->algo-wal
    comp-file-store-.->algo-hash-index
    comp-file-store-.->algo-lsm-tree
    
    comp-filter-.->algo-filter-stop-word
    comp-filter-.->algo-filter-stemmer
    comp-filter-.->algo-filter-collaborative
    comp-filter-.->algo-filter-lowercase
    
    comp-cache-.->comp-memory-cache
    comp-cache-.->comp-file-store
    

```

