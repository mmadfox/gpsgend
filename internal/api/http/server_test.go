package http

// func newServer(t *testing.T) (*Server, *mockgpsgend.MockService) {
// 	ctrl := gomock.NewController(t)
// 	svc := mockgpsgend.NewMockService(ctrl)
// 	srv := NewServer(freePort(), svc)
// 	go func() {
// 		require.NoError(t, srv.Run())
// 	}()
// 	return srv, svc
// }

// func newDevice() *device.Device {
// 	id, _ := uuid.Parse("8eed3937-0545-4fb2-b304-3f472bc3a3eb")
// 	newDevice, err := device.NewBuilder().
// 		ID(id).
// 		Model("test").
// 		Status(device.Stopped).
// 		UserID("userid").
// 		Description("description").
// 		Speed(1, 2, 4).
// 		Battery(1, 100, time.Hour).
// 		Elevation(1, 100, 4).
// 		Offline(1, 300).
// 		Props(map[string]string{"foo": "bar"}).
// 		Build()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return newDevice
// }

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

// func writeGoldenTestData(name string, body []byte) {
// 	filename := "./testdata/" + name + ".json"
// 	os.WriteFile(filename, body, 0755)
// }

// func load(name string) string {
// 	filename := "./testdata/" + name + ".json"
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(data)
// }

// func testdata(filename string) string {
// 	data, err := os.ReadFile("./testdata/" + filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(data)
// }

// func TestServer_newDevice(t *testing.T) {
// 	srv, svcMock := newServer(t)
// 	defer srv.Stop()

// 	type request struct {
// 		body string
// 	}

// 	type response struct {
// 		statusCode int
// 		body       string
// 	}

// 	type goldenFile struct {
// 		enable bool
// 		name   string
// 	}

// 	tests := []struct {
// 		name string
// 		req  request
// 		resp response
// 		gf   goldenFile
// 		init func()
// 	}{
// 		{
// 			name: "should return error when invalid request",
// 			resp: response{
// 				statusCode: http.StatusBadRequest,
// 				body:       load("newDevice_emptyRequest_errorResponse"),
// 			},
// 			gf: goldenFile{
// 				enable: false,
// 				name:   "newDevice_emptyRequest_errorResponse",
// 			},
// 		},
// 		{
// 			name: "should return error when broken JSON",
// 			req: request{
// 				body: `{"model":}`,
// 			},
// 			resp: response{
// 				statusCode: http.StatusBadRequest,
// 				body:       load("newDevice_brokenJSON_errorResponse"),
// 			},
// 			gf: goldenFile{
// 				enable: false,
// 				name:   "newDevice_brokenJSON_errorResponse",
// 			},
// 		},
// 		{
// 			name: "should return error when service return error",
// 			init: func() {
// 				svcErr := errors.New("some error")
// 				svcMock.EXPECT().NewDevice(gomock.Any(), gomock.Any()).Return(nil, svcErr)
// 			},
// 			req: request{
// 				body: load("newDevice_request"),
// 			},
// 			resp: response{
// 				statusCode: http.StatusUnprocessableEntity,
// 				body:       load("newDevice_serviceError_errorResponse"),
// 			},
// 			gf: goldenFile{
// 				enable: false,
// 				name:   "newDevice_serviceError_errorResponse",
// 			},
// 		},
// 		{
// 			name: "should return new device",
// 			init: func() {
// 				svcMock.EXPECT().NewDevice(gomock.Any(), gomock.Any()).Return(newDevice(), nil)
// 			},
// 			req: request{
// 				body: load("newDevice_request"),
// 			},
// 			resp: response{
// 				statusCode: http.StatusOK,
// 				body:       load("newDevice_successResponse"),
// 			},
// 			gf: goldenFile{
// 				enable: false,
// 				name:   "newDevice_successResponse",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.init != nil {
// 				tt.init()
// 			}

// 			req, err := http.NewRequest(http.MethodPost, "/v1/devices", strings.NewReader(tt.req.body))
// 			require.NoError(t, err)

// 			resp, err := srv.Test(req)
// 			require.NoError(t, err)
// 			respBody, err := ioutil.ReadAll(resp.Body)
// 			require.NoError(t, err)

// 			if tt.gf.enable {
// 				writeGoldenTestData(tt.gf.name, respBody)
// 				return
// 			}

// 			require.Equal(t, tt.resp.statusCode, resp.StatusCode)
// 			require.Equal(t, tt.resp.body, string(respBody))
// 			resp.Body.Close()
// 		})
// 	}
// }

// func TestServer_downloadRoute(t *testing.T) {
// 	srv, svcMock := newServer(t)
// 	defer srv.Stop()

// 	originRoutes, err := route.RoutesForChina()
// 	require.NoError(t, err)
// 	routes := make([]*device.Route, len(originRoutes))
// 	for i := 0; i < len(originRoutes); i++ {
// 		routes[i] = device.NewRoute(originRoutes[i])
// 	}
// 	rawRoutes, err := geojson.Encode(originRoutes)
// 	require.NoError(t, err)

// 	tests := []struct {
// 		name           string
// 		mock           func()
// 		wantStatusCode int
// 		want           []byte
// 	}{
// 		{
// 			name: "should return error when gpsgend.Routes return error",
// 			mock: func() {
// 				svcMock.EXPECT().Routes(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
// 			},
// 			wantStatusCode: http.StatusUnprocessableEntity,
// 			want:           []byte(`{"reason":"some error","message":"downloadRoutes","code":422}`),
// 		},
// 		{
// 			name: "should return routes",
// 			mock: func() {
// 				svcMock.EXPECT().Routes(gomock.Any(), gomock.Any()).Return(routes, nil)
// 			},
// 			wantStatusCode: http.StatusOK,
// 			want:           rawRoutes,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mock != nil {
// 				tt.mock()
// 			}

// 			req, err := http.NewRequest(
// 				http.MethodGet,
// 				"/v1/devices/856944ec-23ff-489e-b66e-f6989eecd014/routes/download",
// 				nil,
// 			)
// 			require.NoError(t, err)

// 			resp, err := srv.Test(req)
// 			require.NoError(t, err)
// 			respBody, err := ioutil.ReadAll(resp.Body)
// 			require.NoError(t, err)

// 			if resp.StatusCode == http.StatusOK {
// 				require.Equal(t, resp.Header.Get("Content-Type"), "application/geo+json")
// 				require.Equal(t, resp.Header.Get("Content-Length"), strconv.Itoa(len(rawRoutes)))
// 				require.Equal(t, resp.Header.Get("Content-Disposition"), "attachment; filename=device_856944ec-23ff-489e-b66e-f6989eecd014_routes.geojson")
// 			} else {
// 				require.Equal(t, resp.Header.Get("Content-Type"), fiber.MIMEApplicationJSONCharsetUTF8)
// 			}

// 			require.Equal(t, tt.wantStatusCode, resp.StatusCode)
// 			require.Equal(t, tt.want, respBody)
// 		})
// 	}
// }

// func TestServer_getRoutes(t *testing.T) {
// 	srv, svcMock := newServer(t)
// 	defer srv.Stop()

// 	or1, err := route.China1()
// 	require.NoError(t, err)
// 	or2, err := route.China2()
// 	require.NoError(t, err)
// 	require.NoError(t, err)
// 	routes := make([]*device.Route, 0)
// 	routes = append(routes, device.RouteFrom(uuid.MustParse("91206450-c827-4d03-8351-bbed9b5c0066"), colorful.Color{}, or1))
// 	routes = append(routes, device.RouteFrom(uuid.MustParse("a96cb459-b76f-466b-8b03-9d3cffb9e345"), colorful.Color{}, or2))

// 	tests := []struct {
// 		name           string
// 		mock           func()
// 		wantStatusCode int
// 		wantResponse   string
// 	}{
// 		{
// 			name: "should return routes",
// 			mock: func() {
// 				svcMock.EXPECT().Routes(gomock.Any(), gomock.Any()).Times(1).Return(routes, nil)
// 			},
// 			wantResponse:   testdata("getRoutes_success.json"),
// 			wantStatusCode: http.StatusOK,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mock != nil {
// 				tt.mock()
// 			}

// 			req, err := http.NewRequest(
// 				http.MethodGet,
// 				"/v1/devices/856944ec-23ff-489e-b66e-f6989eecd014/routes",
// 				nil,
// 			)
// 			require.NoError(t, err)

// 			resp, err := srv.Test(req)
// 			require.NoError(t, err)
// 			respBody, err := ioutil.ReadAll(resp.Body)
// 			require.NoError(t, err)

// 			require.Equal(t, tt.wantStatusCode, resp.StatusCode)
// 			require.Equal(t, string(respBody), tt.wantResponse)
// 		})
// 	}
// }

// func TestServer_importRoute(t *testing.T) {
// 	srv, svcMock := newServer(t)
// 	defer srv.Stop()

// 	fsHandler := new(fileServer)
// 	fs := httptest.NewServer(fsHandler)
// 	defer fs.Close()

// 	tests := []struct {
// 		name           string
// 		mock           func()
// 		input          string
// 		wantStatusCode int
// 		matchResponse  string
// 	}{
// 		{
// 			name:           "should return error when url return error",
// 			wantStatusCode: http.StatusBadRequest,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.userSmallFile()
// 				fsHandler.statusCode = http.StatusInternalServerError
// 			},
// 			matchResponse: "failed to download file",
// 		},
// 		{
// 			name:           "should return error when url return small file",
// 			wantStatusCode: http.StatusBadRequest,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.userSmallFile()
// 				fsHandler.statusCode = http.StatusOK
// 			},
// 			matchResponse: "routes data not found. got 7 bytes",
// 		},
// 		{
// 			name:           "should return error when url return large file",
// 			wantStatusCode: http.StatusBadRequest,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.useLargeFile()
// 				fsHandler.statusCode = http.StatusOK
// 			},
// 			matchResponse: "routes data too large",
// 		},
// 		{
// 			name:           "should return error when url return some file",
// 			wantStatusCode: http.StatusBadRequest,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.fileType = ""
// 				fsHandler.statusCode = http.StatusOK
// 			},
// 			matchResponse: "undefined data format",
// 		},
// 		{
// 			name:           "should return error when gpsgend.AddRoutes return error",
// 			wantStatusCode: http.StatusUnprocessableEntity,
// 			input:          testdata("route.geojson"),
// 			mock: func() {
// 				svcMock.EXPECT().
// 					AddRoutes(gomock.Any(), gomock.Any(), gomock.Any()).
// 					Times(1).
// 					Return(errors.New("some error"))
// 			},
// 			matchResponse: "some error",
// 		},
// 		{
// 			name:           "should return routes when url return gpx file",
// 			wantStatusCode: http.StatusOK,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.useGPX()
// 				fsHandler.statusCode = http.StatusOK
// 				svcMock.EXPECT().AddRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
// 			},
// 			matchResponse: "ok",
// 		},
// 		{
// 			name:           "should return routes when url return geojson file",
// 			wantStatusCode: http.StatusOK,
// 			input:          fs.URL,
// 			mock: func() {
// 				fsHandler.useGeoJSON()
// 				fsHandler.statusCode = http.StatusOK
// 				svcMock.EXPECT().AddRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
// 			},
// 			matchResponse: "ok",
// 		},
// 		{
// 			name:           "should return routes",
// 			wantStatusCode: http.StatusOK,
// 			input:          testdata("route.gpx"),
// 			mock: func() {
// 				svcMock.EXPECT().AddRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
// 			},
// 			matchResponse: "ok",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mock != nil {
// 				tt.mock()
// 			}

// 			req, err := http.NewRequest(
// 				http.MethodPost,
// 				"/v1/devices/856944ec-23ff-489e-b66e-f6989eecd014/routes/import",
// 				strings.NewReader(tt.input),
// 			)
// 			require.NoError(t, err)

// 			resp, err := srv.Test(req)
// 			require.NoError(t, err)
// 			respBody, err := ioutil.ReadAll(resp.Body)
// 			require.NoError(t, err)

// 			require.Equal(t, tt.wantStatusCode, resp.StatusCode)
// 			require.Contains(t, string(respBody), tt.matchResponse)
// 		})
// 	}
// }

// func TestServer_uploadRoute(t *testing.T) {
// 	srv, svcMock := newServer(t)
// 	defer srv.Stop()

// 	attachFile := func(filename string, formName string, w *multipart.Writer) {
// 		part, err := w.CreateFormFile(formName, filepath.Base(filename))
// 		require.NoError(t, err)
// 		data, err := os.ReadFile(filename)
// 		require.NoError(t, err)
// 		io.Copy(part, bytes.NewReader(data))
// 		w.Close()
// 	}

// 	tests := []struct {
// 		name           string
// 		mock           func()
// 		attachFile     func(w *multipart.Writer)
// 		wantStatusCode int
// 		wantResponse   []byte
// 	}{
// 		{
// 			name: "should return error when invalid form key",
// 			attachFile: func(w *multipart.Writer) {
// 				attachFile("./testdata/route.gpx", "someForm", w)
// 			},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantResponse:   []byte(`{"reason":"there is no uploaded file associated with the given key","message":"readHeader","code":400}`),
// 		},
// 		{
// 			name: "should return error when file size too small",
// 			attachFile: func(w *multipart.Writer) {
// 				attachFile("./testdata/empty", uploadFilename, w)
// 			},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantResponse:   []byte(`{"reason":"file[\"empty\"] is empty","message":"validateUploadRouteRequest","code":400}`),
// 		},
// 		{
// 			name: "should return error when file size too large",
// 			attachFile: func(w *multipart.Writer) {
// 				filename := filepath.Join(os.TempDir(), "large")
// 				os.WriteFile(filename, []byte(strings.Repeat("s", maxUploadFileSize+1)), 0777)
// 				attachFile(filename, uploadFilename, w)
// 			},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantResponse:   []byte(`{"reason":"file[\"large\"] too large. maximum file size 3000000 bytes","message":"validateUploadRouteRequest","code":400}`),
// 		},
// 		{
// 			name: "should return error when file data xml but not gpx",
// 			attachFile: func(w *multipart.Writer) {
// 				filename := filepath.Join(os.TempDir(), "xml")
// 				os.WriteFile(filename, []byte(`<note> <to>Tove</to> <from>Jani</from> <heading>Reminder</heading> <body>Don't forget me this weekend!</body> </note>`), 0777)
// 				attachFile(filename, uploadFilename, w)
// 			},
// 			wantStatusCode: http.StatusInternalServerError,
// 			wantResponse:   []byte(`{"reason":"expected element type \u003cgpx\u003e but have \u003cnote\u003e","message":"decodeFile","code":500}`),
// 		},
// 		{
// 			name: "should return error when file data json but not geojson",
// 			attachFile: func(w *multipart.Writer) {
// 				filename := filepath.Join(os.TempDir(), "json")
// 				os.WriteFile(filename, []byte(`{"name": "Jane Doe", "favorite-game": "Stardew Valley", "subscriber": false, "name": "John Doe", "favorite-game": "Dragon Quest XI", "subscriber": true }`), 0777)
// 				attachFile(filename, uploadFilename, w)
// 			},
// 			wantStatusCode: http.StatusInternalServerError,
// 			wantResponse:   []byte(`{"reason":"type property not defined","message":"decodeFile","code":500}`),
// 		},
// 		{
// 			name: "should return error when file data is unknown",
// 			attachFile: func(w *multipart.Writer) {
// 				filename := filepath.Join(os.TempDir(), "unknown")
// 				os.WriteFile(filename, []byte(strings.Repeat("s", minUploadFileSize+1)), 0777)
// 				attachFile(filename, uploadFilename, w)
// 			},
// 			wantStatusCode: http.StatusBadRequest,
// 			wantResponse:   []byte(`{"reason":"undefined file \"unknown\" data format","message":"decodeFile","code":400}`),
// 		},
// 		{
// 			name: "should return error when failure gpsgend.AddRoutes",
// 			attachFile: func(w *multipart.Writer) {
// 				attachFile("./testdata/route.geojson", uploadFilename, w)
// 			},
// 			mock: func() {
// 				err := errors.New("some error")
// 				svcMock.EXPECT().AddRoutes(gomock.Any(), gomock.Any(), gomock.All()).
// 					Times(1).
// 					Return(err)
// 			},
// 			wantStatusCode: http.StatusUnprocessableEntity,
// 			wantResponse:   []byte(`{"reason":"some error","message":"addRoutes","code":422}`),
// 		},
// 		{
// 			name: "should return success response",
// 			attachFile: func(w *multipart.Writer) {
// 				attachFile("./testdata/route.geojson", uploadFilename, w)
// 			},
// 			mock: func() {
// 				svcMock.EXPECT().AddRoutes(gomock.Any(), gomock.Any(), gomock.All()).
// 					Times(1).
// 					Return(nil)
// 			},
// 			wantStatusCode: http.StatusOK,
// 			wantResponse:   []byte(`{"ok": true}`),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mock != nil {
// 				tt.mock()
// 			}

// 			buf := new(bytes.Buffer)
// 			w := multipart.NewWriter(buf)
// 			tt.attachFile(w)
// 			req, err := http.NewRequest(
// 				http.MethodPost,
// 				"/v1/devices/856944ec-23ff-489e-b66e-f6989eecd014/routes/upload",
// 				buf,
// 			)
// 			req.Header.Set("Content-Type", w.FormDataContentType())
// 			require.NoError(t, err)

// 			resp, err := srv.Test(req)
// 			require.NoError(t, err)
// 			respBody, err := ioutil.ReadAll(resp.Body)
// 			require.NoError(t, err)

// 			require.Equal(t, tt.wantStatusCode, resp.StatusCode)
// 			require.Equal(t, tt.wantResponse, respBody)
// 		})
// 	}
// }

// type fileServer struct {
// 	fileType   string
// 	statusCode int
// }

// func (fs *fileServer) useGPX() {
// 	fs.fileType = "gpx"
// }

// func (fs *fileServer) useGeoJSON() {
// 	fs.fileType = "geojson"
// }

// func (fs *fileServer) useLargeFile() {
// 	fs.fileType = "large"
// }

// func (fs *fileServer) userSmallFile() {
// 	fs.fileType = "small"
// }

// func (fs *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Disposition", "attachment; filename=file"+fs.fileType)
// 	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
// 	if fs.statusCode > 0 {
// 		w.WriteHeader(fs.statusCode)
// 	}

// 	var data string
// 	switch fs.fileType {
// 	case "gpx":
// 		data = testdata("route.gpx")
// 	case "geojson":
// 		data = testdata("route.geojson")
// 	case "large":
// 		data = strings.Repeat("s", maxUploadFileSize+1)
// 	case "small":
// 		data = strings.Repeat("s", minUploadFileSize-1)
// 	default:
// 		data = strings.Repeat("s", 512)
// 	}

// 	w.Header().Add("Content-Length", strconv.Itoa(len(data)))
// 	io.Copy(w, strings.NewReader(data))
// }
