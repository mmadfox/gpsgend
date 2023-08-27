package middleware

import "github.com/mmadfox/gpsgend/internal/generator"

type Middleware func(generator.Service) generator.Service
