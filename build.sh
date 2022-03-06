# build main
go build -o ./cli/main ./main.go

# build cmds
ls ./cmd | xargs -I {} go build -o ./cli/{} ./cmd/{}