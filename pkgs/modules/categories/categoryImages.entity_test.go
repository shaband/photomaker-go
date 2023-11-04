package categories

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestCategoryImageFake(t *testing.T) {
	categoryImage := CategoryImage{}
	categoryImage.Fake()

	if categoryImage.Image == "" {
		t.Error("Expected Image to be non-empty, got empty string")
	}

	if categoryImage.Order == 0 {
		t.Error("Expected Order to be non-zero, got zero")
	}

	if categoryImage.Category.NameAr == "" {
		t.Error("Expected Category.NameAr to be non-empty, got empty string")
	}

	if categoryImage.Category.NameEn == "" {
		t.Error("Expected Category.NameEn to be non-empty, got empty string")
	}

	if categoryImage.Category.Cover == "" {
		t.Error("Expected Category.Cover to be non-empty, got empty string")
	}
}
