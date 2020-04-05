package postgres

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"labix.org/v2/mgo/bson"
)

func TestPostgres(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}

var onePlusWhitespace = regexp.MustCompile(`\s+`)

func PrepareQueryWhitespace(s string) string {
	ret := strings.TrimSpace(s)
	ret = onePlusWhitespace.ReplaceAllString(ret, `\s+`)
	return ret
}

var sqlEscapeRegexp = regexp.MustCompile(`[.()$*]`)

func PrepareQueryRegex(query string) string {
	matcher := PrepareQueryWhitespace(query)
	matcher = sqlEscapeRegexp.ReplaceAllString(matcher, `\$0`)
	matcher = "^" + matcher + "$"
	return matcher
}

func BeObjectID() sqlMock.Argument { return sqlObjectIDMatcher{} }

type sqlObjectIDMatcher struct{}

func (sqlObjectIDMatcher) Match(v driver.Value) bool {
	switch vv := v.(type) {
	case []uint8:
		return bson.ObjectId(vv).Valid()
	case string:
		return bson.ObjectId(vv).Valid()
	default:
		panic(fmt.Errorf("invalid value type: %T: %+[1]v", v))
	}
}
