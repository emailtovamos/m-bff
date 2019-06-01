package main

import (
	"flag"
	"github.com/emailtovamos/m-bff/bff"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	grpcAddressHighScore := flag.String("address-m-highscore", "localhost:50051", "The grpc server address for highscore service")
	grpcAddressGameEngine := flag.String("address-m-game-engine", "localhost:60051", "The grpc server address for game-engine service")

	serverAddress := flag.String("address-http", ":8081", "HTTP server address")

	flag.Parse()

	gameClient, err := bff.NewGrpcGameServiceClient(*grpcAddressHighScore)
	if err != nil {
		log.Error().Err(err).Msg("Error in creating a client for m-highscore")
	}

	gameEngineClient, err := bff.NewGrpcGameEngineServiceClient(*grpcAddressGameEngine)
	if err != nil {
		log.Error().Err(err).Msg("Error in creating a client for m-game-engine")
	}

	gr := bff.NewGameResource(gameClient, gameEngineClient)

	router := gin.Default()
	router.GET("/geths", gr.GetHighScore)
	router.GET("/seths/:hs", gr.SetHighScore)
	router.GET("/getsize", gr.GetSize)
	router.GET("/setscore/:score", gr.SetScore)

	err = router.Run(*serverAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start bff")
	}

	log.Info().Msgf("Started http-server at %v", *serverAddress)

}
