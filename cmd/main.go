package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	db "database/internal/database"
	cmpt "database/internal/database/compute"
	strg "database/internal/database/storage"
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
		logger.Fatal("cannot start database", zap.Error(err))
	}

	logger.Info("database app started")

Loop:
	for {
		fmt.Print("enter a command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			logger.Fatal("scanner error")
		}
		input := scanner.Text()

		switch input {
		case "stop":
			logger.Info("database app terminated")
			break Loop
		case "":
			fmt.Println("Please enter something !!!")
			continue
		}

		resp, err := database.HandleQuery(ctx, input)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp)
	}
}
