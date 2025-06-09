// Copyright 2025 AI Chat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
