package libs

import "testing"

func TestResponse(t *testing.T) {
	result := Response(nil, 200, "success", nil)

	if result.Error != nil {
		t.Fatal("error :", result.Error)
	}
}
