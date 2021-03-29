package mock

//RepositoryMock ...
type RepositoryMock struct{}

const ConnectMock = 66

//ConnectName ...
func (r RepositoryMock) ConnectName() int { return ConnectMock }
