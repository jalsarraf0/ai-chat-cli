package tui

import "testing"

func TestSpinnerStops(t *testing.T) {
	t.Parallel()
	m := NewModel(5)
	m.spinner = newSpinner()
	m.streaming = true
	// simulate spinner tick
	tm, _ := m.Update(m.spinner.Tick())
	m = tm.(Model)
	// send completion done
	tm, _ = m.Update(llmTokenMsg{Done: true})
	m = tm.(Model)
	if m.streaming {
		t.Fatalf("spinner still visible")
	}
}
