//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package nearestneighbors

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/additional"
	"github.com/semi-technologies/weaviate/entities/search"
	txt2vecmodels "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/additional/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtender(t *testing.T) {
	f := &fakeContextionary{}
	e := NewExtender(f)

	t.Run("with empty results", func(t *testing.T) {
		testData := []search.Result(nil)
		expectedResults := []search.Result(nil)

		res, err := e.Multi(context.Background(), testData, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
	})

	t.Run("with a single result", func(t *testing.T) {
		testData := &search.Result{
			Schema: map[string]interface{}{"name": "item1"},
			Vector: []float32{0.1, 0.3, 0.5},
			AdditionalProperties: map[string]interface{}{
				"classification": &additional.Classification{ // verify it doesn't remove existing additional props
					ID: strfmt.UUID("123"),
				},
			},
		}

		expectedResult := &search.Result{
			Schema: map[string]interface{}{"name": "item1"},
			Vector: []float32{0.1, 0.3, 0.5},
			AdditionalProperties: map[string]interface{}{
				"classification": &additional.Classification{ // verify it doesn't remove existing additional props
					ID: strfmt.UUID("123"),
				},
				"nearestNeighbors": &txt2vecmodels.NearestNeighbors{
					Neighbors: []*txt2vecmodels.NearestNeighbor{
						&txt2vecmodels.NearestNeighbor{
							Concept:  "word1",
							Distance: 1,
						},
						&txt2vecmodels.NearestNeighbor{
							Concept:  "word2",
							Distance: 2,
						},
						&txt2vecmodels.NearestNeighbor{
							Concept:  "word3",
							Distance: 3,
						},
					},
				},
			},
		}

		res, err := e.Single(context.Background(), testData, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResult, res)
	})

	t.Run("with multiple results", func(t *testing.T) {
		vectors := [][]float32{
			[]float32{0.1, 0.2, 0.3},
			[]float32{0.11, 0.22, 0.33},
			[]float32{0.111, 0.222, 0.333},
		}

		testData := []search.Result{
			search.Result{
				Schema: map[string]interface{}{"name": "item1"},
				Vector: vectors[0],
			},
			search.Result{
				Schema: map[string]interface{}{"name": "item2"},
				Vector: vectors[1],
			},
			search.Result{
				Schema: map[string]interface{}{"name": "item3"},
				Vector: vectors[2],
				AdditionalProperties: map[string]interface{}{
					"classification": &additional.Classification{ // verify it doesn't remove existing additional props
						ID: strfmt.UUID("123"),
					},
				},
			},
		}

		expectedResults := []search.Result{
			search.Result{
				Schema: map[string]interface{}{"name": "item1"},
				Vector: vectors[0],
				AdditionalProperties: map[string]interface{}{
					"nearestNeighbors": &txt2vecmodels.NearestNeighbors{
						Neighbors: []*txt2vecmodels.NearestNeighbor{
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word1",
								Distance: 1,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word2",
								Distance: 2,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word3",
								Distance: 3,
							},
						},
					},
				},
			},
			search.Result{
				Schema: map[string]interface{}{"name": "item2"},
				Vector: vectors[1],
				AdditionalProperties: map[string]interface{}{
					"nearestNeighbors": &txt2vecmodels.NearestNeighbors{
						Neighbors: []*txt2vecmodels.NearestNeighbor{
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word4",
								Distance: 0.1,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word5",
								Distance: 0.2,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word6",
								Distance: 0.3,
							},
						},
					},
				},
			},
			search.Result{
				Schema: map[string]interface{}{"name": "item3"},
				Vector: vectors[2],
				AdditionalProperties: map[string]interface{}{
					"classification": &additional.Classification{ // verify it doesn't remove existing additional props
						ID: strfmt.UUID("123"),
					},
					"nearestNeighbors": &txt2vecmodels.NearestNeighbors{
						Neighbors: []*txt2vecmodels.NearestNeighbor{
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word7",
								Distance: 1.1,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word8",
								Distance: 2.2,
							},
							&txt2vecmodels.NearestNeighbor{
								Concept:  "word9",
								Distance: 3.3,
							},
						},
					},
				},
			},
		}

		res, err := e.Multi(context.Background(), testData, nil)
		require.Nil(t, err)
		assert.Equal(t, expectedResults, res)
		assert.Equal(t, f.calledWithVectors, vectors)
	})
}

type fakeContextionary struct {
	calledWithVectors [][]float32
}

func (f *fakeContextionary) MultiNearestWordsByVector(ctx context.Context, vectors [][]float32, k, n int) ([]*txt2vecmodels.NearestNeighbors, error) {
	f.calledWithVectors = vectors
	out := []*txt2vecmodels.NearestNeighbors{
		&txt2vecmodels.NearestNeighbors{
			Neighbors: []*txt2vecmodels.NearestNeighbor{
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word1",
					Distance: 1.0,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word2",
					Distance: 2.0,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "$THING[abc]",
					Distance: 9.99,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word3",
					Distance: 3.0,
					Vector:   nil,
				},
			},
		},

		&txt2vecmodels.NearestNeighbors{
			Neighbors: []*txt2vecmodels.NearestNeighbor{
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word4",
					Distance: 0.1,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word5",
					Distance: 0.2,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word6",
					Distance: 0.3,
					Vector:   nil,
				},
			},
		},

		&txt2vecmodels.NearestNeighbors{
			Neighbors: []*txt2vecmodels.NearestNeighbor{
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word7",
					Distance: 1.1,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word8",
					Distance: 2.2,
					Vector:   nil,
				},
				&txt2vecmodels.NearestNeighbor{
					Concept:  "word9",
					Distance: 3.3,
					Vector:   nil,
				},
			},
		},
	}

	return out[:len(vectors)], nil // return up to three results, but fewer if the input is shorter
}
