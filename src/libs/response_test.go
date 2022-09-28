package libs

import "testing"

func TestResponse(t *testing.T) {
	result := Response(nil, 200, "success", nil)

	if result.IsError != nil {
		t.Fatal("error :", result.IsError)
	}
}
