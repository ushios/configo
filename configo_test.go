package configo

import "testing"

func TestConfigoDev(t *testing.T) {
	config, err := Instance("sample.cfg", "dev")

	if err != nil {
		t.Errorf("Cannot get config (%s)", err)
	}

	host, _ := config.String("host")
	if host != "dev.github.com" {
		t.Errorf("`host` expected (%s) but (%s)", "dev.github.com", host)
	}

	port, _ := config.Int("port")
	if port != 80 {
		t.Errorf("`port` expected (%d) but (%d)", 80, port)
	}

	pi, _ := config.Float("pi")
	if pi != 3.14 {
		t.Errorf("`pi` expected (%f) but (%f)", 3.14, pi)
	}

	prd, _ := config.Bool("prd")
	if prd != false {
		t.Errorf("`prd` expected (%t) but (%t)", false, prd)
	}

}

func TestConfigoPrd(t *testing.T) {
	config, err := Instance("sample.cfg", "prd")

	if err != nil {
		t.Errorf("Cannot get config (%s)", err)
	}

	host, _ := config.String("host")
	if host != "github.com" {
		t.Errorf("`host` expected (%s) but (%s)", "github.com", host)
	}

	port, _ := config.Int("port")
	if port != 443 {
		t.Errorf("`port` expected (%d) but (%d)", 443, port)
	}

	pi, _ := config.Float("pi")
	if pi != 3.1415 {
		t.Errorf("`pi` expected (%f) but (%f)", 3.1415, pi)
	}

	prd, _ := config.Bool("prd")
	if prd != true {
		t.Errorf("`prd` expected (%t) but (%t)", true, prd)
	}
}
