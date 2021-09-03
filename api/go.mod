module github.com/NebuDev14/Diseasy-Peasy

go 1.16

replace github.com/NebuDev14/Diseasy-Peasy/lib => ../lib

replace github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db => ../lib/prisma/db

require (
	github.com/NebuDev14/Diseasy-Peasy/lib v0.0.0-00010101000000-000000000000
	github.com/NebuDev14/Diseasy-Peasy/lib/prisma/db v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.4
)
