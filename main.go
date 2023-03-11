package main

import(
	"BE-JoanaVidon/order-api/usecase"
	"BE-JoanaVidon/order-api/repository"
	"BE-JoanaVidon/order-api/controllers"
	"BE-JoanaVidon/user-api/repository/database"
	"BE-JoanaVidon/user-api/service"
	"BE-JoanaVidon/user-api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main(){
	dns := "adm:Pass123!@/store?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dns))
	if err != nil {
			log.Panic("failed to connect to database")
	}
	
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
	
	//Conectando ao Gin
	r := gin.Default()

	//Routes-User
		r.POST("/users", createHandler.Create)
		r.GET("/users/:id", getHandler.Get)
		r.GET("/users", getAllHandler.GetAll)
		r.PUT("/users/:id/:phone", updateHandler.Update)
		r.DELETE("/users/:id", deleteHandler.Delete)
	
	/* 	exemplos:
		Create - localhost:5000/users
		Get - localhost:5000/users/b9d81829-5de4-4a48-930f-9b435e2bf167
		Get All - localhost:5000/users
		Update - localhost:5000/users/339c92ea-094b-427a-8c51-ab75f18efeb7/4444-2222
		Delete - localhost:5000/users/35bba9db-0e85-4128-ab69-a313971f1d45
	 */

	//Order
	 mySQLrepositoryOrder := repository.NewRepository(DB, err)
	 createOrder := usecase.NewCreateUseCase(mySQLrepositoryOrder)
	 createHandlerOrder := controllers.NewCreateHandler(createOrder)
	 
	 r.POST("/orders/:userID", createHandlerOrder.Create)
		//localhost:5000/orders/94654281-d3ff-4f72-b329-91c794adb22a


	r.Run(":5000")
}