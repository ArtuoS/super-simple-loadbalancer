package domain

type Balancer struct {
	ID      int
	Servers []*Server
}

func NewBalancer() *Balancer {
	return &Balancer{}
}

func (b *Balancer) PushServer(server *Server) {
	b.Servers = append(b.Servers, server)
}

func (b *Balancer) HandleRequest() {
	s := b.getServerWithSmallerCallCounter()
	s.Call()
}

func (b *Balancer) getServerWithSmallerCallCounter() *Server {
	if len(b.Servers) == 1 {
		return b.Servers[0]
	}

	s := b.Servers[0]
	for _, v := range b.Servers {
		if v.callCount < s.CallCount() {
			s = v
		}
	}

	return s
}
