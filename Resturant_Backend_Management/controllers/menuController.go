package controller

import ("github.com/gin-gonic/gin"
		"Resturant_Manager_Go/database"
		"Resturant_Manager_Go/models"
		"context"
		"fmt"
		"net/http"
		"time"

		"github.com/gin-gonic/gin"
		"github.com/go-playground/validator/v10"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/bson/primitive"
		"go.mongodb.org/mongo-driver/mongo")

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
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := foodCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})
		}
		c.JSON(http.StatusOK, food)
	}

	}


func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(menu)
		if validationErr!= nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		 }

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.hex()


		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr!=nil{
			msg := fmt.Sprintf("menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu model.menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id":menuId}


		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date !=nil{
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now())
			msg := "kindly retype the date time"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			defer cancel()
			return
	
		}

		updateObj = append(updateObj,bson.E{"start_date", menu.start_date})
		updateObj = append(updateObj, bson.E{"end_date", menu.end_date})

		if menu.Name != ""{
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}

		if menu.Category != ""{
			updateObj = append(updateObj, bson.E{"name", menu.menu.Category})
		}

		menu.Updated_at, _ =time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		upsert:=true
	
		opt := options.UpdateOptions{
			upsert : &upsert,

		}

		result,err := menu.Collection.UpdateOne(
			ctx, 
			filter,
			bson.D{
				{"$set" , updateObj},
			}, 
			&opt,


			)
			
			if err!= nil {
				msg:= "menu update failed"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			}

			defer cancel()
			msg := "menu update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

			defer cancel()
			c.JSON(http.StatusOK, result)
	}		
}
