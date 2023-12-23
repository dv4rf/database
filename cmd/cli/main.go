package main

import (
	"bufio"
	"context"
	db "database/internal/database"
	cmpt "database/internal/database/compute"
	strg "database/internal/database/storage"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	parser, err := cmpt.NewParser(logger)
	analyzer, err := cmpt.NewAnalyzer(logger)
	storage, err := strg.NewStorage(logger)

	compute, err := cmpt.NewCompute(parser, analyzer, logger)

	database, err := db.NewDatabase(compute, storage, logger)
	if err != nil {
		logger.Fatal("failed to start database", zap.Error(err))
	}

	logger.Info("client app started")

	for {
		fmt.Print("enter a command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			logger.Fatal("scanner error")
		}
		input := scanner.Text()

		if input == "stop" {
			logger.Info("client app terminated")
			break
		}

		handledQuery := database.HandleQuery(ctx, input)
		fmt.Println(handledQuery)
	}
}
