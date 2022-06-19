package main

const BaseKeyFilePath = "f:\\home\\xukun\\Desktop\\zqm-tls-cert\\"

// func main() {

// 	server := network.TLSServer{}
// 	server.CertificateFile = BaseKeyFilePath + "default.pem"
// 	server.PrivateKeyFile = BaseKeyFilePath + "default.key"
// 	server.Port = 6666
// 	server.ServiceFactory = &myServiceFactory{}

// 	err := server.Start()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = server.Join()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ////////////////////////////////////////////////////////////////////////////////

// type myServiceFactory struct{}

// func (inst *myServiceFactory) Create() network.StreamService {
// 	return inst
// }

// func (inst *myServiceFactory) Handle(conn net.Conn) error {

// 	if conn == nil {
// 		return errors.New("conn == nil")
// 	}
// 	defer func() {
// 		conn.Close()
// 	}()

// 	buffer := make([]byte, 256)
// 	builder := bytes.Buffer{}
// 	var err2 error

// 	for {
// 		cnt, err := conn.Read(buffer)
// 		if err != nil {
// 			err2 = err
// 			break
// 		}
// 		builder.Write(buffer[0:cnt])
// 		log.Println("handle bytes: cnt=", cnt)
// 	}
// 	log.Println("handle bytes: str=", builder.String())
// 	return err2
// }
