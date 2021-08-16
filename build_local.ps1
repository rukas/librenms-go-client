go test -v ./...
go build -o go-librenms
cp .\go-librenms $home/go/src/ -force
cp  .\* C:\Source\terraform-provider-librenms\vendor\go-librenms\ -force