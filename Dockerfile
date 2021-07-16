# Frontend Build
FROM node:16-alpine AS frontend
WORKDIR /app/
ADD . .
RUN npm install && npm run build

# Backend Build
FROM golang:1.16-alpine AS backend
WORKDIR /app
ADD . .
RUN go install ./...

# Production Image
FROM alpine:3.11
COPY --from=frontend /app/dist ./dist
COPY --from=backend /go/bin/backend ./backend
CMD ["./backend"]