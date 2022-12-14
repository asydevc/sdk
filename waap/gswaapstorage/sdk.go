// author: asydevc <asydev@163.com>
// date: 2021-03-05
package gswaapstorage

import "github.com/asydevc/sdk"

const (
	Name = "gs-waap-storage"
)

// LogAttack 攻击日志存储
func LogAttack(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/log/attack").SetBody(body).Run(ctx)
}
