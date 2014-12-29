package cassert

import (
	"github.com/smartystreets/goconvey/convey/assertions"
	"github.com/smartystreets/goconvey/convey/reporting"
	"testing"
)

var (
	errorTemplate   = "* %s \nLine %d: - %v \n%s\n"
	failureTemplate = "* %s \nLine %d:\n%s\n"
)

type assertion func(actual interface{}, expected ...interface{}) string
type sofunc func(actual interface{}, assert assertion, expected ...interface{})

func report(t *testing.T, r *reporting.AssertionResult) {
	if r.Error != nil {
		t.Errorf(errorTemplate, r.File, r.Line, r.Error, r.StackTrace)
	} else if r.Failure != "" {
		t.Errorf(failureTemplate, r.File, r.Line, r.Failure)
	}
}

func NewSo(t *testing.T) sofunc {
	return func(actual interface{}, assert assertion, expected ...interface{}) {
		So(t, actual, assert, expected...)
	}
}

func So(t *testing.T, actual interface{}, assert assertion, expected ...interface{}) {
	str := assert(actual, expected...)
	result := reporting.NewFailureReport(str)
	report(t, result)

	// out := reporting.NewPrinter(reporting.NewConsole())
	// reporter := reporting.NewProblemReporter(out)
	// defer reporter.EndStory()

	// reporter.Report(result)
}

var (
	ShouldEqual          = assertions.ShouldEqual
	ShouldNotEqual       = assertions.ShouldNotEqual
	ShouldAlmostEqual    = assertions.ShouldAlmostEqual
	ShouldNotAlmostEqual = assertions.ShouldNotAlmostEqual
	ShouldResemble       = assertions.ShouldResemble
	ShouldNotResemble    = assertions.ShouldNotResemble
	ShouldPointTo        = assertions.ShouldPointTo
	ShouldNotPointTo     = assertions.ShouldNotPointTo
	ShouldBeNil          = assertions.ShouldBeNil
	ShouldNotBeNil       = assertions.ShouldNotBeNil
	ShouldBeTrue         = assertions.ShouldBeTrue
	ShouldBeFalse        = assertions.ShouldBeFalse
	ShouldBeZeroValue    = assertions.ShouldBeZeroValue

	ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	ShouldBeLessThan             = assertions.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	ShouldBeBetween              = assertions.ShouldBeBetween
	ShouldNotBeBetween           = assertions.ShouldNotBeBetween
	ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
	ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual

	ShouldContain    = assertions.ShouldContain
	ShouldNotContain = assertions.ShouldNotContain
	ShouldBeIn       = assertions.ShouldBeIn
	ShouldNotBeIn    = assertions.ShouldNotBeIn
	ShouldBeEmpty    = assertions.ShouldBeEmpty
	ShouldNotBeEmpty = assertions.ShouldNotBeEmpty

	ShouldStartWith           = assertions.ShouldStartWith
	ShouldNotStartWith        = assertions.ShouldNotStartWith
	ShouldEndWith             = assertions.ShouldEndWith
	ShouldNotEndWith          = assertions.ShouldNotEndWith
	ShouldBeBlank             = assertions.ShouldBeBlank
	ShouldNotBeBlank          = assertions.ShouldNotBeBlank
	ShouldContainSubstring    = assertions.ShouldContainSubstring
	ShouldNotContainSubstring = assertions.ShouldNotContainSubstring

	ShouldPanic        = assertions.ShouldPanic
	ShouldNotPanic     = assertions.ShouldNotPanic
	ShouldPanicWith    = assertions.ShouldPanicWith
	ShouldNotPanicWith = assertions.ShouldNotPanicWith

	ShouldHaveSameTypeAs    = assertions.ShouldHaveSameTypeAs
	ShouldNotHaveSameTypeAs = assertions.ShouldNotHaveSameTypeAs
	ShouldImplement         = assertions.ShouldImplement
	ShouldNotImplement      = assertions.ShouldNotImplement

	ShouldHappenBefore         = assertions.ShouldHappenBefore
	ShouldHappenOnOrBefore     = assertions.ShouldHappenOnOrBefore
	ShouldHappenAfter          = assertions.ShouldHappenAfter
	ShouldHappenOnOrAfter      = assertions.ShouldHappenOnOrAfter
	ShouldHappenBetween        = assertions.ShouldHappenBetween
	ShouldHappenOnOrBetween    = assertions.ShouldHappenOnOrBetween
	ShouldNotHappenOnOrBetween = assertions.ShouldNotHappenOnOrBetween
	ShouldHappenWithin         = assertions.ShouldHappenWithin
	ShouldNotHappenWithin      = assertions.ShouldNotHappenWithin
	ShouldBeChronological      = assertions.ShouldBeChronological
)
