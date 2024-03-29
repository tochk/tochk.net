// Code generated by qtc from "article.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line article.qtpl:1
package templates

//line article.qtpl:1
import "tochkru-golang/internal/app/datastruct"

//line article.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line article.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line article.qtpl:2
func StreamArticlePage(qw422016 *qt422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers, topTags []datastruct.Tags) {
//line article.qtpl:2
	qw422016.N().S(`
`)
//line article.qtpl:3
	qw422016.N().S(Header(article.Title))
//line article.qtpl:3
	qw422016.N().S(`
`)
//line article.qtpl:4
	qw422016.N().S(Menu(topTags))
//line article.qtpl:4
	qw422016.N().S(`
`)
//line article.qtpl:5
	qw422016.N().S(FullArticle(article, tags, teamMembers))
//line article.qtpl:5
	qw422016.N().S(`
`)
//line article.qtpl:6
	qw422016.N().S(Footer())
//line article.qtpl:6
	qw422016.N().S(`
`)
//line article.qtpl:7
}

//line article.qtpl:7
func WriteArticlePage(qq422016 qtio422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers, topTags []datastruct.Tags) {
//line article.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line article.qtpl:7
	StreamArticlePage(qw422016, article, tags, teamMembers, topTags)
//line article.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:7
}

//line article.qtpl:7
func ArticlePage(article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers, topTags []datastruct.Tags) string {
//line article.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line article.qtpl:7
	WriteArticlePage(qb422016, article, tags, teamMembers, topTags)
//line article.qtpl:7
	qs422016 := string(qb422016.B)
//line article.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line article.qtpl:7
	return qs422016
//line article.qtpl:7
}

//line article.qtpl:9
func StreamFullArticle(qw422016 *qt422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line article.qtpl:9
	qw422016.N().S(`
<article id='block_news'>
    <div id='block_in_block_news'>
    <div id='block_in_block_news_2'>
    <div id='block_news_head' style='font-size:20px;'><a href='/article/`)
//line article.qtpl:13
	qw422016.N().D(article.ID)
//line article.qtpl:13
	qw422016.N().S(`'><h1 style="font-size: 20px">`)
//line article.qtpl:13
	qw422016.E().S(article.Title)
//line article.qtpl:13
	qw422016.N().S(`</h1></a></div>
    <div id='block_news_content' style='font-size:16px;'>`)
//line article.qtpl:14
	qw422016.N().S(article.Text)
//line article.qtpl:14
	qw422016.N().S(`</div>
    </div></div>
    <div id='block_news_info'>
    <div id='block_news_info_2'>
    <div id='block_date' style='font-size: 20px;'>`)
//line article.qtpl:18
	qw422016.E().S(article.CreatedAt.Format("2006.01.02"))
//line article.qtpl:18
	qw422016.N().S(`</div>
    <div id='block_data'>
    <author id='data_author'> Author<br><text_green>`)
//line article.qtpl:20
	for _, teamMemberID := range article.TeamMembersIDs {
//line article.qtpl:20
		qw422016.E().S(teamMembers[teamMemberID].Name)
//line article.qtpl:20
		qw422016.N().S(`<br>`)
//line article.qtpl:20
	}
//line article.qtpl:20
	qw422016.N().S(`<br></text_green></author>
    <div id='data_category'> Tags<br><text_green>`)
//line article.qtpl:21
	for _, tagID := range article.TagsIDs {
//line article.qtpl:21
		qw422016.E().S(tags[tagID].Name)
//line article.qtpl:21
		qw422016.N().S(`<br>`)
//line article.qtpl:21
	}
//line article.qtpl:21
	qw422016.N().S(`<br></text_green></div>
    </div>
    <div id='block_info_footer' style='font-size: 20px;'> <a href='/article/`)
//line article.qtpl:23
	qw422016.N().D(article.ID)
//line article.qtpl:23
	qw422016.N().S(`'>read more</a></div>
    </div></div>
</article>
`)
//line article.qtpl:26
}

//line article.qtpl:26
func WriteFullArticle(qq422016 qtio422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line article.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line article.qtpl:26
	StreamFullArticle(qw422016, article, tags, teamMembers)
//line article.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line article.qtpl:26
}

//line article.qtpl:26
func FullArticle(article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) string {
//line article.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line article.qtpl:26
	WriteFullArticle(qb422016, article, tags, teamMembers)
//line article.qtpl:26
	qs422016 := string(qb422016.B)
//line article.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line article.qtpl:26
	return qs422016
//line article.qtpl:26
}
