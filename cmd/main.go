package main

import (
	"fmt"

	"github.com/saravanasai/durableflow/internal/runtime"
)

func main() {

	fmt.Println("Hello durable flow")

	rt := runtime.NewRuntime()

	rt.RegisterWorkflow("purchase", &runtime.PurchaseWorkflow{})

	rt.RegisterActivity("debit", &runtime.DebitUser{})

	rt.Start("purchase", []any{"Saravana Sai", "1000"})
}
