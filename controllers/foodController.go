package controllers

import (
	"context"
	"fmt"
	"github.com/bostigger/restaurant-management-api/database"
	"github.com/bostigger/restaurant-management-api/helpers"
	"github.com/bostigger/restaurant-management-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var foodCollections *mongo.Collection = database.GetCollection(database.Client, "food")
var menuCollections *mongo.Collection = database.GetCollection(database.Client, "menu")

var validate = validator.New()

func CreateFood(c *gin.Context) {
	log.Printf("%v", c)
	err := helpers.CheckUserAccess(c, "admin")
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var menu models.Menu
	var newFood models.Food

	err = c.BindJSON(&newFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var validateErr = validate.Struct(newFood)
	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr.Error()})
		return
	}
	newFood.ID = primitive.NewObjectID()
	newFood.FoodId = newFood.ID.Hex()
	newFood.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newFood.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	checkMenu := menuCollections.FindOne(ctx, bson.M{"menuId": *newFood.MenuId})
	err = checkMenu.Decode(&menu)
	if err != nil {
		msg := fmt.Sprintf("No menu found")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": msg})
		return
	}

	result, err := foodCollections.InsertOne(ctx, newFood)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, result)

}

func GetFood(c *gin.Context) {
	foodId := c.Params.ByName("food_id")

	//foodId, err := primitive.ObjectIDFromHex(food)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	fmt.Println(foodId)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var gottenfood models.Food

	res := foodCollections.FindOne(ctx, bson.M{"foodId": foodId})
	err := res.Decode(&gottenfood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gottenfood)
}

func UpdateFood(c *gin.Context) {
	foodId := c.Params.ByName("food_id")
	//foodId, err := primitive.ObjectIDFromHex(fId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var newFoodInfo models.Food

	err := c.BindJSON(&newFoodInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateFields := bson.M{}
	newChangesValue := reflect.ValueOf(newFoodInfo)
	for i := 0; i < newChangesValue.NumField(); i++ {
		field := newChangesValue.Type().Field(i)
		newFieldValue := newChangesValue.Field(i)

		bsonTag := field.Tag.Get("bson")
		bsonFieldName := strings.Split(bsonTag, ",")[0]

		if bsonFieldName == "-" || bsonFieldName == "_id" || bsonFieldName == "" {
			continue
		}
		if !newFieldValue.IsValid() || reflect.DeepEqual(newFieldValue.Interface(), reflect.Zero(newFieldValue.Type()).Interface()) {
			continue
		}
		updateFields[bsonFieldName] = newFieldValue.Interface()
	}
	opts := options2.FindOneAndUpdate().SetReturnDocument(options2.After)

	res := foodCollections.FindOneAndUpdate(ctx, bson.M{"foodId": foodId}, bson.M{"$set": updateFields}, opts)

	var updatedFood models.Food
	err = res.Decode(&updatedFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedFood)
}

func GetAllFoods(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var foods []bson.M

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}
	page, err1 := strconv.Atoi(c.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{"$match", bson.D{}}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "null"},
			{"total_count", bson.D{
				{"$sum", 1},
			}},
			{
				"data", bson.D{
					{"$push", "$$ROOT"},
				}},
		}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, groupStage, projectStage}

	cursor, err := foodCollections.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = cursor.All(ctx, &foods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(foods) < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There is no food here"})
		return
	}
	c.JSON(http.StatusOK, foods[0])

}

func DeleteFood(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	foodId := c.Params.ByName("food_id")
	fid, err := primitive.ObjectIDFromHex(foodId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	foodCollections.FindOneAndDelete(ctx, bson.M{"_id": fid})
	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})

}

func ReturnError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
