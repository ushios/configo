package configo

import "testing"

func TestConfigoDev(t *testing.T) {
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

	port, err := config.Int("port")
	if err != nil {
		t.Errorf("Cannot got port key value(%s)", err)
	}

	if port != 80 {
		t.Errorf("`port` expected (%d) but (%d)", 80, port)
	}
}

func TestConfigoPrd(t *testing.T) {
	config, err := Instance("prd")

	if err != nil {
		t.Errorf("Cannot get config (%s)", err)
	}

	host, err := config.String("host")
	if err != nil {
		t.Errorf("Cannot got host key value (%s)", err)
	}

	if host != "github.com" {
		t.Errorf("`host` expected (%s) but (%s)", "github.com", host)
	}

	port, err := config.Int("port")
	if err != nil {
		t.Errorf("Cannot got port key value(%s)", err)
	}

	if port != 443 {
		t.Errorf("`port` expected (%d) but (%d)", 443, port)
	}
}
