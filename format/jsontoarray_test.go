package format

import (
	"testing"

	"github.com/trivago/gollum/core"
	"github.com/trivago/tgo/ttesting"
)

func TestJSONToArray(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.JSONToArray")
	config.Override("Fields", []interface{}{
		"foo",
		"bar",
	})

	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*JSONToArray)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("{\"foo\":\"value1\",\"bar\":\"value2\"}"),
		0, core.InvalidStreamID)

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	expect.Equal("value1,value2", string(msg.Data()))
}

func TestJSONToArrayApplyTo(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.JSONToArray")
	config.Override("ApplyTo", "foo")
	config.Override("Fields", []interface{}{
		"foo",
		"bar",
	})

	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*JSONToArray)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("{\"test\":\"value\"}"),
		0, core.InvalidStreamID)
	msg.MetaData().SetValue("foo", []byte("{\"foo\":\"value1\",\"bar\":\"value2\"}"))

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	expect.Equal("value1,value2", msg.MetaData().GetValueString("foo"))
	expect.Equal("{\"test\":\"value\"}", msg.String())
}