package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context) {

}

func GetEntryById(c *gin.Context) {

}

// 2
func GetEntries(c *gin.Context) {
	// 如果在 100 秒內操作沒有完成，上下文對象將被取消，從而通知程序停止該操作。
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)
}

func GetEntriesByIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}

func UpdateIngredient(c *gin.Context) {

}

// 1 most easily
func DeleteEntry(c *gin.Context) {
	// 透過 .ByName() 方法，它返回指定名稱的參數的值,並儲存在entryID變數中
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	// 超時處理
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// 從名為 entryCollection 的 MongoDB 集合中，刪除 _id 欄位等於 docID 變數的文件。
	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
