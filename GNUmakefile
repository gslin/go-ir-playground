#
.DEFAULT:	all
.PHONY:		all clean test

#
DIST?=	dist/
GO?=	go

#
all:: dist/ir-bm25 dist/ir-tfidf
	@true

dist/ir-bm25:: cmd/ir-bm25/* internal/**
	"${GO}" build -o "${DIST}" ./cmd/ir-bm25

dist/ir-tfidf:: cmd/ir-tfidf/* internal/**
	"${GO}" build -o "${DIST}" ./cmd/ir-tfidf

clean::
	rm -rf "${DIST}"

test::
	go test ./...
