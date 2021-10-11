package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}

//https://blog.csdn.net/weixin_40165163/article/details/90740155?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522163376631716780265468242%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=163376631716780265468242&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-1-90740155.pc_search_ecpm_flag&utm_term=golang+%E8%A3%85%E9%A5%B0%E5%99%A8%E6%A8%A1%E5%BC%8F&spm=1018.2226.3001.4187