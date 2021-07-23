package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vitorath/codebank/infrastructure/grpc/server"
	"github.com/vitorath/codebank/infrastructure/kafka"
	"github.com/vitorath/codebank/infrastructure/repository"
	"github.com/vitorath/codebank/usecase"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDB()
	defer db.Close()

	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionalUseCase(db, producer)
	serveGRPC(processTransactionUseCase)
}

func setupTransactionalUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	usecase := usecase.NewUseCaseTransaction(transactionRepository)
	usecase.KafkaProducer = producer
	return usecase
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("oassword"),
		os.Getenv("dbname"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}

func serveGRPC(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Running gRPC Server")
	grpcServer.Serve()
}

// Create Credit Card
// cc := domain.NewCreditCard()
// cc.Number = "1234"
// cc.Name = "Vitor"
// cc.ExpirationMonth = 7
// cc.ExpirationYear = 2021
// cc.CVV = 123
// cc.Limit = 1000
// cc.Balance = 0

// repo := repository.NewTransactionRepositoryDB(db)
// err := repo.CreateCreditCard(*cc)
// if err != nil {
// 	fmt.Println(err)
// }
