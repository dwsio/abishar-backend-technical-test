package services

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceConnection struct {
}

func InitServicesConn(
	certFile string,
) *ServiceConnection {
	var err error
	var creds credentials.TransportCredentials
	if certFile != "" {
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		creds = insecure.NewCredentials()
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))

	services := &ServiceConnection{}

	return services
}

func initGrpcClientConn(address string, name string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if address == "" {
		return nil, fmt.Errorf("%s address is empty", name)
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed connect to %s: %v", name, err)
	}

	logrus.Println(fmt.Sprintf("[service - connection] %s State: %s", name, conn.GetState().String()))
	logrus.Println(fmt.Sprintf("[service - connection] %s Connected, on %s", name, address))

	return conn, nil
}

func (s *ServiceConnection) CloseAllServicesConn() {

}