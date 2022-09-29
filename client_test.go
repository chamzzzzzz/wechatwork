package wechatwork

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	client := Client{
		CorpId:     os.Getenv("WECHATWORK_TEST_CORP_ID"),
		CorpSecret: os.Getenv("WECHATWORK_TEST_CORP_SECRET"),
	}

	users, err := client.GetDepartmentUsers(1)
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		t.Log(user)
	}
	t.Log(len(users))
}
