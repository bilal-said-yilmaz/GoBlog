# Go base image
FROM golang:1.19-alpine as builder

# Çalışma dizini oluştur
WORKDIR /app

# Gerekli dosyaları çalışma dizinine kopyala
COPY go.mod .
COPY go.sum .
RUN go mod  \
    download

COPY . .

# Uygulamayı inşa et
RUN go build -o myapp .

# Minimal base image
FROM alpine:latest

WORKDIR /root/

# Gerekli sertifikaları ekle
RUN apk --no-cache add ca-certificates

# Uygulamayı ve config dosyalarını kopyala
COPY --from=builder /app/myapp .

# Uygulamayı başlat
CMD ["./myapp"]