package routes

import (
	"github.com/gin-gonic/gin"
	"restuarant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:Invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices/create", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/update", controllers.UpdateInvoice())
}