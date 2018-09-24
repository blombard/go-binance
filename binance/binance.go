/*
   binance.go
       Wrapper for the Binance Exchange API

   Authors:
       Pat DePippo  <patrick.depippo@dcrypt.io>
       Matthew Woop <matthew.woop@dcrypt.io>

   To Do:

*/
package binance2

//"errors"

const (
	BaseUrl = "https://api.binance.com"
	//BaseUrl = "http://localhost:3000"
)

type Binance struct {
	client *Client
}

/*
func handleErr(r jsonResponse) error {

    if !r.Success {
        return errors.New(r.Message)
    }

    return nil
}
*/
func New(key, secret string) *Binance {
	client := NewClient(key, secret)
	return &Binance{client}
}
