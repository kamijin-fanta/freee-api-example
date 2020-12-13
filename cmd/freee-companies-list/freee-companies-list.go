package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kamijin-fanta/freee-api-example/pkg/freee"
)

func main() {
	accessToken := flag.String("access-token", "", "freee access token")
	flag.Parse()

	apiConfig := freee.NewConfiguration()
	client := freee.NewAPIClient(apiConfig)

	ctx := context.TODO()
	ctx = context.WithValue(ctx, freee.ContextAccessToken, *accessToken)

	res, _, err := client.CompaniesApi.GetCompanies(ctx).Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("get %d company\n", len(res.Companies))
	for _, company := range res.Companies {
		name := ""
		if company.DisplayName.IsSet() {
			name = *company.DisplayName.Get()
		}
		fmt.Printf("  id=%d name=%s role=%s\n", company.Id, name, company.Role)
	}
}
