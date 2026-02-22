package frontend

import (
	"forge/internal/generator"
	"forge/internal/system"
)

func Next(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(spec("next", map[string]string{
		"package.json": "{\n  \"name\": \"next-app\",\n  \"version\": \"1.0.0\",\n  \"private\": true,\n  \"scripts\": {\"dev\": \"next dev\", \"build\": \"next build\", \"start\": \"next start\", \"lint\": \"next lint\"},\n  \"dependencies\": {\"next\": \"^15.0.0\", \"react\": \"^18.3.0\", \"react-dom\": \"^18.3.0\"}\n}\n",
		"app/layout.tsx": "export default function RootLayout({ children }: { children: React.ReactNode }) {\n  return (<html lang='en'><body>{children}</body></html>);\n}\n",
		"app/page.tsx":   "export default function Home() { return <main>Forge Next app</main>; }\n",
		"next.config.mjs": "const nextConfig = {};\nexport default nextConfig;\n",
		"tsconfig.json":   "{\"compilerOptions\": {\"target\": \"ES2020\", \"strict\": true, \"moduleResolution\": \"bundler\"}}\n",
	}, []string{"npm install", "npm run dev"}))
}

func Vite(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(spec("vite", map[string]string{
		"package.json": "{\n  \"name\": \"vite-app\",\n  \"version\": \"1.0.0\",\n  \"private\": true,\n  \"scripts\": {\"dev\": \"vite\", \"build\": \"vite build\", \"preview\": \"vite preview\"},\n  \"dependencies\": {\"react\": \"^18.3.0\", \"react-dom\": \"^18.3.0\"},\n  \"devDependencies\": {\"typescript\": \"^5.6.0\", \"vite\": \"^5.4.0\"}\n}\n",
		"src/main.tsx": "import React from 'react';\nimport ReactDOM from 'react-dom/client';\n\nReactDOM.createRoot(document.getElementById('root')!).render(<h1>Forge Vite app</h1>);\n",
		"index.html": "<!doctype html><html><head><meta charset='UTF-8' /><meta name='viewport' content='width=device-width, initial-scale=1.0' /><title>Vite App</title></head><body><div id='root'></div><script type='module' src='/src/main.tsx'></script></body></html>\n",
		"tsconfig.json": "{\"compilerOptions\": {\"target\": \"ES2020\", \"jsx\": \"react-jsx\"}}\n",
		"vite.config.ts": "import { defineConfig } from 'vite';\n\nexport default defineConfig({});\n",
	}, []string{"npm install", "npm run dev"}))
}

func spec(name string, files map[string]string, next []string) generator.Spec {
	return generator.Spec{
		Name:      name,
		Category:  "frontend",
		Tools:     []string{"git", "node"},
		Files:     files,
		Next:      next,
		Bootstrap: []system.Command{},
	}
}
