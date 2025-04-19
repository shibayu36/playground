// src/index.ts
import { McpServer } from '@modelcontextprotocol/sdk/server/mcp.js'
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js'

const server = new McpServer({
  name: 'server-time',
  version: '0.0.1',
  capabilities: {
    resources: {},
    tools: {},
  },
})

server.tool(
  'get-server-time',
  'サーバーの現在時刻を取得します。',
  {},
  async () => {
    return {
      content: [
        {
          type: 'text',
          text: `サーバーの現在時刻は ${new Date().toLocaleString()} です。`,
        },
      ],
    }
  },
)

async function main() {
  const transport = new StdioServerTransport()
  await server.connect(transport)
  console.log('Server is running...')
}

main().catch((error) => {
  console.error('Error starting server:', error)
  process.exit(1)
})
