package models

import "github.com/omenstudio/fastgql/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
