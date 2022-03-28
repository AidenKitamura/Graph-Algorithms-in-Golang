package behavioral

import "fmt"

type networkResponse struct {
	status int
	header string
	body   string
}

func (n networkResponse) GetStatus() int {
	return n.status
}

func (n networkResponse) GetHeader() string {
	return n.header
}

func (n networkResponse) GetBody() string {
	return n.body
}

type NetworkResponse interface {
	GetStatus() int
	GetHeader() string
	GetBody() string
}

type Parser interface {
	Parse(n NetworkResponse)
	SetNext(p Parser)
}

type statusParser struct {
	next Parser
}

func (s *statusParser) Parse(n NetworkResponse) {
	switch n.GetStatus() {
	case 404:
		fmt.Println("Cannot find page, abort.")
	default:
		fmt.Printf("Other status: %d\n", n.GetStatus())
		if s.next != nil {
			s.next.Parse(n)
		}
	}
}

func (s *statusParser) SetNext(p Parser) {
	s.next = p
}

type headerParser struct {
	next Parser
}

func (h *headerParser) Parse(n NetworkResponse) {
	switch n.GetHeader() {
	case "Text", "CSS", "JS":
		fmt.Printf("Got Header: %s\n", n.GetHeader())
		if h.next != nil {
			h.next.Parse(n)
		}
	default:
		fmt.Printf("Unknown header, abort.\n")
	}
}

func (h *headerParser) SetNext(p Parser) {
	h.next = p
}

type bodyParser struct {
	next Parser
}

func (b *bodyParser) Parse(n NetworkResponse) {
	fmt.Println(n.GetBody())
}

func (b *bodyParser) SetNext(p Parser) {
	b.next = p
}
