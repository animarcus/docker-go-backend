# Docker Containers for running a database and interfacing with it using Go
This repo facilitates the use of a database server running inside a docker container. There is also code to be able to connect to it using golang. 

For now, the go code for connecting to the database is running outside of the docker container for testing purposes. When ready for production one can create a go container and have everything dockerized.

I spent a lot of time figuring this out so I hope I can save someone's time because a repo like this is what I would have needed to start out.
