include $(GOROOT)/src/Make.inc

TARG=vmath

GOFILES=\
	doc.go \
	f32.go \
	m3.go \
	m4.go \
	v3.go \
	v4.go \
	p3.go \
	t3.go \

include $(GOROOT)/src/Make.pkg

pretty:
	gofmt -tabs=false -tabwidth=4 -w=true *.go
