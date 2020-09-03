# Api
## Routes
### Users
        POST:   users                   🡺 create a user 
        PUT:    users/by-id/{id}        🡺 modify a user 
        PUT:    users/by-name/{id}      🡺 modify a user by username   
        GET:    users                   🡺 get all users 
        GET:    users/by-id/{id}        🡺 get user info by id 
        GET:    users/by-name/{name}    🡺 get user info by username 
        DELETE: users/by-id/{id}        🡺 hide user in db (soft del) 

### Worlds
        POST:   worlds                  🡺 create a world
        PUT:    worlds/{name}           🡺 modify a world
        GET:    worlds                  🡺 get world list by active users
        GET:    worlds/popular          🡺 get world list by all-time popularity
        GET:    worlds/by/{name}        🡺 get world list the creator's name
        GET:    worlds/{name}           🡺 get a world
        DELETE: worlds/{name}           🡺 hide a world in db (soft del) 

## Technology
        Lang:   Go 1.15
        DB:     Gorm, MySQL
        Routes: Gorilla mux router

