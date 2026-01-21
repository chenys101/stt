# ui/Dockerfile
FROM node:18-alpine AS builder
WORKDIR /app

# Install deps
COPY package.json package-lock.json* ./
RUN npm ci

# Copy source and build
COPY . .
# 注入前端运行时变量，例如 VITE_API_BASE
ARG VITE_API_BASE
ENV VITE_API_BASE=${VITE_API_BASE}
RUN npm run build

# Production image served by nginx
FROM nginx:stable-alpine AS runner
COPY --from=builder /app/dist /usr/share/nginx/html
# 可添加自定义 nginx 配置
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]