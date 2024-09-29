package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// 数组的测试哦
type ListTestSuit struct {
	suite.Suite
}

// 测试插入方法
func (s *ListTestSuit) TestListInsert() {

}

func TestListSuit(t *testing.T) {
	suite.Run(t, new(ListTestSuit))
}
