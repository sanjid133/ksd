package server

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	"fmt"
	"github.com/sanjid133/ksd/apis/sho/v1"
	"github.com/sanjid133/ksd/pkg/server"
	"io"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"net"
	"github.com/sanjid133/ksd/pkg/operator"
	"github.com/spf13/pflag"

)

const defaultEtcdPathPrefix = "/registry/ksd.kubernetes.io"

type KsdServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	StdOut io.Writer
	StdErr io.Writer
}

func NewKsdServerOptions(out, errOut io.Writer) *KsdServerOptions {
	o := &KsdServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, server.Codecs.LegacyCodec(v1.SchemeGroupVersion)),
		StdOut:             out,
		StdErr:             errOut,
	}
	return o
}

func (o *KsdServerOptions) AddFlags(fs *pflag.FlagSet) {
	o.RecommendedOptions.AddFlags(fs)
	//o.OperatorOptions.AddFlags(fs)
}

func (o *KsdServerOptions) Complete() error {
	return nil
}

func (o *KsdServerOptions) Config() (*server.KsdConfig, error) {

	//here add certs
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}
	serverConfig := genericapiserver.NewRecommendedConfig(server.Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	operatorConfig := operator.NewOperatorConfig("", "")
	config := &server.KsdConfig{
		GenericConfig: serverConfig,
		OperatorConfig:     operatorConfig,
	}
	return config, nil

}

func (o KsdServerOptions) Validate(args []string) error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o KsdServerOptions) RunKsdServer(stopCh <-chan struct{}) error {
	fmt.Println("1...................")
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	err = server.GenericAPIServer.PrepareRun().Run(stopCh)
	fmt.Println(err)

	cntrl, err := config.OperatorConfig.New()
	if err != nil {
		return err
	}

	stop := make(chan struct{})
	go cntrl.Run(2, stop)
	return nil
}
