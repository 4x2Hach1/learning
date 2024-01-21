package loggers

import (
	"context"
	"errors"

	"github.com/4x2Hach1/learning/gql-go/api/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func SetErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	utils.LoggingInfo()("error-resolver-log", "err", err)

	var gqlerr *gqlerror.Error
	if errors.As(err, &gqlerr) {
		return gqlerr
	}

	return nil
}

func SetRecoverFunc(ctx context.Context, v interface{}) error {
	utils.LoggingWarn()("error-recover-log", "err", v)
	return errors.New("error")
}
