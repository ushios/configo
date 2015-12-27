package configo

import "testing"

func TestConfigo(t *testing.T) {
	config, err := Instance("dev")

	if err != nil {
		t.Errorf("Cannot get config (%s)", err)
	}

	host, err := config.String("host")
	if err != nil {
		t.Errorf("Cannot got host key value (%s)", err)
	}

	if host != "dev.github.com" {
		t.Errorf("`host` expected (%s) but (%s)", "dev.github.com", host)
	}
}
