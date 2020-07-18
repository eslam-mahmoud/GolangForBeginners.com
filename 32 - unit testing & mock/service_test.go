package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	// t.Run("valid case", func(t *testing.T) {
	// 	result := add(1, 9)
	// 	assert.Equal(t, 10, result)
	// })

	// var tests = []struct {
	// 	a, b int
	// 	want int
	// }{
	// 	{1, 1, 2},
	// 	{1, -1, 0},
	// 	{0, 0, 0},
	// }

	// for _, tt := range tests {
	// 	testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
	// 	t.Run(testname, func(t *testing.T) {
	// 		result := add(tt.a, tt.b)
	// 		assert.Equal(t, tt.want, result)
	// 	})
	// }
}

func TestNewUSer(t *testing.T) {
	// t.Run("valid case", func(t *testing.T) {
	// 	u := newUser("eslam", 10)
	// 	assert.NotNil(t, u)
	// })
	// t.Run("invalid case", func(t *testing.T) {
	// 	u := newUser("eslam", 10)
	// 	assert.Equal(t, 100, u.age)
	// })
}

type mockProduct struct {
	mock.Mock
}

func (p *mockProduct) buy() error {
	args := p.Called()
	return args.Error(0)
}
func (p *mockProduct) sell() error {
	args := p.Called()
	return args.Error(0)
}
func TestBuyProduct(t *testing.T) {
	u := newUser("eslam", 10)
	assert.NotNil(t, u)

	t.Run("test success buy", func(t *testing.T) {
		p := new(mockProduct)
		// set up expectations
		p.On("buy").Return(nil)

		result := u.buyProduct(p)
		assert.NoError(t, result)
	})
	t.Run("test fail buy", func(t *testing.T) {
		p := new(mockProduct)
		// set up expectations
		e := errors.New("error buy producut")
		p.On("buy").Return(e)

		result := u.buyProduct(p)
		assert.Error(t, result)
		assert.EqualError(t, e, result.Error())
	})
}
