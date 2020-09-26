package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sue445/primap/server/config"
	"github.com/sue445/primap/server/db"
	"github.com/sue445/primap/server/prismdb"
	"io/ioutil"
	"net/http"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// QueueSaveShopHandler returns handler of /job/queue/save_shop
func QueueSaveShopHandler(w http.ResponseWriter, r *http.Request) {
	err := queueSaveShopHandler(r)

	if err != nil {
		handleError(w, err)
		return
	}

	fmt.Fprint(w, "ok")
}

func queueSaveShopHandler(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return errors.WithStack(err)
	}

	var m PubSubMessage
	err = json.Unmarshal(body, &m)

	if err != nil {
		return errors.WithStack(err)
	}

	var shop prismdb.Shop
	err = json.Unmarshal(m.Message.Data, &shop)

	if err != nil {
		return errors.WithStack(err)
	}

	return saveShop(r.Context(), config.GetProjectID(), &shop)
}

func saveShop(ctx context.Context, projectID string, shop *prismdb.Shop) error {
	dao := db.NewShopDao(projectID)

	entity, err := dao.LoadOrCreateShop(shop.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	entity.Prefecture = shop.Prefecture
	entity.Series = shop.Series

	err = entity.UpdateAddressWithLocation(ctx, shop.Address)
	if err != nil {
		return errors.WithStack(err)
	}

	err = dao.SaveShop(entity)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
