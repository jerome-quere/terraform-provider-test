# Terraform provider test

This repo aim to trigger an issue when StateFunc in used on a property inside a Set. (https://github.com/hashicorp/terraform/issues/21641)

### Step to reproduce
```
env TF_ACC=1 TF_LOG=WARN go test -v ./test
```

The If statement phone_book.go:70 is trigger but it should not.


