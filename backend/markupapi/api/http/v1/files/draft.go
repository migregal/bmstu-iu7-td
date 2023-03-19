package files

import (
	"github.com/gothing/draft"

	"markup2/markupapi/api/http/v1/files/add"
	"markup2/markupapi/api/http/v1/files/del"
	"markup2/markupapi/api/http/v1/files/get"
	"markup2/markupapi/api/http/v1/files/upd"
)


var Service = draft.Compose(
	"Файлы",
	add.Endpoint,
	get.Endpoint,
	del.Endpoint,
	upd.Endpoint,
)
