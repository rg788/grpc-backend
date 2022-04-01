package main

import (
	pb "grpc-backend/gen/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type PORTINPUT struct {
	ID      string `json:"id" binding:"required"`
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

	xyz(cc)

}

func xyz(cc pb.PortServiceClient) {

	g := gin.Default()
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

	if err := g.Run(":5050"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	g.Run(":5050")
}
