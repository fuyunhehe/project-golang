package ping

import "testing"

func TestPing(t *testing.T) {
	r := Ping()
	t.Errorf("%++v", r)
}
