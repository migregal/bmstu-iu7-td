package files

import (
	"github.com/gothing/draft"

	"markup2/markupapi/api/http/v1/files/upd"
)


var Service = draft.Compose(
	"Файлы",
	upd.Endpoint,
)
