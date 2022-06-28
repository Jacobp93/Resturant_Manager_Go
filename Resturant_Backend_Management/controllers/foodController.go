package controller

import (
	"Resturant_Manager_Go/database"
	"Resturant_Manager_Go/models"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page-1) = recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{({"$match" , bson.D({})})}
		groupStage := bson.D{{"$group" , bson.{{"id", "null"}}} , {"total_count", bson.D{{"$sum , 1"}}},{"data", bson.D{{"spush" ,"$$ROOT"}}} }}}
		ProjectStage := bson.D{

		{"$project" , bson.D}{
			{"_id" , 0}
			{"total_count" , 1}
			{"food_items" ,bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerpage}}}},
		}

		}
	}

	result, err := foodcollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupstage, projectStage

	})
	defer cancel()
	if err != nil{
		c.JSON{http.StatusInternalServerError, gin.h("error: error occured while listing foood items ")}
	}
	var allfoods []bson.H
	result.All(ctix, &allFoods); err != nil{
		log.fatal(err)
	}
	c.JSON(http.StatusOK, allFoods[0])
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}


func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(contect.Background(), 100*time.Seconds)
		var menu models.Menu
		var food models.Food

		foodId := c.Param("food_id")

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D 

		if food.Name != nil {
			updateObj = append(updateObj,bson.E{"name", food.Name} )
		}

		if food.Price != nil{
			updateObj = append(updateObj, bson.E{"price", food.Price})
		}

		if food.Food_image !nil{
			updateObj = append(updateObj, bson.E{"food_Image", food.image})
		}

		if food.Menu_id != nil{
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food_Menu_id}).Decode(&menu)
			defer cancel()
			if err !=nil{
				msg:= fmt.Sprintf("message was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", food.Price})
		}

		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", food.Updated_at})
	
		upsert := true
		filter := bson.M{"food_id:" foodID}

		opt := option.UpdateOptions{
			upsert: &upsert,
		}
	
		foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj}
			}, 
			&opt,
			)

			if err!=nil {
				msg := fmt.Sprint("food item update failed ")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			c.JSON(http.StatusOK, result)
	}
}