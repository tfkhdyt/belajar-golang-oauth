package index

import "github.com/gofiber/fiber/v2"

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (i *IndexHandler) Index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(`<body>
      <body>
        <a href='/auth/login/github'>
          <button>
            Login with GitHub
          </button>
        </a>
      </body>
    </body>`)
}
