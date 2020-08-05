# ssoacehgo

Client SDK untuk bahasa pemograman Go (a.k.a golang)

Contoh penggunaan
```go
package login

import (
	"fmt"

	"github.com/fuadarradhi/ssoacehgo"
)

func GetToken(c *x.Ctx) {

	//get token dari post atau get, sesuai kebutuhan
	token := c.FormValue("token")

	sso, err := ssoacehgo.NewSSOAcehClient("/path/dimana/file/disimpan/sso_secure.json")
  
	if err == nil {
		res, _ := sso.ParseToken(token)
    
		// gunakan hasil parse
		fmt.Println(res.ID)
		fmt.Println(res.Nama)
		fmt.Println(res.Email)
		fmt.Println(res.EmailAlternatif)
		fmt.Println(res.TelegramID)
		fmt.Println(res.HP)
		fmt.Println(res.NIK)
		fmt.Println(res.NIP)
    
		// jika anda membutuhkan avatar
		fmt.Println(fmt.Sprintf("https://sso.acehprov.go.id/static/avatar/%s",res.Avatar))
    
		// jam token ini digenerate
		fmt.Println(res.DateTime)
    
	} else {
		fmt.Print(err)
	}
}
```
