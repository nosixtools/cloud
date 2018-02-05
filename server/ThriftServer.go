package server



type ThriftServer struct {
	host string
	port int
	group string
	version string
	weight string
	isStarted bool
}

func(ts *ThriftServer) Start() error  {
	return nil
}

func (ts *ThriftServer) Stop() error  {
	return nil
}

func (ts *ThriftServer) IsStart()  {

}

