package controllers

import (
	"context"
	"github.com/bostigger/restaurant-management-api/helpers"
	"github.com/bostigger/restaurant-management-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func CreateMenu(c *gin.Context) {
	err := helpers.CheckUserAccess(c, "admin")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You cant make this operation"})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var newMenu models.Menu
	err = c.BindJSON(&newMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMenu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newMenu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newMenu.ID = primitive.NewObjectID()
	newMenu.MenuId = newMenu.ID.Hex()

	validateErr := validate.Struct(newMenu)
	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr.Error()})
		return
	}
	res, err := menuCollections.InsertOne(ctx, newMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func GetMenus(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var menus []bson.M

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 1
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
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
			{"data", bson.D{
				{"$push", "$$ROOT"},
			}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_times", bson.D{
				{"slice", []interface{}{"$data", startIndex, recordPerPage}},
			}},
		}},
	}
	pipeline := mongo.Pipeline{matchStage, groupStage, projectStage}

	cursor, err := menuCollections.Aggregate(ctx, pipeline)
	err = cursor.All(ctx, &menus)
	if err != nil {
		return
	}

	if len(menus) < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No menus found"})
		return
	}
	c.JSON(http.StatusOK, menus[0])
}

func GetMenu(c *gin.Context) {
	var menuId = c.Param("menu_id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var menu models.Menu
	var err = menuCollections.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No menu with id found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)

}

func UpdateMenu(c *gin.Context) {
	menuId := c.Param("menu_id")
	if menuId == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Id"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var menu models.Menu

	err := c.BindJSON(&menu)
	if err != nil {
		return
	}

	updatedFields := bson.M{}
	newMenuValues := reflect.ValueOf(menu)

	for i := 0; i < newMenuValues.NumField(); i++ {
		field := newMenuValues.Type().Field(i)
		newFieldFieldValue := newMenuValues.Field(i)

		bsonTag := field.Tag.Get("bson")
		bsonTagFieldName := strings.Split(bsonTag, ",")[0]
		if bsonTagFieldName == "-" || bsonTagFieldName == "_id" || bsonTagFieldName == "" {
			continue
		}
		if !newFieldFieldValue.IsValid() || reflect.DeepEqual(newFieldFieldValue.Interface(), reflect.Zero(newFieldFieldValue.Type()).Interface()) {
			continue
		}
		updatedFields[bsonTagFieldName] = newFieldFieldValue.Interface()

	}
	opts := options2.FindOneAndUpdate().SetReturnDocument(options2.After)

	var updatedMenu bson.M
	err = menuCollections.FindOneAndUpdate(ctx, bson.M{"menuId": menuId}, bson.M{"$set": updatedFields}, opts).Decode(&updatedMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMenu)

}

func DeleteMenu(c *gin.Context) {
	menuId := c.Param("menu_id")
	if menuId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No menu Id was passed"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	res := menuCollections.FindOneAndDelete(ctx, bson.M{"menuId": menuId})

	if res.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Err()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
