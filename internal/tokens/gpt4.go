/*
Copyright © 2025 KubeRocketAI Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tokens

import (
	"context"
	"fmt"

	"github.com/tiktoken-go/tokenizer"
)

// GPT4Calculator implements TokenCalculator for GPT-4 using tiktoken
type GPT4Calculator struct {
	encoding tokenizer.Codec
}

// NewGPT4Calculator creates a new GPT-4 token calculator
func NewGPT4Calculator() (*GPT4Calculator, error) {
	encoding, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GPT-4 tokenizer: %w", err)
	}

	return &GPT4Calculator{
		encoding: encoding,
	}, nil
}

// CalculateTokens calculates the number of tokens for the given text using GPT-4 tokenization
func (g *GPT4Calculator) CalculateTokens(ctx context.Context, text string) (int, error) {
	if text == "" {
		return 0, nil
	}

	tokensCount, err := g.encoding.Count(text)
	if err != nil {
		return 0, fmt.Errorf("failed to count tokens: %w", err)
	}

	return tokensCount, nil
}
