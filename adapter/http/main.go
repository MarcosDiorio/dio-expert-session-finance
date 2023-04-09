package http

import (
	"net/http"

	"github.com/MarcosDiorio/dio-expert-session-finance/adapter/http/actuator"
	"github.com/MarcosDiorio/dio-expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() error {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(":8080", nil)
}
