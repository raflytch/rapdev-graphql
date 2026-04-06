package graphqldelivery

import (
	"encoding/json"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

type requestBody struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func NewHandler(schema graphql.Schema) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params requestBody

		if c.Method() == fiber.MethodGet {
			params = parseGetParams(c)
		} else {
			if err := c.BodyParser(&params); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"errors": []fiber.Map{{"message": "invalid request body"}},
				})
			}
		}

		if params.Query == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": []fiber.Map{{"message": "query is required"}},
			})
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			OperationName:  params.OperationName,
			VariableValues: params.Variables,
			Context:        c.UserContext(),
		})

		if result.HasErrors() {
			return c.Status(fiber.StatusOK).JSON(result)
		}

		return c.JSON(result)
	}
}

func NewPlaygroundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(playgroundHTML)
	}
}

func parseGetParams(c *fiber.Ctx) requestBody {
	var params requestBody
	params.Query = c.Query("query")
	params.OperationName = c.Query("operationName")

	varsRaw := c.Query("variables")
	if varsRaw != "" {
		decoded, err := url.QueryUnescape(varsRaw)
		if err == nil {
			json.Unmarshal([]byte(decoded), &params.Variables)
		}
	}

	return params
}

const playgroundHTML = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Rapdev Portfolio - GraphQL Playground</title>
  <link rel="stylesheet" href="https://unpkg.com/graphiql@3.0.6/graphiql.min.css" />
</head>
<body style="margin: 0; overflow: hidden;">
  <div id="graphiql" style="height: 100vh;"></div>
  <script crossorigin src="https://unpkg.com/react@18/umd/react.production.min.js"></script>
  <script crossorigin src="https://unpkg.com/react-dom@18/umd/react-dom.production.min.js"></script>
  <script crossorigin src="https://unpkg.com/graphiql@3.0.6/graphiql.min.js"></script>
  <script>
    const fetcher = GraphiQL.createFetcher({ url: '/graphql' });
    ReactDOM.createRoot(document.getElementById('graphiql')).render(
      React.createElement(GraphiQL, { fetcher: fetcher })
    );
  </script>
</body>
</html>`
