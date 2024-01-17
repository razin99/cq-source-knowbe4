package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-knowbe4/client"
)

type AccountRiskScoreHistory struct {
	RiskScore float64   `json:"risk_score"`
	Date      time.Time `json:"date"`
}

func AccountRiskScoreHistoryTable() *schema.Table {
	return &schema.Table{
		Name:      "knowbe4_account_risk_score_history",
		Resolver:  fetchAccountRiskScoreHistory,
		Transform: transformers.TransformWithStruct(&AccountRiskScoreHistory{}, transformers.WithPrimaryKeys("Date")),
	}
}

func fetchAccountRiskScoreHistory(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var page uint64 = 1
	path := "/v1/account/risk_score_history"
	query := url.Values{}
	query.Set("full", "true")

	for {
		query.Set("page", fmt.Sprint(page))
		response, err := cl.KnowBe4.NewRequest(ctx, path, query)
		if err != nil {
			return err
		}

		var riskScoreHistory []AccountRiskScoreHistory
		err = func() error {
			defer response.Body.Close()
			return json.NewDecoder(response.Body).Decode(&riskScoreHistory)
		}()
		if err != nil {
			return err
		}

		if len(riskScoreHistory) == 0 {
			break
		}

		for _, enrollment := range riskScoreHistory {
			res <- enrollment
		}

		page++
	}

	return nil
}
