package cerberus

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	sdkBls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/signer/bls/types"

	v1 "github.com/Layr-Labs/cerberus-api/pkg/api/v1"
)

type Config struct {
	URL          string
	PublicKeyHex string

	// Optional: in case if your signer uses local keystore
	Password string

	EnableTLS       bool
	TLSCertFilePath string

	SigningTimeout time.Duration
}

type Signer struct {
	signerClient v1.SignerClient
	kmsClient    v1.KeyManagerClient
	pubKeyHex    string
	password     string
}

func New(cfg Config) (Signer, error) {
	opts := make([]grpc.DialOption, 0)
	if cfg.EnableTLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.TLSCertFilePath, "")
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(cfg.URL, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	signerClient := v1.NewSignerClient(conn)
	kmsClient := v1.NewKeyManagerClient(conn)
	return Signer{
		signerClient: signerClient,
		kmsClient:    kmsClient,
		pubKeyHex:    cfg.PublicKeyHex,
		password:     cfg.Password,
	}, nil
}

func (s Signer) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 32 {
		return nil, types.ErrInvalidMessageLength
	}

	resp, err := s.signerClient.SignGeneric(ctx, &v1.SignGenericRequest{
		Data:        msg,
		PublicKeyG1: s.pubKeyHex,
		Password:    s.password,
	})
	if err != nil {
		return nil, err
	}

	return resp.Signature, nil
}

func (s Signer) SignG1(ctx context.Context, msg []byte) ([]byte, error) {
	if len(msg) != 64 {
		return nil, types.ErrInvalidMessageLength
	}

	resp, err := s.signerClient.SignG1(ctx, &v1.SignG1Request{
		Data:        msg,
		PublicKeyG1: s.pubKeyHex,
		Password:    s.password,
	})
	if err != nil {
		return nil, err
	}

	return resp.Signature, nil
}

func (s Signer) GetOperatorId() (string, error) {
	pkBytes, err := hex.DecodeString(s.pubKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode BLS public key: %w", err)
	}
	pubkey := new(sdkBls.G1Point)
	publicKey := pubkey.Deserialize(pkBytes)
	return publicKey.GetOperatorID(), nil
}

func (s Signer) GetPublicKeyG1() string {
	return s.pubKeyHex
}

func (s Signer) GetPublicKeyG2() string {
	resp, err := s.kmsClient.GetKeyMetadata(context.Background(), &v1.GetKeyMetadataRequest{
		PublicKeyG1: s.pubKeyHex,
	})
	if err != nil {
		log.Fatalf("could not get key metadata from cerberus: %v", err)
	}

	return resp.PublicKeyG2
}
