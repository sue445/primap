package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/server/prismdb"
	"github.com/sue445/primap/server/testutil"
	"testing"

	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func Test_publishShop(t *testing.T) {
	srv := pstest.NewServer()
	defer srv.Close()

	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	if !assert.NoError(t, err) {
		return
	}

	defer conn.Close()

	ctx := context.Background()
	projectID := testutil.TestProjectID()
	client, err := pubsub.NewClient(ctx, projectID, option.WithGRPCConn(conn))
	if !assert.NoError(t, err) {
		return
	}

	_, err = client.CreateTopic(ctx, topicID)
	if !assert.NoError(t, err) {
		return
	}

	shop := &prismdb.Shop{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	err = publishShop(ctx, client, shop)

	assert.NoError(t, err)
}
