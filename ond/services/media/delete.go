package media

import (
	"context"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"net/http"
)

func (i impl) Delete(ctx context.Context, fileID string) *errors.ErrResponse {
	endpoint := fmt.Sprintf(resourceURL, "/"+fileID)

	_, respErr := i.client.Do(ctx, i.opts, http.MethodDelete, endpoint, nil)

	return respErr
}
