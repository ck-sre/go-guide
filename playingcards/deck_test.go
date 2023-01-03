package main

func TestNewDeck(t *testing.T) {
	d := newDec

	if len(d) != 40 {
		t.Errorf("expected 40, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected Ace, but got %v", d[0])
	}

	if d[len(d)-1] != "Ten of Hearts" {
		t.Errorf("expected last element ten of heards,	 but got %v", d[len(d)-1])
	}
}

func SaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
