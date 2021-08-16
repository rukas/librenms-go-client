go test -v ./...
go build -o librenms-go-client
cp .\librenms-go-client $home/go/src/ -force
cp  .\* C:\Source\terraform-provider-librenms\vendor\librenms-go-client\ -force