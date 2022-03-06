package json

import (
	"io/fs"
	"io/ioutil"

	"github.com/bitly/go-simplejson"
)

func GetJsonValue(path string, branch ...string) (*simplejson.Json, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	json, err := simplejson.NewJson(buf)
	if err != nil {
		return nil, err
	}

	return json.GetPath(branch...), nil
}

func SetJsonValue(path string, val interface{}, branch ...string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// read json
	json, err := simplejson.NewJson(buf)
	if err != nil {
		return err
	}

	// set json value
	json.SetPath(branch, val)

	bytes, err := json.EncodePretty()
	if err != nil {
		return err
	}

	// write to disk
	ioutil.WriteFile(path, bytes, fs.ModeAppend)
	return nil
}
