package apiserver

type apiserver struct {
}

func new() *APIServer {
	return &APIServer{}
}

func (s *APIServer) Start() error {
	return nil
}
