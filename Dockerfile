FROM go1.24.2 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 5. アプリケーションをビルド
# -oオプションで出力ファイル名を指定
RUN go build -o my-blog-api .

# 6. ポートを公開
EXPOSE 8080

# 7. コンテナ起動時に実行するコマンド
CMD ["./my-blog-api"]