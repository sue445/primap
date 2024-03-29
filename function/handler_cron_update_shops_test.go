package primap

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/primap/prishops"
	"github.com/sue445/primap/testutil"
	"os"
	"testing"

	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func Test_getAndPublishShops(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://prismdb.takanakahiko.me/sparql",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/all_shops.json")))

	ctx := context.Background()
	projectID := os.Getenv("PUBSUB_PROJECT_ID")

	client, err := pubsub.NewClient(ctx, projectID)
	if !assert.NoError(t, err) {
		return
	}

	_, err = client.CreateTopic(ctx, topicID)
	if !assert.NoError(t, err) {
		return
	}

	err = getAndPublishShops(ctx, projectID)
	assert.NoError(t, err)
}

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

	shop := &prishops.Shop{
		Name:       "ＭＥＧＡドン・キホーテＵＮＹ名張",
		Prefecture: "三重県",
		Address:    "三重県名張市下比奈知黒田3100番地の1",
		Series:     []string{"prichan"},
	}

	topic := client.Topic(topicID)
	err = publishShop(ctx, topic, shop)

	assert.NoError(t, err)
}
