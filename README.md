# terraspot

Go CLI scanning instances to spot resources managed outside Terraform.

## Instances


### GitLab

```hcl
instance "vault" {
    url = "gitlab.example.com"
}
```

### Vault

```hcl
instance "vault" {
    url = "vault.example.com"
}
```

## Terraform projects

```hcl
terraform "example" {
    path = "~/example/example"
    workspace = "default"
}
```
