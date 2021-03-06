package invdendpoint

import (
	"encoding/json"
	"github.com/ActiveState/invoiced-go/invdutil"
	"testing"
)

func TestUnMarshalEventObject(t *testing.T) {
	s := `{
    "id": 1228003,
    "type": "transaction.created",
    "data": {
        "object": {
            "amount": 55,
            "created_at": 1451500772,
            "currency": "usd",
            "customer": 15455,
            "date": 1451500771,
            "fee": 0,
            "gateway": null,
            "gateway_id": null,
            "id": 212047,
            "invoice": 196539,
            "metadata": [],
            "method": "other",
            "notes": null,
            "parent_transaction": null,
            "status": "succeeded",
            "theme": null,
            "type": "payment",
            "pdf_url": "https:\/\/dundermifflin.invoiced.com\/payments\/59FHO96idoXFeiBDu1y5Zggg\/pdf",
            "metadata": {}
        }
    }
}`

	so := new(Event)

	err := json.Unmarshal([]byte(s), so)

	if err != nil {
		t.Fatal(err)
	}

	if so.Id != 1228003 {
		t.Fatal("Event id is incorrect")
	}

	if so.Type != "transaction.created" {
		t.Fatal("Event type is incorrect")
	}

	object := `{
	       "object": {
	           "amount": 55,
	           "created_at": 1451500772,
	           "currency": "usd",
	           "customer": 15455,
	           "date": 1451500771,
	           "fee": 0,
	           "gateway": null,
	           "gateway_id": null,
	           "id": 212047,
	           "invoice": 196539,
	           "metadata": [],
	           "method": "other",
	           "notes": null,
	           "parent_transaction": null,
	           "status": "succeeded",
	           "theme": null,
	           "type": "payment",
	           "pdf_url": "https:\/\/dundermifflin.invoiced.com\/payments\/59FHO96idoXFeiBDu1y5Zggg\/pdf",
	           "metadata": {}
	       }
	   }`

	equal, err := invdutil.JsonEqual(object, string(so.Data))

	if err != nil {
		t.Fatal(err)
	}

	if !equal {
		t.Fatal("Event object is incorrect")
	}

}
