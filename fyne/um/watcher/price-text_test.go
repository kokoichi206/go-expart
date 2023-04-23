package main

import "testing"

func TestApp_getPriceText(t *testing.T) {
	open, _, _ := testApp.getPriceText()

	if open.Text != "open: $2005.265 USD" {
		t.Errorf("open.Text = %s; want open: $2005.265 USD", open.Text)
	}
}
