// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line index.qtpl:1
package templates

//line index.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line index.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line index.qtpl:1
func StreamIndexPage(qw422016 *qt422016.Writer) {
//line index.qtpl:1
	qw422016.N().S(`
index
`)
//line index.qtpl:3
}

//line index.qtpl:3
func WriteIndexPage(qq422016 qtio422016.Writer) {
//line index.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line index.qtpl:3
	StreamIndexPage(qw422016)
//line index.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line index.qtpl:3
}

//line index.qtpl:3
func IndexPage() string {
//line index.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line index.qtpl:3
	WriteIndexPage(qb422016)
//line index.qtpl:3
	qs422016 := string(qb422016.B)
//line index.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line index.qtpl:3
	return qs422016
//line index.qtpl:3
}