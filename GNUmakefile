#
.DEFAULT:	all
.PHONY:		all clean

#
DIST?=	dist/
GO?=	go

#
all: dist/ir-tfidf
	@true

dist/ir-tfidf: cmd/ir-tfidf/*
	"${GO}" build -o "${DIST}" ./cmd/ir-tfidf

clean:
	rm -rf "${DIST}"
