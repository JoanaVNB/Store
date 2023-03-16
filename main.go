package main

import (
	"BE-JoanaVidon/order-api/controllers"
	"BE-JoanaVidon/order-api/repository"
	"BE-JoanaVidon/order-api/usecase"
	"BE-JoanaVidon/order-api/domain"
	domainU "BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/handlers"
	"BE-JoanaVidon/user-api/repository/database"
	"BE-JoanaVidon/user-api/repository/elasticSearch"
	"BE-JoanaVidon/user-api/service"
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main(){
	dns := "root:secret@tcp(localhost:3306)/store?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
			log.Panic("failed to connect to database")
	}


	DB.AutoMigrate(&domainU.User{})
	DB.AutoMigrate(&domain.Order{})
	

	mySQLrepository := database.NewRepository(DB, err)

	createUC := service.NewCreateUseCase(mySQLrepository)
	createHandler := handlers.NewCreateHandler(createUC)

	getUC := service.NewGetUseCase(mySQLrepository)
	getHandler := handlers.NewGetHandler(getUC)

	getAllUC := service.NewGetAllUseCase(mySQLrepository)
	getAllHandler := handlers.NewGetAllHandler(getAllUC)

	updateUC := service.NewUpdateUseCase(mySQLrepository)
	updateHandler := handlers.NewUpdateHandler(updateUC)

	deleteUC := service.NewDeleteUseCase(mySQLrepository)
	deleteHandler := handlers.NewDeleteHandler(deleteUC)
	
//Elastic Search
	ctx := context.Background()

	ctx = elasticSearch.LoadUsersFromFile(ctx)
	ctx = elasticSearch.ConnectionWithElasticSearch(ctx)
	elasticSearch.IndexUsersAsDocuments(ctx)
	elasticSearch.QueryUserByDocumentID(ctx)
	elasticSearch.QueryUsersByEmail(ctx, "unique@gmail.com")

	r := gin.Default()

	//Routes-User
		r.POST("/users", createHandler.Create)
		r.GET("/users/:id", getHandler.Get)
		r.GET("/users", getAllHandler.GetAll)
		r.PUT("/users/:id/:phone", updateHandler.Update)
		r.DELETE("/users/:id", deleteHandler.Delete)
	
	//Order
	 mySQLrepositoryOrder := repository.NewRepository(DB, err)
	 createOrder := usecase.NewCreateUseCase(mySQLrepositoryOrder)
	 createHandlerOrder := controllers.NewCreateHandler(createOrder)
	 
	 r.POST("/orders/:userID", createHandlerOrder.Create)
	
	 r.Run(":5000")
}

	/* 	exemplos:
	Users:
		Create - localhost:5000/users
		Get - localhost:5000/users/b9d81829-5de4-4a48-930f-9b435e2bf167
		Get All - localhost:5000/users
		Update - localhost:5000/users/339c92ea-094b-427a-8c51-ab75f18efeb7/4444-2222
		Delete - localhost:5000/users/35bba9db-0e85-4128-ab69-a313971f1d45
	Orders:
		POST - localhost:5000/orders/94654281-d3ff-4f72-b329-91c794adb22a
	 */