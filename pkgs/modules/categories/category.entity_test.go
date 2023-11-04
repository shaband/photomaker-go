package categories

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestCategoryFake(t *testing.T) {
	category := Category{}
	category.Fake()

	if category.NameAr == "" {
		t.Error("Expected NameAr to be non-empty, got empty string")
	}

	if category.NameEn == "" {
		t.Error("Expected NameEn to be non-empty, got empty string")
	}

	if category.Cover == "" {
		t.Error("Expected Cover to be non-empty, got empty string")
	}
}
