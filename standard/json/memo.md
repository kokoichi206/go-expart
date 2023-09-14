## MarshalIndent

Marshal + Indent らしい、ほう

``` go
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	b, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
```

## StructTag

Tag の管理はこいつ

https://pkg.go.dev/reflect#StructTag

この辺で Tag の値使ってる

https://go.dev/src/encoding/json/encode.go#1116:~:text=tag%20%3A%3D%20sf.Tag.Get(%22json

```
typeEncoder > newTypeEncoder > newStructEncoder > cachedTypeFields > typeFields
```
                 