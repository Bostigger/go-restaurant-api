package controllers

import (
	"context"
	"github.com/bostigger/restaurant-management-api/database"
	"github.com/bostigger/restaurant-management-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
