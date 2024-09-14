package gen

import (
	"bytes"
	"testing"

	"github.com/Callahari/zusi-go/tcp/message/fahrpult"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	sub := fahrpult.DataProg{}

	res, err := Reflect(sub)
	assert.NoError(t, err)

	buf := bytes.NewBufferString("")
	err = Generate(*res, buf)
	assert.NoError(t, err)
}

func TestFullFile(t *testing.T) {
	sub := fahrpult.DataProg{}

	res, err := Reflect(sub)
	assert.NoError(t, err)

	buf := bytes.NewBufferString("")
	err = GenerateReaderFile([]Message{*res}, buf)
	assert.NoError(t, err)
}
