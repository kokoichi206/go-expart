package main

import "testing"

func TestApp_getToolBar(t *testing.T) {
	toolbar := testApp.getToolBar()

	if len(toolbar.Items) != 4 {
		t.Errorf("len(toolbar.Items) = %d; want 4", len(toolbar.Items))
	}
}
