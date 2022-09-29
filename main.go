package main

import (
	"unconvicted/db"
	"unconvicted/logger"
	"unconvicted/routes"
	"unconvicted/utils"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	//go jobs.StartJobs()
	routes.RunAllRoutes()
}
