package controllers

import (
	"context"
	"github.com/bostigger/restaurant-management-api/database"
	"github.com/bostigger/restaurant-management-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var tablesCollection = database.GetCollection(database.Client, "table")

func CreateTable(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var newTable models.Table

	newTable.ID = primitive.NewObjectID()
	newTable.TableId = newTable.ID.Hex()
	newTable.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newTable.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err := c.BindJSON(&newTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	validateErr := validate.Struct(newTable)
	if validateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validateErr.Error()})
		return
	}
	table, err := tablesCollection.InsertOne(ctx, &newTable)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, table)

}

func GetTable(c *gin.Context) {
	tableId := c.Param("table_id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var table models.Table

	result := tablesCollection.FindOne(ctx, bson.M{"tableId": tableId})
	err := result.Decode(&table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)

}

func UpdateTable(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	tableId := c.Param("table_id")

	var newTableValues models.Table
	err := c.BindJSON(&newTableValues)
	if err != nil {
		return
	}
	updatedFields := bson.M{}
	newChangesValues := reflect.ValueOf(newTableValues)
	for i := 0; i < newChangesValues.NumField(); i++ {
		field := newChangesValues.Type().Field(i)
		fieldValue := newChangesValues.Field(i)

		bsonTag := field.Tag.Get("bson")
		bsonFieldName := strings.Split(bsonTag, ",")[0]
		if bsonFieldName == "-" || bsonFieldName == "_" || bsonFieldName == "" {
			continue
		}
		if !fieldValue.IsValid() || reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			continue
		}
		updatedFields[bsonFieldName] = fieldValue.Interface()

	}
	opts := options2.FindOneAndUpdate().SetReturnDocument(options2.After)

	res := tablesCollection.FindOneAndUpdate(ctx, bson.M{"tableId": tableId}, bson.M{"$set": updatedFields}, opts)
	var newTable models.Table
	err = res.Decode(&newTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTable)
}
