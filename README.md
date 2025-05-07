# 🧩 sp - Snippet CLI Tool

`sp` es una CLI minimalista y útil para guardar, visualizar e **inyectar fragmentos de código (snippets)** en archivos. Ideal para automatizar la reutilización de bloques de texto o código en proyectos.

---

## 🚀 Características

- Crear nuevos snippets desde la terminal
- Listar todos los snippets existentes
- Mostrar el contenido de un snippet
- Inyectar un snippet dentro de un archivo (o crear uno nuevo si no existe)
- Almacenar los snippets localmente en tu sistema

---

## 📦 Instalación

Compílalo desde el código fuente:

```bash
go build -o sp
mv sp /usr/local/bin/
```

---

## 🔰 Uso básico

### 1. Crear un snippet

```bash
sp new dockerignore "vendor/\n*.log\ndocker-compose.override.yml"

```

### 2. Listar snippets

```bash
sp list
```

### 3. Ver contenido de un snippet

```bash
sp show dockerignore
```

### 4, Crear o actualizar un archivo con el snippet

```bash
sp inject dockerignore .dockerignore
```
