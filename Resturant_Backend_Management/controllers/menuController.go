package controller

import "github.com/gin-gonic/gin"

var menuCollection *mongo.Collection = database.OpenCollection(database.Client , "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = Context.WithTimeout(Context.Background(), 100*time.Second)
		result, err := menuCollection.Find(Context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the menu itself"})
		}
		var allMenus []bson.m
		if err = result.ALL(ctx, &allMenus); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
