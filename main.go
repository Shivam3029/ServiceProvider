package main

import (
	   "github.com/gin-gonic/gin"
	   "net/http"
	   "strconv"
	   //"reflect"
	   "log"
)
type ServiceProvider struct{
	ID int `json:"id"`
	Name string `json:"name"`
}


var serviceproviders []ServiceProvider
var sp1 []ServiceProvider

func main() {
router := gin.Default()
serviceproviders=append(serviceproviders, ServiceProvider{ID: 1,Name :"SP1"},
						ServiceProvider{ID:2,Name:"SP2"},
						ServiceProvider{ID:3,Name:"SP3"},
						ServiceProvider{ID:4,Name:"SP4"},
						ServiceProvider{ID:5,Name:"SP5"})
v1 := router.Group("/serviceproviders")
 {
  v1.POST("/", addServiceProvider)
  v1.GET("/", getServiceProviders)
  v1.GET("/:id", getServiceProvider)
  v1.PATCH("/", updateServiceProvider)
  v1.DELETE("/:id", deleteServiceProvider)
 }
 router.Run()
}
func getServiceProviders(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": serviceproviders})
}

func getServiceProvider(c *gin.Context){
	//var serviceprovider ServiceProvider
	serviceproviderID := c.Param("id")
	//log.Println(reflect.TypeOf(serviceproviderID))
	i,_:=strconv.Atoi(serviceproviderID)
	for _,serviceprovider:=range serviceproviders{
		if serviceprovider.ID==i{
			c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceprovider})

		}
	}
	//i,_:=strconv.Atoi(serviceproviderID)
	//log.Println(reflect.TypeOf(i))
	
}
func addServiceProvider(c *gin.Context) {
	var serviceprovider ServiceProvider
	c.BindJSON(&serviceprovider)
	log.Println(serviceprovider)
	//serviceproviders=append(serviceproviders,serviceprovider)
	for _,item:=range serviceproviders{
		if item.ID==serviceprovider.ID{
			c.JSON(409,gin.H{"status":409,"message":"service provider already exists"})
			return
		}
	}
	serviceproviders=append(serviceproviders,serviceprovider)
	c.JSON(http.StatusOK,gin.H{
		"id":serviceprovider.ID,
		"name":serviceprovider.Name,
	})
	//serviceproviders=append(serviceproviders,serviceprovider)
	//c.JSON(http.StatusOK,gin.H{
	//	"id":serviceprovider.ID,
	//	"name":serviceprovider.Name,
	//})
	
}
func updateServiceProvider(c *gin.Context){

	var serviceprovider ServiceProvider
	c.BindJSON(&serviceprovider)
	
	for i,item:=range serviceproviders{
		if item.ID==serviceprovider.ID{
			serviceproviders[i]=serviceprovider
		}
	}
	//serviceproviders=append(serviceproviders,serviceprovider)
	c.JSON(http.StatusOK,gin.H{
		"id":serviceprovider.ID,
		"name":serviceprovider.Name,
	})
	//c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceproviders})
}
func deleteServiceProvider(c *gin.Context){
	serviceproviderID := c.Param("id")
	id,_:=strconv.Atoi(serviceproviderID)
	//id, _:=strconv.Atoi(params["id"])
	for i,item:=range serviceproviders{
		if item.ID==id{
			serviceproviders=append(serviceproviders[:i],serviceproviders[i+1:]...)
		}
	}
	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":serviceproviders})
	//json.NewEncoder(w).Encode(serviceproviders)
	//log.Println("deleting a service provider")
}