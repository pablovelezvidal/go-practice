package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("expected deck len %d, got %d", 52, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card of Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("expected last card of King of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// remove any previous not deleted file
	os.Remove("_decktesting.txt")

	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("expected deck len %d, got %d", 52, len(d))
	}

	d.SaveToFile("_decktesting.txt")

	d2 := NewDeckFromFile("_decktesting.txt")

	if len(d2) != 52 {
		t.Errorf("expected deck len %d, got %d", 52, len(d2))
	}
}
