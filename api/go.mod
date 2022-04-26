module github.com/NebuDev14/Diseasy-Peasy

go 1.16

replace github.com/NebuDev14/Diseasy-Peasy/lib => ../lib

replace github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db => ../lib/prisma/db

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/prisma/prisma-client-go v0.14.0 // indirect
)
