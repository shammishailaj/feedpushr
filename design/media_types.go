package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Feed is the feed resource media type.
var Feed = MediaType("application/vnd.feedpushr.feed.v1+json", func() {
	Description("A RSS feed")
	TypeName("Feed")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", String, "ID of feed (MD5 of the xmlUrl)", func() {
			Example("5bfb841c028281c0051828c115fd1f50")
		})
		Attribute("xmlUrl", String, "URL of the XML feed", func() {
			Example("http://www.hashicorp.com/feed.xml")
		})
		Attribute("htmlUrl", String, "URL of the feed website", func() {
			Example("http://www.hashicorp.com/blog")
		})
		Attribute("hubUrl", String, "URL of the PubSubHubbud hub", func() {
			Example("http://pubsubhubbub.appspot.com")
		})
		Attribute("text", String, "Text attribute of the Feed", func() {
			Example("RSS Feed")
		})
		Attribute("title", String, "Title of the Feed", func() {
			Example("Hashicorp Blog")
		})
		Attribute("status", String, "Aggregation status", func() {
			Enum("running", "stopped")
		})
		Attribute("lastCheck", DateTime, "Last aggregation pass")
		Attribute("nextCheck", DateTime, "Next aggregation pass")
		Attribute("errorMsg", String, "Last aggregation error")
		Attribute("errorCount", Integer, "Number of consecutive aggregation errors")
		Attribute("cdate", DateTime, "Date of creation")
		Attribute("mdate", DateTime, "Date of modification")

		Required("id", "xmlUrl", "title", "cdate", "mdate")
	})

	View("default", func() {
		Attribute("id")
		Attribute("xmlUrl")
		Attribute("htmlUrl")
		Attribute("hubUrl")
		Attribute("title")
		Attribute("text")
		Attribute("status")
		Attribute("lastCheck")
		Attribute("nextCheck")
		Attribute("errorMsg")
		Attribute("errorCount")
		Attribute("cdate")
		Attribute("mdate")
	})

	View("tiny", func() {
		Description("tiny is the view used to list feeds")
		Attribute("id")
		Attribute("xmlUrl")
		Attribute("title")
		Attribute("cdate")
	})

	View("link", func() {
		Attribute("id")
		Attribute("xmlUrl")
	})
})
