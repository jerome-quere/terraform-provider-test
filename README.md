# Terraform provider test

This repo aim to trigger a warning i should not get on terraform 0.12.1.

### Step to reproduce
```
env TF_ACC=1 TF_LOG=WARN go test -v ./test
```

Warning that I get
```
2019/06/11 20:53:13 [WARN] Provider "test" produced an unexpected new value for test_server.base, but we are tolerating it because it is using the legacy plugin SDK.
    The following problems may be the cause of any confusing errors from downstream operations:
      - .address: was null, but now cty.StringVal("")

```


