package template

import (
	"github.com/mohuishou/go-design-pattern/13_template/example_csdn"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sms_Send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 1239999)
	assert.NoError(t, err)
}


func TestTemplatePattern(t *testing.T) {
	cricket := new(example_csdn.Cricket)
	cricket.Play()

	football := new(example_csdn.Football)
	football.Play()
}