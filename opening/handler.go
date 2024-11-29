package opening

type OpeningHandler struct {
	Usecase OpeningUsecase
}

// NewOpeningHandler initializes and returns a OpeningHandler instance
func NewOpeningHandler(usecase OpeningUsecase) *OpeningHandler {
	return &OpeningHandler{
		Usecase: usecase,
	}
}
