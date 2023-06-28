package http

// func freePort() int {
// 	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
// 	if err != nil {
// 		panic(err)
// 	}
// 	l, err := net.ListenTCP("tcp", addr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer l.Close()
// 	return l.Addr().(*net.TCPAddr).Port
// }

// func newServer(t *testing.T) (*Server, *mockservice.MockService, *bytes.Buffer) {
// 	ctrl := gomock.NewController(t)
// 	buf := bytes.NewBuffer(nil)
// 	logger := slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{}))
// 	svc := mockservice.NewMockService(ctrl)
// 	query := mockdevice.NewMockQuery(ctrl)
// 	publisher := publisher.New(logger)
// 	srv := NewServer(svc, query, publisher, logger)
// 	go func() {
// 		addr := fmt.Sprintf(":%d", freePort())
// 		require.NoError(t, srv.Listen(addr))
// 	}()
// 	return srv, svc, buf
// }
