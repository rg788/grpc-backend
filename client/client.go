package main

import (
	pb "grpc-backend/gen/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type PORTINPUT struct {
	ID      int64  `json:"id" binding:"required"`
	NAME    string `json:"name" binding:"required"`
	CODE    string `json:"code" binding:"required"`
	CITY    string `json:"city" binding:"required"`
	STATE   string `json:"state" binding:"required"`
	COUNTRY string `json:"country" binding:"required"`
}

func main() {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cc := pb.NewPortServiceClient(conn)

	endPoints(cc)

}

func endPoints(cc pb.PortServiceClient) {

	g := gin.Default()

	//Creating

	g.POST("/v1/ports", func(ctx *gin.Context) {

		var input PORTINPUT
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		req := &pb.CreatePortRequest{
			Port: &pb.Port{
				Id: input.ID, Name: input.NAME, Code: input.CODE,
				City: input.CITY, State: input.STATE, Country: input.COUNTRY,
			},
		}

		if response, err := cc.CreatePort(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Create": response,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

	})
	//Updating
	g.PUT("/v1/ports/:id", func(ctx *gin.Context) {

		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		var input PORTINPUT
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &pb.UpdatePortRequest{
			Port: &pb.Port{
				Id: id, Name: input.NAME, Code: input.CODE,
				City: input.CITY, State: input.STATE, Country: input.COUNTRY,
			},
		}

		if response, err := cc.UpdatePort(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Result": response,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

	})
	//Retreive
	g.GET("/v1/ports/:id", func(ctx *gin.Context) {

		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		var input PORTINPUT
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &pb.RetrievePortRequest{PortId: id}

		if response, err := cc.RetreivePort(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"id":      response.Id,
				"name":    response.Name,
				"code":    response.Code,
				"city":    response.City,
				"state":   response.State,
				"country": response.Country,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

	})

	// Deletion
	g.DELETE("/v1/ports/:id", func(ctx *gin.Context) {

		id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
		var input PORTINPUT
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &pb.DeletePortResquest{
			PortId: id,
		}

		if response, err := cc.DeletePort(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Result": response,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

	})
	//Pagination

	g.GET("/v1/ports", func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
		count, _ := strconv.Atoi(ctx.DefaultQuery("count", "10"))

		var input PORTINPUT
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		req := &pb.ListPortRequest{
			Page:  int32(page),
			Count: int32(count),
		}
		if _, err := cc.ListPort(ctx, req); err == nil {
			/* ctx.JSON(http.StatusOK, gin.H{
				"id":      response.Id,
				"name":    response.Name,
				"code":    response.Code,
				"city":    response.City,
				"state":   response.State,
				"country": response.Country,
			}) */
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})


			//receving streams
			// rstream,_ := cc.ListPort(ctx, req)
			// rstream.Recv()



		}

	})

	if err := g.Run(":5050"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
