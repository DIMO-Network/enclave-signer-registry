package app

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/DIMO-Network/enclave-bridge/pkg/certs"
	"github.com/DIMO-Network/enclave-bridge/pkg/certs/acme"
	"github.com/DIMO-Network/enclave-bridge/pkg/client"
	"github.com/DIMO-Network/enclave-bridge/pkg/wellknown"
	"github.com/DIMO-Network/enclave-signer-registry/internal/client/devlicense"
	"github.com/DIMO-Network/enclave-signer-registry/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
)

func CreateEnclaveWebServer(ctx context.Context, logger *zerolog.Logger, clientPort uint32, settings *config.Settings) (*fiber.App, *tls.Config, error) {
	// Setup HTTP client
	httpClient := client.NewHTTPClient(clientPort, nil)
	var certFunc certs.GetCertificateFunc
	var tlsConfig *tls.Config
	if settings.TLS.Enabled {
		certPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to generate cert private key: %w", err)
		}
		tlsConfig, err = setupTLSConfig(settings, certPrivateKey, httpClient, logger)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to setup TLS config: %w", err)
		}
		certFunc = tlsConfig.GetCertificate
	}
	devLicenseClient, err := devlicense.NewClient(ctx, settings, httpClient)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create dev license client: %w", err)
	}
	// Setup the controller with all its dependencies
	ctrl, err := NewController(settings, devLicenseClient)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup controller: %w", err)
	}

	wellKnownCtrl, err := wellknown.NewController(nil, certFunc)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup well known controller: %w", err)
	}

	app := createApp(logger, ctrl, wellKnownCtrl)

	return app, tlsConfig, nil
}

func createApp(logger *zerolog.Logger, ctrl *Controller, wellKnownCtrl *wellknown.Controller) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return ErrorHandler(c, err, logger)
		},
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	}))

	app.Use(func(c *fiber.Ctx) error {
		userCtx := logger.With().Str("httpPath", strings.TrimPrefix(c.Path(), "/")).
			Str("httpMethod", c.Method()).Logger().WithContext(c.UserContext())
		c.SetUserContext(userCtx)
		return c.Next()
	})

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", HealthCheck)
	app.Get("/developer-license", ctrl.GetDeveloperLicense)
	app.Post("/add-signer", ctrl.AddSigner)
	wellknown.RegisterRoutes(app, wellKnownCtrl)
	return app
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(ctx *fiber.Ctx) error {
	res := map[string]any{
		"data": "Server is up and running",
	}

	return ctx.JSON(res)
}

// ErrorHandler custom handler to log recovered errors using our logger and return json instead of string.
func ErrorHandler(ctx *fiber.Ctx, err error, logger *zerolog.Logger) error {
	code := fiber.StatusInternalServerError // Default 500 statuscode
	message := "Internal error."

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	// don't log not found errors
	if code != fiber.StatusNotFound {
		logger.Err(err).Int("httpStatusCode", code).
			Str("httpPath", strings.TrimPrefix(ctx.Path(), "/")).
			Str("httpMethod", ctx.Method()).
			Msg("caught an error from http request")
	}

	return ctx.Status(code).JSON(codeResp{Code: code, Message: message})
}

type codeResp struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// setupTLSConfig configures TLS settings including certificate management
func setupTLSConfig(settings *config.Settings, certPrivateKey *ecdsa.PrivateKey, httpClient *http.Client, logger *zerolog.Logger) (*tls.Config, error) {
	if settings.TLS.LocalCerts.CertFile != "" && settings.TLS.LocalCerts.KeyFile != "" {
		certFunc, err := certs.GetCertificatesFromSettings(&settings.TLS.LocalCerts)
		if err != nil {
			return nil, fmt.Errorf("failed to get certificates from settings: %w", err)
		}
		return &tls.Config{
			MinVersion:     tls.VersionTLS12,
			GetCertificate: certFunc,
		}, nil
	}
	certLogger := logger.With().Str("component", "acme").Logger()
	// Configure our ACME cert manager and get a certificate using ACME!
	certService, err := acme.NewCertManager(&acme.CertManagerConfig{
		ACMEConfig: &settings.TLS.ACMEConfig,
		Key:        certPrivateKey,
		HTTPClient: httpClient,
		Logger:     &certLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create ACME cert manager: %w", err)
	}

	err = certService.Start(context.TODO(), logger)
	if err != nil {
		logger.Err(err).Msg("failed to start ACME cert manager")
	}
	return &tls.Config{
		MinVersion:     tls.VersionTLS12,
		GetCertificate: certService.GetCertificate,
		NextProtos:     []string{"http/1.1", "acme-tls/1"},
	}, nil
}
