#
.DEFAULT:	all
.PHONY:		all clean

#
DIST?=	dist/
GO?=	go

#
all: dist/ir-bm25 dist/ir-tfidf
	@true

dist/ir-bm25: cmd/ir-bm25/*
	"${GO}" build -o "${DIST}" ./cmd/ir-bm25

dist/ir-tfidf: cmd/ir-tfidf/*
	"${GO}" build -o "${DIST}" ./cmd/ir-tfidf

clean:
	rm -rf "${DIST}"
