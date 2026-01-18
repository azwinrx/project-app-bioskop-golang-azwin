package wire

import (
	"net/http"
	"project-app-bioskop-golang-azwin/internal/adaptor"
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/internal/middleware"
	"project-app-bioskop-golang-azwin/internal/usecase"
	"project-app-bioskop-golang-azwin/pkg/database"
	"project-app-bioskop-golang-azwin/pkg/utils"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type App struct {
	Router *chi.Mux
	Logger *zap.Logger
}

func InitializeApp(db database.PgxIface, logger *zap.Logger, config utils.Configuration) *App {
	// Initialize validator
	validate := validator.New()

	// Repositories
	userRepo := repository.NewUsersRepository(db, logger)
	cinemaRepo := repository.NewCinemasRepository(db, logger)
	seatRepo := repository.NewSeatsRepository(db, logger)
	showtimeRepo := repository.NewShowtimesRepository(db, logger)
	bookingRepo := repository.NewBookingsRepository(db, logger)
	bookingSeatRepo := repository.NewBookingSeatsRepository(db, logger)
	paymentMethodRepo := repository.NewPaymentMethodsRepository(db, logger)
	paymentRepo := repository.NewPaymentsRepository(db, logger)

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, logger)
	cinemaUsecase := usecase.NewCinemaUsecase(cinemaRepo, logger)
	bookingUsecase := usecase.NewBookingUsecase(
		bookingRepo,
		bookingSeatRepo,
		seatRepo,
		showtimeRepo,
		cinemaRepo,
		logger,
	)
	paymentUsecase := usecase.NewPaymentUsecase(
		paymentRepo,
		paymentMethodRepo,
		bookingRepo,
		logger,
	)

	// Adaptors - aggregated
	adaptors := adaptor.NewAdaptor(
		authUsecase,
		cinemaUsecase,
		bookingUsecase,
		paymentUsecase,
		validate,
		logger,
	)

	// Setup Router
	router := setupRouter(adaptors, logger, userRepo)

	return &App{
		Router: router,
		Logger: logger,
	}
}

func setupRouter(
	adaptors *adaptor.Adaptor,
	logger *zap.Logger,
	userRepo repository.UsersRepository,
) *chi.Mux {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.LoggingMiddleware(logger))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Public routes - Authentication
		r.Post("/register", adaptors.AuthAdaptor.Register)
		r.Post("/login", adaptors.AuthAdaptor.Login)

		// Public routes - Cinemas
		r.Get("/cinemas", adaptors.CinemaAdaptor.GetAllCinemas)
		r.Get("/cinemas/{cinemaId}", adaptors.CinemaAdaptor.GetCinemaByID)

		// Public routes - Seats availability
		r.Get("/cinemas/{cinemaId}/seats", adaptors.BookingAdaptor.GetSeatsAvailability)

		// Public routes - Payment methods
		r.Get("/payment-methods", adaptors.PaymentAdaptor.GetPaymentMethods)

		// Protected routes - require authentication
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(logger, userRepo))

			// Auth
			r.Post("/logout", adaptors.AuthAdaptor.Logout)

			// Bookings
			r.Post("/booking", adaptors.BookingAdaptor.CreateBooking)
			r.Get("/user/bookings", adaptors.BookingAdaptor.GetUserBookings)

			// Payments
			r.Post("/pay", adaptors.PaymentAdaptor.ProcessPayment)
		})
	})

	return r
}
