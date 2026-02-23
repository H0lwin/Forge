package backend

import (
	"forge/internal/generator"
	"forge/internal/system"
)

func Django(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(baseSpec("django", map[string]string{
		"manage.py":               "#!/usr/bin/env python\nimport os\nimport sys\n\nif __name__ == '__main__':\n    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'config.settings.dev')\n    from django.core.management import execute_from_command_line\n    execute_from_command_line(sys.argv)\n",
		"config/settings/base.py": "import os\nfrom pathlib import Path\nimport environ\n\nenv = environ.Env(DEBUG=(bool, False))\nBASE_DIR = Path(__file__).resolve().parent.parent.parent\nenviron.Env.read_env(os.path.join(BASE_DIR, '.env'))\n\nSECRET_KEY = env('SECRET_KEY', default='change-me')\nDEBUG = env('DEBUG')\nALLOWED_HOSTS = env.list('ALLOWED_HOSTS', default=['*'])\n\nINSTALLED_APPS = [\n    'django.contrib.admin',\n    'django.contrib.auth',\n    'django.contrib.contenttypes',\n    'django.contrib.sessions',\n    'django.contrib.messages',\n    'django.contrib.staticfiles',\n    'apps.core',\n]\n\nMIDDLEWARE = [\n    'django.middleware.security.SecurityMiddleware',\n    'django.contrib.sessions.middleware.SessionMiddleware',\n    'django.middleware.common.CommonMiddleware',\n    'django.middleware.csrf.CsrfViewMiddleware',\n    'django.contrib.auth.middleware.AuthenticationMiddleware',\n    'django.contrib.messages.middleware.MessageMiddleware',\n    'django.middleware.clickjacking.XFrameOptionsMiddleware', \n]\n\nROOT_URLCONF = 'config.urls'\n\nTEMPLATES = [{\n    'BACKEND': 'django.template.backends.django.DjangoTemplates',\n    'DIRS': [BASE_DIR / 'templates'],\n    'APP_DIRS': True,\n    'OPTIONS': {\n        'context_processors': [\n            'django.template.context_processors.debug',\n            'django.template.context_processors.request',\n            'django.contrib.auth.context_processors.auth',\n            'django.contrib.messages.context_processors.messages',\n        ],\n    },\n}]\n\nWSGI_APPLICATION = 'config.wsgi.application'\n\nDATABASES = {\n    'default': env.db('DATABASE_URL', default=f'sqlite:///{BASE_DIR}/db.sqlite3')\n}\n\nSTATIC_URL = 'static/'\nSTATIC_ROOT = BASE_DIR / 'staticfiles'\nSTATICFILES_DIRS = [BASE_DIR / 'static']\n",
		"config/settings/dev.py":  "from .base import *\nDEBUG = True\n",
		"config/settings/prod.py": "from .base import *\nDEBUG = False\n",
		"config/urls.py":          "from django.contrib import admin\nfrom django.urls import path\n\nurlpatterns = [\n    path('office/', admin.site.urls),\n]\n",
		"config/wsgi.py":          "import os\nfrom django.core.wsgi import get_wsgi_application\nos.environ.setdefault('DJANGO_SETTINGS_MODULE', 'config.settings.prod')\napplication = get_wsgi_application()\n",
		"apps/core/__init__.py":   "",
		"templates/.gitkeep":      "",
		"static/.gitkeep":         "",
		"requirements.txt":        "Django>=5.0\ndjango-environ>=0.11.2\n",
		".gitignore":              ".venv/\n__pycache__/\n.env\ndb.sqlite3\nstaticfiles/\n",
		".env":                    "DEBUG=True\nSECRET_KEY=dev-secret-key-123\nALLOWED_HOSTS=*\n",
	}, []string{"python manage.py runserver"}))
}

func FastAPI(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(baseSpec("fastapi", map[string]string{
		"app/main.py":      "from fastapi import FastAPI\n\napp = FastAPI(title='FastAPI App')\n\n@app.get('/health')\ndef health():\n    return {'status': 'ok'}\n",
		"requirements.txt": "fastapi\nuvicorn[standard]\npydantic-settings\n",
		"alembic.ini":      "[alembic]\nscript_location = alembic\n",
		".gitignore":       ".venv/\n__pycache__/\n.env\n",
	}, []string{"uvicorn app.main:app --reload"}))
}

func Flask(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(baseSpec("flask", map[string]string{
		"app/__init__.py":  "from flask import Flask\n\ndef create_app():\n    app = Flask(__name__)\n\n    @app.get('/health')\n    def health():\n        return {'status': 'ok'}\n\n    return app\n",
		"run.py":           "from app import create_app\napp = create_app()\n",
		"requirements.txt": "flask\n",
		".gitignore":       ".venv/\n__pycache__/\n.env\n",
	}, []string{"flask --app run.py run --debug"}))
}

func Express(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(baseSpec("express", map[string]string{
		"src/index.js": "const express = require('express');\nconst app = express();\napp.get('/health', (_, res) => res.json({status: 'ok'}));\napp.listen(3000);\n",
		"package.json": "{\n  \"name\": \"express-app\",\n  \"version\": \"1.0.0\",\n  \"scripts\": {\"dev\": \"node src/index.js\", \"start\": \"node src/index.js\"},\n  \"dependencies\": {\"express\": \"^4.21.2\"}\n}\n",
		".gitignore":   "node_modules/\n.env\n",
	}, []string{"npm run dev"}))
}

func NestJS(fs system.FileSystem, ex system.Executor) generator.Generator {
	b := generator.NewGenericBuilder(fs, ex)
	return b.New(baseSpec("nestjs", map[string]string{
		"src/main.ts":   "import { NestFactory } from '@nestjs/core';\nimport { Module } from '@nestjs/common';\n\n@Module({})\nclass AppModule {}\n\nasync function bootstrap() {\n  const app = await NestFactory.create(AppModule);\n  await app.listen(3000);\n}\nbootstrap();\n",
		"package.json":  "{\n  \"name\": \"nestjs-app\",\n  \"version\": \"1.0.0\",\n  \"scripts\": {\"start:dev\": \"ts-node src/main.ts\"},\n  \"dependencies\": {\"@nestjs/common\": \"^10.0.0\", \"@nestjs/core\": \"^10.0.0\", \"reflect-metadata\": \"^0.2.0\", \"rxjs\": \"^7.8.0\"}\n}\n",
		"tsconfig.json": "{\"compilerOptions\": {\"target\": \"ES2020\", \"module\": \"commonjs\"}}\n",
	}, []string{"npm run start:dev"}))
}

func baseSpec(name string, files map[string]string, next []string) generator.Spec {
	tools := []string{"git"}
	if name == "express" || name == "nestjs" {
		tools = append(tools, "node")
	} else {
		tools = append(tools, "python3|python")
	}
	return generator.Spec{
		Name:      name,
		Category:  "backend",
		Tools:     tools,
		Files:     files,
		Next:      next,
		Bootstrap: []system.Command{},
	}
}
