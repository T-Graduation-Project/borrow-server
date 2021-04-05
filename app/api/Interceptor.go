package api

//func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
//	return func(ctx context.Context, req server.Request, resp interface{}) error {
//		meta, ok := metadata.FromIncomingContext(ctx)
//		if !ok {
//			return errors.New("no auth meta-data found in request")
//		}
//
//		// Note this is now uppercase (not entirely sure why this is...)
//		token := meta["Token"]
//
//		// Auth here
//		authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
//		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
//			Token: token,
//		})
//		log.Println("Auth Resp:", authResp)
//		if err != nil {
//			return err
//		}
//		err = fn(ctx, req, resp)
//		return err
//	}
//}

//func AuthInterceptor(
//	ctx context.Context, req interface{},
//	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	conn, err := grpc.Dial(authAddress, grpc.WithInsecure())
//	if err != nil {
//		service.log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := userPb.NewUserClient(conn)
//	// 从 metadate 中取出 token
//	md, _ := metadata.FromIncomingContext(ctx)
//	var token string
//	if val, ok := md["token"]; ok {
//		token = val[0]
//	} else {
//		service.log.Fatalf("md error")
//	}
//	// auth
//	r, err := c.Auth(ctx, &userPb.AuthReq{Token: token})
//	if err != nil {
//		service.log.Fatalf("could not authenticate: %v", err)
//	}
//	service.log.Info(r)
//	return handler(ctx, req)
//}
