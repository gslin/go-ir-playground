#
.DEFAULT:	all
.PHONY:		all clean test

#
DIST?=	dist/
GO?=	go

#
all:: dist/ir-bm25 dist/ir-tfidf
	@true

dist/%:: cmd/%/* internal/**
	"${GO}" build -o "${DIST}" ./cmd/$(notdir $@)

clean::
	rm -rf "${DIST}"

test::
	go test ./...
