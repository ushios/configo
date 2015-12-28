package configo

import (
	"bufio"
	"bytes"
	"testing"
)

func TestConfigoDev(t *testing.T) {
	assetByte, err := Asset("sample.cfg")
	if err != nil {
		t.Errorf("Asset got error (%s)", err)
	}
	buffer := bytes.NewBuffer(assetByte)

	config, err := Instance(bufio.NewReader(buffer), "dev")
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
	assetByte, err := Asset("sample.cfg")
	if err != nil {
		t.Errorf("Asset got error (%s)", err)
	}
	buffer := bytes.NewBuffer(assetByte)

	config, err := Instance(bufio.NewReader(buffer), "prd")
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

func TestMerge(t *testing.T) {
	assetByte, err := Asset("sample.cfg")
	if err != nil {
		t.Errorf("Asset got error (%s)", err)
	}
	buffer := bytes.NewBuffer(assetByte)

	config, err := Instance(bufio.NewReader(buffer), "dev")
	if err != nil {
		t.Errorf("Cannot get config (%s)", err)
	}

	databaseByte, err := Asset("database.cfg")
	if err != nil {
		t.Errorf("Asset got error (%s)", err)
	}
	databaseBuffer := bytes.NewBuffer(databaseByte)

	err = config.Merge(bufio.NewReader(databaseBuffer))
	if err != nil {
		t.Errorf("Cannot merge config (%s)", err)
	}

	dbhost, _ := config.String("dbhost")
	if dbhost != "devdb.local" {
		t.Errorf("`dbhost` expected (%s) but (%s)", "devdb.local", dbhost)
	}

	host, _ := config.String("host")
	if host != "dev.github.com" {
		t.Errorf("`host` expected (%s) but (%s)", "dev.github.com", host)
	}

	port, _ := config.Int("port")
	if port != 3306 {
		t.Errorf("`port` expected (%d) but (%d)", 3306, port)
	}
}
