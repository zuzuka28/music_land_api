package track

type Service struct {
	*createService
	*deleteService
	*fetchService
	*searchService
}

func NewService(r Repository, fs FileStorage) *Service {
	return &Service{
		createService: newCreateService(r, fs),
		deleteService: newDeleteService(r, fs),
		fetchService:  newFetchService(r),
		searchService: newSearchService(r),
	}
}
