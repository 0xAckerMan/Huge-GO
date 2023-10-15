package tests

import (
	"testing"

	"github.com/0xAckerMan/Huge-GO/Go-Almost_a_circle/models"
)

func TestNewBaseWithID(t *testing.T){
    base := models.NewBase(43)
    if base.Id != 43 {
        t.Errorf("Expected ID to be 43 but got %d\n", base.Id)
    }
}

func TestNewBaseWithNegative(t *testing.T){
    base := models.NewBase(-10)
    if base.Id != 1 {
    t.Errorf("Expected ID to be -10 but got %d\n",base.Id)
    }
    
}

func TestNewBaseWith0(t *testing.T){
    base := models.NewBase(0)
    if base.Id != 2 {
        t.Errorf("Expected ID to be 43 but got %d\n", base.Id)
    }
}
