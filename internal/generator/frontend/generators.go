package frontend

import (
	"forge/internal/generator"
	"forge/internal/system"
)

func Next(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(spec("next", map[string]string{
		"package.json": "{\n  \"name\": \"next-app\",\n  \"version\": \"1.0.0\",\n  \"scripts\": {\"dev\": \"next dev\"},\n  \"dependencies\": {\"next\": \"^15.0.0\", \"react\": \"^18.3.0\", \"react-dom\": \"^18.3.0\"}\n}\n",
		"app/page.tsx": "export default function Home() { return <main>Forge Next app</main>; }\n",
		"tsconfig.json": "{\"compilerOptions\": {\"target\": \"ES2020\"}}\n",
	}, []string{"npm run dev"}))
}

func Vite(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(spec("vite", map[string]string{
		"package.json": "{\n  \"name\": \"vite-app\",\n  \"version\": \"1.0.0\",\n  \"scripts\": {\"dev\": \"vite\"},\n  \"dependencies\": {\"react\": \"^18.3.0\", \"react-dom\": \"^18.3.0\"},\n  \"devDependencies\": {\"vite\": \"^5.4.0\"}\n}\n",
		"src/main.tsx": "import React from 'react';\nimport ReactDOM from 'react-dom/client';\nReactDOM.createRoot(document.getElementById('root')!).render(<h1>Forge Vite app</h1>);\n",
		"index.html": "<!doctype html><html><body><div id='root'></div><script type='module' src='/src/main.tsx'></script></body></html>\n",
	}, []string{"npm run dev"}))
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
