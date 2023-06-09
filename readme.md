cmd/: 这个目录包含程序的入口文件（如 main.go），用于启动整个应用程序。

internal/: 这个目录存放您项目的内部代码。按照 Go 社区推荐的项目布局，将代码放在 internal 目录下可以确保这些代码不会被其他外部项目导入和使用。

internal/api/: 这个目录包含与 API 相关的代码，如处理 HTTP 请求和响应的逻辑。

internal/api/e: 存放状态码和msg信息

internal/api/handlers/: 这里存放 API 的处理函数。每个文件代表一个功能模块，例如： user.go: 处理与用户相关的请求。

internal/api/middleware/: 这里存放 API 的中间件，用于处理跨越多个处理函数的通用逻辑。

internal/api/routes/: 这里存放 API 的路由配置。

internal/config/: 这个目录包含与项目配置相关的代码。

internal/database/: 这个目录包含与数据库连接和操作相关的代码。

internal/models/: 这个目录包含定义数据模型的文件。每个文件定义了与特定领域相关的数据结构。

internal/services/: 这个目录包含服务层的代码，用于处理业务逻辑。

internal/utils/: 这个目录包含一些通用的工具函数和库。

hash.go: 提供密码哈希和验证功能。

jwt.go: 用于生成和验证 JWT（JSON Web Token）。

validation.go: 提供一些输入验证功能。

go.mod: Go 项目的模块定义文件。它定义了项目的依
