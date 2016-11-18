package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/kapmahc/lotus/web"
)

//Mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	//guest
	//member
	// admin
	ag := rt.Group("/shop",
		p.Jwt.CurrentUserHandler(true),
		p.Jwt.MustAdminHandler(),
	)

	ag.GET("/countries", web.JSON(p.countriesIndex))
	ag.POST("/countries", web.JSON(p.countriesCreate))
	ag.GET("/countries/:id", web.JSON(p.countriesShow))
	ag.POST("/countries/:id",
		p.Jwt.CurrentUserHandler(true), p.Jwt.MustAdminHandler(),
		web.JSON(p.countriesUpdate),
	)
	ag.DELETE("/countries/:id", web.JSON(p.countriesDestroy))

	ag.GET("/states", web.JSON(p.statesIndex))
	ag.POST("/states", web.JSON(p.statesCreate))
	ag.GET("/states/:id", web.JSON(p.statesShow))
	ag.POST("/states/:id", web.JSON(p.statesUpdate))
	ag.DELETE("/states/:id", web.JSON(p.statesDestroy))

	ag.GET("/payment-methods", web.JSON(p.paymentMethodsIndex))
	ag.POST("/payment-methods", web.JSON(p.paymentMethodsCreate))
	ag.GET("/payment-methods/:id", web.JSON(p.paymentMethodsShow))
	ag.POST("/payment-methods/:id", web.JSON(p.paymentMethodsUpdate))
	ag.DELETE("/payment-methods/:id", web.JSON(p.paymentMethodsDestroy))

	ag.GET("/shipping-methods", web.JSON(p.shippingMethodsIndex))
	ag.POST("/shipping-methods", web.JSON(p.shippingMethodsCreate))
	ag.GET("/shipping-methods/:id", web.JSON(p.shippingMethodsShow))
	ag.POST("/shipping-methods/:id", web.JSON(p.shippingMethodsUpdate))
	ag.DELETE("/shipping-methods/:id", web.JSON(p.shippingMethodsDestroy))
}
