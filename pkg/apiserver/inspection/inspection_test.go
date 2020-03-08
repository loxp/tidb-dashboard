package inspection

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	. "github.com/pingcap/check"
)

func TestT(t *testing.T) {
	CustomVerboseFlag = true
	TestingT(t)
}

var _ = Suite(&testInspectionSuite{})

type testInspectionSuite struct{}

func (t *testInspectionSuite) TestInspection(c *C) {
	cli, err := gorm.Open("mysql", "root:@tcp(172.16.5.40:4009)/test?charset=utf8&parseTime=True&loc=Local")
	c.Assert(err, IsNil)
	defer cli.Close()

	startTime1 := "2020-03-08 02:23:00"
	endTime1 := "2020-03-08 02:27:00"

	startTime2 := "2020-03-08 02:30:00"
	endTime2 := "2020-03-08 02:34:00"

	is := &clusterInspection{
		referStartTime: startTime1,
		referEndTime:   endTime1,

		startTime: startTime2,
		endTime:   endTime2,
		db:        cli,
	}
	_, err = is.inspectForAffectByBigQuery()
	c.Assert(err, IsNil)
}