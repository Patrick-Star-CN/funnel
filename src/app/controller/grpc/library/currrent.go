package library

import (
	"context"
	"fmt"
	"funnel/app/errors"
	"funnel/app/service/libraryService"
	rpc "funnel/rpc"
)

type LibraryRpc struct{}

func (LibraryRpc) LibraryCurrent(ctx context.Context, in *rpc.LibraryCurrentRequest) (*rpc.LibraryCurrentReply, error) {

	user, err := libraryService.GetUser(in.Username, in.Password)
	if err != nil {

		return nil, err
	}
	books, err := libraryService.GetCurrentBorrow(user)

	if err == errors.ERR_Session_Expired {
		user, err = libraryService.GetUser(in.Username, in.Password)
		if err != nil {
			return nil, err
		}
		books, err = libraryService.GetCurrentBorrow(user)
	}

	if err != nil {
		return nil, err
	}
	fmt.Print(books)
	return nil, err

}
