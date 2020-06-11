package data

import (
	"context"
	"fmt"
	"jobBoardApi/models"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SavePosting - saves job posting
func SavePosting(c echo.Context) error {
	job1 := models.Job{JobTitle: "Job1", Company: "Company1", Location: "Location1", Remote: "Remote1", Job: "Job1"}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := GetCollection().InsertOne(ctx, job1)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusCreated, insertResult)
}

// GetPostings - get all job postings
func GetPostings(c echo.Context) error {
	var results []*models.Job

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := GetCollection().Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		var element models.Job
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &element)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	cursor.Close(ctx)

	return c.JSON(http.StatusCreated, results)
}

// UpdatePosting - updates a single posting
func UpdatePosting(c echo.Context) error {
	requestData := new(models.Job)
	if err := c.Bind(requestData); err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", id)
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{
		"$set": bson.M{
			"jobTitle": requestData.JobTitle,
			"company":  requestData.Company,
			"location": requestData.Location,
			"remote":   requestData.Remote,
			"job":      requestData.Job,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := GetCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return c.JSON(http.StatusCreated, result)
}

// DeletePosting - deletes a single posting
func DeletePosting(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", id)
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := GetCollection().DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	return c.JSON(http.StatusCreated, result)
}
