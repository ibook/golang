package main

import (
	"fmt"
	"encoding/json"
	"strconv"
)

type Currency struct{
	TokenName 		string	`json:"TokenName"`
	TokenSymbol 	string	`json:"TokenSymbol"`
	TotalSupply 	uint	`json:"TotalSupply"`
}

type Token struct {
	Currency		map[string]Currency	`json:"Currency"`
}

func (token *Token) transfer (_from *Account, _to *Account, _currency string, _value uint) bool{
	if(_from.Frozen || _to.Frozen) {
		return false
	}
	if(_from.BalanceOf[_currency] >= _value){
		_from.BalanceOf[_currency] -= _value;
		_to.BalanceOf[_currency] += _value;
	}
	return true
}
func (token *Token) initialSupply(_name string, _symbol string, _supply uint, _account *Account){
	token.Currency[_symbol] = Currency{TokenName: _name, TokenSymbol: _symbol, TotalSupply: _supply};
	_account.BalanceOf[_symbol] = _supply
}

func (token *Token) mint(_currency string, _amount uint, _account *Account) {
	
	cur := token.Currency[_currency]
	cur.TotalSupply += _amount;
	token.Currency[_currency] = cur
	_account.BalanceOf[_currency] = _amount;
	
}
func (token *Token) burn(_currency string, _amount uint, _account *Account) {
	if(token.Currency[_currency].TotalSupply >= _amount){
		cur := token.Currency[_currency]
		cur.TotalSupply -= _amount;
		token.Currency[_currency] = cur
		_account.BalanceOf[_currency] -= _amount;
	}
}

type Account struct {
	Name			string	`json:"Name"`
	Frozen			bool	`json:"Frozen"`
	BalanceOf		map[string]uint	`json:"BalanceOf"`
	
}
func (account *Account) balance (_currency string) uint{
	return account.BalanceOf[_currency]
}

func main(){
	token := &Token{Currency: map[string]Currency{}}

	coinbase := &Account{
		Name: "Coinbase",
		Frozen: false,
		BalanceOf: map[string]uint{}}

	token.initialSupply("水果币","Apple",10000, coinbase)
	token.initialSupply("积分币","PPC",10000, coinbase)

	token1, _ := json.Marshal(token)
	fmt.Println(string(token1))

	token.mint("Apple", 10000, coinbase) 

	token2, _ := json.Marshal(token)
	fmt.Println(string(token2))

	token.burn("Apple", 500, coinbase)

	tokenJson, _ := json.Marshal(token)
	fmt.Println(string(tokenJson))

	fmt.Println(strconv.Itoa(int(coinbase.balance("Apple"))))

	

	transfer()
	frozen()
}

func transfer(){
	fmt.Println("transfer -----")
	account1 := &Account{
		Name: "Neo",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	token := &Token{}
	token.transfer(account1,account2,"RMB", 500)

	from1, _ := json.Marshal(account1)
	fmt.Println(string(from1))

	to1, _ := json.Marshal(account2)
	fmt.Println(string(to1))

}

func frozen(){
	fmt.Println("Frozen -----")
	account1 := &Account{
		Name: "Neo",
		Frozen: true,
		BalanceOf: map[string]uint{"RMB":1000}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	token := &Token{}
	token.transfer(account1,account2,"RMB", 500)

	from1, _ := json.Marshal(account1)
	fmt.Println(string(from1))

	to1, _ := json.Marshal(account2)
	fmt.Println(string(to1))

}