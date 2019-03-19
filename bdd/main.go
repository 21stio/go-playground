package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/metamatex/metamate/pkg/types"
	"github.com/metamatex/metamate/pkg/utils/ptr"
	"github.com/21stio/go-playground/bdd/bdd"
)

func main() {
	s := bdd.Spec{}

	s.Describe("selecting", func(d *bdd.Description) {
		d.Context("selecting fields", func(c *bdd.Context) {
			req := types.GetProductsRequest{
				Select: &types.GetProductsResponseSelect{
					Products: &types.ProductsSelect{
						Info: &types.InfoSelect{
							Name: &types.TextSelect{
								Value: ptr.Bool(true),
							},
							Description: &types.TextSelect{
								Value: ptr.Bool(true),
							},
						},
					},
				},
			}

			c.Input = req

			c.It("MUST only return those fields", func(t *bdd.It) {
				t.Expected = types.GetProductsResponse{
					Products: []types.Product{
						types.Product{
							Info: &types.Info{
								Name: &types.Text{
									Value: ptr.String("chio chips"),
								},
								Description: &types.Text{
									Value: ptr.String("super crunchy"),
								},
							},
						},
					},
				}

				var a *bdd.It
				assert.Equal(t, 1, 1)
				require.Nil(t, a)
			})
		})
	})

	s.Run()
	s.Print(1)
}
